import { describe, expect, it } from 'vitest';
import { defaults, validWidgets } from './dashboard-layout';
import { parseKillTypes } from './dashboard-data';

describe('dashboard model', () => {
	it('accepts the default layout and rejects duplicate or out of bounds widgets', () => {
		expect(validWidgets(defaults)).toBe(true);
		expect(validWidgets([{ ...defaults[0], x: 11 }])).toBe(false);
		expect(validWidgets([defaults[0], { ...defaults[0] }])).toBe(false);
	});

	it('normalizes kill counts from array and map snapshots', () => {
		expect(
			parseKillTypes([
				{ type: 'allies', count: '3' },
				{ name: 'rivals', kills: 0 }
			])
		).toEqual([{ type: 'allies', count: 3 }]);
		expect(parseKillTypes({ allies: '4', rivals: -1, civilian: 'bad' })).toEqual([
			{ type: 'allies', count: 4 }
		]);
	});
});
