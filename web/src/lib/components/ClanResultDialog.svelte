<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import Check from '@lucide/svelte/icons/check';
	import ShieldCheck from '@lucide/svelte/icons/shield-check';
	import { m } from '$lib/paraglide/messages.js';
	import type { AppSettings } from '$lib/api/generated';
	import {
		memberName,
		memberRole,
		orderedResultDetails,
		type SearchResult
	} from './searchResultDetails';

	let {
		open = $bindable(false),
		result,
		settings
	}: { open?: boolean; result: SearchResult; settings: AppSettings | null } = $props();
</script>

<Dialog.Root bind:open>
	<Dialog.Content class="max-h-[85vh] overflow-y-auto sm:max-w-xl">
		<Dialog.Header>
			<div class="flex h-5 items-center gap-2">
				<Dialog.Title class="leading-5">{result.name}</Dialog.Title>
				{#if result.data.verified === true}
					<span
						class="grid size-5 shrink-0 place-items-center text-primary"
						aria-label={m['searchBar.verified']()}
					>
						<Check class="size-4" aria-hidden="true" />
					</span>
				{/if}
				{#if result.data.permanent === true}
					<span
						class="grid size-5 shrink-0 place-items-center text-primary"
						aria-label={m['searchBar.permanent']()}
					>
						<ShieldCheck class="size-4" aria-hidden="true" />
					</span>
				{/if}
			</div>
			<Dialog.Description>{m['searchBar.clan']()}</Dialog.Description>
		</Dialog.Header>
		<dl class="grid gap-3 sm:grid-cols-2">
			{#each orderedResultDetails('clan', result.data, settings) as detail (detail.key)}
				<div class="min-w-0 rounded-md bg-muted p-3">
					<dt class="text-xs text-muted-foreground">{detail.label}</dt>
					<dd class="mt-1 text-sm break-words">
						{#if detail.key === 'members' && Array.isArray(detail.raw)}
							{#if detail.raw.length}
								<ul class="space-y-1">
									{#each detail.raw as member, index (`${memberName(member)}-${index}`)}
										<li>
											{memberName(member)}
											<span class="text-muted-foreground">— {memberRole(member)}</span>
										</li>
									{/each}
								</ul>
							{:else}—{/if}
						{:else}
							{detail.value}
						{/if}
					</dd>
				</div>
			{/each}
		</dl>
	</Dialog.Content>
</Dialog.Root>
