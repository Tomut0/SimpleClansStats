<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import * as Table from '$lib/components/ui/table';
	import type { Clan } from './types';
	let { clans }: { clans: Clan[] } = $props();
	const topClans = $derived(
		[...clans].sort((a, b) => Number(b.balance ?? 0) - Number(a.balance ?? 0)).slice(0, 5)
	);
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{m['pages.dashboard.widget_top']()}</header>
	<div class="min-h-0 flex-1 overflow-auto px-4 pb-4">
		<Table.Root class="text-xs">
			<Table.Header class="bg-muted/50">
				<Table.Row>
					<Table.Head class="h-8 px-2 py-1 font-normal text-muted-foreground">#</Table.Head>
					<Table.Head class="h-8 px-2 py-1 font-normal text-muted-foreground">
						{m['pages.dashboard.table_clan']()}
					</Table.Head>
					<Table.Head class="h-8 px-2 py-1 text-right font-normal text-muted-foreground">
						{m['pages.dashboard.table_balance']()}
					</Table.Head>
					<Table.Head class="h-8 px-2 py-1 text-right font-normal text-muted-foreground">
						{m['pages.dashboard.table_players']()}
					</Table.Head>
					<Table.Head class="h-8 px-2 py-1 text-right font-normal text-muted-foreground">
						{m['pages.dashboard.table_kd']()}
					</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each topClans as clan, index (`${clan.tag ?? clan.name}-${index}`)}
					{@const players =
						clan.playerCount ??
						clan.memberCount ??
						(Array.isArray(clan.members) ? clan.members.length : null)}
					<Table.Row>
						<Table.Cell class="p-2">{index + 1}</Table.Cell>
						<Table.Cell class="p-2 font-medium">
							{String(clan.tag ?? clan.name ?? '—')}
						</Table.Cell>
						<Table.Cell class="p-2 text-right">
							{Number(clan.balance ?? 0).toLocaleString()}
						</Table.Cell>
						<Table.Cell class="p-2 text-right">
							{players === null ? '—' : Number(players).toLocaleString()}
						</Table.Cell>
						<Table.Cell class="p-2 text-right">
							{clan.kdRatio ?? clan.kd ?? clan.kdr ?? '—'}
						</Table.Cell>
					</Table.Row>
				{:else}
					<Table.Row>
						<Table.Cell colspan={5} class="h-24 text-center text-muted-foreground">
							{m['pages.dashboard.no_clans']()}
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>
</section>
