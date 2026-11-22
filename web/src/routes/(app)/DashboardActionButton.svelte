<script lang="ts">
	import { Tooltip } from 'bits-ui';
	import type { LucideProps } from '@lucide/svelte';
	import type { Component } from 'svelte';

	let {
		label,
		icon: Icon,
		onclick,
		primary = false
	}: {
		label: string;
		icon: Component<LucideProps>;
		onclick: () => void;
		primary?: boolean;
	} = $props();
</script>

<Tooltip.Root>
	<Tooltip.Trigger>
		{#snippet child({ props })}
			<button
				{...props}
				type="button"
				class={[
					'grid size-9 place-items-center rounded-md',
					primary ? 'bg-primary text-primary-foreground' : 'border text-foreground'
				]}
				aria-label={label}
				{onclick}
			>
				<Icon class="size-4" aria-hidden="true" />
			</button>
		{/snippet}
	</Tooltip.Trigger>
	<Tooltip.Content
		class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md"
	>
		{label}
	</Tooltip.Content>
</Tooltip.Root>
