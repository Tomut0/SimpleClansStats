import { describe, expect, it } from 'vitest';
import { getLocale, setLocale } from '#lib/paraglide/runtime.js';
import {
	defaultWidgets,
	serializeLayout,
	validWidgets,
	type WidgetLayout
} from './dashboard-layout';
import { clanStatusSlices, counterValue, makeWidgetData, parseKillTypes } from './dashboard-data';
import { dashboardLabel, dashboardWidgetLabel } from './dashboard-labels';

describe('dashboard model', () => {
	it('accepts the default layout and rejects duplicate or out of bounds widgets', () => {
		expect(validWidgets(defaultWidgets)).toBe(true);
		expect(validWidgets([{ ...defaultWidgets[0], x: 11 }])).toBe(false);
		expect(validWidgets([defaultWidgets[0], { ...defaultWidgets[0] }])).toBe(false);
	});

	it('requires the same widget ID format for counters and other widgets', () => {
		const counter = defaultWidgets[0];
		expect(validWidgets([{ ...counter, id: 'clans' }])).toBe(false);
		expect(validWidgets([{ ...counter, id: 'counter-12345678-1234-1234-1234-123456789abc' }])).toBe(
			false
		);
	});

	it('normalizes kill counts from array and map snapshots', () => {
		expect(
			parseKillTypes([{ type: 'allies', count: '3' }, null, { name: 'rivals', kills: 0 }])
		).toEqual([{ type: 'allies', count: 3 }]);
		expect(parseKillTypes({ allies: '4', rivals: -1, civilian: 'bad' })).toEqual([
			{ type: 'allies', count: 4 }
		]);
	});

	it('uses collector data and keeps widget arrays safe to read', () => {
		const data = makeWidgetData({
			state: {
				clans: [{ status: 'unverified', balance: 25 }, { permanent: true }, null],
				extensions: {
					scCollectorV1: {
						onlinePlayers: [{ name: 'Alice' }],
						wars: 'invalid',
						killStats: { ally: 3 }
					}
				}
			},
			events: [],
			history: [],
			clanBalanceChanges: []
		});

		expect(data.wars).toEqual([]);
		expect(counterValue('onlinePlayers', data)).toBe(1);
		expect(counterValue('balance', data)).toBe(25);
		expect(counterValue('kills', data)).toBe(3);
		expect(clanStatusSlices(data).map(({ count }) => count)).toEqual([0, 1, 1]);
	});

	it('saves the table row limit with the updated position', () => {
		const table: WidgetLayout = {
			id: 'widget-12345678-1234-1234-1234-123456789abc',
			type: 'table',
			field: 'wars',
			limit: 20,
			x: 0,
			y: 0,
			w: 3,
			h: 2,
			autoPosition: true
		};
		const layout = serializeLayout(
			[table, defaultWidgets[0]],
			[{ id: 'unknown' }, { id: table.id, x: 4, y: 2, w: 5, h: 3 }]
		);
		expect(JSON.parse(JSON.stringify(layout))).toEqual({
			version: 1,
			widgets: [{ id: table.id, type: 'table', field: 'wars', limit: 20, x: 4, y: 2, w: 5, h: 3 }]
		});
		expect(validWidgets(layout.widgets)).toBe(true);
	});

	it.each([
		{ locale: 'en', kills: 'Kills by type', online: 'Online players', timeSeries: 'Time series' },
		{
			locale: 'ru',
			kills: 'Убийства по типам',
			online: 'Игроки онлайн',
			timeSeries: 'Изменение во времени'
		}
	] as const)(
		'looks up dashboard labels in $locale',
		async ({ locale, kills, online, timeSeries }) => {
			const previousLocale = getLocale();
			const position = { x: 0, y: 0, w: 2, h: 2 };
			await setLocale(locale, { reload: false });
			try {
				expect(dashboardLabel('timeSeries')).toBe(timeSeries);
				expect(
					dashboardWidgetLabel({
						...position,
						id: 'widget-123',
						type: 'distribution',
						field: 'killTypes'
					})
				).toBe(kills);
				expect(
					dashboardWidgetLabel({
						...position,
						id: 'widget-124',
						type: 'counter',
						field: 'onlinePlayers',
						format: 'number'
					})
				).toBe(online);
			} finally {
				await setLocale(previousLocale, { reload: false });
			}
		}
	);
});
