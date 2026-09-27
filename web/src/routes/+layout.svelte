<script lang="ts">
	import type { Pathname } from '$app/types';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { locales, localizeHref } from '$lib/paraglide/runtime';
	import './layout.css';
	import { ModeWatcher } from 'mode-watcher';
	import { APP_NAME } from '$app/env/public';

	let { children } = $props();
</script>

<svelte:head>
	<title>
		{page.data.title ? `${page.data.title} - ${APP_NAME}` : APP_NAME}
	</title>
</svelte:head>

<ModeWatcher />

{@render children()}

<div style="display:none">
	{#each locales as locale (locale)}
		<a href={resolve(localizeHref(page.url.pathname, { locale }) as Pathname)}>{locale}</a>
	{/each}
</div>
