<script lang="ts">
	import RowsWidget from './RowsWidget.svelte';
	import { m } from '$lib/paraglide/messages.js';
	let { changes }: { changes: Record<string, unknown>[] } = $props();
</script>
<RowsWidget title={m['pages.dashboard.gainers_losers_type']()} rows={changes} fields={['tag', 'balanceChange']} />
