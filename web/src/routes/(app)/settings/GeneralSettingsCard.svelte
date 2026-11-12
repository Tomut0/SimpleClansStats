<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import LogoPreview from '$lib/components/LogoPreview.svelte';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Switch } from '$lib/components/ui/switch/index.js';
	import { APP_NAME } from '$app/env/public';
	import { asset } from '$app/paths';
	import type { AppSettings } from '$lib/api/generated';

	let {
		id,
		title,
		settings = null,
		logoUrl = asset('/favicon.svg'),
		disabled = true
	}: {
		id: string;
		title: string;
		settings?: AppSettings | null;
		logoUrl?: string;
		disabled?: boolean;
	} = $props();
</script>

<fieldset class="min-w-0" {disabled}>
	{#key settings}
		<Card.Root {id}>
			<Card.Header>
				<Card.Title>{title}</Card.Title>
			</Card.Header>
			<Separator />
			<Card.Content class="flex flex-col gap-4">
				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="logo-upload">
							{m['pages.settings.categories.general.logo_title']()}
						</Field.Label>
						<Field.Description>
							{m['pages.settings.categories.general.logo_desc']()}
						</Field.Description>
					</Field.Content>

					<LogoPreview {logoUrl} />
				</Field.Field>

				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="app-name">
							{m['pages.settings.categories.general.name_title']()}
						</Field.Label>
						<Field.Description>
							{m['pages.settings.categories.general.name_desc']()}
						</Field.Description>
					</Field.Content>

					<Input
						id="app-name"
						name="appName"
						value={settings?.appName || APP_NAME}
						required
						maxlength={100}
						class="max-w-xs"
					/>
				</Field.Field>

				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="player-avatar-url"
							>{m['pages.settings.categories.general.player_avatar_url_title']()}</Field.Label
						>
						<Field.Description
							>{m['pages.settings.categories.general.player_avatar_url_desc']()}</Field.Description
						>
					</Field.Content>
					<Input
						id="player-avatar-url"
						name="playerAvatarUrl"
						value={settings?.playerAvatarUrl ?? 'https://mc-heads.net/avatar/{username}/32'}
						required
						maxlength={500}
						class="max-w-lg"
					/>
				</Field.Field>

				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="remove-author">
							{m['pages.settings.categories.general.remove_author_title']()}
						</Field.Label>
						<Field.Description>
							{m['pages.settings.categories.general.remove_author_description']()}
						</Field.Description>
					</Field.Content>

					<Switch
						id="remove-author"
						name="removeAuthor"
						checked={settings?.removeAuthor ?? false}
					/>
				</Field.Field>

				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="remove-telemetry">
							{m['pages.settings.categories.general.remove_telemetry_title']()}
						</Field.Label>
						<Field.Description>
							{m['pages.settings.categories.general.remove_telemetry_description']()}
						</Field.Description>
					</Field.Content>
					<Switch
						id="remove-telemetry"
						name="removeTelemetry"
						checked={settings?.removeTelemetry ?? false}
					/>
				</Field.Field>
			</Card.Content>
		</Card.Root>
	{/key}
</fieldset>
