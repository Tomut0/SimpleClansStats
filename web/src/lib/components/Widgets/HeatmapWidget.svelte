<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { Tooltip } from 'bits-ui';
	let { activity }: { activity: unknown[] } = $props();
	const values = $derived(Array.from({ length: 168 }, (_, i) => Number(activity[i] ?? 0)));
	const peak = $derived(Math.max(1, ...values));
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">
		{m['pages.dashboard.heatmap_type']()}
	</header>
	{#if activity.length}<div
			class="grid flex-1 grid-flow-col grid-rows-7 gap-1 p-4"
			role="img"
			aria-label={m['pages.dashboard.heatmap_type']()}
		>
			{#each values as value, index (index)}<Tooltip.Root
					><Tooltip.Trigger
						>{#snippet child({ props })}<span
								{...props}
								class="block min-h-2 min-w-2 rounded-sm bg-primary"
								style={`opacity: ${0.12 + (0.88 * value) / peak}`}
							></span>{/snippet}</Tooltip.Trigger
					><Tooltip.Content
						class="z-50 rounded-md bg-foreground px-3 py-1.5 text-xs text-background shadow-md"
						>{index}: {value}</Tooltip.Content
					></Tooltip.Root
				>{/each}
		</div>{:else}<div class="flex flex-1 items-center justify-center text-sm text-muted-foreground">
			{m['pages.dashboard.no_widget_data']()}
		</div>{/if}
</section>
