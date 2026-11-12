import { getLocale } from '$lib/paraglide/runtime.js';
import { m } from '$lib/paraglide/messages.js';
import type { AppSettings } from '$lib/api/generated';

export type SearchResult = {
	type: 'clan' | 'player';
	name: string;
	data: Record<string, unknown>;
};

export type ResultDetail = { key: string; label: string; raw: unknown; value: string };

const hiddenFields = new Set([
	'id',
	'uuid',
	'externalId',
	'playerId',
	'clanExternalId',
	'leaderExternalId',
	'memberExternalIds',
	'allyExternalIds',
	'rivalExternalIds',
	'pastClanExternalIds',
	'verified',
	'permanent'
]);
const messages = m as unknown as Record<string, () => string>;
const dateFields = new Set([
	'foundedAt',
	'createdAt',
	'firstSeen',
	'lastSeen',
	'lastSeenAt',
	'joinedCurrentClanAt',
	'joinedClanAt',
	'joinedAt',
	'leftAt'
]);

function fieldLabel(key: string) {
	if (key === 'details' || key === 'killsByType' || key === 'killStats') return m['searchBar.fields.killsByType']();
	if (key === 'colorHex' || key === 'colorTag') return m['searchBar.fields.colorTag']();
	return (
		messages[`searchBar.fields.${key}`]?.() ??
		messages[`pages.dashboard.${key}`]?.() ??
		messages[`searchBar.${key}`]?.() ??
		key
			.replace(/([a-z\d])([A-Z])/g, '$1 $2')
			.replace(/[_-]+/g, ' ')
			.replace(/^./, (c) => c.toUpperCase())
	);
}

const fieldOrder: Record<SearchResult['type'], string[][]> = {
	player: [
		['name'], ['role', 'currentRole'], ['kdr', 'kdRatio', 'kd'], ['kills'], ['deaths'],
		['details', 'killsByType', 'killStats'], ['joinedClanAt', 'joinedCurrentClanAt'], ['lastSeenAt', 'lastSeen']
	],
	clan: [
		['name'], ['description'], ['colorHex', 'colorTag'], ['tag'], ['members'], ['balance'],
		['kdr', 'kdRatio', 'kd'], ['kills'], ['deaths'], ['foundedAt', 'createdAt']
	]
};

export function orderedResultDetails(
	type: SearchResult['type'],
	data: Record<string, unknown>,
	settings: AppSettings | null
): ResultDetail[] {
	const details = resultDetails(data, settings);
	const order = fieldOrder[type];
	return details.sort((a, b) => {
		const indexA = order.findIndex((aliases) => aliases.includes(a.key));
		const indexB = order.findIndex((aliases) => aliases.includes(b.key));
		return (indexA < 0 ? order.length : indexA) - (indexB < 0 ? order.length : indexB);
	});
}

export function formatDateTime(value: unknown, settings: AppSettings | null) {
	if (typeof value !== 'string') return String(value ?? '—');
	const date = new Date(value);
	if (Number.isNaN(date.getTime())) return String(value);
	const locale = getLocale();
	const timeZone = settings?.timeZone || 'Europe/Moscow';
	const parts = Object.fromEntries(
		new Intl.DateTimeFormat(locale, {
			timeZone,
			year: 'numeric',
			month: '2-digit',
			day: '2-digit'
		})
			.formatToParts(date)
			.map(({ type, value }) => [type, value])
	);
	const dateText = (settings?.dateFormat || 'DD.MM.YYYY')
		.replace('YYYY', parts.year)
		.replace('MM', parts.month)
		.replace('DD', parts.day);
	const timeText = new Intl.DateTimeFormat(locale, {
		timeZone,
		hour: '2-digit',
		minute: '2-digit',
		hourCycle: settings?.timeFormat === '12h' ? 'h12' : 'h23'
	}).format(date);
	return `${dateText}, ${timeText}`;
}

export function formatRole(role: unknown) {
	if (typeof role !== 'string') return String(role ?? '');
	return messages[`searchBar.roles.${role.toLowerCase()}`]?.() ?? role;
}

function formatNested(value: unknown, settings: AppSettings | null): string {
	if (Array.isArray(value)) return value.map((item) => formatNested(item, settings)).join(', ');
	if (value && typeof value === 'object') {
		return Object.entries(value as Record<string, unknown>)
			.filter(([key]) => !hiddenFields.has(key))
			.map(([key, item]) => `${fieldLabel(key)}: ${formatValue(key, item, settings)}`)
			.join(' · ');
	}
	return String(value ?? '—');
}

function formatValue(key: string, value: unknown, settings: AppSettings | null): string {
	if (dateFields.has(key)) return formatDateTime(value, settings);
	if (key === 'role' || key === 'currentRole') return formatRole(value);
	if (Array.isArray(value) || (value && typeof value === 'object'))
		return formatNested(value, settings);
	if (typeof value === 'boolean') return value ? m['searchBar.yes']() : m['searchBar.no']();
	return String(value ?? '—');
}

export function resultDetails(
	data: Record<string, unknown>,
	settings: AppSettings | null
): ResultDetail[] {
	return Object.entries(data)
		.filter(
			([key, value]) =>
				!hiddenFields.has(key) &&
				!(key === 'rank' && (value == null || value === '')) &&
				!(['allies', 'rivals'].includes(key) && Array.isArray(value) && value.length === 0)
		)
		.map(([key, raw]) => ({
			key,
			label: fieldLabel(key),
			raw,
			value: formatValue(key, raw, settings)
		}));
}

export function memberName(value: unknown) {
	return value && typeof value === 'object' && 'name' in value ? String(value.name) : String(value);
}

export function memberRole(value: unknown) {
	return value && typeof value === 'object' && 'role' in value
		? formatRole(value.role)
		: formatRole('member');
}

export function isLeader(data: Record<string, unknown>) {
	return String(data.role ?? data.currentRole ?? '').toLowerCase() === 'leader';
}
