<script lang="ts">
  import type { Component, Snippet } from 'svelte';
  import type { LucideProps } from '@lucide/svelte';

  let {
    icon: Icon,
    title,
    description,
    children,
    actions
  }: {
    icon?: Component<LucideProps>;
    title?: string;
    description?: string;
    children?: Snippet;
    actions?: Snippet;
  } = $props();
</script>

<section class="flex items-center justify-between">
  <div class="flex items-center gap-2 pt-8 pb-4">
    {#if Icon}
      <Icon class="h-10 w-10 text-foreground" />
    {/if}

    <div class="flex flex-col">
      <h1 class="text-xl font-medium text-foreground">{title}</h1>
      {#if description}
        <p class="text-muted-foreground">{description}</p>
      {/if}
    </div>

    {@render children?.()}
  </div>

  <div>
    {@render actions?.()}
  </div>
</section>