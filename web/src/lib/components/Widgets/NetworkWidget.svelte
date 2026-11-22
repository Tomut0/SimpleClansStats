<script lang="ts">
	import { m } from '#lib/paraglide/messages.js';
	let { clans }: { clans: Record<string, unknown>[] } = $props();
	const perRing = 12;
	const ringHeight = 220;
	const centerX = 400;
	const radiusX = 330;
	const radiusY = 82;
	const rings = $derived(Math.ceil(clans.length / perRing));
	const viewHeight = $derived(Math.max(ringHeight, rings * ringHeight));
	const nodes = $derived(
		clans.map((clan, index) => {
			const ringIndex = Math.floor(index / perRing);
			const indexInRing = index % perRing;
			const clansInRing = Math.min(perRing, clans.length - ringIndex * perRing);
			const angle = (2 * Math.PI * indexInRing) / clansInRing;
			const centerY = ringHeight / 2 + ringIndex * ringHeight;

			return {
				key: String(clan.externalId ?? index),
				label: String(clan.tag ?? clan.name ?? '—'),
				x: clans.length === 1 ? centerX : centerX + radiusX * Math.cos(angle),
				y: centerY + radiusY * Math.sin(angle),
				clan
			};
		})
	);
	const links = $derived(
		nodes.flatMap((node, index) => {
			const allies = Array.isArray(node.clan.allyExternalIds) ? node.clan.allyExternalIds : [];
			const rivals = Array.isArray(node.clan.rivalExternalIds) ? node.clan.rivalExternalIds : [];
			return nodes.slice(index + 1).flatMap((target) => {
				if (allies.includes(target.key)) return [{ from: node, to: target, color: '#22c55e' }];
				if (rivals.includes(target.key)) return [{ from: node, to: target, color: '#ef4444' }];
				return [];
			});
		})
	);
</script>

<section
	class="flex h-full min-h-0 flex-col overflow-hidden rounded-lg border bg-card text-card-foreground"
>
	<header class="widget-header px-4 py-3 text-sm font-medium text-muted-foreground">
		{m['pages.dashboard.widget_network']()}
	</header>
	{#if nodes.length}
		<div class="min-h-0 flex-1 overflow-auto">
			<svg
				class="block h-auto w-full"
				viewBox={`0 0 800 ${viewHeight}`}
				role="img"
				aria-label={m['pages.dashboard.widget_network']()}
			>
				{#each links as link (`${link.from.key}-${link.to.key}`)}<line
						x1={link.from.x}
						y1={link.from.y}
						x2={link.to.x}
						y2={link.to.y}
						stroke={link.color}
						stroke-width="2.5"
						stroke-opacity="0.7"
					/>{/each}
				{#each nodes as node (node.key)}<rect
						x={node.x - 44}
						y={node.y - 16}
						width="88"
						height="32"
						rx="16"
						class="fill-card stroke-border"
						stroke-width="1.5"
					/><text
						x={node.x}
						y={node.y + 4}
						text-anchor="middle"
						class="fill-foreground"
						font-size="12"
						font-weight="600"
						>{node.label.length > 12 ? `${node.label.slice(0, 11)}…` : node.label}</text
					>{/each}
			</svg>
		</div>
		<div class="flex justify-center gap-4 py-2 text-xs text-muted-foreground">
			<span class="flex items-center gap-1"
				><i class="size-2 rounded-full bg-green-500"></i>{m[
					'pages.dashboard.network_allies'
				]()}</span
			><span class="flex items-center gap-1"
				><i class="size-2 rounded-full bg-red-500"></i>{m['pages.dashboard.network_rivals']()}</span
			>
		</div>
	{:else}
		<div class="flex flex-1 items-center justify-center text-sm text-muted-foreground">
			{m['pages.dashboard.no_clans']()}
		</div>
	{/if}
</section>
