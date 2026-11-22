<script lang="ts">
	import { m } from '#lib/paraglide/messages.js';
	import { API_URL, APP_NAME, PAGE_SEPARATOR } from '$app/env/public';
	import { formatDateTime } from '#lib/date-time.js';
	import type { PageProps } from './$types';
	import { serverState } from '#lib/states/serverState.svelte.js';
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import PageHeader from '#lib/components/PageHeader.svelte';
	import * as Empty from '#lib/components/ui/empty/index.js';
	import { Spinner } from '#lib/components/ui/spinner/index.js';
	import DashboardGrid, {
		type DashboardActions
	} from '#lib/components/Widgets/DashboardGrid.svelte';
	import Plus from '@lucide/svelte/icons/plus';
	import SlidersHorizontal from '@lucide/svelte/icons/sliders-horizontal';
	import RotateCcw from '@lucide/svelte/icons/rotate-ccw';
	import Save from '@lucide/svelte/icons/save';
	import DashboardActionButton from './DashboardActionButton.svelte';

	const { data }: PageProps = $props();
	let editing = $state(false);
	let dashboardActions = $state<DashboardActions | null>(null);
	let receivedSnapshot = $state(false);
	const selectedServer = $derived(
		data.servers.find((server) => server.id === serverState.selectedId)
	);
	const firstAvailableServer = $derived(data.servers.find((server) => !server.disabled));

	$effect.pre(() => {
		if (selectedServer && !selectedServer.disabled) return;
		serverState.selectedId = firstAvailableServer?.id ?? '';
	});

	const waitingForSnapshot = $derived(
		selectedServer?.lastIngestAt === null && !selectedServer.disabled && !receivedSnapshot
	);

	const updatedAt = $derived(
		selectedServer?.lastIngestAt
			? formatDateTime(new Date(selectedServer.lastIngestAt), data.settings)
			: null
	);
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
					<span class="inline-flex items-center gap-2"
						><span
							class={[
								'size-1.5 rounded-full',
								selectedServer.online ? 'bg-emerald-500' : 'bg-muted-foreground'
							]}
							aria-hidden="true"
						></span>{selectedServer.online ? m['server.online']() : m['server.offline']()}</span
					>
					<span class="h-3 w-px bg-border" aria-hidden="true"></span>
					<span
						>{m['server.updated']()}
						{#if updatedAt && selectedServer.lastIngestAt}<time
								datetime={selectedServer.lastIngestAt}
								class="text-foreground">{updatedAt}</time
							>{:else}
							{m['server.never_updated']()}{/if}</span
					>
					{#if !selectedServer.disabled && data.user?.role === 'admin' && dashboardActions}
						{#if editing}
							<DashboardActionButton
								label={m['pages.dashboard.add_widget_button']()}
								icon={Plus}
								onclick={dashboardActions.addWidget}
							/>
							<DashboardActionButton
								label={m['pages.dashboard.restore_default']()}
								icon={RotateCcw}
								onclick={dashboardActions.restoreDefaults}
							/>
							<DashboardActionButton
								label={m['pages.dashboard.save_dashboard']()}
								icon={Save}
								onclick={dashboardActions.saveDashboard}
								primary
							/>
						{:else}
							<DashboardActionButton
								label="Customize dashboard"
								icon={SlidersHorizontal}
								onclick={dashboardActions.beginEditing}
							/>
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
			<div class={{ invisible: waitingForSnapshot }}>
				<DashboardGrid
					serverId={selectedServer.id}
					apiUrl={API_URL}
					avatarUrlTemplate={data.settings?.playerAvatarUrl ??
						'https://mc-heads.net/avatar/{username}/32'}
					bind:editing
					onActionsReady={(actions) => (dashboardActions = actions)}
					onSnapshot={(received) => (receivedSnapshot = received)}
				/>
			</div>
		{/key}
	{/if}
</main>
