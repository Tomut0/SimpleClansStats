import type { DashboardWidget } from '#lib/api/generated/index.js';
import { m } from '#lib/paraglide/messages.js';

const labels: Partial<Record<`pages.dashboard.${'widget' | 'counter'}_${string}`, () => string>> =
	m;

const widgetLabelIds: Partial<Record<DashboardWidget['type'], string>> = {
	leaderboard: 'top',
	activityFeed: 'actions',
	onlineSnapshot: 'online'
};

export function dashboardLabel(id: string, kind: 'widget' | 'counter' = 'widget') {
	return labels[`pages.dashboard.${kind}_${id}`]?.() ?? id;
}

export function dashboardWidgetLabel(widget: DashboardWidget) {
	if (widget.type === 'counter' || widget.type === 'timeSeries') {
		return dashboardLabel(widget.field, 'counter');
	}

	if (widget.type === 'distribution' || widget.type === 'table') {
		return dashboardLabel(widget.field);
	}

	return dashboardLabel(widgetLabelIds[widget.type] ?? widget.type);
}
