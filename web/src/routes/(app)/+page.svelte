<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { getLocale } from '$lib/paraglide/runtime.js';
	import { APP_NAME, PAGE_SEPARATOR } from '$app/env/public';
	import { API_URL } from '$app/env/public';
	import type { PageProps } from './$types';
	import { serverState } from '$lib/states/serverState.svelte';
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import PageHeader from '$lib/components/PageHeader.svelte';
	import * as Empty from '$lib/components/ui/empty';
	import { Spinner } from '$lib/components/ui/spinner';
	import DashboardGrid from '$lib/components/Widgets/DashboardGrid.svelte';
	import Plus from '@lucide/svelte/icons/plus';
	import SlidersHorizontal from '@lucide/svelte/icons/sliders-horizontal';
	import RotateCcw from '@lucide/svelte/icons/rotate-ccw';
	import Save from '@lucide/svelte/icons/save';
	import { Tooltip } from 'bits-ui';

	const { data }: PageProps = $props();
	let editing = $state(false);
	let addWidget = $state(() => {});
	let restoreDefaults = $state(() => {});
	let saveDashboard = $state(() => {});
	const selectedServer = $derived(
		data.servers.find((server) => server.id === serverState.selectedId)
	);

	$effect.pre(() => {
		if (selectedServer && !selectedServer.disabled) return;
		serverState.selectedId = data.servers.find((server) => !server.disabled)?.id ?? '';
	});

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
				<div class="flex flex-wrap items-center justify-end gap-x-2 gap-y-1 text-sm text-muted-foreground">
					<span class="inline-flex items-center gap-2"><span class={['size-1.5 rounded-full', selectedServer.online ? 'bg-emerald-500' : 'bg-muted-foreground']} aria-hidden="true"></span>{selectedServer.online ? m['server.online']() : m['server.offline']()}</span>
					<span class="h-3 w-px bg-border" aria-hidden="true"></span>
					<span>{m['server.updated']()}{#if updatedAt && selectedServer.lastIngestAt}<time datetime={selectedServer.lastIngestAt} class="text-foreground">{updatedAt}</time>{:else}{m['server.never_updated']()}{/if}</span>
					{#if selectedServer && !selectedServer.disabled}
						{#if editing}
							<Tooltip.Root><Tooltip.Trigger><button class="grid size-9 place-items-center rounded-md border text-foreground" aria-label={m['pages.dashboard.add_widget_button']()} onclick={addWidget}><Plus class="size-4" aria-hidden="true" /></button></Tooltip.Trigger><Tooltip.Content class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md">{m['pages.dashboard.add_widget_button']()}</Tooltip.Content></Tooltip.Root>
							<Tooltip.Root><Tooltip.Trigger><button class="grid size-9 place-items-center rounded-md border text-foreground" aria-label={m['pages.dashboard.restore_default']()} onclick={restoreDefaults}><RotateCcw class="size-4" aria-hidden="true" /></button></Tooltip.Trigger><Tooltip.Content class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md">{m['pages.dashboard.restore_default']()}</Tooltip.Content></Tooltip.Root>
							<Tooltip.Root><Tooltip.Trigger><button class="grid size-9 place-items-center rounded-md bg-primary text-primary-foreground" aria-label={m['pages.dashboard.save_dashboard']()} onclick={saveDashboard}><Save class="size-4" aria-hidden="true" /></button></Tooltip.Trigger><Tooltip.Content class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md">{m['pages.dashboard.save_dashboard']()}</Tooltip.Content></Tooltip.Root>
						{:else}
						<Tooltip.Root><Tooltip.Trigger><button class="grid size-9 place-items-center rounded-md border text-foreground" aria-label="Customize dashboard" onclick={() => (editing = true)}><SlidersHorizontal class="size-4" aria-hidden="true" /></button></Tooltip.Trigger><Tooltip.Content class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md">Customize dashboard</Tooltip.Content></Tooltip.Root>
						{/if}
					{/if}
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
	{#if selectedServer && !selectedServer.disabled}
		{#key selectedServer.id}
			<DashboardGrid serverId={selectedServer.id} apiUrl={API_URL} avatarUrlTemplate={data.settings?.playerAvatarUrl ?? 'https://mc-heads.net/avatar/{username}/32'} bind:editing bind:onAddWidget={addWidget} bind:onRestoreDefaults={restoreDefaults} bind:onSaveDashboard={saveDashboard} />
		{/key}
	{/if}
</main>
