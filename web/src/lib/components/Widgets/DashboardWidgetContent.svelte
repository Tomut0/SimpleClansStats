<script lang="ts">
	import type { DashboardWidget } from '$lib/api/generated';
	import { m } from '$lib/paraglide/messages.js';
	import type { ConfigurableWidgetField, CounterField, WidgetData } from './types';
	import { counterFieldFor, isCounter } from './dashboard-layout';
	import { clanStatusSlices, counterValue, killSlices, tableRows } from './dashboard-data';
	import { configuredTitle, widgetTitle } from './dashboard-labels';
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
	const kills = $derived(killSlices(data));
	const statuses = $derived(clanStatusSlices(data));
</script>

{#if isCounter(widget)}
	<CountWidget
		title={widgetTitle(widget.id, counterFieldFor(widget))}
		value={counterValue(counterFieldFor(widget), data)}
		format={widget.format ?? 'number'}
	/>
{:else if widget.type === 'distribution'}
	<DonutWidget
		title={configuredTitle(widget)}
		items={widget.field === 'clanStatus' ? statuses : kills}
		centerLabel={widget.field === 'clanStatus'
			? m['pages.dashboard.widget_clans']()
			: m['pages.dashboard.kills']()}
		emptyMessage={widget.field === 'clanStatus'
			? m['pages.dashboard.no_clans']()
			: m['pages.dashboard.no_kill_data']()}
	/>
{:else if widget.type === 'leaderboard' || widget.id === 'top'}
	<TopClansWidget clans={data.clans} />
{:else if widget.type === 'table'}
	{@const table = tableRows(widget.field as ConfigurableWidgetField, data)}
	<RowsWidget
		title={configuredTitle(widget)}
		rows={table.rows.slice(0, widget.limit ?? 10)}
		fields={table.fields}
	/>
{:else if widget.type === 'activityFeed' || widget.id === 'actions'}
	<LatestActionsWidget events={data.events} />
{:else if widget.type === 'onlineSnapshot' || widget.id === 'online'}
	<OnlinePlayersWidget players={data.onlinePlayers} {avatarUrlTemplate} />
{:else if widget.type === 'timeSeries'}
	<TimeSeriesWidget
		title={configuredTitle(widget)}
		metric={widget.field as CounterField}
		history={data.history}
	/>
{:else if widget.type === 'peakOnline'}
	<PeakOnlineWidget title={m['pages.dashboard.peak_online_type']()} history={data.history} />
{:else if widget.type === 'warCenter'}
	{@const table = tableRows('wars', data)}
	<RowsWidget title={configuredTitle(widget)} rows={table.rows} fields={table.fields} />
{:else if widget.type === 'clanRoster'}
	{@const table = tableRows('clanPlayers', data)}
	<RowsWidget title={configuredTitle(widget)} rows={table.rows} fields={table.fields} />
{:else if widget.type === 'economy'}
	<EconomyWidget clans={data.clans} />
{:else if widget.type === 'killFeed'}
	<KillFeedWidget events={data.events} />
{:else if widget.type === 'pendingInvites'}
	{@const table = tableRows('pendingInvites', data)}
	<RowsWidget title={configuredTitle(widget)} rows={table.rows} fields={table.fields} />
{:else if widget.type === 'claims'}
	{@const table = tableRows('claims', data)}
	<RowsWidget title={configuredTitle(widget)} rows={table.rows} fields={table.fields} />
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
	{@const table = tableRows('clanBalanceChange', data)}
	<RowsWidget title={configuredTitle(widget)} rows={table.rows} fields={table.fields} />
{:else if widget.id === 'kills'}
	<DonutWidget
		title={widgetTitle(widget.id)}
		items={kills}
		centerLabel={m['pages.dashboard.kills']()}
		emptyMessage={m['pages.dashboard.no_kill_data']()}
	/>
{:else if widget.id === 'status'}
	<DonutWidget
		title={widgetTitle(widget.id)}
		items={statuses}
		centerLabel={m['pages.dashboard.widget_clans']()}
		emptyMessage={m['pages.dashboard.no_clans']()}
	/>
{/if}
