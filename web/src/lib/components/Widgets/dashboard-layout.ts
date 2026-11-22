import type { DashboardLayout, DashboardWidget } from '$lib/api/generated';
import type { GridStackWidget } from 'gridstack';
import type {
	ConfigurableWidgetField,
	ConfigurableWidgetType,
	CounterField,
	CounterFormat
} from './types';

export type WidgetLayout = DashboardWidget & {
	autoPosition?: boolean;
};

export type WidgetConfiguration = Pick<DashboardWidget, 'type' | 'field' | 'format' | 'limit'>;

export const layoutVersion = 1;
export const maxWidgets = 24;
export const tableLimits = [5, 10, 20] as const;

const defaultCounterFields = ['clans', 'players', 'wars', 'balance'] as const;

const defaultWidgetConfigurations: Omit<DashboardWidget, 'id'>[] = [
	...defaultCounterFields.map((field, column) => ({
		type: 'counter' as const,
		field,
		format: 'number' as const,
		x: column * 2,
		y: 0,
		w: 2,
		h: 2
	})),
	{ type: 'distribution', field: 'killTypes', x: 0, y: 2, w: 4, h: 3 },
	{ type: 'distribution', field: 'clanStatus', x: 4, y: 2, w: 4, h: 3 },
	{ type: 'activityFeed', field: 'recentActions', x: 8, y: 0, w: 4, h: 5 },
	{ type: 'leaderboard', field: 'clanBalance', x: 0, y: 5, w: 8, h: 3 },
	{ type: 'onlineSnapshot', field: 'onlinePlayers', x: 8, y: 5, w: 4, h: 3 }
];

export function createWidgetId(uuid: string = crypto.randomUUID()) {
	return `widget-${uuid}`;
}

// Keep default IDs stable so a saved layout and a reset use the same widget identities.
export const defaultWidgets: WidgetLayout[] = defaultWidgetConfigurations.map((widget, index) => {
	const suffix = String(index + 1).padStart(12, '0');
	const id = createWidgetId(`00000000-0000-4000-8000-${suffix}`);
	return { ...widget, id };
});

export const counterFields: CounterField[] = [...defaultCounterFields, 'onlinePlayers', 'kills'];
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
function validWidgetConfiguration({ id, type, field, format, limit }: DashboardWidget) {
	if (typeof id !== 'string' || !/^widget-[0-9a-fA-F-]{36}$/.test(id)) return false;
	if (!Object.hasOwn(configurableFields, type) || !configurableFields[type].includes(field)) {
		return false;
	}

	const hasValidFormat = format !== undefined && counterFormats.includes(format);
	if (type === 'counter') {
		return hasValidFormat && limit === undefined;
	}

	if (format !== undefined && !hasValidFormat) return false;
	return limit === undefined || (type === 'table' && tableLimits.includes(limit));
}

function validWidgetPosition({ x, y, w, h }: DashboardWidget) {
	return (
		[x, y, w, h].every(Number.isInteger) &&
		x >= 0 &&
		y >= 0 &&
		y <= 10000 &&
		w >= 1 &&
		h >= 1 &&
		h <= 12 &&
		x + w <= 12
	);
}

export function validWidgets(widgets: DashboardWidget[]) {
	return (
		Array.isArray(widgets) &&
		widgets.length > 0 &&
		widgets.length <= maxWidgets &&
		widgets.every(
			(widget) => widget && validWidgetConfiguration(widget) && validWidgetPosition(widget)
		) &&
		new Set(widgets.map(({ id }) => id)).size === widgets.length
	);
}

export function serializeLayout(
	widgets: DashboardWidget[],
	positions: GridStackWidget[]
): DashboardLayout {
	const widgetsById = new Map(widgets.map((widget) => [widget.id, widget]));
	const saved: DashboardWidget[] = [];
	for (const { id, x = 0, y = 0, w = 1, h = 1 } of positions) {
		const widget = widgetsById.get(String(id));
		if (!widget) continue;
		const { type, field, format, limit } = widget;
		saved.push({ id: widget.id, type, field, format, limit, x, y, w, h });
	}
	return { version: layoutVersion, widgets: saved };
}
