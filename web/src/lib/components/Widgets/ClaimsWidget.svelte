<script lang="ts">
	import RowsWidget from './RowsWidget.svelte';
	let { claims, title }: { claims: Record<string, unknown>[]; title: string } = $props();
</script>
<RowsWidget {title} rows={claims} fields={['name', 'owner', 'world', 'size']} />
