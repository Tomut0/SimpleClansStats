<script lang="ts">
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import SearchIcon from '@lucide/svelte/icons/search';
	import { m } from '$lib/paraglide/messages.js';
</script>

<InputGroup.Root class="w-64">
	<InputGroup.Input placeholder={m['searchBar.placeholder']()} />
	<InputGroup.Addon>
		<SearchIcon />
	</InputGroup.Addon>
</InputGroup.Root>
