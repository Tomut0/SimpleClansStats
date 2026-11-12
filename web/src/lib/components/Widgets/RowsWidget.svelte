<script lang="ts">
	import * as Table from '$lib/components/ui/table';
	import { m } from '$lib/paraglide/messages.js';
	let {
		title,
		rows,
		fields
	}: { title: string; rows: Record<string, unknown>[]; fields: string[] } = $props();
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{title}</header>
	<div class="min-h-0 flex-1 overflow-auto px-4 pb-4">
		<Table.Root class="text-xs">
			<Table.Header class="bg-muted/50"
				><Table.Row
					>{#each fields as field}<Table.Head class="h-8 px-2 font-normal text-muted-foreground"
							>{field}</Table.Head
						>{/each}</Table.Row
				></Table.Header
			>
			<Table.Body>
				{#each rows as row, index (`${row.id ?? row.uuid ?? row.name ?? row.tag ?? index}`)}
					<Table.Row
						>{#each fields as field}<Table.Cell class="p-2">{String(row[field] ?? '—')}</Table.Cell
							>{/each}</Table.Row
					>
				{:else}<Table.Row
						><Table.Cell colspan={fields.length} class="h-24 text-center text-muted-foreground"
							>{m['pages.dashboard.no_widget_data']()}</Table.Cell
						></Table.Row
					>{/each}
			</Table.Body>
		</Table.Root>
	</div>
</section>
