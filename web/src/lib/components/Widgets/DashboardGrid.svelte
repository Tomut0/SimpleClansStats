<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { SvelteSet } from 'svelte/reactivity';
	import { GridStack, type GridStackWidget } from 'gridstack';
	import Plus from '@lucide/svelte/icons/plus';
	import Trash2 from '@lucide/svelte/icons/trash-2';
	import type { DashboardData, DashboardLayout } from '$lib/api/generated';
	import 'gridstack/dist/gridstack.min.css';
	import { m } from '$lib/paraglide/messages.js';
	import * as Dialog from '$lib/components/ui/dialog';
	import DashboardWidgetContent from './DashboardWidgetContent.svelte';
	import {
		configurableFields,
		configurableTypes,
		defaults,
		isCounter,
		validWidgets,
		type WidgetLayout
	} from './dashboard-layout';
	import { makeWidgetData } from './dashboard-data';
	import { configuredFieldTitle, configuredTitle, configuredTypeTitle } from './dashboard-labels';
	import type {
		ConfigurableWidgetField,
		ConfigurableWidgetType,
		CounterFormat,
		WidgetData
	} from './types';

	let {
		serverId,
		apiUrl,
		avatarUrlTemplate,
		editing = $bindable(false),
		onActionsReady = () => {},
		onSnapshot = () => {}
	}: {
		serverId: string;
		apiUrl: string;
		avatarUrlTemplate: string;
		editing?: boolean;
		onActionsReady?: (actions: {
			addWidget: () => void;
			restoreDefaults: () => void;
			saveDashboard: () => void;
		}) => void;
		onSnapshot?: (received: boolean) => void;
	} = $props();

	let widgets = $state<WidgetLayout[]>(defaults.map((widget) => ({ ...widget })));
	let gridElement: HTMLDivElement;
	let grid: GridStack | undefined;
	let data = $state<DashboardData | null>(null);
	let saveError = $state(false);
	let addWidgetOpen = $state(false);
	let configuringType = $state<ConfigurableWidgetType | null>(null);
	let selectedField = $state<ConfigurableWidgetField>('clans');
	let counterFormat = $state<CounterFormat>('number');
	let tableLimit = $state(10);
	let hoverTile = $state<{ x: number; y: number; width: number; height: number } | null>(null);
	let mounted = new SvelteSet<string>();
	let available = $derived(widgets.length >= 24 ? [] : configurableTypes);
	let widgetData = $derived<WidgetData>(makeWidgetData(data));

	onMount(() => {
		let source: EventSource | undefined;
		let disposed = false;
		onSnapshot(false);
		onActionsReady({
			addWidget: openWidgetCatalog,
			restoreDefaults,
			saveDashboard: () => void saveDashboard()
		});
		const initialize = async () => {
			try {
				const response = await fetch(`${apiUrl}/dashboard/layout`, { credentials: 'include' });
				if (response.ok) {
					const saved = (await response.json()) as { version: number; widgets: WidgetLayout[] };
					if (saved.version === 1 && saved.widgets.length && validWidgets(saved.widgets))
						widgets = saved.widgets;
				} else throw new Error('layout unavailable');
			} catch {
				// Keep the built-in layout if the shared layout cannot be loaded.
			}
			await tick();
			if (disposed) return;
			connectGrid();
			source = new EventSource(
				`${apiUrl}/dashboard/events?serverId=${encodeURIComponent(serverId)}`,
				{ withCredentials: true }
			);
			source.addEventListener('snapshot', (event) => {
				data = JSON.parse((event as MessageEvent<string>).data) as DashboardData;
				onSnapshot(true);
			});
		};
		void initialize();
		return () => {
			disposed = true;
			source?.close();
			grid?.destroy(false);
		};
	});

	$effect(() => {
		const count = widgets.length;
		const editingNow = editing;
		if (!grid) return;
		grid.setStatic(!editingNow);
		if (!count && !editingNow) return;
		void tick().then(() => {
			for (const node of gridElement.querySelectorAll<HTMLElement>('.grid-stack-item')) {
				const id = node.dataset.widgetId;
				if (id && !mounted.has(id)) {
					const widget = widgets.find((item) => item.id === id);
					if (widget) grid?.makeWidget(node, gridOptions(widget));
					mounted.add(id);
				}
			}
		});
	});

	function addConfiguredWidget() {
		if (widgets.length >= 24 || !configuringType) return;
		const isCount = configuringType === 'counter';
		widgets = [
			...widgets,
			{
				id: `${isCount ? 'counter' : 'widget'}-${crypto.randomUUID()}`,
				type: configuringType,
				field: selectedField,
				...(isCount ? { format: counterFormat } : {}),
				...(configuringType === 'table' ? { limit: tableLimit } : {}),
				w: isCount ? 2 : 3,
				h: 2,
				minW: isCount ? 2 : 3,
				minH: 2
			} as WidgetLayout
		];
		configuringType = null;
		addWidgetOpen = false;
	}
	function configureWidgetType(type: ConfigurableWidgetType) {
		configuringType = type;
		selectedField = configurableFields[type][0];
		counterFormat = 'number';
		tableLimit = 10;
	}
	function openWidgetCatalog() {
		configuringType = null;
		addWidgetOpen = true;
	}
	function setEditing(value: boolean) {
		editing = value;
		grid?.setStatic(!value);
	}
	function moveAddTile(event: PointerEvent) {
		if (!editing || event.pointerType === 'touch') return;
		if (event.target instanceof Element && event.target.closest('.add-widget-hover')) return;
		if (
			event.buttons ||
			(event.target instanceof Element && event.target.closest('.remove-widget, .widget-header'))
		) {
			hoverTile = null;
			return;
		}
		if (!grid) return;
		const columns = grid.getColumn();
		const rows = grid.getRow();
		const width = Math.min(3, columns);
		const height = 2;
		if (rows < height) {
			hoverTile = null;
			return;
		}
		const cell = grid.getCellFromPixel({ left: event.pageX, top: event.pageY }, true);
		const cellWidth = gridElement.clientWidth / columns;
		const cellHeight = grid.getCellHeight(true);
		const x = Math.max(0, Math.min(cell.x, columns - width));
		const y = Math.max(0, Math.min(cell.y, rows - height));
		if (!grid.isAreaEmpty(x, y, width, height)) {
			hoverTile = null;
			return;
		}
		hoverTile = {
			x: x * cellWidth,
			y: y * cellHeight,
			width: width * cellWidth,
			height: height * cellHeight
		};
	}
	function removeWidget(id: string) {
		const node = gridElement.querySelector<HTMLElement>(`[data-widget-id="${id}"]`);
		if (node && grid) grid.removeWidget(node, false);
		mounted.delete(id);
		widgets = widgets.filter((widget) => widget.id !== id);
	}
	function restoreDefaults() {
		grid?.destroy(false);
		grid = undefined;
		mounted.clear();
		widgets = defaults.map((widget) => ({ ...widget }));
		void tick().then(() => {
			connectGrid();
			grid?.setStatic(!editing);
		});
	}
	async function saveDashboard() {
		const layout = serializeLayout();
		try {
			const response = await fetch(`${apiUrl}/dashboard/layout`, {
				method: 'PUT',
				credentials: 'include',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(layout)
			});
			if (!response.ok) throw new Error('save failed');
			saveError = false;
			setEditing(false);
		} catch {
			saveError = true;
		}
	}
	function gridOptions(widget: WidgetLayout): GridStackWidget {
		return { ...widget, minW: isCounter(widget) ? 2 : 3, minH: 2 };
	}
	function connectGrid() {
		for (const widget of widgets) {
			const node = gridElement.querySelector<HTMLElement>(`[data-widget-id="${widget.id}"]`);
			if (node) {
				const options = gridOptions(widget);
				for (const [key, value] of Object.entries({
					'gs-id': widget.id,
					'gs-x': widget.x,
					'gs-y': widget.y,
					'gs-w': widget.w,
					'gs-h': widget.h,
					'gs-min-w': options.minW,
					'gs-min-h': options.minH
				}))
					node.setAttribute(key, String(value));
			}
		}
		grid = GridStack.init(
			{
				column: 12,
				cellHeight: 76,
				margin: 8,
				staticGrid: !editing,
				draggable: { handle: '.widget-header' }
			},
			gridElement
		);
		for (const node of gridElement.querySelectorAll<HTMLElement>('.grid-stack-item'))
			mounted.add(node.dataset.widgetId!);
	}
	function serializeLayout(): DashboardLayout {
		const saved = grid?.save(false, false);
		const items = Array.isArray(saved) ? (saved as GridStackWidget[]) : [];
		return {
			version: 1,
			widgets: items
				.filter((item) => widgets.some((widget) => widget.id === String(item.id)))
				.map((item) => {
					const widget = widgets.find((candidate) => candidate.id === String(item.id))!;
					return {
						id: widget.id,
						type: widget.type,
						field: widget.field,
						format: widget.format,
						x: item.x ?? 0,
						y: item.y ?? 0,
						w: item.w ?? 1,
						h: item.h ?? 1
					};
				})
		};
	}
