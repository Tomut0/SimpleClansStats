<script lang="ts">
	import { API_URL } from '$app/env/public';
	import BrandIcon from '$lib/components/BrandIcon.svelte';
	import ThemeSwitcher from '$lib/components/ThemeSwitcher.svelte';
	import * as Avatar from '$lib/components/ui/avatar/index.js';
	import Button from './ui/button/button.svelte';
	import { m } from '$lib/paraglide/messages.js';
	import SearchBar from './SearchBar.svelte';
	import ServerSelector from './ServerSelector.svelte';
	import Separator from './ui/separator/separator.svelte';
	import type { AppSettings, Server, User } from '$lib/api/generated';
	import { refreshAll } from '$app/navigation';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import { resolve } from '$app/paths';
	import { locales, setLocale } from '$lib/paraglide/runtime';
	import { getLocale } from '$lib/paraglide/runtime';
	import Languages from '@lucide/svelte/icons/languages';

	let {
		servers = [],
		user,
		settings
	}: { servers: Server[]; user: User | null; settings: AppSettings | null } = $props();

	let errorMessage = $state('');
	let isSubmitting = $state(false);

	async function logout() {
		try {
			isSubmitting = true;

			const res = await fetch(`${API_URL}/auth/logout`, {
				credentials: 'include',
				method: 'POST'
			});

			if (res.ok) {
				refreshAll();
			}
		} catch {
			errorMessage = m['errors.http_error']();
		} finally {
			isSubmitting = false;
		}
	}
</script>

<header>
	<div class="flex items-center justify-between pb-4">
		<BrandIcon />
		<div class="flex space-x-2">
			<ServerSelector {servers} />
			<SearchBar {settings} />
			{#if settings?.multilingual}
				<DropdownMenu.Root>
					<DropdownMenu.Trigger>
						<Button variant="outline" size="icon" aria-label={m['language.change']()}>
							<Languages />
						</Button>
					</DropdownMenu.Trigger>
					<DropdownMenu.Content align="end">
						{#each locales as locale (locale)}
							<DropdownMenu.Item onclick={() => setLocale(locale)}>
								<span class={`fi fi-${locale === 'en' ? 'gb' : locale}`} aria-hidden="true"></span>
								{locale.toUpperCase()}{getLocale() === locale ? ' ✓' : ''}
							</DropdownMenu.Item>
						{/each}
					</DropdownMenu.Content>
				</DropdownMenu.Root>
			{/if}
			<ThemeSwitcher />
			{#if errorMessage}
				<p class="text-sm text-destructive" role="alert">{errorMessage}</p>
			{/if}

			{#if user?.email}
				<DropdownMenu.Root>
					<DropdownMenu.Trigger
						><Avatar.Root class="cursor-pointer">
							<Avatar.Image
								src="https://api.dicebear.com/10.x/planets/svg?seed={user.email}"
								alt="{user.email}'s Avatar"
							/>
							<Avatar.Fallback>{user.username}</Avatar.Fallback>
						</Avatar.Root></DropdownMenu.Trigger
					>
					<DropdownMenu.Content class="w-40" align="start">
						<DropdownMenu.Group>
							<DropdownMenu.Item>
								<a href={resolve('/(app)/settings')}>{m['pages.settings.title']()}</a>
							</DropdownMenu.Item>
							<DropdownMenu.Item>
								<button onclick={logout} disabled={isSubmitting}>
									{m['auth.logout']()}
								</button>
							</DropdownMenu.Item>
						</DropdownMenu.Group>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
			{:else}
				<Button href="/login">{m['auth.login']()}</Button>
			{/if}
		</div>
	</div>

	<Separator />
</header>
