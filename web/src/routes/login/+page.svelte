<script>
	import BrandIcon from '$lib/components/BrandIcon.svelte';
	import LoginForm from '$lib/components/LoginForm.svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { resolve } from '$app/paths';
</script>

<div class="flex h-screen w-full flex-col items-center justify-center gap-8 px-4">
	<BrandIcon />
	<LoginForm />
	<a href={resolve('')}>{m['pages.login.backToHome']()}</a>
</div>
