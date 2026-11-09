<script lang="ts">
	import { onDestroy } from 'svelte';
	import { m } from '$lib/paraglide/messages.js';
	import { Input } from '$lib/components/ui/input/index.js';

	let { logoUrl }: { logoUrl: string } = $props();
	let previewUrl = $state<string>();

	function selectLogo(event: Event) {
		const input = event.currentTarget as HTMLInputElement;
		const file = input.files?.[0];
		if (previewUrl) URL.revokeObjectURL(previewUrl);
		previewUrl = file && ['image/png', 'image/jpeg'].includes(file.type)
			? URL.createObjectURL(file)
			: undefined;
	}

	onDestroy(() => {
		if (previewUrl) URL.revokeObjectURL(previewUrl);
	});
</script>

<div
	class="relative size-18 shrink-0 rounded-xl has-focus-visible:ring-2 has-focus-visible:ring-ring"
>
	<img
		src={previewUrl ?? logoUrl}
		alt={m['pages.settings.categories.general.logo_title']()}
		class="size-full rounded-xl bg-muted object-contain p-2"
	/>
	<Input
		id="logo-upload"
		name="logo"
		type="file"
		accept="image/png,image/jpeg"
		aria-label={m['pages.settings.categories.general.logo_title']()}
		class="absolute inset-0 size-full cursor-pointer opacity-0"
		onchange={selectLogo}
	/>
</div>
