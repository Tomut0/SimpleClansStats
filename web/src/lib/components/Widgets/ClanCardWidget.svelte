<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import type { Clan } from './types';
	let { clans }: { clans: Clan[] } = $props();
	const clan = $derived(clans[0]);
</script>

<section
	class="flex h-full flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">
		{m['pages.dashboard.clan_card_type']()}
	</header>
	{#if clan}<div class="grid flex-1 content-center gap-2 p-4">
			<strong class="text-xl">{String(clan.tag ?? clan.name ?? '—')}</strong><span
				>{String(clan.name ?? '')}</span
			><span
				>{m['pages.dashboard.table_balance']()}: {Number(clan.balance ?? 0).toLocaleString()}</span
			><span>{String(clan.status ?? (clan.verified ? m['pages.dashboard.verified']() : '—'))}</span>
		</div>{:else}<div class="flex flex-1 items-center justify-center text-sm text-muted-foreground">
			{m['pages.dashboard.no_clans']()}
		</div>{/if}
</section>
