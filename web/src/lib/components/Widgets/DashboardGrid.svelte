<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { GridStack, type GridStackWidget } from 'gridstack';
	import Plus from '@lucide/svelte/icons/plus';
	import Trash2 from '@lucide/svelte/icons/trash-2';
	import type { DashboardData, DashboardLayout, DashboardWidget } from '$lib/api/generated';
	import 'gridstack/dist/gridstack.min.css';
	import CountWidget from './CountWidget.svelte';
	import DonutWidget from './DonutWidget.svelte';
	import LatestActionsWidget from './LatestActionsWidget.svelte';
	import TopClansWidget from './TopClansWidget.svelte';
	import OnlinePlayersWidget from './OnlinePlayersWidget.svelte';
	import TimeSeriesWidget from './TimeSeriesWidget.svelte';
	import PeakOnlineWidget from './PeakOnlineWidget.svelte';
	import WarCenterWidget from './WarCenterWidget.svelte';
	import ClanRosterWidget from './ClanRosterWidget.svelte';
	import EconomyWidget from './EconomyWidget.svelte';
	import KillFeedWidget from './KillFeedWidget.svelte';
	import PendingInvitesWidget from './PendingInvitesWidget.svelte';
	import ClaimsWidget from './ClaimsWidget.svelte';
	import ClanCardWidget from './ClanCardWidget.svelte';
	import PlayerCardWidget from './PlayerCardWidget.svelte';
	import NetworkWidget from './NetworkWidget.svelte';
	import CompareWidget from './CompareWidget.svelte';
	import HeatmapWidget from './HeatmapWidget.svelte';
	import GainersLosersWidget from './GainersLosersWidget.svelte';
	import RowsWidget from './RowsWidget.svelte';
	import type {
		ConfigurableWidgetField,
		ConfigurableWidgetType,
		CounterField,
		CounterFormat,
		DonutSlice,
		WidgetData
	} from './types';
	import { m } from '$lib/paraglide/messages.js';
	import * as Dialog from '$lib/components/ui/dialog';

	type WidgetLayout = GridStackWidget & DashboardWidget;
	let { serverId, apiUrl, avatarUrlTemplate, editing = $bindable(false), onAddWidget = $bindable(() => {}), onRestoreDefaults = $bindable(() => {}), onSaveDashboard = $bindable(() => {}) }: { serverId: string; apiUrl: string; avatarUrlTemplate: string; editing?: boolean; onAddWidget?: () => void; onRestoreDefaults?: () => void; onSaveDashboard?: () => void } = $props();
	const defaults: WidgetLayout[] = [
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
	const catalog = [
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
	const legacyCounterFields: Record<string, CounterField> = {
		clans: 'clans',
		players: 'players',
		wars: 'wars',
		balance: 'balance'
	};
	const counterFields: CounterField[] = [
		'clans',
		'players',
		'wars',
		'balance',
		'onlinePlayers',
		'kills'
	];
	const counterFormats: CounterFormat[] = ['number', 'compact', 'plain'];
	const configurableTypes: ConfigurableWidgetType[] = [
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
	const configurableFields: Record<ConfigurableWidgetType, ConfigurableWidgetField[]> = {
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
	let widgets = $state<WidgetLayout[]>(defaults);
	let gridElement: HTMLDivElement;
	let grid: GridStack | undefined;
	let data = $state<DashboardData | null>(null);
	let saveError = $state(false);
	let addWidgetOpen = $state(false);
	let configuringType = $state<ConfigurableWidgetType | null>(null);
	let selectedField = $state<ConfigurableWidgetField>('clans');
	let counterFormat = $state<CounterFormat>('number');
	let tableLimit = $state(10);
	let hoverTile = $state<{ x: number; y: number; width: number; height: number } | null>(null);
	let mounted = new Set<string>();
	let storageKey = $derived('simpleclansstats.dashboard.v1');
	let available = $derived(widgets.length >= 24 ? [] : configurableTypes);
	const snapshotState = $derived(data?.state ?? {});
	const collector = $derived(
		(snapshotState.extensions as Record<string, Record<string, unknown>> | undefined)
			?.scCollectorV1 ?? {}
	);
	const widgetData = $derived<WidgetData>({
		clans: (snapshotState.clans as Record<string, unknown>[] | undefined) ?? [],
		players: (snapshotState.players as Record<string, unknown>[] | undefined) ?? [],
		wars: (collector.wars as Record<string, unknown>[] | undefined) ?? [],
		onlinePlayers: (collector.onlinePlayers as Record<string, unknown>[] | undefined) ?? [],
		events: data?.events ?? [],
		history: data?.history ?? [],
		pendingInvites: (collector.pendingInvites as Record<string, unknown>[] | undefined) ?? [],
		claims: (collector.claims as Record<string, unknown>[] | undefined) ?? [],
		clanBalanceChanges: data?.clanBalanceChanges ?? [],
		playerActivity: (collector.playerActivity as unknown[] | undefined) ?? [],
		killTypes: parseKillTypes(
			collector.killsByType ?? collector.killStats ?? snapshotState.killsByType
		)
	});
	const killSlices = $derived<DonutSlice[]>(
		widgetData.killTypes.map((item) => ({
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
		}))
	);
	const clanStatuses = [
		{ id: 'verified', label: m['pages.dashboard.verified'](), color: '#54b4ef' },
		{ id: 'unverified', label: m['pages.dashboard.unverified'](), color: '#00c7b1' },
		{ id: 'permanent', label: m['pages.dashboard.permanent'](), color: '#ffcc00' }
	];
	const clanStatusSlices = $derived<DonutSlice[]>(
		clanStatuses.map((status) => ({
			label: status.label,
			color: status.color,
			count: widgetData.clans.filter((clan) => classifyClanStatus(clan) === status.id).length
		}))
	);
	function classifyClanStatus(clan: Record<string, unknown>) {
		const status = String(clan.status ?? '').toLowerCase();
		if (status.includes('permanent') || clan.permanent === true) return 'permanent';
		if (status.includes('unverified') || clan.verified === false) return 'unverified';
		if (status.includes('verified') || clan.verified === true) return 'verified';
		return '';
	}
	function parseKillTypes(value: unknown) {
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
	function isCounter(widget: DashboardWidget) {
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
		return (
			!widget.type && !widget.field && !widget.format && Boolean(legacyCounterFields[widget.id])
		);
	}
	function validConfiguredWidget(widget: DashboardWidget) {
		if (!widget.type || !supportedWidgetTypes.includes(widget.type as ConfigurableWidgetType))
			return false;
		const type = widget.type as ConfigurableWidgetType;
		const idPattern = type === 'counter' ? /^counter-[0-9a-f-]{36}$/i : /^widget-[0-9a-f-]{36}$/i;
		const validId =
			Boolean(legacyCounterFields[widget.id] && type === 'counter') || idPattern.test(widget.id);
		return validId && configurableFields[type].includes(widget.field as ConfigurableWidgetField) &&
			(type !== 'table' || [5, 10, 20].includes(widget.limit ?? 10));
	}
	function counterFieldFor(widget: DashboardWidget): CounterField {
		return (
			(isCounterField(widget.field) && widget.field) || legacyCounterFields[widget.id] || 'clans'
		);
	}
	function counterValue(field: CounterField) {
		if (field === 'clans') return widgetData.clans.length;
		if (field === 'players') return widgetData.players.length;
		if (field === 'wars') return widgetData.wars.length;
		if (field === 'onlinePlayers') return widgetData.onlinePlayers.length;
		if (field === 'kills') return widgetData.killTypes.reduce((sum, kill) => sum + kill.count, 0);
		return widgetData.clans.reduce((sum, clan) => sum + Number(clan.balance ?? 0), 0);
	}
	function widgetTitle(id: string, field?: CounterField) {
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
	function configuredTitle(widget: DashboardWidget) {
		if (isCounter(widget)) return widgetTitle(widget.id, counterFieldFor(widget));
		if (widget.type === 'table') return configuredFieldTitle(widget.field as ConfigurableWidgetField);
		if (widget.type === 'distribution')
			return widgetTitle(widget.field === 'clanStatus' ? 'status' : 'kills');
		if (widget.type === 'leaderboard') return widgetTitle('top');
		if (widget.type === 'activityFeed') return widgetTitle('actions');
		if (widget.type === 'onlineSnapshot') return widgetTitle('online');
		if (widget.type === 'timeSeries') return widgetTitle(widget.field ?? 'clans', widget.field as CounterField);
		if (widget.type === 'peakOnline') return m['pages.dashboard.peak_online_type']();
		if (widget.type === 'warCenter') return m['pages.dashboard.war_center_type']();
		if (widget.type === 'clanRoster') return m['pages.dashboard.clan_roster_type']();
		if (widget.type === 'economy') return m['pages.dashboard.economy_type']();
		if (widget.type === 'killFeed') return m['pages.dashboard.kill_feed_type']();
		if (widget.type === 'pendingInvites') return m['pages.dashboard.pending_invites_type']();
		if (widget.type === 'claims') return m['pages.dashboard.claims_type']();
		if (widget.type === 'clanCard') return m['pages.dashboard.clan_card_type']();
		if (widget.type === 'playerCard') return m['pages.dashboard.player_card_type']();
		if (widget.type === 'network') return m['pages.dashboard.network_type']();
		if (widget.type === 'compare') return m['pages.dashboard.compare_type']();
		if (widget.type === 'heatmap') return m['pages.dashboard.heatmap_type']();
		if (widget.type === 'gainersLosers') return m['pages.dashboard.gainers_losers_type']();
		return widgetTitle(widget.id);
	}
	function configuredTypeTitle(type: ConfigurableWidgetType) {
		if (type === 'counter') return m['pages.dashboard.counter_type']();
		if (type === 'table') return m['pages.dashboard.table_type']();
		if (type === 'distribution') return m['pages.dashboard.distribution_type']();
		if (type === 'leaderboard') return m['pages.dashboard.leaderboard_type']();
		if (type === 'activityFeed') return m['pages.dashboard.activity_feed_type']();
		if (type === 'onlineSnapshot') return m['pages.dashboard.online_snapshot_type']();
		if (type === 'timeSeries') return m['pages.dashboard.time_series_type']();
		if (type === 'peakOnline') return m['pages.dashboard.peak_online_type']();
		if (type === 'warCenter') return m['pages.dashboard.war_center_type']();
		if (type === 'clanRoster') return m['pages.dashboard.clan_roster_type']();
		if (type === 'economy') return m['pages.dashboard.economy_type']();
		if (type === 'killFeed') return m['pages.dashboard.kill_feed_type']();
		if (type === 'pendingInvites') return m['pages.dashboard.pending_invites_type']();
		if (type === 'claims') return m['pages.dashboard.claims_type']();
		if (type === 'clanCard') return m['pages.dashboard.clan_card_type']();
		if (type === 'playerCard') return m['pages.dashboard.player_card_type']();
		if (type === 'network') return m['pages.dashboard.network_type']();
		if (type === 'compare') return m['pages.dashboard.compare_type']();
		if (type === 'heatmap') return m['pages.dashboard.heatmap_type']();
		return m['pages.dashboard.gainers_losers_type']();
	}
	function configuredFieldTitle(field: ConfigurableWidgetField) {
		if (counterFields.includes(field as CounterField))
			return widgetTitle(field, field as CounterField);
		if (field === 'recentActions') return widgetTitle('actions');
		if (field === 'onlinePlayers') return widgetTitle('onlinePlayers', 'onlinePlayers');
		if (field === 'clanBalance') return widgetTitle('top');
		if (field === 'wars') return m['pages.dashboard.war_center_type']();
		if (field === 'clanPlayers') return m['pages.dashboard.clan_roster_type']();
		if (field === 'recentKills') return m['pages.dashboard.kill_feed_type']();
		if (field === 'pendingInvites') return m['pages.dashboard.pending_invites_type']();
		if (field === 'claims') return m['pages.dashboard.claims_type']();
		if (field === 'clan') return m['pages.dashboard.clan_card_type']();
		if (field === 'player') return m['pages.dashboard.player_card_type']();
		if (field === 'clanRelations') return m['pages.dashboard.network_type']();
		if (field === 'clanMetrics') return m['pages.dashboard.compare_type']();
		if (field === 'playerActivity') return m['pages.dashboard.heatmap_type']();
		if (field === 'clanBalanceChange') return m['pages.dashboard.gainers_losers_type']();
		if (field === 'killTypes') return widgetTitle('kills');
		return widgetTitle('status');
	}
	function tableRows(field: ConfigurableWidgetField) {
		if (field === 'wars') return { rows: widgetData.wars, fields: ['name', 'status', 'startedAt'] };
		if (field === 'clanPlayers') return { rows: widgetData.players, fields: ['name', 'clanTag', 'rank'] };
		if (field === 'pendingInvites') return { rows: widgetData.pendingInvites, fields: ['playerName', 'clanTag', 'createdAt'] };
		return { rows: widgetData.clanBalanceChanges, fields: ['tag', 'balanceChange'] };
	}

	onMount(() => {
		let source: EventSource | undefined;
		let disposed = false;
		const initialize = async () => {
			try {
				const response = await fetch(`${apiUrl}/dashboard/layout`, { credentials: 'include' });
				if (response.ok) {
					const saved = (await response.json()) as { version: number; widgets: WidgetLayout[] };
					if (saved.version === 1 && saved.widgets.length && validWidgets(saved.widgets))
						widgets = saved.widgets;
				} else throw new Error('layout unavailable');
			} catch {
				try {
					const saved = localStorage.getItem(storageKey);
					if (saved) {
						const parsed = JSON.parse(saved) as WidgetLayout[];
						if (validWidgets(parsed)) widgets = parsed;
					}
				} catch {
					localStorage.removeItem(storageKey);
				}
			}
			await tick();
			if (disposed) return;
			connectGrid();
			source = new EventSource(
				`${apiUrl}/dashboard/events?serverId=${encodeURIComponent(serverId)}`,
				{ withCredentials: true }
			);
			source.addEventListener('snapshot', (event) => {
				data = JSON.parse((event as MessageEvent<string>).data) as DashboardData;
			});
		};
		void initialize();
		return () => {
			disposed = true;
			source?.close();
			grid?.destroy(false);
		};
	});

	$effect(() => {
		const count = widgets.length;
		const editingNow = editing;
		if (!grid || (!count && !editingNow)) return;
		void tick().then(() => {
			for (const node of gridElement.querySelectorAll<HTMLElement>('.grid-stack-item')) {
				const id = node.dataset.widgetId;
				if (id && !mounted.has(id)) {
					const widget = widgets.find((item) => item.id === id);
					if (widget) grid?.makeWidget(node, gridOptions(widget));
					mounted.add(id);
				}
			}
		});
	});

	function addConfiguredWidget() {
		if (widgets.length >= 24) return;
		if (!configuringType) return;
		const isCount = configuringType === 'counter';
		widgets = [
			...widgets,
			{
				id: `${isCount ? 'counter' : 'widget'}-${crypto.randomUUID()}`,
				type: configuringType,
				field: selectedField,
				...(isCount ? { format: counterFormat } : {}),
				...(configuringType === 'table' ? { limit: tableLimit } : {}),
				w: isCount ? 2 : 3,
				h: 2,
				minW: isCount ? 2 : 3,
				minH: 2
			} as WidgetLayout
		];
		configuringType = null;
		addWidgetOpen = false;
	}
	function configureWidgetType(type: ConfigurableWidgetType) {
		configuringType = type;
		selectedField = configurableFields[type][0];
		counterFormat = 'number';
		tableLimit = 10;
	}
	function openWidgetCatalog() {
		configuringType = null;
		addWidgetOpen = true;
	}
	function setEditing(value: boolean) {
		editing = value;
		grid?.setStatic(!value);
	}
	onMount(() => {
		onAddWidget = openWidgetCatalog;
		onRestoreDefaults = restoreDefaults;
		onSaveDashboard = () => void saveDashboard();
	});
	function moveAddTile(event: PointerEvent) {
		if (!editing || event.pointerType === 'touch') return;
		if (event.target instanceof Element && event.target.closest('.add-widget-hover')) return;
		if (
			event.buttons ||
			(event.target instanceof Element && event.target.closest('.remove-widget, .widget-header'))
		) {
			hoverTile = null;
			return;
		}
		if (!grid) return;
		const columns = grid.getColumn();
		const rows = grid.getRow();
		const width = Math.min(3, columns);
		const height = 2;
		if (rows < height) {
			hoverTile = null;
			return;
		}
		const cell = grid.getCellFromPixel({ left: event.pageX, top: event.pageY }, true);
		const cellWidth = gridElement.clientWidth / columns;
		const cellHeight = grid.getCellHeight(true);
		const x = Math.max(0, Math.min(cell.x, columns - width));
		const y = Math.max(0, Math.min(cell.y, rows - height));
		if (!grid.isAreaEmpty(x, y, width, height)) {
			hoverTile = null;
			return;
		}
		hoverTile = {
			x: x * cellWidth,
			y: y * cellHeight,
			width: width * cellWidth,
			height: height * cellHeight
		};
	}
	function removeWidget(id: string) {
		const node = gridElement.querySelector<HTMLElement>(`[data-widget-id="${id}"]`);
		if (node && grid) grid.removeWidget(node, false);
		mounted.delete(id);
		widgets = widgets.filter((widget) => widget.id !== id);
		void tick().then(() =>
			localStorage.setItem(storageKey, JSON.stringify(serializeLayout().widgets))
		);
	}
	function restoreDefaults() {
		grid?.destroy(false);
		grid = undefined;
		mounted.clear();
		widgets = defaults.map((widget) => ({ ...widget }));
		localStorage.removeItem(storageKey);
		void tick().then(() => {
			connectGrid();
			grid?.setStatic(!editing);
		});
	}
	async function saveDashboard() {
		const layout = serializeLayout();
		try {
			const response = await fetch(`${apiUrl}/dashboard/layout`, {
				method: 'PUT',
				credentials: 'include',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(layout)
			});
			if (!response.ok) throw new Error('save failed');
			localStorage.setItem(storageKey, JSON.stringify(layout.widgets));
			saveError = false;
			setEditing(false);
		} catch {
			saveError = true;
		}
	}
	function validWidgets(value: WidgetLayout[]): value is WidgetLayout[] {
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
							: catalog.includes(item.id) && !item.field && !item.format) &&
					Number.isInteger(item.x) &&
					Number.isInteger(item.y) &&
					Number.isInteger(item.w) &&
					Number.isInteger(item.h) &&
					item.x >= 0 &&
					item.x <= 11 &&
					item.y >= 0 &&
					item.y <= 10000 &&
					item.w >= 1 &&
					item.w <= 12 &&
					item.h >= 1 &&
					item.h <= 12 &&
					item.x + item.w <= 12
			) &&
			new Set(value.map((item) => item.id)).size === value.length
		);
	}
	function gridOptions(widget: WidgetLayout): GridStackWidget {
		const minW = isCounter(widget) ? 2 : 3;
		return { ...widget, minW, minH: 2 };
	}
	function connectGrid() {
		for (const widget of widgets) {
			const node = gridElement.querySelector<HTMLElement>(`[data-widget-id="${widget.id}"]`);
			if (node)
				for (const [key, value] of Object.entries({
					'gs-id': widget.id,
					'gs-x': widget.x,
					'gs-y': widget.y,
					'gs-w': widget.w,
					'gs-h': widget.h,
					'gs-min-w': gridOptions(widget).minW,
					'gs-min-h': gridOptions(widget).minH
				}))
					node.setAttribute(key, String(value));
		}
		grid = GridStack.init(
			{
				column: 12,
				cellHeight: 76,
				margin: 8,
				staticGrid: !editing,
				draggable: { handle: '.widget-header' }
			},
			gridElement
		);
		for (const node of gridElement.querySelectorAll<HTMLElement>('.grid-stack-item'))
			mounted.add(node.dataset.widgetId!);
		grid.on('change', () => {
			if (editing) localStorage.setItem(storageKey, JSON.stringify(serializeLayout().widgets));
		});
	}
	function serializeLayout(): DashboardLayout {
		const saved = grid?.save(false, false);
		const items = Array.isArray(saved) ? (saved as GridStackWidget[]) : [];
		return {
			version: 1,
			widgets: items
				.filter((item) => widgets.some((widget) => widget.id === String(item.id)))
				.map((item) => {
					const widget = widgets.find((candidate) => candidate.id === String(item.id))!;
					return {
						id: widget.id,
						type: widget.type,
						field: widget.field,
						format: widget.format,
						x: item.x ?? 0,
						y: item.y ?? 0,
						w: item.w ?? 1,
						h: item.h ?? 1
					};
				})
		};
	}
</script>

<div class="mb-4">
	{#if saveError}<div class="text-sm text-muted-foreground" role="status">
			{m['pages.dashboard.save_failed']()}
		</div>{/if}
</div>

<div
	class="grid-stack"
	class:grid-stack-editing={editing}
	bind:this={gridElement}
	role="region"
	aria-label={m['pages.dashboard.title']()}
	onpointermove={moveAddTile}
	onpointerleave={() => (hoverTile = null)}
>
	{#each widgets as widget (widget.id)}
		<div class="grid-stack-item" data-widget-id={widget.id}>
			<div class="grid-stack-item-content relative">
				{#if isCounter(widget)}<CountWidget
						title={widgetTitle(widget.id, counterFieldFor(widget))}
						value={counterValue(counterFieldFor(widget))}
						format={widget.format ?? 'number'}
					/>
				{:else if widget.type === 'distribution'}<DonutWidget
						title={configuredTitle(widget)}
						items={widget.field === 'clanStatus' ? clanStatusSlices : killSlices}
						centerLabel={widget.field === 'clanStatus'
							? m['pages.dashboard.widget_clans']()
							: m['pages.dashboard.kills']()}
						emptyMessage={widget.field === 'clanStatus'
							? m['pages.dashboard.no_clans']()
							: m['pages.dashboard.no_kill_data']()}
					/>
				{:else if widget.type === 'leaderboard'}<TopClansWidget clans={widgetData.clans} />
				{:else if widget.type === 'table'}{@const table = tableRows(widget.field as ConfigurableWidgetField)}<RowsWidget title={configuredTitle(widget)} rows={table.rows.slice(0, widget.limit ?? 10)} fields={table.fields} />
				{:else if widget.type === 'activityFeed'}<LatestActionsWidget events={widgetData.events} />
				{:else if widget.type === 'onlineSnapshot'}<OnlinePlayersWidget
						players={widgetData.onlinePlayers}
						{avatarUrlTemplate}
					/>
				{:else if widget.type === 'timeSeries'}<TimeSeriesWidget
					title={configuredTitle(widget)} metric={widget.field as CounterField} history={widgetData.history}
				/>
				{:else if widget.type === 'peakOnline'}<PeakOnlineWidget
					title={m['pages.dashboard.peak_online_type']()} history={widgetData.history}
				/>
				{:else if widget.type === 'warCenter'}<WarCenterWidget wars={widgetData.wars} title={configuredTitle(widget)} />
				{:else if widget.type === 'clanRoster'}<ClanRosterWidget players={widgetData.players} title={configuredTitle(widget)} />
				{:else if widget.type === 'economy'}<EconomyWidget clans={widgetData.clans} />
				{:else if widget.type === 'killFeed'}<KillFeedWidget events={widgetData.events} />
				{:else if widget.type === 'pendingInvites'}<PendingInvitesWidget invites={widgetData.pendingInvites} title={configuredTitle(widget)} />
				{:else if widget.type === 'claims'}<ClaimsWidget claims={widgetData.claims} title={configuredTitle(widget)} />
				{:else if widget.type === 'clanCard'}<ClanCardWidget clans={widgetData.clans} />
				{:else if widget.type === 'playerCard'}<PlayerCardWidget players={widgetData.players} />
				{:else if widget.type === 'network'}<NetworkWidget clans={widgetData.clans} />
				{:else if widget.type === 'compare'}<CompareWidget clans={widgetData.clans} />
				{:else if widget.type === 'heatmap'}<HeatmapWidget activity={widgetData.playerActivity} />
				{:else if widget.type === 'gainersLosers'}<GainersLosersWidget changes={widgetData.clanBalanceChanges} />
				{:else if widget.id === 'kills'}<DonutWidget
						title={widgetTitle(widget.id)}
						items={killSlices}
						centerLabel={m['pages.dashboard.kills']()}
						emptyMessage={m['pages.dashboard.no_kill_data']()}
					/>
				{:else if widget.id === 'status'}<DonutWidget
						title={widgetTitle(widget.id)}
						items={clanStatusSlices}
						centerLabel={m['pages.dashboard.widget_clans']()}
						emptyMessage={m['pages.dashboard.no_clans']()}
					/>
				{:else if widget.id === 'actions'}<LatestActionsWidget events={widgetData.events} />
				{:else if widget.id === 'top'}<TopClansWidget clans={widgetData.clans} />
				{:else if widget.id === 'online'}<OnlinePlayersWidget
						players={widgetData.onlinePlayers}
						{avatarUrlTemplate}
					/>{/if}
				{#if editing && widgets.length > 1}<button
						class="remove-widget text-destructive-foreground absolute top-2 right-2 z-30 grid size-7 place-items-center rounded bg-destructive hover:bg-destructive/90"
						aria-label={m['pages.dashboard.remove_widget']({
							widget: configuredTitle(widget)
						})}
						onclick={() => removeWidget(widget.id)}
						><Trash2 class="size-4" aria-hidden="true" /></button
					>{/if}
			</div>
		</div>
	{/each}
	{#if editing && hoverTile}
		<button
			class="add-widget-hover flex flex-col items-center justify-center gap-2 rounded-lg border border-dashed bg-background/95 text-muted-foreground hover:border-primary hover:text-foreground"
			style:left="{hoverTile.x}px"
			style:top="{hoverTile.y}px"
			style:width="{hoverTile.width}px"
			style:height="{hoverTile.height}px"
			aria-label={m['pages.dashboard.add_widget_button']()}
			onclick={openWidgetCatalog}
		>
			<Plus class="size-7" aria-hidden="true" />
			<span>{m['pages.dashboard.add_widget_button']()}</span>
		</button>
	{/if}
</div>

<Dialog.Root bind:open={addWidgetOpen}>
	<Dialog.Content class="max-w-lg">
		<Dialog.Header>
			<Dialog.Title
				>{configuringType
					? configuredTypeTitle(configuringType)
					: m['pages.dashboard.add_widget_title']()}</Dialog.Title
			>
			<Dialog.Description>{m['pages.dashboard.add_widget_description']()}</Dialog.Description>
		</Dialog.Header>
		{#if configuringType}
			<div class="grid gap-4">
				<label class="grid gap-2 text-sm font-medium">
					{m['pages.dashboard.counter_field_label']()}
					<select class="h-10 rounded-md border bg-background px-3" bind:value={selectedField}>
						{#each configurableFields[configuringType] as field}
							<option value={field}>{configuredFieldTitle(field)}</option>
						{/each}
					</select>
				</label>
				{#if configuringType === 'counter'}
					<label class="grid gap-2 text-sm font-medium">
						{m['pages.dashboard.counter_format_label']()}
						<select class="h-10 rounded-md border bg-background px-3" bind:value={counterFormat}>
							<option value="number">{m['pages.dashboard.counter_format_number']()}</option>
							<option value="compact">{m['pages.dashboard.counter_format_compact']()}</option>
							<option value="plain">{m['pages.dashboard.counter_format_plain']()}</option>
						</select>
					</label>
				{/if}
				{#if configuringType === 'table'}
					<label class="grid gap-2 text-sm font-medium">
						{m['pages.dashboard.table_limit_label']()}
						<select class="h-10 rounded-md border bg-background px-3" bind:value={tableLimit}>
							<option value={5}>5</option><option value={10}>10</option><option value={20}>20</option>
						</select>
					</label>
				{/if}
				<div class="flex justify-between gap-2">
					<button
						class="rounded-md border px-3 py-2 text-sm"
						onclick={() => (configuringType = null)}
					>
						{m['pages.dashboard.counter_back']()}
					</button>
					<button
						class="rounded-md bg-primary px-3 py-2 text-sm text-primary-foreground"
						onclick={addConfiguredWidget}
					>
						{m['pages.dashboard.add_configured_widget']()}
					</button>
				</div>
			</div>
		{:else if available.length}
			<div class="grid gap-2 sm:grid-cols-2">
				{#each available as type}
					<button
						class="flex items-center gap-3 rounded-md border p-3 text-left hover:bg-accent"
						onclick={() => configureWidgetType(type)}
					>
						<Plus class="size-4 shrink-0" aria-hidden="true" />
						<span>{configuredTypeTitle(type)}</span>
					</button>
				{/each}
			</div>
		{:else}
			<p class="text-sm text-muted-foreground">{m['pages.dashboard.widget_limit_reached']()}</p>
		{/if}
	</Dialog.Content>
</Dialog.Root>

<style>
	.grid-stack-editing :global(.widget-header) {
		cursor: move;
	}
	.add-widget-hover {
		position: absolute;
		z-index: 20;
	}
</style>
