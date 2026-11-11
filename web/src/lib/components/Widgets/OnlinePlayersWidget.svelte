<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import type { Player } from './types';
	let { players, avatarUrlTemplate }: { players: Player[]; avatarUrlTemplate: string } = $props();
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">
		{m['pages.dashboard.widget_online']()}
	</header>
	<ul class="min-h-0 flex-1 space-y-2 overflow-auto px-4 pb-4">
		{#each players as player}<li class="flex items-center gap-3">
				{#if player.name}<img
					src={avatarUrlTemplate.replace('{username}', encodeURIComponent(String(player.name)))}
					alt=""
					aria-hidden="true"
					loading="lazy"
					referrerpolicy="no-referrer"
					class="size-8 rounded bg-muted [image-rendering:pixelated]"
					onerror={(event) => ((event.currentTarget as HTMLImageElement).hidden = true)}
				/>{/if}
				<span class="min-w-0 flex-1 truncate">{String(player.name ?? '—')}</span><span>{String(player.clanTag ?? '')}</span>
			</li>{:else}<li class="text-sm text-muted-foreground">
				{m['pages.dashboard.no_players_online']()}
			</li>{/each}
	</ul>
</section>
