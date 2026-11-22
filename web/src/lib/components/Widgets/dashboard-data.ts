import type { DashboardData } from '$lib/api/generated';
import { m } from '$lib/paraglide/messages.js';
import type {
	Clan,
	ConfigurableWidgetField,
	CounterField,
	DonutSlice,
	KillTypeCount,
	WidgetData
} from './types';

export function makeWidgetData(data: DashboardData | null): WidgetData {
	const state = data?.state;
	const collector = readCollectorExtension(state?.extensions);
	const { playerActivity, killsByType, killStats } = collector;

	return {
		clans: records(state?.clans),
		players: records(state?.players),
		wars: records(collector.wars),
		onlinePlayers: records(collector.onlinePlayers),
		pendingInvites: records(collector.pendingInvites),
		claims: records(collector.claims),
		playerActivity: Array.isArray(playerActivity) ? playerActivity : [],
		killTypes: parseKillTypes(killsByType ?? killStats ?? state?.killsByType),
		events: data?.events ?? [],
		history: data?.history ?? [],
		clanBalanceChanges: data?.clanBalanceChanges ?? []
	};
}

function readCollectorExtension(extensions: unknown): Record<string, unknown> {
	const collector = isRecord(extensions) ? extensions.scCollectorV1 : undefined;
	return isRecord(collector) ? collector : {};
}

function isRecord(value: unknown): value is Record<string, unknown> {
	return value !== null && typeof value === 'object' && !Array.isArray(value);
}

function records(value: unknown): Record<string, unknown>[] {
	return Array.isArray(value) ? value.filter(isRecord) : [];
}

export function parseKillTypes(value: unknown): KillTypeCount[] {
	let counts: KillTypeCount[];
	if (Array.isArray(value)) {
		counts = value.filter(isRecord).map(({ type, name, count, kills }) => ({
			type: String(type ?? name ?? 'other'),
			count: Number(count ?? kills ?? 0)
		}));
	} else if (isRecord(value)) {
		counts = Object.entries(value).map(([type, count]) => ({ type, count: Number(count) }));
	} else {
		return [];
	}

	return counts.filter(({ count }) => Number.isFinite(count) && count > 0);
}

const allyKill = { label: m['pages.dashboard.allies'], color: '#54b4ef' };
const rivalKill = { label: m['pages.dashboard.rivals'], color: '#00c7b1' };
const killStyles = new Map([
	['ally', allyKill],
	['allies', allyKill],
	['rival', rivalKill],
	['rivals', rivalKill],
	['civilian', { label: m['pages.dashboard.civilian'], color: '#ffcc00' }]
]);

export function killSlices(widgetData: WidgetData): DonutSlice[] {
	return widgetData.killTypes.map(({ type, count }) => {
		const style = killStyles.get(type.toLowerCase());
		return { label: style?.label() ?? type, count, color: style?.color ?? '#64748b' };
	});
}

const clanStatuses = [
	{ id: 'verified', label: m['pages.dashboard.verified'], color: '#54b4ef' },
	{ id: 'unverified', label: m['pages.dashboard.unverified'], color: '#00c7b1' },
	{ id: 'permanent', label: m['pages.dashboard.permanent'], color: '#ffcc00' }
] as const;

export function clanStatusSlices(widgetData: WidgetData): DonutSlice[] {
	const counts = { verified: 0, unverified: 0, permanent: 0 };
	for (const clan of widgetData.clans) {
		const status = classifyClanStatus(clan);
		if (status) counts[status] += 1;
	}
	return clanStatuses.map(({ id, label, color }) => ({
		id,
		label: label(),
		color,
		count: counts[id]
	}));
}

function classifyClanStatus(clan: Clan) {
	const status = String(clan.status ?? '').toLowerCase();
	if (status.includes('permanent') || clan.permanent === true) return 'permanent';
	if (status.includes('unverified') || clan.verified === false) return 'unverified';
	if (status.includes('verified') || clan.verified === true) return 'verified';
}

export function counterValue(field: CounterField, widgetData: WidgetData) {
	switch (field) {
		case 'kills':
			return widgetData.killTypes.reduce((sum, kill) => sum + kill.count, 0);
		case 'balance':
			return widgetData.clans.reduce((sum, clan) => sum + Number(clan.balance ?? 0), 0);
		default:
			return widgetData[field].length;
	}
}

export function tableRows(field: ConfigurableWidgetField, widgetData: WidgetData) {
	switch (field) {
		case 'wars':
			return { rows: widgetData.wars, fields: ['name', 'status', 'startedAt'] };
		case 'clanPlayers':
			return { rows: widgetData.players, fields: ['name', 'clanTag', 'rank'] };
		case 'pendingInvites':
			return { rows: widgetData.pendingInvites, fields: ['playerName', 'clanTag', 'createdAt'] };
		case 'claims':
			return { rows: widgetData.claims, fields: ['name', 'owner', 'world', 'size'] };
		default:
			return { rows: widgetData.clanBalanceChanges, fields: ['tag', 'balanceChange'] };
	}
}
