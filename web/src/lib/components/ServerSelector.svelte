<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js';
	import ServerIcon from '@lucide/svelte/icons/server';
	import { m } from '$lib/paraglide/messages.js';
	import type { Server } from '$lib/api/generated';
	import { serverState } from '$lib/states/serverState.svelte';

	let { servers = [] }: { servers: Server[] } = $props();

	const items = $derived(
		servers.map((s) => ({
			value: s.id,
			label: s.displayName,
			disabled: s.disabled ?? true
		}))
	);

	const selectedValue = $derived(
		items.some((item) => item.value === serverState.selectedId) ? serverState.selectedId : ''
	);
</script>

{#if items.length > 0}
	<Select.Root
		type="single"
		name="server"
		{items}
		value={selectedValue}
		onValueChange={(next) => (serverState.selectedId = next)}
	>
		<Select.Trigger class="min-w-50">
			<ServerIcon class="h-4 w-4 text-muted-foreground" />
			<Select.Value placeholder={m['serverSelector.placeholder']()} />
		</Select.Trigger>
		<Select.Content>
			<Select.Group>
				{#each items as item (item.value)}
					<Select.Item value={item.value} label={item.label} disabled={item.disabled}>
						{item.label}
					</Select.Item>
				{/each}
			</Select.Group>
		</Select.Content>
	</Select.Root>
{/if}
