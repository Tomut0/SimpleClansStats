<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import type { PageProps } from './$types';
	import { APP_NAME, PAGE_SEPARATOR } from '$app/env/public';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>{m['dashboard.title']()}{PAGE_SEPARATOR}{APP_NAME}</title>
</svelte:head>

<div class="px-16 py-12">
	<Header servers={data.servers} />

	<main>
		<section class="flex items-center justify-between">
			<div class="flex items-center gap-2 py-8">
				<LayoutDashboardIcon class="h-10 w-10 text-foreground" />
				<div class="flex flex-col">
					<h1 class="text-xl font-medium text-foreground">{m['dashboard.title']()}</h1>
					<p class="text-muted-foreground">{m['dashboard.description']()}</p>
				</div>
			</div>
			<div>
				<div class=""></div>
				<p>{m['server.online']()}</p>
			</div>
		</section>
	</main>

	<Footer />
</div>
