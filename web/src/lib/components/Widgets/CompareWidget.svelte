<script lang="ts">
	import { BarChart } from 'layerchart';
	import * as Chart from '$lib/components/ui/chart';
	import { m } from '$lib/paraglide/messages.js';
	let { clans }: { clans: Record<string, unknown>[] } = $props();
	const chartData = $derived(
		clans
			.slice(0, 8)
			.map((clan) => ({
				name: String(clan.tag ?? clan.name ?? '—'),
				balance: Number(clan.balance ?? 0),
				players: Number(clan.playerCount ?? clan.memberCount ?? 0)
			}))
	);
	const config = {
		balance: { label: m['pages.dashboard.table_balance'](), color: 'var(--primary)' },
		players: { label: m['pages.dashboard.table_players'](), color: 'var(--chart-2)' }
	};
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">
		{m['pages.dashboard.compare_type']()}
	</header>
	{#if chartData.length}<Chart.Container
			{config}
			class="min-h-0 flex-1 p-4"
			aria-label={m['pages.dashboard.compare_type']()}
			><BarChart
				data={chartData}
				x="name"
				y="balance"
				series={[{ key: 'balance' }, { key: 'players' }]}
				seriesLayout="group">{#snippet tooltip()}<Chart.Tooltip />{/snippet}</BarChart
			></Chart.Container
		>{:else}<div class="flex flex-1 items-center justify-center text-sm text-muted-foreground">
			{m['pages.dashboard.no_clans']()}
		</div>{/if}
</section>
