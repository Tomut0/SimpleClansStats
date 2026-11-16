<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { API_URL } from '$app/env/public';
	import { toast } from 'svelte-sonner';
	import * as Card from '$lib/components/ui/card/index.js';
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Spinner } from '$lib/components/ui/spinner/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { AppSettings, ThemePresetRequest } from '$lib/api/generated';

	const defaultPreset = 'b2pklPa1A';
	let {
		id,
		title,
		settings,
		disabled = true
	}: {
		id: string;
		title: string;
		settings: AppSettings | null;
		disabled?: boolean;
	} = $props();
	let preset = $derived(settings?.theme ?? defaultPreset);
	let applying = $state(false);

	async function applyTheme(code = preset) {
		if (disabled || applying) return;
		applying = true;
		try {
			const body: ThemePresetRequest = { preset: code };
			const response = await fetch(`${API_URL}/settings/theme`, {
				method: 'POST',
				credentials: 'include',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(body)
			});
			if (!response.ok) {
				throw new Error(`Theme request failed (${response.status})`);
			}
			toast.success(m['pages.settings.categories.theme.applied']());
			await new Promise((resolve) => setTimeout(resolve, 500));
			window.location.reload();
		} catch {
			applying = false;
			toast.error(m['pages.settings.categories.theme.apply_failed']());
		}
	}
</script>

<fieldset class="min-w-0" {disabled}>
	<Card.Root {id}>
		<Card.Header><Card.Title>{title}</Card.Title></Card.Header>
		<Separator />
		<Card.Content class="flex flex-col gap-4">
			<Field.Field>
				<Field.Content>
					<Field.Label for="theme-preset"
						>{m['pages.settings.categories.theme.preset']()}</Field.Label
					>
					<Field.Description>
						{m['pages.settings.categories.theme.description_before']()}
						<a
							href="https://shadcn-svelte.com/create/"
							target="_blank"
							rel="noopener noreferrer"
							class="underline underline-offset-4">shadcn-svelte</a
						>
						{m['pages.settings.categories.theme.description_after']()}
					</Field.Description>
				</Field.Content>
				<input
					id="theme-preset"
					class="h-9 rounded-md border border-input bg-background px-3 text-sm"
					bind:value={preset}
					maxlength="32"
					autocomplete="off"
				/>
			</Field.Field>
			<div class="flex flex-wrap gap-2">
				<Button type="button" disabled={disabled || applying} onclick={() => applyTheme()}>
					{m['pages.settings.categories.theme.apply']()}
				</Button>
				<Button
					type="button"
					variant="secondary"
					disabled={disabled || applying || settings?.theme === defaultPreset}
					onclick={() => {
						preset = defaultPreset;
						applyTheme(defaultPreset);
					}}
				>
					{m['pages.settings.categories.theme.reset']()}
				</Button>
			</div>
		</Card.Content>
	</Card.Root>
</fieldset>

<Dialog.Root open={applying}>
	<Dialog.Content
		class="justify-items-center gap-3 text-center"
		showCloseButton={false}
		escapeKeydownBehavior="ignore"
		interactOutsideBehavior="ignore"
	>
		<Spinner class="size-8 motion-reduce:animate-none" aria-hidden="true" />
		<Dialog.Title>{m['pages.settings.categories.theme.applying']()}</Dialog.Title>
		<Dialog.Description
			>{m['pages.settings.categories.theme.applying_description']()}</Dialog.Description
		>
	</Dialog.Content>
</Dialog.Root>
