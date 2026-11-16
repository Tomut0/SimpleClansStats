<script lang="ts">
	import { API_URL } from '$app/env/public';
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import SearchIcon from '@lucide/svelte/icons/search';
	import { m } from '$lib/paraglide/messages.js';
	import type { AppSettings } from '$lib/api/generated';
	import { serverState } from '$lib/states/serverState.svelte';
	import ClanResultDialog from './ClanResultDialog.svelte';
	import PlayerResultDialog from './PlayerResultDialog.svelte';
	import type { SearchResult } from './searchResultDetails';

	let { settings }: { settings: AppSettings | null } = $props();

	let query = $state('');
	let results = $state<SearchResult[]>([]);
	let selected = $state<SearchResult | null>(null);
	let open = $state(false);
	let loading = $state(false);
	let controller: AbortController | undefined;
	let debounce: ReturnType<typeof setTimeout> | undefined;

	$effect(() => {
		const value = query;
		const serverId = serverState.selectedId;
		controller?.abort();
		if (debounce) clearTimeout(debounce);
		if (value.trim().length < 2 || !serverId) {
			results = [];
			loading = false;
			return;
		}
		loading = true;
		debounce = setTimeout(() => void search(value, serverId), 250);
		return () => {
			if (debounce) clearTimeout(debounce);
			controller?.abort();
		};
	});

	async function search(value: string, serverId: string) {
		const requestController = (controller = new AbortController());
		try {
			const params = new URLSearchParams({ q: value, serverId, limit: '10' });
			const response = await fetch(`${API_URL}/search?${params}`, {
				credentials: 'omit',
				signal: requestController.signal
			});
			const matches = response.ok ? await response.json() : [];
			if (!requestController.signal.aborted) results = matches;
		} catch {
			if (!requestController.signal.aborted) results = [];
		} finally {
			if (!requestController.signal.aborted) loading = false;
		}
	}
</script>

<div class="relative w-64">
	<InputGroup.Root>
		<InputGroup.Input
			placeholder={m['searchBar.placeholder']()}
			value={query}
			autocomplete="off"
			oninput={(event) => (query = event.currentTarget.value)}
		/>
		<InputGroup.Addon><SearchIcon /></InputGroup.Addon>
	</InputGroup.Root>
	{#if query.trim().length >= 2}
		<div
			class="absolute top-full right-0 left-0 z-40 mt-1 max-h-80 overflow-y-auto rounded-md border bg-popover p-1 text-popover-foreground shadow-md"
		>
			{#if loading}
				<p class="p-3 text-sm text-muted-foreground">{m['searchBar.loading']()}</p>
			{:else if results.length}
				{#each results as result (`${result.type}-${result.name}`)}
					<button
						class="flex w-full items-center justify-between rounded-sm px-3 py-2 text-left text-sm hover:bg-accent"
						onclick={() => {
							selected = result;
							open = true;
						}}
					>
						<span class="truncate">{result.name}</span>
						<span class="ml-3 text-xs text-muted-foreground"
							>{result.type === 'clan' ? m['searchBar.clan']() : m['searchBar.player']()}</span
						>
					</button>
				{/each}
			{:else}
				<p class="p-3 text-sm text-muted-foreground">{m['searchBar.noResults']()}</p>
			{/if}
		</div>
	{/if}
</div>

{#if selected?.type === 'clan'}
	<ClanResultDialog bind:open result={selected} {settings} />
{:else if selected?.type === 'player'}
	<PlayerResultDialog bind:open result={selected} {settings} />
{/if}
