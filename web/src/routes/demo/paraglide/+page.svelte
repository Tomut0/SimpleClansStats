<script lang="ts">
	import { setLocale } from '$lib/paraglide/runtime';
</script>

<h1>Paraglide language demo</h1>

<div>
	<button onclick={() => setLocale('en')}>en</button>
	<button onclick={() => setLocale('ru')}>ru</button>
</div>

<p>
	If you use VSCode, install the

	<a
		href="https://marketplace.visualstudio.com/items?itemName=inlang.vs-code-extension"
		target="_blank">Sherlock i18n extension</a
	>

	for a better i18n experience.
</p>
