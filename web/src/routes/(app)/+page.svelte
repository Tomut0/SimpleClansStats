<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { APP_NAME, PAGE_SEPARATOR } from '$app/env/public';
	import { page } from '$app/state';
	import LayoutDashboardIcon from '@lucide/svelte/icons/layout-dashboard';
	import PageHeader from '$lib/components/PageHeader.svelte'
</script>

<svelte:head>
	<title>{m['pages.dashboard.title']()}{PAGE_SEPARATOR}{page.data.settings?.appName || APP_NAME}</title>
</svelte:head>

<main>
    <PageHeader
        icon={LayoutDashboardIcon}
        title={m['pages.dashboard.title']()}
        description={m['pages.dashboard.description']()}
   	>
        {#snippet actions()}
          <p>{m['server.online']()}</p>
        {/snippet}
    </PageHeader>
</main>
