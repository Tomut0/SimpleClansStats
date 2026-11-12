<script lang="ts">
	import { PieChart } from 'layerchart';
	import * as Chart from '$lib/components/ui/chart';
	import type { DonutSlice } from './types';
	let {
		title,
		items,
		centerLabel,
		emptyMessage
	}: { title: string; items: DonutSlice[]; centerLabel: string; emptyMessage: string } = $props();
	const total = $derived(items.reduce((sum, item) => sum + item.count, 0));
	const visibleItems = $derived(items.filter((item) => item.count > 0));
	const chartData = $derived(
		visibleItems.map((item, index) => ({
			category: `slice-${index}`,
			value: item.count,
			color: item.color
		}))
	);
	const chartConfig = $derived(
		Object.fromEntries(
			visibleItems.map((item, index) => [
				`slice-${index}`,
				{ label: item.label, color: item.color }
			])
		) as Chart.ChartConfig
	);
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{title}</header>
	{#if total}
		<div class="flex min-h-0 flex-1 items-center justify-center gap-5 overflow-auto p-4">
			<div class="relative size-36 shrink-0">
				<Chart.Container
					config={chartConfig}
					class="absolute inset-0 aspect-square h-full w-full"
					aria-label={title}
				>
					<PieChart
						data={chartData}
						key="category"
						value="value"
						c="color"
						innerRadius={0.62}
						padAngle={0.015}
					>
						{#snippet tooltip()}
							<Chart.Tooltip />
						{/snippet}
					</PieChart>
				</Chart.Container>
				<div
					class="pointer-events-none absolute inset-0 z-10 grid place-content-center text-center"
					aria-hidden="true"
				>
					<strong>{total.toLocaleString()}</strong><span class="text-xs">{centerLabel}</span>
				</div>
			</div>
			<ul class="space-y-2">
				{#each items as item (item.label)}
					<li class="flex items-center gap-2 text-sm">
						<span class="size-2.5 rounded-full" style:background={item.color}></span>
						<span>{item.label}</span>
						<span class="ml-auto pl-3">{item.count}</span>
					</li>
				{/each}
			</ul>
		</div>
	{:else}
		<div class="flex flex-1 items-center justify-center p-4 text-sm text-muted-foreground">
			{emptyMessage}
		</div>
	{/if}
</section>
