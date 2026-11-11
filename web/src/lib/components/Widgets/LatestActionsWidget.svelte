<script lang="ts">
	import Swords from '@lucide/svelte/icons/swords';
	import UserRoundPlus from '@lucide/svelte/icons/user-round-plus';
	import ShieldCheck from '@lucide/svelte/icons/shield-check';
	import Wallet from '@lucide/svelte/icons/wallet';
	import UserRoundMinus from '@lucide/svelte/icons/user-round-minus';
	import Activity from '@lucide/svelte/icons/activity';
	import { getLocale } from '$lib/paraglide/runtime.js';
	import { m } from '$lib/paraglide/messages.js';
	import type { DashboardEvent } from './types';
	let { events, title = m['pages.dashboard.widget_actions']() }: { events: DashboardEvent[]; title?: string } = $props();

	function eventIcon(event: DashboardEvent) {
		const type = `${event.type ?? ''} ${event.message ?? ''}`.toLowerCase();
		if (/war/.test(type)) return { Icon: Swords, color: 'bg-rose-500/15 text-rose-400' };
		if (/join|invite/.test(type)) return { Icon: UserRoundPlus, color: 'bg-emerald-500/15 text-emerald-400' };
		if (/verif|status/.test(type)) return { Icon: ShieldCheck, color: 'bg-blue-500/15 text-blue-400' };
		if (/balance|deposit|withdraw|money/.test(type)) return { Icon: Wallet, color: 'bg-amber-500/15 text-amber-400' };
		if (/leave|kick|quit/.test(type)) return { Icon: UserRoundMinus, color: 'bg-rose-500/15 text-rose-400' };
		return { Icon: Activity, color: 'bg-muted text-muted-foreground' };
	}

	function eventDate(event: DashboardEvent) {
		const value = event.timestamp ?? event.occurredAt ?? event.createdAt ?? event.at;
		if (typeof value !== 'string' && typeof value !== 'number') return null;
		const date = new Date(value);
		return Number.isNaN(date.getTime()) ? null : date;
	}

	function relativeTime(date: Date) {
		const seconds = Math.round((date.getTime() - Date.now()) / 1000);
		const units: [Intl.RelativeTimeFormatUnit, number][] = [
			['year', 31536000], ['month', 2592000], ['day', 86400], ['hour', 3600], ['minute', 60]
		];
		const [unit, size] = units.find(([, size]) => Math.abs(seconds) >= size) ?? ['second', 1];
		return new Intl.RelativeTimeFormat(getLocale(), { numeric: 'auto' }).format(Math.round(seconds / size), unit);
	}
</script>

<section class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground">
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{title}</header>
	<ul class="min-h-0 flex-1 space-y-1 overflow-auto px-4 pb-4">
		{#each events as event, index (`${event.id ?? event.timestamp ?? index}`)}
			{@const icon = eventIcon(event)}
			{@const Icon = icon.Icon}
			{@const date = eventDate(event)}
			<li class="flex min-w-0 items-center gap-3 py-2">
				<span class={`grid size-10 shrink-0 place-items-center rounded-xl ${icon.color}`}><Icon class="size-5" aria-hidden="true" /></span>
				<div class="min-w-0 flex-1">
					<p class="truncate text-sm">{String(event.message ?? event.type ?? JSON.stringify(event))}</p>
					{#if date}<time class="text-sm text-muted-foreground" datetime={date.toISOString()}>{relativeTime(date)}</time>{/if}
				</div>
			</li>
		{:else}<li class="text-sm text-muted-foreground">{m['pages.dashboard.no_recent_actions']()}</li>{/each}
	</ul>
</section>
