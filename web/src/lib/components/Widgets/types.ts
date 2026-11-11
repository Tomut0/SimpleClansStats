export type Clan = Record<string, unknown>;
export type Player = Record<string, unknown>;
export type DashboardEvent = Record<string, unknown>;
export type KillTypeCount = { type: string; count: number };
export type DonutSlice = { label: string; count: number; color: string };
export type CounterField = 'clans' | 'players' | 'wars' | 'balance' | 'onlinePlayers' | 'kills';
export type CounterFormat = 'number' | 'compact' | 'plain';
export type ConfigurableWidgetType =
	'counter' | 'table' | 'leaderboard' | 'activityFeed' | 'onlineSnapshot' | 'distribution' | 'timeSeries' | 'peakOnline' | 'warCenter' | 'clanRoster' | 'economy' | 'killFeed' | 'pendingInvites' | 'claims' | 'clanCard' | 'playerCard' | 'network' | 'compare' | 'heatmap' | 'gainersLosers';
export type ConfigurableWidgetField =
	CounterField | 'clanBalance' | 'recentActions' | 'killTypes' | 'clanStatus' | 'wars' | 'clanPlayers' | 'recentKills' | 'pendingInvites' | 'claims' | 'clan' | 'player' | 'clanRelations' | 'clanMetrics' | 'playerActivity' | 'clanBalanceChange';

export type WidgetData = {
	clans: Clan[];
	players: Player[];
	wars: Record<string, unknown>[];
	onlinePlayers: Player[];
	events: DashboardEvent[];
	killTypes: KillTypeCount[];
	history: { at: string; clans: number; players: number; wars: number; balance: number; onlinePlayers: number; kills: number }[];
	pendingInvites: Record<string, unknown>[];
	claims: Record<string, unknown>[];
	clanBalanceChanges: Record<string, unknown>[];
	playerActivity: unknown[];
};
