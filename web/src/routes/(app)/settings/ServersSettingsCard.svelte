<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { API_URL } from '$app/env/public';
	import { invalidate } from '$app/navigation';
	import { toast } from 'svelte-sonner';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import AddServerDialog from '$lib/components/AddServerDialog.svelte';
	import * as Empty from '$lib/components/ui/empty/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import CircleAlert from '@lucide/svelte/icons/circle-alert';
	import ServerIcon from '@lucide/svelte/icons/server';
	import type { AppSettings, Server } from '$lib/api/generated';

	let {
		id,
		title,
		servers = [],
		canManage = false,
		settings = null,
		disabled = true
	}: {
		id: string;
		title: string;
		servers?: Server[];
		canManage?: boolean;
		settings?: AppSettings | null;
		disabled?: boolean;
	} = $props();

	let deleting = $state(false);

	async function deleteServer(server: Server) {
		if (
			deleting ||
			!window.confirm(
				m['pages.settings.categories.servers.delete_confirm']({ name: server.displayName })
			)
		)
			return;

		deleting = true;
		try {
			const response = await fetch(`${API_URL}/servers/${server.id}`, {
				method: 'DELETE',
				credentials: 'include'
			});
			if (response.status !== 204) throw new Error('Server deletion failed');
			await invalidate('app:servers');
		} catch {
			toast.error(m['pages.settings.categories.servers.delete_failed']());
		} finally {
			deleting = false;
		}
	}
</script>

<Card.Root {id}>
	<Card.Header class="flex items-center justify-between">
		<Card.Title>{title}</Card.Title>
		{#if canManage}
			<Card.Action>
				<AddServerDialog />
			</Card.Action>
		{/if}
	</Card.Header>
	<Separator />
	<Card.Content class="flex flex-col gap-4">
		<Field.Field orientation="horizontal">
			<Field.Content>
				<Field.Label for="server-online-timeout">
					{m['pages.settings.categories.servers.online_timeout']()}
				</Field.Label>
				<Field.Description
					>{m['pages.settings.categories.servers.online_timeout_desc']()}</Field.Description
				>
			</Field.Content>
			<Input
				id="server-online-timeout"
				name="serverOnlineTimeoutMinutes"
				type="number"
				value={settings?.serverOnlineTimeoutMinutes}
				min={1}
				max={1440}
				required
				{disabled}
				class="max-w-28"
			/>
		</Field.Field>
		<Field.Field orientation="horizontal">
			<Field.Content>
				<Field.Label for="ingest-retention">
					{m['pages.settings.categories.servers.retention']()}
				</Field.Label>
				<Field.Description
					>{m['pages.settings.categories.servers.retention_desc']()}</Field.Description
				>
			</Field.Content>
			<Input
				id="ingest-retention"
				name="ingestRetentionDays"
				type="number"
				value={settings?.ingestRetentionDays}
				min={0}
				max={36500}
				required
				{disabled}
				class="max-w-28"
			/>
		</Field.Field>
		<Separator />
		{#if servers.length === 0}
			<Empty.Root>
				<Empty.Header>
					<Empty.Media variant="icon">
						<CircleAlert />
					</Empty.Media>
					<Empty.Title>{m['pages.settings.categories.servers.empty.title']()}</Empty.Title>
					<Empty.Description>
						{m['pages.settings.categories.servers.empty.description']()}
					</Empty.Description>
				</Empty.Header>
			</Empty.Root>
		{:else}
			<ul class="grid gap-2">
				{#each servers as server (server.id)}
					<li class="flex items-center justify-between gap-2">
						<span class="flex items-center gap-2"><ServerIcon class="w-4 h-4 text-foreground"/> {server.displayName}</span>
						<div class="flex items-center gap-2">
							<span class="text-sm text-muted-foreground">
								{server.online
									? m['pages.settings.categories.servers.online']()
									: m['pages.settings.categories.servers.offline']()}
							</span>
							{#if canManage}
								<Button
									type="button"
									variant="ghost"
									size="sm"
									class="text-destructive"
									disabled={deleting}
									onclick={() => deleteServer(server)}
								>
									{m['pages.settings.categories.servers.delete']()}
								</Button>
							{/if}
						</div>
					</li>
				{/each}
			</ul>
		{/if}
	</Card.Content>
</Card.Root>
