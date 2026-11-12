<script lang="ts">
	import LatestActionsWidget from './LatestActionsWidget.svelte';
	import { m } from '$lib/paraglide/messages.js';
	let { events }: { events: Record<string, unknown>[] } = $props();
	const kills = $derived(events.filter((event) => /kill|death/i.test(String(event.type ?? ''))));
</script>

<LatestActionsWidget events={kills} title={m['pages.dashboard.kill_feed_type']()} />
