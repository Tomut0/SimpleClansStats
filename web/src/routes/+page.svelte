<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import config from '$lib/config';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<svelte:head>
	<title>{m['dashboard.title']()}{config.pageSeparator}{config.appName}</title>
</svelte:head>

<Header />

<main>
    
</main>

<Footer />