<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import type { Player } from './types';
	let { players }: { players: Player[] } = $props();
	const player = $derived(players[0]);
</script>
<section class="flex h-full flex-col overflow-hidden rounded-lg border bg-card text-card-foreground">
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{m['pages.dashboard.player_card_type']()}</header>
	{#if player}<div class="grid flex-1 content-center gap-2 p-4"><strong class="text-xl">{String(player.name ?? '—')}</strong><span>{String(player.clanTag ?? '—')}</span><span>{String(player.rank ?? '—')}</span><span>{m['pages.dashboard.widget_online']()}: {String(player.online ?? '—')}</span></div>{:else}<div class="flex flex-1 items-center justify-center text-sm text-muted-foreground">{m['pages.dashboard.no_widget_data']()}</div>{/if}
</section>
