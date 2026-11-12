<script lang="ts">
	import Clipboard from '@lucide/svelte/icons/clipboard';
	import { API_URL } from '$app/env/public';
	import { invalidate } from '$app/navigation';
	import { toast } from 'svelte-sonner';
	import type { ServerRegistration } from '$lib/api/generated';
	import { m } from '$lib/paraglide/messages.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const id = $props.id();
	let displayName = $state('');
	let saving = $state(false);
	let credentials = $state<{ endpoint: string; apiKey: string; command: string } | null>(null);

	async function register(event: SubmitEvent) {
		event.preventDefault();
		if (saving || !displayName.trim()) return;

		saving = true;
		try {
			const response = await fetch(`${API_URL}/servers`, {
				method: 'POST',
				credentials: 'include',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ displayName: displayName.trim() })
			});
			if (response.status !== 201) throw new Error('Server registration failed');
			const { apiKey }: ServerRegistration = await response.json();
			const endpoint = new URL(`${API_URL}/ingest`, window.location.origin).href;
			credentials = {
				endpoint,
				apiKey,
				command: `/scstats config apply --endpoint ${endpoint} --api-key ${apiKey}`
			};
			displayName = '';
		} catch {
			toast.error(m['pages.settings.categories.servers.dialog.create_failed']());
			return;
		} finally {
			saving = false;
		}

		try {
			await invalidate('app:servers');
		} catch {
			toast.error(m['pages.settings.categories.servers.dialog.refresh_failed']());
		}
	}

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

<Dialog.Root
	onOpenChange={(open) => {
		if (!open && !saving) credentials = null;
	}}
>
	<Dialog.Trigger>
		{#snippet child({ props })}
			<Button type="button" {...props}>{m['pages.settings.categories.servers.add_server']()}</Button
			>
		{/snippet}
	</Dialog.Trigger>
	<Dialog.Content
		showCloseButton={false}
		escapeKeydownBehavior={saving || credentials ? 'ignore' : 'close'}
		interactOutsideBehavior={saving || credentials ? 'ignore' : 'close'}
		class="max-h-[calc(100dvh-2rem)] gap-5 overflow-y-auto rounded-2xl border border-border p-5 sm:max-w-136"
	>
		<Dialog.Header>
			<Dialog.Title class="text-xl">
				{m['pages.settings.categories.servers.dialog.title']()}
			</Dialog.Title>
			<Dialog.Description class="text-base">
				{#if credentials}
					{m['pages.settings.categories.servers.dialog.description']()}
					<span class="font-semibold text-primary dark:text-primary-foreground">SCCollector</span>.
				{:else}
					{m['pages.settings.categories.servers.dialog.register_description']()}
				{/if}
			</Dialog.Description>
		</Dialog.Header>
		{#if credentials}
			<p class="text-sm text-muted-foreground">
				{m['pages.settings.categories.servers.dialog.key_once']()}
			</p>
			{@render copyField(
				'endpoint',
				m['pages.settings.categories.servers.dialog.endpoint_title'](),
				credentials.endpoint
			)}
			{@render copyField(
				'api-key',
				m['pages.settings.categories.servers.dialog.api_key_title'](),
				credentials.apiKey
			)}
			<Separator />
			{@render copyField(
				'command',
				m['pages.settings.categories.servers.dialog.command_title'](),
				credentials.command,
				m['pages.settings.categories.servers.dialog.command']()
			)}
			<Dialog.Close>
				{#snippet child({ props })}
					<Button type="button" {...props}
						>{m['pages.settings.categories.servers.dialog.done']()}</Button
					>
				{/snippet}
			</Dialog.Close>
		{:else}
			<form class="grid gap-4" onsubmit={register}>
				<div class="grid gap-2">
					<Label for={`${id}-display-name`}
						>{m['pages.settings.categories.servers.dialog.server_name']()}</Label
					>
					<Input
						id={`${id}-display-name`}
						bind:value={displayName}
						required
						maxlength={100}
						disabled={saving}
					/>
				</div>
				<Button
					type="submit"
					disabled={saving || !displayName.trim()}
					class="h-10 w-full rounded-xl text-base"
				>
					{m['pages.settings.categories.servers.dialog.submit']()}
				</Button>
			</form>
		{/if}
	</Dialog.Content>
</Dialog.Root>
