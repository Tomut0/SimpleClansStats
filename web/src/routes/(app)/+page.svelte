<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { getLocale } from '$lib/paraglide/runtime.js';
	import { APP_NAME, PAGE_SEPARATOR } from '$app/env/public';
	import { invalidate } from '$app/navigation';
	import type { PageProps } from './$types';
	import { serverState } from '$lib/states/serverState.svelte';
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import PageHeader from '$lib/components/PageHeader.svelte';
	import * as Empty from '$lib/components/ui/empty';
	import { Spinner } from '$lib/components/ui/spinner';

	const { data }: PageProps = $props();
	const selectedServer = $derived(
		data.servers.find((server) => server.id === serverState.selectedId)
	);
	const waitingForSnapshot = $derived(
		selectedServer?.lastIngestAt === null && !selectedServer.disabled
	);
	const updatedAt = $derived.by(() => {
		if (!selectedServer?.lastIngestAt) return null;

		const locale = getLocale();
		const timeZone = data.settings?.timeZone ?? 'Europe/Moscow';
		const receivedAt = new Date(selectedServer.lastIngestAt);
		const dateFormatter = new Intl.DateTimeFormat(locale, {
			timeZone,
			year: 'numeric',
			month: '2-digit',
			day: '2-digit'
		});
		const { year, month, day } = Object.fromEntries(
			dateFormatter.formatToParts(receivedAt).map(({ type, value }) => [type, value])
		);
		const date = (data.settings?.dateFormat ?? 'DD.MM.YYYY')
			.replace('YYYY', year)
			.replace('MM', month)
			.replace('DD', day);
		const time = new Intl.DateTimeFormat(locale, {
			timeZone,
			hour: '2-digit',
			minute: '2-digit',
			hourCycle: data.settings?.timeFormat === '12h' ? 'h12' : 'h23'
		}).format(receivedAt);
		return `${date}, ${time}`;
	});

	$effect(() => {
		if (!selectedServer || selectedServer.disabled) return;

		const timer = setInterval(() => {
			void invalidate('app:servers').catch(() => {});
		}, 10_000);
		return () => clearInterval(timer);
	});
</script>

<svelte:head>
	<title>{m['pages.dashboard.title']()}{PAGE_SEPARATOR}{data.settings?.appName || APP_NAME}</title>
</svelte:head>

<main>
	<PageHeader
		icon={LayoutDashboardIcon}
		title={m['pages.dashboard.title']()}
		description={m['pages.dashboard.description']()}
	>
		{#snippet actions()}
			{#if selectedServer}
				<div
					class="flex flex-wrap items-center justify-end gap-x-2 gap-y-1 text-sm text-muted-foreground"
				>
					<span class="inline-flex items-center gap-2">
						<span
							class={[
								'size-1.5 rounded-full',
								selectedServer.online ? 'bg-emerald-500' : 'bg-muted-foreground'
							]}
							aria-hidden="true"
						></span>
						{selectedServer.online ? m['server.online']() : m['server.offline']()}
					</span>
					<span class="h-3 w-px bg-border" aria-hidden="true"></span>
					<span>
						{m['server.updated']()}
						{#if updatedAt && selectedServer.lastIngestAt}
							<time datetime={selectedServer.lastIngestAt} class="text-foreground">{updatedAt}</time
							>
						{:else}
							{m['server.never_updated']()}
						{/if}
					</span>
				</div>
			{/if}
		{/snippet}
	</PageHeader>
	{#if waitingForSnapshot}
		<Empty.Root class="min-h-[65vh]" role="status">
			<Empty.Header>
				<Empty.Media>
					<Spinner role="presentation" aria-hidden="true" class="motion-reduce:animate-none" />
				</Empty.Media>
				<Empty.Title>{m['pages.dashboard.waiting_title']()}</Empty.Title>
				<Empty.Description>{m['pages.dashboard.waiting_description']()}</Empty.Description>
			</Empty.Header>
		</Empty.Root>
	{/if}
</main>
