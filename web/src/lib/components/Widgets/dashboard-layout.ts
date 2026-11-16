import type { DashboardWidget } from '$lib/api/generated';
import type {
	ConfigurableWidgetField,
	ConfigurableWidgetType,
	CounterField,
	CounterFormat
} from './types';

export type WidgetLayout = DashboardWidget & {
	x?: number;
	y?: number;
	w?: number;
	h?: number;
	minW?: number;
	minH?: number;
};

export const defaults: WidgetLayout[] = [
	{
		id: 'clans',
		type: 'counter',
		field: 'clans',
		format: 'number',
		x: 0,
		y: 0,
		w: 2,
		h: 2,
		minW: 2,
		minH: 2
	},
	{
		id: 'players',
		type: 'counter',
		field: 'players',
		format: 'number',
		x: 2,
		y: 0,
		w: 2,
		h: 2,
		minW: 2,
		minH: 2
	},
	{
		id: 'wars',
		type: 'counter',
		field: 'wars',
		format: 'number',
		x: 4,
		y: 0,
		w: 2,
		h: 2,
		minW: 2,
		minH: 2
	},
	{
		id: 'balance',
		type: 'counter',
		field: 'balance',
		format: 'number',
		x: 6,
		y: 0,
		w: 2,
		h: 2,
		minW: 2,
		minH: 2
	},
	{ id: 'kills', x: 0, y: 2, w: 4, h: 3, minW: 3, minH: 2 },
	{ id: 'status', x: 4, y: 2, w: 4, h: 3, minW: 3, minH: 2 },
	{ id: 'actions', x: 8, y: 0, w: 4, h: 5, minW: 3, minH: 2 },
	{ id: 'top', x: 0, y: 5, w: 8, h: 3, minW: 3, minH: 2 },
	{ id: 'online', x: 8, y: 5, w: 4, h: 3, minW: 3, minH: 2 }
];

export const legacyCounterFields: Record<string, CounterField> = {
	clans: 'clans',
	players: 'players',
	wars: 'wars',
	balance: 'balance'
};
export const counterFields: CounterField[] = [
	'clans',
	'players',
	'wars',
	'balance',
	'onlinePlayers',
	'kills'
];
export const counterFormats: CounterFormat[] = ['number', 'compact', 'plain'];
export const configurableTypes: ConfigurableWidgetType[] = [
	'counter',
	'table',
	'distribution',
	'leaderboard',
	'activityFeed',
	'onlineSnapshot',
	'timeSeries',
	'peakOnline',
	'economy',
	'killFeed',
	'claims',
	'clanCard',
	'playerCard',
	'network',
	'compare',
	'heatmap'
];
export const configurableFields: Record<ConfigurableWidgetType, ConfigurableWidgetField[]> = {
	counter: counterFields,
	table: ['wars', 'clanPlayers', 'pendingInvites', 'clanBalanceChange'],
	distribution: ['killTypes', 'clanStatus'],
	leaderboard: ['clanBalance'],
	activityFeed: ['recentActions'],
	onlineSnapshot: ['onlinePlayers'],
	timeSeries: counterFields,
	peakOnline: ['onlinePlayers'],
	warCenter: ['wars'],
	clanRoster: ['clanPlayers'],
	economy: ['clanBalance'],
	killFeed: ['recentKills'],
	pendingInvites: ['pendingInvites'],
	claims: ['claims'],
	clanCard: ['clan'],
	playerCard: ['player'],
	network: ['clanRelations'],
	compare: ['clanMetrics'],
	heatmap: ['playerActivity'],
	gainersLosers: ['clanBalanceChange']
};
const supportedWidgetTypes = Object.keys(configurableFields) as ConfigurableWidgetType[];
const legacyWidgetIds = [
	'clans',
	'players',
	'wars',
	'balance',
	'kills',
	'status',
	'actions',
	'top',
	'online'
];

export function isCounter(widget: DashboardWidget) {
	return widget.type === 'counter' || (!widget.type && Boolean(legacyCounterFields[widget.id]));
}
function isCounterField(value: unknown): value is CounterField {
	return typeof value === 'string' && counterFields.includes(value as CounterField);
}
function validCounter(widget: DashboardWidget) {
	if (widget.type === 'counter') {
		const legacyField = legacyCounterFields[widget.id];
		const validId = legacyField
			? legacyField === widget.field
			: /^counter-[0-9a-f-]{36}$/i.test(widget.id);
		return Boolean(
			widget.format &&
			validId &&
			isCounterField(widget.field) &&
			counterFormats.includes(widget.format)
		);
	}
	return !widget.type && !widget.field && !widget.format && Boolean(legacyCounterFields[widget.id]);
}
function validConfiguredWidget(widget: DashboardWidget) {
	if (!widget.type || !supportedWidgetTypes.includes(widget.type as ConfigurableWidgetType))
		return false;
	const type = widget.type as ConfigurableWidgetType;
	const idPattern = type === 'counter' ? /^counter-[0-9a-f-]{36}$/i : /^widget-[0-9a-f-]{36}$/i;
	const validId =
		Boolean(legacyCounterFields[widget.id] && type === 'counter') || idPattern.test(widget.id);
	return (
		validId &&
		configurableFields[type].includes(widget.field as ConfigurableWidgetField) &&
		(type !== 'table' || [5, 10, 20].includes(widget.limit ?? 10))
	);
}

export function validWidgets(value: WidgetLayout[]): value is WidgetLayout[] {
	return (
		Array.isArray(value) &&
		value.length > 0 &&
		value.length <= 24 &&
		value.every(
			(item) =>
				(isCounter(item)
					? validCounter(item)
					: item.type
						? validConfiguredWidget(item)
						: legacyWidgetIds.includes(item.id) && !item.field && !item.format) &&
				Number.isInteger(item.x) &&
				Number.isInteger(item.y) &&
				Number.isInteger(item.w) &&
				Number.isInteger(item.h) &&
				item.x! >= 0 &&
				item.x! <= 11 &&
				item.y! >= 0 &&
				item.y! <= 10000 &&
				item.w! >= 1 &&
				item.w! <= 12 &&
				item.h! >= 1 &&
				item.h! <= 12 &&
				item.x! + item.w! <= 12
		) &&
		new Set(value.map((item) => item.id)).size === value.length
	);
}

export function counterFieldFor(widget: DashboardWidget): CounterField {
	return (
		(isCounterField(widget.field) && widget.field) || legacyCounterFields[widget.id] || 'clans'
	);
}
