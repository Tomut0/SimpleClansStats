<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { APP_NAME } from '$app/env/public';
	import { page } from '$app/state';
	import { version } from '$app/env';
</script>

<footer class="flex items-center justify-between py-4">
	<p>{page.data.settings?.appName || APP_NAME}ㆍ{version ?? 'unknown'}</p>
	<p class="text-right">
		{m['footer.copyright']()}
		{#if !page.data.settings?.removeAuthor}
			<br />{m['footer.author']()}
		{/if}
	</p>
</footer>
