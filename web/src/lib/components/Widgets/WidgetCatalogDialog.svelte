<script lang="ts">
	import Plus from '@lucide/svelte/icons/plus';
	import type { DashboardWidget } from '#lib/api/generated/index.js';
	import * as Dialog from '#lib/components/ui/dialog/index.js';
	import { m } from '#lib/paraglide/messages.js';
	import {
		configurableFields,
		configurableTypes,
		counterFields,
		tableLimits,
		type WidgetConfiguration
	} from './dashboard-layout';
	import { dashboardLabel } from './dashboard-labels';
	import type { ConfigurableWidgetField, ConfigurableWidgetType, CounterFormat } from './types';

	let {
		open = $bindable(false),
		limitReached,
		onAdd
	}: {
		open?: boolean;
		limitReached: boolean;
		onAdd: (configuration: WidgetConfiguration) => void;
	} = $props();

	let configuringType = $state<ConfigurableWidgetType | null>(null);
	let selectedField = $state<ConfigurableWidgetField>('clans');
	let counterFormat = $state<CounterFormat>('number');
	let tableLimit = $state<NonNullable<DashboardWidget['limit']>>(10);

	function configureWidgetType(type: ConfigurableWidgetType) {
		configuringType = type;
		selectedField = configurableFields[type][0];
		counterFormat = 'number';
		tableLimit = 10;
	}

	function addConfiguredWidget() {
		if (limitReached || !configuringType) return;
		onAdd({
			type: configuringType,
			field: selectedField,
			format: configuringType === 'counter' ? counterFormat : undefined,
			limit: configuringType === 'table' ? tableLimit : undefined
		});
		configuringType = null;
		open = false;
	}
</script>

<Dialog.Root
	bind:open
	onOpenChange={(isOpen) => {
		if (!isOpen) configuringType = null;
	}}
>
	<Dialog.Content class="max-w-lg">
		<Dialog.Header>
			<Dialog.Title
				>{configuringType
					? dashboardLabel(configuringType)
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
							{@const kind = counterFields.some((counter) => counter === field)
								? 'counter'
								: 'widget'}
							<option value={field}>{dashboardLabel(field, kind)}</option>
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
							{#each tableLimits as limit (limit)}
								<option value={limit}>{limit}</option>
							{/each}
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
		{:else if !limitReached}
			<div class="grid gap-2 sm:grid-cols-2">
				{#each configurableTypes as type (type)}
					<button
						class="flex items-center gap-3 rounded-md border p-3 text-left hover:bg-accent"
						onclick={() => configureWidgetType(type)}
					>
						<Plus class="size-4 shrink-0" aria-hidden="true" />
						<span>{dashboardLabel(type)}</span>
					</button>
				{/each}
			</div>
		{:else}
			<p class="text-sm text-muted-foreground">{m['pages.dashboard.widget_limit_reached']()}</p>
		{/if}
	</Dialog.Content>
</Dialog.Root>
