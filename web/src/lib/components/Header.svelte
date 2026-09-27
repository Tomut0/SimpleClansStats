<script lang="ts">
	import BrandIcon from '$lib/components/BrandIcon.svelte';
	import ThemeSwitcher from '$lib/components/ThemeSwitcher.svelte';
	import Button from './ui/button/button.svelte';
	import { m } from '$lib/paraglide/messages.js';
	import SearchBar from './SearchBar.svelte';
	import ServerSelector from './ServerSelector.svelte';
	import Separator from './ui/separator/separator.svelte';
	import type { Server } from '$lib/api/generated';

	let { servers = [] }: { servers: Server[] } = $props();
</script>

<header>
	<div class="flex items-center justify-between pb-4">
		<BrandIcon />
		<div class="flex space-x-2">
			<ServerSelector {servers} />
			<SearchBar />
			<ThemeSwitcher />
			<Button href="/login">{m['auth.login']()}</Button>
		</div>
	</div>

	<Separator />
</header>
