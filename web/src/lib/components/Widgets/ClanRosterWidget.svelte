<script lang="ts">
	import RowsWidget from './RowsWidget.svelte';
	let { players, title }: { players: Record<string, unknown>[]; title: string } = $props();
</script>
<RowsWidget {title} rows={players} fields={['name', 'clanTag', 'rank']} />