</script>

<div class="mb-4">
	{#if saveError}<div class="text-sm text-muted-foreground" role="status">
			{m['pages.dashboard.save_failed']()}
		</div>{/if}
</div>

<div
	class="grid-stack"
	class:grid-stack-editing={editing}
	bind:this={gridElement}
	role="region"
	aria-label={m['pages.dashboard.title']()}
	onpointermove={moveAddTile}
	onpointerleave={() => (hoverTile = null)}
>
	{#each widgets as widget (widget.id)}
		<div class="grid-stack-item" data-widget-id={widget.id}>
			<div class="grid-stack-item-content relative">
				<DashboardWidgetContent {widget} data={widgetData} {avatarUrlTemplate} />
				{#if editing && widgets.length > 1}
					<button
						class="remove-widget text-destructive-foreground absolute top-2 right-2 z-30 grid size-7 place-items-center rounded bg-destructive hover:bg-destructive/90"
						aria-label={m['pages.dashboard.remove_widget']({ widget: configuredTitle(widget) })}
						onclick={() => removeWidget(widget.id)}
						><Trash2 class="size-4" aria-hidden="true" /></button
					>
				{/if}
			</div>
		</div>
	{/each}
	{#if editing && hoverTile}
		<button
			class="add-widget-hover flex flex-col items-center justify-center gap-2 rounded-lg border border-dashed bg-background/95 text-muted-foreground hover:border-primary hover:text-foreground"
			style:left="{hoverTile.x}px"
			style:top="{hoverTile.y}px"
			style:width="{hoverTile.width}px"
			style:height="{hoverTile.height}px"
			aria-label={m['pages.dashboard.add_widget_button']()}
			onclick={openWidgetCatalog}
		>
			<Plus class="size-7" aria-hidden="true" />
			<span>{m['pages.dashboard.add_widget_button']()}</span>
		</button>
	{/if}
</div>

<Dialog.Root bind:open={addWidgetOpen}>
	<Dialog.Content class="max-w-lg">
		<Dialog.Header>
			<Dialog.Title
				>{configuringType
					? configuredTypeTitle(configuringType)
					: m['pages.dashboard.add_widget_title']()}</Dialog.Title
			>
			<Dialog.Description>{m['pages.dashboard.add_widget_description']()}</Dialog.Description>
		</Dialog.Header>
		{#if configuringType}
			<div class="grid gap-4">
				<label class="grid gap-2 text-sm font-medium">
					{m['pages.dashboard.counter_field_label']()}
					<select class="h-10 rounded-md border bg-background px-3" bind:value={selectedField}>
						{#each configurableFields[configuringType] as field (field)}
							<option value={field}>{configuredFieldTitle(field)}</option>
						{/each}
					</select>
				</label>
				{#if configuringType === 'counter'}
					<label class="grid gap-2 text-sm font-medium">
						{m['pages.dashboard.counter_format_label']()}
						<select class="h-10 rounded-md border bg-background px-3" bind:value={counterFormat}>
							<option value="number">{m['pages.dashboard.counter_format_number']()}</option>
							<option value="compact">{m['pages.dashboard.counter_format_compact']()}</option>
							<option value="plain">{m['pages.dashboard.counter_format_plain']()}</option>
						</select>
					</label>
				{/if}
				{#if configuringType === 'table'}
					<label class="grid gap-2 text-sm font-medium">
						{m['pages.dashboard.table_limit_label']()}
						<select class="h-10 rounded-md border bg-background px-3" bind:value={tableLimit}>
							<option value={5}>5</option><option value={10}>10</option><option value={20}
								>20</option
							>
						</select>
					</label>
				{/if}
				<div class="flex justify-between gap-2">
					<button
						class="rounded-md border px-3 py-2 text-sm"
						onclick={() => (configuringType = null)}
					>
						{m['pages.dashboard.counter_back']()}
					</button>
					<button
						class="rounded-md bg-primary px-3 py-2 text-sm text-primary-foreground"
						onclick={addConfiguredWidget}
					>
						{m['pages.dashboard.add_configured_widget']()}
					</button>
				</div>
			</div>
		{:else if available.length}
			<div class="grid gap-2 sm:grid-cols-2">
				{#each available as type (type)}
					<button
						class="flex items-center gap-3 rounded-md border p-3 text-left hover:bg-accent"
						onclick={() => configureWidgetType(type)}
					>
						<Plus class="size-4 shrink-0" aria-hidden="true" />
						<span>{configuredTypeTitle(type)}</span>
					</button>
				{/each}
			</div>
		{:else}
			<p class="text-sm text-muted-foreground">{m['pages.dashboard.widget_limit_reached']()}</p>
		{/if}
	</Dialog.Content>
</Dialog.Root>

<style>
	.grid-stack-editing :global(.widget-header) {
		cursor: move;
	}
	.add-widget-hover {
		position: absolute;
		z-index: 20;
	}
</style>
