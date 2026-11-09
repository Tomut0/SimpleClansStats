<script lang="ts">
    import { Toaster } from "$lib/components/ui/sonner/index.js";
    import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	import type { LayoutProps } from './$types';

	let { data, children }: LayoutProps = $props();
</script>

<Toaster />

<div class="px-16 py-12">
	<Header servers={data.servers} user={data.user} />
	{@render children()}
	<Footer />
</div>
