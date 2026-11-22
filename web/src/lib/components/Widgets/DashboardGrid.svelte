<script module lang="ts">
	export type DashboardActions = {
		beginEditing: () => void;
		addWidget: () => void;
		restoreDefaults: () => void;
		saveDashboard: () => void;
	};
</script>

<script lang="ts">
	import { tick, untrack } from 'svelte';
	import { GridStack, type GridStackWidget } from 'gridstack';
	import Plus from '@lucide/svelte/icons/plus';
	import Trash2 from '@lucide/svelte/icons/trash-2';
	import type { DashboardData, DashboardLayout } from '$lib/api/generated';
	import 'gridstack/dist/gridstack.min.css';
	import { m } from '$lib/paraglide/messages.js';
	import WidgetCatalogDialog from './WidgetCatalogDialog.svelte';
	import DashboardWidgetContent from './DashboardWidgetContent.svelte';
	import {
		createWidgetId,
		defaultWidgets,
		layoutVersion,
		maxWidgets,
		serializeLayout,
		validWidgets,
		type WidgetConfiguration,
		type WidgetLayout
	} from './dashboard-layout';
	import { makeWidgetData } from './dashboard-data';
	import { dashboardWidgetLabel } from './dashboard-labels';
	import type { WidgetData } from './types';

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
		onActionsReady?: (actions: DashboardActions | null) => void;
		onSnapshot?: (received: boolean) => void;
	} = $props();

	let widgets = $state<WidgetLayout[]>([...defaultWidgets]);
	let grid: GridStack | undefined;
	let data = $state<DashboardData | null>(null);
	let saveError = $state(false);
	let addWidgetOpen = $state(false);
	let hoverTile = $state<{ x: number; y: number; width: number; height: number } | null>(null);
	const limitReached = $derived(widgets.length >= maxWidgets);
	const widgetData = $derived<WidgetData>(makeWidgetData(data));

	function mountDashboard(container: HTMLDivElement) {
		const controller = new AbortController();
		let source: EventSource | undefined;

		async function initializeDashboard() {
			const savedWidgets = await loadSavedWidgets(controller.signal);
			if (controller.signal.aborted) return;
			if (savedWidgets) widgets = savedWidgets;
			await tick();
			if (controller.signal.aborted) return;

			connectGrid(container);
			onActionsReady({
				beginEditing: () => setEditing(true),
				addWidget: openWidgetCatalog,
				restoreDefaults: () => void restoreDefaults(),
				saveDashboard: () => void saveDashboard()
			});
			source = new EventSource(
				`${apiUrl}/dashboard/events?serverId=${encodeURIComponent(serverId)}`,
				{ withCredentials: true }
			);
			source.addEventListener('snapshot', (event) => {
				data = JSON.parse((event as MessageEvent<string>).data) as DashboardData;
				onSnapshot(true);
			});
		}

		untrack(() => {
			onSnapshot(false);
			void initializeDashboard();
		});
		return () => {
			controller.abort();
			source?.close();
			disconnectGrid();
			onActionsReady(null);
		};
	}

	async function loadSavedWidgets(signal: AbortSignal) {
		try {
			const response = await fetch(`${apiUrl}/dashboard/layout`, {
				credentials: 'include',
				signal
			});
			if (!response.ok) return;
			const saved = (await response.json()) as DashboardLayout | null;
			if (saved?.version === layoutVersion && validWidgets(saved.widgets)) return saved.widgets;
		} catch {
			// Keep the built-in layout if the shared layout cannot be loaded.
		}
	}

	function setEditing(value: boolean) {
		editing = value;
		grid?.setStatic(!value);
	}

	async function addConfiguredWidget(configuration: WidgetConfiguration) {
		if (limitReached) return;
		const widget: WidgetLayout = {
			...configuration,
			id: createWidgetId(),
			x: 0,
			y: 0,
			w: configuration.type === 'counter' ? 2 : 3,
			h: 2,
			autoPosition: true
		};
		widgets = [...widgets, widget];
		await tick();

		if (!grid) return;
		const element = findWidgetElement(grid.el, widget.id);
		if (element) grid.makeWidget(element, gridOptions(widget));
	}

	function openWidgetCatalog() {
		addWidgetOpen = true;
	}

	function moveAddTile(event: PointerEvent) {
		if (!grid || !editing || event.pointerType === 'touch') return;
		const target = event.target instanceof Element ? event.target : null;
		if (target?.closest('.add-widget-hover')) return;
		if (event.buttons || target?.closest('.remove-widget, .widget-header')) {
			hoverTile = null;
			return;
		}
		const columns = grid.getColumn();
		const rows = grid.getRow();
		const width = Math.min(3, columns);
		const height = 2;
		if (rows < height) {
			hoverTile = null;
			return;
		}
		const cell = grid.getCellFromPixel({ left: event.pageX, top: event.pageY }, true);
		const x = Math.max(0, Math.min(cell.x, columns - width));
		const y = Math.max(0, Math.min(cell.y, rows - height));
		if (!grid.isAreaEmpty(x, y, width, height)) {
			hoverTile = null;
			return;
		}
		const cellWidth = grid.el.clientWidth / columns;
		const cellHeight = grid.getCellHeight(true);
		hoverTile = {
			x: x * cellWidth,
			y: y * cellHeight,
			width: width * cellWidth,
			height: height * cellHeight
		};
	}

	function findWidgetElement(container: HTMLElement, id: string) {
		return container.querySelector<HTMLElement>(`[data-widget-id="${id}"]`);
	}

	function removeWidget(id: string) {
		if (!grid) return;
		const element = findWidgetElement(grid.el, id);
		if (element) grid.removeWidget(element, false);
		widgets = widgets.filter((widget) => widget.id !== id);
	}

	async function restoreDefaults() {
		const container = grid?.el;
		if (!container) return;
		disconnectGrid();
		widgets = [...defaultWidgets];
		await tick();
		if (container.isConnected) connectGrid(container);
	}

	async function saveDashboard() {
		if (!grid) return;
		const positions = grid.save(false, false);
		const layout = serializeLayout(widgets, Array.isArray(positions) ? positions : []);
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
		const { id, x, y, w, h, autoPosition } = widget;
		return { id, x, y, w, h, autoPosition, minW: widget.type === 'counter' ? 2 : 3, minH: 2 };
	}

	function connectGrid(container: HTMLElement) {
		// GridStack reads the saved geometry from the elements rendered by Svelte.
		for (const widget of widgets) {
			const element = findWidgetElement(container, widget.id);
			if (!element) continue;
			const { id, x, y, w, h, minW, minH } = gridOptions(widget);
			const attributes = {
				'gs-id': id,
				'gs-x': x,
				'gs-y': y,
				'gs-w': w,
				'gs-h': h,
				'gs-min-w': minW,
				'gs-min-h': minH
			};
			for (const [name, value] of Object.entries(attributes)) {
				element.setAttribute(name, String(value));
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
			container
		);
	}

	function disconnectGrid() {
		grid?.destroy(false);
		grid = undefined;
	}
</script>

<div class="mb-4">
	{#if saveError}<div class="text-sm text-muted-foreground" role="status">
			{m['pages.dashboard.save_failed']()}
		</div>{/if}
</div>

<div
	class="grid-stack"
	data-editing={editing}
	{@attach mountDashboard}
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
						aria-label={m['pages.dashboard.remove_widget']({
							widget: dashboardWidgetLabel(widget)
						})}
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

<WidgetCatalogDialog
	bind:open={addWidgetOpen}
	{limitReached}
	onAdd={(configuration) => void addConfiguredWidget(configuration)}
/>

<style>
	.grid-stack[data-editing='true'] :global(.widget-header) {
		cursor: move;
	}
	.add-widget-hover {
		position: absolute;
		z-index: 20;
	}
</style>
