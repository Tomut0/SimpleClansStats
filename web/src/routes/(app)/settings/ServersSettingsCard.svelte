<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import AddServerDialog from '$lib/components/AddServerDialog.svelte';
	import * as Empty from '$lib/components/ui/empty/index.js';
	import CircleAlert from '@lucide/svelte/icons/circle-alert';

	let { id, title, empty = false }: { id: string; title: string; empty?: boolean } = $props();
</script>

<Card.Root {id}>
	<Card.Header class="flex items-center justify-between">
		<Card.Title>{title}</Card.Title>
		<Card.Action>
			<AddServerDialog />
		</Card.Action>
	</Card.Header>
	<Separator />
	<Card.Content class="flex flex-col gap-4">
		{#if empty}
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
		{/if}
	</Card.Content>
</Card.Root>
