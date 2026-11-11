<script lang="ts">
	import { AreaChart } from 'layerchart';
	import * as Chart from '$lib/components/ui/chart';
	import { m } from '$lib/paraglide/messages.js';
	type Point = { at: string; clans: number; players: number; wars: number; balance: number; onlinePlayers: number; kills: number };
	let { title, metric, history }: { title: string; metric: keyof Omit<Point, 'at'>; history: Point[] } = $props();
	const chartData = $derived(history.map((point) => ({ at: new Date(point.at).toLocaleTimeString(), value: point[metric] })));
	const config = $derived({ value: { label: title, color: 'var(--primary)' } });
</script>

<section class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground">
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{title}</header>
	{#if chartData.length > 1}
		<Chart.Container {config} class="min-h-0 flex-1 p-4" aria-label={title}>
			<AreaChart data={chartData} x="at" y="value">
				{#snippet tooltip()}<Chart.Tooltip />{/snippet}
			</AreaChart>
		</Chart.Container>
	{:else}
		<div class="flex flex-1 items-center justify-center p-4 text-sm text-muted-foreground">{m['pages.dashboard.history_unavailable']()}</div>
	{/if}
</section>
