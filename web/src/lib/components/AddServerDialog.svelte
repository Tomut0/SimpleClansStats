<script lang="ts">
	import Clipboard from '@lucide/svelte/icons/clipboard';
	import { toast } from 'svelte-sonner';
	import { m } from '$lib/paraglide/messages.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	// shortcut: example configuration until the server registration API returns credentials.
	const endpoint = 'https://ststats.example.com/api/v1/ingest';
	const apiKey = 'sc-...';
	const command = `/scstats config apply --endpoint ${endpoint} --api-key ${apiKey}`;
	const id = $props.id();

	async function copy(value: string) {
		try {
			await navigator.clipboard.writeText(value);
			toast.success(m['pages.settings.categories.servers.dialog.copied']());
		} catch {
			toast.error(m['pages.settings.categories.servers.dialog.copy_failed']());
		}
	}
</script>

{#snippet copyField(name: string, label: string, value: string, hint?: string)}
	<div class="grid min-w-0 gap-2.5">
		<Label
			for={`${id}-${name}`}
			class={hint ? 'text-base font-normal text-muted-foreground' : 'text-base'}
		>
			{hint ?? label}
		</Label>
		<InputGroup.Root class="h-10 rounded-xl border-border">
			<InputGroup.Input
				id={`${id}-${name}`}
				{value}
				readonly
				class="min-w-0 truncate text-base text-muted-foreground md:text-base"
			/>
			<InputGroup.Addon align="inline-end" class="pr-2">
				<InputGroup.Button
					size="icon-sm"
					class="text-muted-foreground"
					aria-label={m['pages.settings.categories.servers.dialog.copy']({ label })}
					onclick={() => copy(value)}
				>
					<Clipboard class="size-5" />
				</InputGroup.Button>
			</InputGroup.Addon>
		</InputGroup.Root>
	</div>
{/snippet}

<Dialog.Root>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<Button {...props}>{m['pages.settings.categories.servers.add_server']()}</Button>
		{/snippet}
	</Dialog.Trigger>
	<Dialog.Content
		showCloseButton={false}
		class="max-h-[calc(100dvh-2rem)] gap-5 overflow-y-auto rounded-2xl border border-border p-5 sm:max-w-[34rem]"
	>
		<Dialog.Header>
			<Dialog.Title class="text-xl">
				{m['pages.settings.categories.servers.dialog.title']()}
			</Dialog.Title>
			<Dialog.Description class="text-base">
				{m['pages.settings.categories.servers.dialog.description']()}
				<span class="font-semibold text-primary dark:text-primary-foreground">SCCollector</span>.
			</Dialog.Description>
		</Dialog.Header>
		{@render copyField(
			'endpoint',
			m['pages.settings.categories.servers.dialog.endpoint_title'](),
			endpoint
		)}
		{@render copyField(
			'api-key',
			m['pages.settings.categories.servers.dialog.api_key_title'](),
			apiKey
		)}
		<Separator />
		{@render copyField(
			'command',
			m['pages.settings.categories.servers.dialog.command_title'](),
			command,
			m['pages.settings.categories.servers.dialog.command']()
		)}
		<Button disabled class="h-10 w-full rounded-xl text-base">
			{m['pages.settings.categories.servers.dialog.submit']()}
		</Button>
	</Dialog.Content>
</Dialog.Root>
