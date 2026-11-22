import type { DashboardData, DashboardWidget } from '$lib/api/generated';

export type Clan = Record<string, unknown>;
export type Player = Record<string, unknown>;
export type DashboardEvent = DashboardData['events'][number];
export type KillTypeCount = { type: string; count: number };
export type DonutSlice = { label: string; count: number; color: string };
export type CounterField = 'clans' | 'players' | 'wars' | 'balance' | 'onlinePlayers' | 'kills';
export type CounterFormat = NonNullable<DashboardWidget['format']>;
export type ConfigurableWidgetType = DashboardWidget['type'];
export type ConfigurableWidgetField = DashboardWidget['field'];

export type WidgetData = {
	clans: Clan[];
	players: Player[];
	wars: Record<string, unknown>[];
	onlinePlayers: Player[];
	events: DashboardEvent[];
	killTypes: KillTypeCount[];
	history: DashboardData['history'];
	pendingInvites: Record<string, unknown>[];
	claims: Record<string, unknown>[];
	clanBalanceChanges: DashboardData['clanBalanceChanges'];
	playerActivity: unknown[];
};
