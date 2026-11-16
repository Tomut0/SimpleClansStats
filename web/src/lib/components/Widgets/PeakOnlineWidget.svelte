<script lang="ts">
	import TimeSeriesWidget from './TimeSeriesWidget.svelte';
	type Point = {
		at: string;
		clans: number;
		players: number;
		wars: number;
		balance: number;
		onlinePlayers: number;
		kills: number;
	};
	let { history, title }: { history: Point[]; title: string } = $props();
</script>

<TimeSeriesWidget {history} {title} metric="onlinePlayers" />
