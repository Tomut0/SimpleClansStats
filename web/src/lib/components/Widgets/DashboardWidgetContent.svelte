<script lang="ts">
	import type { DashboardWidget } from '$lib/api/generated';
	import { m } from '$lib/paraglide/messages.js';
	import type { ConfigurableWidgetField, CounterField, WidgetData } from './types';
	import { clanStatusSlices, counterValue, killSlices, tableRows } from './dashboard-data';
	import { dashboardWidgetLabel } from './dashboard-labels';
	import ClanCardWidget from './ClanCardWidget.svelte';
	import CompareWidget from './CompareWidget.svelte';
	import CountWidget from './CountWidget.svelte';
	import DonutWidget from './DonutWidget.svelte';
	import EconomyWidget from './EconomyWidget.svelte';
	import HeatmapWidget from './HeatmapWidget.svelte';
	import KillFeedWidget from './KillFeedWidget.svelte';
	import LatestActionsWidget from './LatestActionsWidget.svelte';
	import NetworkWidget from './NetworkWidget.svelte';
	import OnlinePlayersWidget from './OnlinePlayersWidget.svelte';
	import PeakOnlineWidget from './PeakOnlineWidget.svelte';
	import PlayerCardWidget from './PlayerCardWidget.svelte';
	import RowsWidget from './RowsWidget.svelte';
	import TimeSeriesWidget from './TimeSeriesWidget.svelte';
	import TopClansWidget from './TopClansWidget.svelte';

	let {
		widget,
		data,
		avatarUrlTemplate
	}: {
		widget: DashboardWidget;
		data: WidgetData;
		avatarUrlTemplate: string;
	} = $props();
	const title = $derived(dashboardWidgetLabel(widget));
	const kills = $derived(killSlices(data));
	const statuses = $derived(clanStatusSlices(data));
</script>

{#snippet table(field: ConfigurableWidgetField, limit?: number)}
	{@const result = tableRows(field, data)}
	<RowsWidget {title} rows={result.rows.slice(0, limit)} fields={result.fields} />
{/snippet}

{#snippet distribution(showClanStatus: boolean)}
	<DonutWidget
		{title}
		items={showClanStatus ? statuses : kills}
		centerLabel={showClanStatus
			? m['pages.dashboard.widget_clans']()
			: m['pages.dashboard.kills']()}
		emptyMessage={showClanStatus
			? m['pages.dashboard.no_clans']()
			: m['pages.dashboard.no_kill_data']()}
	/>
{/snippet}

{#if widget.type === 'counter'}
	<CountWidget
		{title}
		value={counterValue(widget.field as CounterField, data)}
		format={widget.format ?? 'number'}
	/>
{:else if widget.type === 'distribution'}
	{@render distribution(widget.field === 'clanStatus')}
{:else if widget.type === 'leaderboard'}
	<TopClansWidget clans={data.clans} />
{:else if widget.type === 'table'}
	{@render table(widget.field, widget.limit ?? 10)}
{:else if widget.type === 'activityFeed'}
	<LatestActionsWidget events={data.events} />
{:else if widget.type === 'onlineSnapshot'}
	<OnlinePlayersWidget players={data.onlinePlayers} {avatarUrlTemplate} />
{:else if widget.type === 'timeSeries'}
	<TimeSeriesWidget {title} metric={widget.field as CounterField} history={data.history} />
{:else if widget.type === 'peakOnline'}
	<PeakOnlineWidget {title} history={data.history} />
{:else if widget.type === 'warCenter'}
	{@render table('wars')}
{:else if widget.type === 'clanRoster'}
	{@render table('clanPlayers')}
{:else if widget.type === 'economy'}
	<EconomyWidget clans={data.clans} />
{:else if widget.type === 'killFeed'}
	<KillFeedWidget events={data.events} />
{:else if widget.type === 'pendingInvites'}
	{@render table('pendingInvites')}
{:else if widget.type === 'claims'}
	{@render table('claims')}
{:else if widget.type === 'clanCard'}
	<ClanCardWidget clans={data.clans} />
{:else if widget.type === 'playerCard'}
	<PlayerCardWidget players={data.players} />
{:else if widget.type === 'network'}
	<NetworkWidget clans={data.clans} />
{:else if widget.type === 'compare'}
	<CompareWidget clans={data.clans} />
{:else if widget.type === 'heatmap'}
	<HeatmapWidget activity={data.playerActivity} />
{:else if widget.type === 'gainersLosers'}
	{@render table('clanBalanceChange')}
{/if}
