<script lang="ts">
	import RowsWidget from './RowsWidget.svelte';
	let { wars, title }: { wars: Record<string, unknown>[]; title: string } = $props();
</script>
<RowsWidget {title} rows={wars} fields={['name', 'status', 'startedAt']} />
