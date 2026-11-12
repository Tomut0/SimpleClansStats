import type { DashboardWidget } from '$lib/api/generated';
import { m } from '$lib/paraglide/messages.js';
import type { ConfigurableWidgetField, ConfigurableWidgetType, CounterField } from './types';
import { counterFields, counterFieldFor, isCounter, legacyCounterFields } from './dashboard-layout';

export function widgetTitle(id: string, field?: CounterField) {
	const counter = field ?? legacyCounterFields[id];
	if (counter === 'clans') return m['pages.dashboard.counter_clans']();
	if (counter === 'players') return m['pages.dashboard.counter_players']();
	if (counter === 'wars') return m['pages.dashboard.counter_wars']();
	if (counter === 'balance') return m['pages.dashboard.counter_balance']();
	if (counter === 'onlinePlayers') return m['pages.dashboard.counter_online']();
	if (counter === 'kills') return m['pages.dashboard.counter_kills']();
	if (id === 'clans') return m['pages.dashboard.widget_clans']();
	if (id === 'players') return m['pages.dashboard.widget_players']();
	if (id === 'wars') return m['pages.dashboard.widget_wars']();
	if (id === 'balance') return m['pages.dashboard.widget_balance']();
	if (id === 'kills') return m['pages.dashboard.widget_kills']();
	if (id === 'status') return m['pages.dashboard.widget_status']();
	if (id === 'actions') return m['pages.dashboard.widget_actions']();
	if (id === 'top') return m['pages.dashboard.widget_top']();
	if (id === 'online') return m['pages.dashboard.widget_online']();
	return id;
}

export function configuredTitle(widget: DashboardWidget) {
	if (isCounter(widget)) return widgetTitle(widget.id, counterFieldFor(widget));
	if (widget.type === 'table') return configuredFieldTitle(widget.field as ConfigurableWidgetField);
	if (widget.type === 'distribution')
		return widgetTitle(widget.field === 'clanStatus' ? 'status' : 'kills');
	if (widget.type === 'leaderboard') return widgetTitle('top');
	if (widget.type === 'activityFeed') return widgetTitle('actions');
	if (widget.type === 'onlineSnapshot') return widgetTitle('online');
	if (widget.type === 'timeSeries')
		return widgetTitle(widget.field ?? 'clans', widget.field as CounterField);
	if (widget.type) return configuredTypeTitle(widget.type as ConfigurableWidgetType);
	return widgetTitle(widget.id);
}

export function configuredTypeTitle(type: ConfigurableWidgetType) {
	const titles: Record<ConfigurableWidgetType, () => string> = {
		counter: () => m['pages.dashboard.counter_type'](),
		table: () => m['pages.dashboard.table_type'](),
		distribution: () => m['pages.dashboard.distribution_type'](),
		leaderboard: () => m['pages.dashboard.leaderboard_type'](),
		activityFeed: () => m['pages.dashboard.activity_feed_type'](),
		onlineSnapshot: () => m['pages.dashboard.online_snapshot_type'](),
		timeSeries: () => m['pages.dashboard.time_series_type'](),
		peakOnline: () => m['pages.dashboard.peak_online_type'](),
		warCenter: () => m['pages.dashboard.war_center_type'](),
		clanRoster: () => m['pages.dashboard.clan_roster_type'](),
		economy: () => m['pages.dashboard.economy_type'](),
		killFeed: () => m['pages.dashboard.kill_feed_type'](),
		pendingInvites: () => m['pages.dashboard.pending_invites_type'](),
		claims: () => m['pages.dashboard.claims_type'](),
		clanCard: () => m['pages.dashboard.clan_card_type'](),
		playerCard: () => m['pages.dashboard.player_card_type'](),
		network: () => m['pages.dashboard.network_type'](),
		compare: () => m['pages.dashboard.compare_type'](),
		heatmap: () => m['pages.dashboard.heatmap_type'](),
		gainersLosers: () => m['pages.dashboard.gainers_losers_type']()
	};
	return titles[type]();
}

export function configuredFieldTitle(field: ConfigurableWidgetField) {
	if (counterFields.includes(field as CounterField))
		return widgetTitle(field, field as CounterField);
	const titles: Partial<Record<ConfigurableWidgetField, () => string>> = {
		recentActions: () => widgetTitle('actions'),
		clanBalance: () => widgetTitle('top'),
		clanPlayers: () => m['pages.dashboard.clan_roster_type'](),
		recentKills: () => m['pages.dashboard.kill_feed_type'](),
		pendingInvites: () => m['pages.dashboard.pending_invites_type'](),
		claims: () => m['pages.dashboard.claims_type'](),
		clan: () => m['pages.dashboard.clan_card_type'](),
		player: () => m['pages.dashboard.player_card_type'](),
		clanRelations: () => m['pages.dashboard.network_type'](),
		clanMetrics: () => m['pages.dashboard.compare_type'](),
		playerActivity: () => m['pages.dashboard.heatmap_type'](),
		clanBalanceChange: () => m['pages.dashboard.gainers_losers_type'](),
		killTypes: () => widgetTitle('kills'),
		clanStatus: () => widgetTitle('status')
	};
	return titles[field]?.() ?? widgetTitle('status');
}
