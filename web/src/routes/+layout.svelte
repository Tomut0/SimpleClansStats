<script lang="ts">
	import type { Pathname } from '$app/types';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { locales, localizeHref } from '$lib/paraglide/runtime';
	import './layout.css';
	import config from '$lib/config';
	import { ModeWatcher } from 'mode-watcher';

	let { children } = $props();
</script>

<svelte:head>
	<title>
		{page.data.title ? `${page.data.title} - ${config.appName}` : config.appName}
	</title>

	<link rel="icon" href="legacy-icon.ico" sizes="any" />
	<link rel="icon" type="image/svg+xml" href="modern-icon.svg" />
	<link rel="apple-touch-icon" href="apple-touch-icon.png" />
</svelte:head>

<ModeWatcher />

{@render children()}

<div style="display:none">
	{#each locales as locale (locale)}
		<a href={resolve(localizeHref(page.url.pathname, { locale }) as Pathname)}>{locale}</a>
	{/each}
</div>
