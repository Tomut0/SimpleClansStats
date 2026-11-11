<script lang="ts">
	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages.js';
	import type { Clan } from './types';
	let { clans }: { clans: Clan[] } = $props();
	const data = $derived(
		[...clans]
			.sort((a, b) => Number(b.balance ?? 0) - Number(a.balance ?? 0))
			.slice(0, 8)
			.map((clan) => ({
				name: String(clan.tag ?? clan.name ?? '—'),
				balance: Number(clan.balance ?? 0)
			}))
	);
	const maxBalance = $derived(Math.max(1, ...data.map(({ balance }) => balance)));
</script>

<section class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground">
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{m['pages.dashboard.economy_type']()}</header>
	{#if data.length}
		<div class="min-h-0 flex-1 overflow-auto px-4 pb-4">
			<Table.Root>
				<Table.Header><Table.Row><Table.Head>{m['pages.dashboard.table_clan']()}</Table.Head><Table.Head class="text-right">{m['pages.dashboard.table_balance']()}</Table.Head></Table.Row></Table.Header>
				<Table.Body>
					{#each data as clan (clan.name)}
						<Table.Row>
							<Table.Cell class="w-2/3 font-medium">{clan.name}<div class="mt-1 h-1.5 rounded-full bg-muted"><div class="h-full rounded-full bg-primary" style={`width: ${Math.max(2, clan.balance / maxBalance * 100)}%`}></div></div></Table.Cell>
							<Table.Cell class="text-right tabular-nums">{clan.balance.toLocaleString(undefined, { maximumFractionDigits: 2 })}</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		</div>
	{:else}<div class="flex flex-1 items-center justify-center text-sm text-muted-foreground">{m['pages.dashboard.no_clans']()}</div>{/if}
</section>
