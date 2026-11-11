<script lang="ts">
	import type { CounterFormat } from './types';
	let {
		title,
		value,
		format = 'number'
	}: { title: string; value: number; format?: CounterFormat } = $props();
	const displayValue = $derived(
		format === 'compact'
			? value.toLocaleString(undefined, { notation: 'compact', maximumFractionDigits: 1 })
			: format === 'plain'
				? Math.trunc(value).toString()
				: value.toLocaleString()
	);
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">{title}</header>
	<div class="flex min-h-0 flex-1 items-center px-4 pb-4">
		<strong class="text-3xl">{displayValue}</strong>
	</div>
</section>
