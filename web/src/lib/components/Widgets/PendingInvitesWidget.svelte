<script lang="ts">
	import RowsWidget from './RowsWidget.svelte';
	let { invites, title }: { invites: Record<string, unknown>[]; title: string } = $props();
</script>
<RowsWidget {title} rows={invites} fields={['playerName', 'clanTag', 'createdAt']} />
