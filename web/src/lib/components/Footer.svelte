<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { APP_NAME } from '$app/env/public';
	import { version } from '$app/env';
</script>

<footer class="flex items-center justify-between py-4">
	<p>{APP_NAME}ㆍ{version ?? 'unknown'}</p>
	<p class="text-right">
		{m['dashboard.footer.copyright']()}<br />
		{m['dashboard.footer.author']()}
	</p>
</footer>
