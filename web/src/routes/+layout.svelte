<script lang="ts">
	import type { Path } from '$app/types';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { locales, localizeHref } from '$lib/paraglide/runtime';
	import './layout.css';
	import 'flag-icons/css/flag-icons.min.css';
	import { ModeWatcher } from 'mode-watcher';
	import { API_URL, APP_NAME } from '$app/env/public';
	import { Tooltip } from 'bits-ui';

	let { children } = $props();
	const appName = $derived(page.data.settings?.appName || APP_NAME);
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href={`${API_URL}/settings/theme.css?v=${page.data.settings?.theme ?? ''}`}
	/>
	<title>
		{page.data.title ? `${page.data.title} - ${appName}` : appName}
	</title>
</svelte:head>

<ModeWatcher />

<Tooltip.Provider>
	{@render children()}
</Tooltip.Provider>

<div style="display:none">
	{#each locales as locale (locale)}
		<a href={resolve(localizeHref(page.url.pathname, { locale }) as Path)}>{locale}</a>
	{/each}
</div>
