import type { DashboardData } from '$lib/api/generated';
import { m } from '$lib/paraglide/messages.js';
import type { ConfigurableWidgetField, CounterField, DonutSlice, WidgetData } from './types';

export function makeWidgetData(data: DashboardData | null): WidgetData {
	const state = data?.state ?? {};
	const collector =
		(state.extensions as Record<string, Record<string, unknown>> | undefined)?.scCollectorV1 ?? {};
	return {
		clans: (state.clans as Record<string, unknown>[] | undefined) ?? [],
		players: (state.players as Record<string, unknown>[] | undefined) ?? [],
		wars: (collector.wars as Record<string, unknown>[] | undefined) ?? [],
		onlinePlayers: (collector.onlinePlayers as Record<string, unknown>[] | undefined) ?? [],
		events: data?.events ?? [],
		history: data?.history ?? [],
		pendingInvites: (collector.pendingInvites as Record<string, unknown>[] | undefined) ?? [],
		claims: (collector.claims as Record<string, unknown>[] | undefined) ?? [],
		clanBalanceChanges: data?.clanBalanceChanges ?? [],
		playerActivity: (collector.playerActivity as unknown[] | undefined) ?? [],
		killTypes: parseKillTypes(collector.killsByType ?? collector.killStats ?? state.killsByType)
	};
}

export function parseKillTypes(value: unknown) {
	if (Array.isArray(value))
		return value
			.map((item) => ({
				type: String(item.type ?? item.name ?? 'other'),
				count: Number(item.count ?? item.kills ?? 0)
			}))
			.filter((item) => Number.isFinite(item.count) && item.count > 0);
	if (value && typeof value === 'object')
		return Object.entries(value)
			.map(([type, count]) => ({ type, count: Number(count) }))
			.filter((item) => Number.isFinite(item.count) && item.count > 0);
	return [];
}

export function killSlices(widgetData: WidgetData): DonutSlice[] {
	return widgetData.killTypes.map((item) => ({
		label:
			(
				{
					allies: m['pages.dashboard.allies'](),
					ally: m['pages.dashboard.allies'](),
					rivals: m['pages.dashboard.rivals'](),
					rival: m['pages.dashboard.rivals'](),
					civilian: m['pages.dashboard.civilian']()
				} as Record<string, string>
			)[item.type.toLowerCase()] ?? item.type,
		count: item.count,
		color:
			(
				{
					allies: '#54b4ef',
					ally: '#54b4ef',
					rivals: '#00c7b1',
					rival: '#00c7b1',
					civilian: '#ffcc00'
				} as Record<string, string>
			)[item.type.toLowerCase()] ?? '#64748b'
	}));
}

export function clanStatusSlices(widgetData: WidgetData): DonutSlice[] {
	const statuses = [
		{ id: 'verified', label: m['pages.dashboard.verified'](), color: '#54b4ef' },
		{ id: 'unverified', label: m['pages.dashboard.unverified'](), color: '#00c7b1' },
		{ id: 'permanent', label: m['pages.dashboard.permanent'](), color: '#ffcc00' }
	];
	return statuses.map((status) => ({
		...status,
		count: widgetData.clans.filter((clan) => classifyClanStatus(clan) === status.id).length
	}));
}

function classifyClanStatus(clan: Record<string, unknown>) {
	const status = String(clan.status ?? '').toLowerCase();
	if (status.includes('permanent') || clan.permanent === true) return 'permanent';
	if (status.includes('unverified') || clan.verified === false) return 'unverified';
	if (status.includes('verified') || clan.verified === true) return 'verified';
	return '';
}

export function counterValue(field: CounterField, widgetData: WidgetData) {
	if (field === 'clans') return widgetData.clans.length;
	if (field === 'players') return widgetData.players.length;
	if (field === 'wars') return widgetData.wars.length;
	if (field === 'onlinePlayers') return widgetData.onlinePlayers.length;
	if (field === 'kills') return widgetData.killTypes.reduce((sum, kill) => sum + kill.count, 0);
	return widgetData.clans.reduce((sum, clan) => sum + Number(clan.balance ?? 0), 0);
}

export function tableRows(field: ConfigurableWidgetField, widgetData: WidgetData) {
	if (field === 'wars') return { rows: widgetData.wars, fields: ['name', 'status', 'startedAt'] };
	if (field === 'clanPlayers')
		return { rows: widgetData.players, fields: ['name', 'clanTag', 'rank'] };
	if (field === 'pendingInvites')
		return { rows: widgetData.pendingInvites, fields: ['playerName', 'clanTag', 'createdAt'] };
	if (field === 'claims')
		return { rows: widgetData.claims, fields: ['name', 'owner', 'world', 'size'] };
	return { rows: widgetData.clanBalanceChanges, fields: ['tag', 'balanceChange'] };
}
