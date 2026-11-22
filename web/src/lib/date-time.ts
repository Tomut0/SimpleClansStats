import type { AppSettings } from '#lib/api/generated/index.js';
import { getLocale } from '#lib/paraglide/runtime.js';

type DateTimeSettings = Pick<AppSettings, 'timeZone' | 'dateFormat' | 'timeFormat'>;

export function formatDateTime(date: Date, settings: DateTimeSettings | null): string {
	const locale = getLocale();
	const timeZone = settings?.timeZone || 'Europe/Moscow';
	const dateFormatter = new Intl.DateTimeFormat(locale, {
		timeZone,
		year: 'numeric',
		month: '2-digit',
		day: '2-digit'
	});
	const { year, month, day } = Object.fromEntries(
		dateFormatter.formatToParts(date).map(({ type, value }) => [type, value])
	);
	const dateText = (settings?.dateFormat || 'DD.MM.YYYY')
		.replace('YYYY', year)
		.replace('MM', month)
		.replace('DD', day);
	const timeText = new Intl.DateTimeFormat(locale, {
		timeZone,
		hour: '2-digit',
		minute: '2-digit',
		hourCycle: settings?.timeFormat === '12h' ? 'h12' : 'h23'
	}).format(date);
	return `${dateText}, ${timeText}`;
}
