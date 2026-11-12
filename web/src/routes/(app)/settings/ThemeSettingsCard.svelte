<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import { API_URL } from '$app/env/public';
	import { invalidate } from '$app/navigation';
	import { toast } from 'svelte-sonner';
	import * as Card from '$lib/components/ui/card/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { AppSettings, ThemePresetRequest } from '$lib/api/generated';

	const defaultPreset = 'b2pklPa1A';
	let { id, title, settings, disabled = true }: {
		id: string;
		title: string;
		settings: AppSettings | null;
		disabled?: boolean;
	} = $props();
	let preset = $state(defaultPreset);
	let applying = $state(false);

	$effect(() => {
		preset = settings?.theme ?? defaultPreset;
	});

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
				toast.error(m['pages.settings.categories.theme.apply_failed']());
				return;
			}
			await invalidate('app:settings');
			toast.success(m['pages.settings.categories.theme.applied']());
		} catch {
			toast.error(m['pages.settings.categories.theme.apply_failed']());
		} finally {
			applying = false;
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
					<Field.Label for="theme-preset">{m['pages.settings.categories.theme.preset']()}</Field.Label>
					<Field.Description>
						{m['pages.settings.categories.theme.description_before']()}
						<a
							href="https://shadcn-svelte.com/create/"
							target="_blank"
							rel="noopener noreferrer"
							class="underline underline-offset-4"
						>shadcn-svelte</a
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
					{applying ? m['pages.settings.categories.theme.applying']() : m['pages.settings.categories.theme.apply']()}
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
