<script lang="ts">
	import { APP_NAME } from '$app/env/public';
	import { page } from '$app/state';
</script>

<div class="flex items-center gap-4">
	<img src={page.data.logoUrl} alt="" class="size-12 object-contain" />
	<span class="scroll-m-20 text-3xl font-bold tracking-tight"
		>{page.data.settings?.appName || APP_NAME}</span
	>
</div>
