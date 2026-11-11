<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { API_URL, APP_NAME, PAGE_SEPARATOR } from '$app/env/public';
	import { invalidate } from '$app/navigation';
	import Settings from '@lucide/svelte/icons/settings';
	import PageHeader from '$lib/components/PageHeader.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { resolve } from '$app/paths';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';
	import type { PageProps } from './$types';
	import * as Card from '$lib/components/ui/card/index.js';
	import ChevronRight from '@lucide/svelte/icons/chevron-right';
	import GeneralSettingsCard from './GeneralSettingsCard.svelte';
	import ServersSettingsCard from './ServersSettingsCard.svelte';
	import TimezoneSettingsCard from './TimezoneSettingsCard.svelte';

	const { data }: PageProps = $props();
	const serversEmpty = $derived(data.servers?.length === 0);

	const canSave = $derived(data.settings !== null && data.user?.role === 'admin');
	let saving = $state(false);
	const disabled = $derived(!canSave || saving);

	const sections = $derived([
		{
			id: 'general',
			title: m['pages.settings.categories.general.title'](),
			component: GeneralSettingsCard,
			props: { settings: data.settings, logoUrl: data.logoUrl, disabled }
		},
		{
			id: 'servers',
			title: m['pages.settings.categories.servers.title'](),
			component: ServersSettingsCard,
			props: {
				servers: data.servers,
				canManage: data.user?.role === 'admin',
				settings: data.settings,
				disabled
			}
		},
		{
			id: 'timezone',
			title: m['pages.settings.categories.timezone.title'](),
			component: TimezoneSettingsCard,
			props: { settings: data.settings, disabled }
		}
	]);

	async function saveSettings(event: SubmitEvent) {
		event.preventDefault();
		if (!canSave || saving) return;

		const body = new FormData(event.currentTarget as HTMLFormElement);
		body.set('removeAuthor', String(body.has('removeAuthor')));
		const logo = body.get('logo');
		if (logo instanceof File && !logo.name) body.delete('logo');

		saving = true;
		try {
			const response = await fetch(`${API_URL}/settings`, {
				method: 'POST',
				credentials: 'include',
				body
			});
			if (!response.ok) {
				toast.error(
					response.status === 400 || response.status === 413
						? m['pages.settings.invalid_settings']()
						: m['pages.settings.save_failed']()
				);
				return;
			}
			await invalidate('app:settings');
			toast.success(m['pages.settings.saved']());
		} catch {
			toast.error(m['pages.settings.save_failed']());
		} finally {
			saving = false;
		}
	}

	function notifyIfServersEmpty() {
		if (serversEmpty) {
			toast.error(m['errors.servers_empty']());
		}
	}

	onMount(() => {
		notifyIfServersEmpty();
	});
</script>

<svelte:head>
	<title>{m['pages.settings.title']()}{PAGE_SEPARATOR}{data.settings?.appName || APP_NAME}</title>
</svelte:head>

<main>
	<PageHeader
		icon={Settings}
		title={m['pages.settings.title']()}
		description={m['pages.settings.description']()}
	>
		{#snippet actions()}
			<div class="flex gap-2">
				<Button type="submit" form="app-settings" {disabled}>
					{m['pages.settings.save']()}
				</Button>
				<Button variant="secondary" href={resolve('/')} onclick={notifyIfServersEmpty}>
					{m['pages.settings.backToDashboard']()}
				</Button>
			</div>
		{/snippet}
	</PageHeader>

	<div class="grid grid-cols-4 gap-4">
		<div class="sticky top-8 flex flex-col gap-4 self-start">
			{#each sections as section (section.id)}
				<Card.Root>
					<Card.Content>
						<!-- eslint-disable-next-line svelte/no-navigation-without-resolve -->
						<a href={`#${section.id}`} class="flex justify-between">
							{section.title}
							<ChevronRight class="h-4 w-4" />
						</a>
					</Card.Content>
				</Card.Root>
			{/each}
		</div>

		<div class="col-span-3">
			{#if !data.settings}
				<p class="mb-4 text-destructive" role="alert">{m['pages.settings.load_failed']()}</p>
			{/if}
			<form id="app-settings" class="space-y-4" onsubmit={saveSettings}>
				{#each sections as section (section.id)}
					<section.component id={section.id} title={section.title} {...section.props} />
				{/each}
			</form>
		</div>
	</div>
</main>
