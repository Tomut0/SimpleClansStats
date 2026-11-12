<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import Check from '@lucide/svelte/icons/check';
	import Crown from '@lucide/svelte/icons/crown';
	import { Tooltip } from 'bits-ui';
	import { m } from '$lib/paraglide/messages.js';
	import type { AppSettings } from '$lib/api/generated';
	import {
		formatRole,
		isLeader,
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
			<div class="flex items-center gap-2">
				{#if result.name}
					<img
						src={(settings?.playerAvatarUrl || 'https://mc-heads.net/avatar/{username}/32').replace(
							'{username}',
							encodeURIComponent(result.name)
						)}
						alt=""
						aria-hidden="true"
						referrerpolicy="no-referrer"
						class="size-8 shrink-0 rounded bg-muted [image-rendering:pixelated]"
						onerror={(event) => ((event.currentTarget as HTMLImageElement).hidden = true)}
					/>
				{/if}
				<Dialog.Title class="leading-5">{result.name}</Dialog.Title>
				{#if result.data.verified === true}
					<span
						class="grid size-5 shrink-0 place-items-center text-primary"
						aria-label={m['searchBar.verified']()}
					>
						<Check class="size-4" aria-hidden="true" />
					</span>
				{/if}
				{#if isLeader(result.data)}
					<Tooltip.Root>
						<Tooltip.Trigger>
							<button
								type="button"
								class="grid size-5 shrink-0 place-items-center p-0 leading-none text-amber-500"
								aria-label={formatRole('leader')}
							>
								<Crown class="size-4" aria-hidden="true" />
							</button>
						</Tooltip.Trigger>
						<Tooltip.Content
							class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md"
						>
							{formatRole('leader')}
						</Tooltip.Content>
					</Tooltip.Root>
				{/if}
			</div>
			<Dialog.Description>{m['searchBar.player']()}</Dialog.Description>
		</Dialog.Header>
		<dl class="grid gap-3 sm:grid-cols-2">
			{#each orderedResultDetails('player', result.data, settings) as detail (detail.key)}
				<div class="min-w-0 rounded-md bg-muted p-3">
					<dt class="text-xs text-muted-foreground">{detail.label}</dt>
					<dd class="mt-1 text-sm break-words">{detail.value}</dd>
				</div>
			{/each}
		</dl>
	</Dialog.Content>
</Dialog.Root>
