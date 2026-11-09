<script lang="ts">
	import { m } from '$lib/paraglide/messages.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { tick, untrack } from 'svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Field from '$lib/components/ui/field/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import * as Select from '$lib/components/ui/select/index.js';
	import { CalendarDays, Clock3, Globe2 } from '@lucide/svelte';
	import type { AppSettings } from '$lib/api/generated';

	let { id, title, settings = null, disabled = true }: {
		id: string;
		title: string;
		settings?: AppSettings | null;
		disabled?: boolean;
	} = $props();

	const timeZones = ['UTC', ...Intl.supportedValuesOf('timeZone')].sort();
	const dateFormats = ['DD.MM.YYYY', 'YYYY-MM-DD', 'MM/DD/YYYY'].map((value) => ({
		value,
		label: value
	}));
	const timeFormats = ['24h', '12h'].map((value) => ({ value, label: value }));
	let loadedSettings = untrack(() => settings);
	let timeZone = $state(loadedSettings?.timeZone ?? 'Europe/Moscow');
	let dateFormat = $state<string>(loadedSettings?.dateFormat ?? 'DD.MM.YYYY');
	let timeFormat = $state<string>(loadedSettings?.timeFormat ?? '24h');
	let timezoneOpen = $state(false);
	let timezoneTrigger = $state<HTMLButtonElement | null>(null);

	$effect.pre(() => {
		if (loadedSettings === settings) return;
		loadedSettings = settings;
		timeZone = settings?.timeZone ?? 'Europe/Moscow';
		dateFormat = settings?.dateFormat ?? 'DD.MM.YYYY';
		timeFormat = settings?.timeFormat ?? '24h';
		timezoneOpen = false;
	});

	function selectTimeZone(value: string) {
		timeZone = value;
		timezoneOpen = false;
		tick().then(() => timezoneTrigger?.focus());
	}
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
						<Field.Label for="time-zone"
							>{m['pages.settings.categories.timezone.timezone_title']()}</Field.Label
						>
						<Field.Description>
							{m['pages.settings.categories.timezone.timezone_desc']()}
						</Field.Description>
					</Field.Content>
					<input type="hidden" name="timeZone" value={timeZone} />
					<Popover.Root bind:open={timezoneOpen}>
						<Popover.Trigger bind:ref={timezoneTrigger}>
							{#snippet child({ props })}
								<Button
									{...props}
									id="time-zone"
									variant="ghost"
									role="combobox"
									aria-expanded={timezoneOpen}
									aria-label={m['pages.settings.categories.timezone.timezone_title']()}
									class="w-full justify-start bg-input/50 font-normal text-muted-foreground hover:bg-input/50 aria-expanded:bg-input/50 sm:w-64"
								>
									<Globe2 />
									{timeZone}
								</Button>
							{/snippet}
						</Popover.Trigger>
						<Popover.Content class="w-64 p-0">
							<Command.Root>
								<Command.Input placeholder={m['pages.settings.categories.timezone.search']()} />
								<Command.List>
									<Command.Empty>
										{m['pages.settings.categories.timezone.no_results']()}
									</Command.Empty>
									<Command.Group>
										{#each timeZones as zone (zone)}
											<Command.Item
												value={zone}
												data-checked={timeZone === zone}
												onSelect={() => selectTimeZone(zone)}
											>
												{zone}
											</Command.Item>
										{/each}
									</Command.Group>
								</Command.List>
							</Command.Root>
						</Popover.Content>
					</Popover.Root>
				</Field.Field>

				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="date-format"
							>{m['pages.settings.categories.timezone.date_format_title']()}</Field.Label
						>
						<Field.Description
							>{m['pages.settings.categories.timezone.date_format_desc']()}</Field.Description
						>
					</Field.Content>
					<Select.Root type="single" name="dateFormat" items={dateFormats} bind:value={dateFormat}>
						<Select.Trigger id="date-format" class="w-full justify-start font-normal sm:w-64">
							<CalendarDays />
							<Select.Value />
						</Select.Trigger>
						<Select.Content>
							{#each dateFormats as format (format.value)}
								<Select.Item value={format.value} label={format.label}>
									{format.label}
								</Select.Item>
							{/each}
						</Select.Content>
					</Select.Root>
				</Field.Field>

				<Field.Field orientation="horizontal">
					<Field.Content>
						<Field.Label for="time-format"
							>{m['pages.settings.categories.timezone.time_format_title']()}</Field.Label
						>
						<Field.Description
							>{m['pages.settings.categories.timezone.time_format_desc']()}</Field.Description
						>
					</Field.Content>
					<Select.Root type="single" name="timeFormat" items={timeFormats} bind:value={timeFormat}>
						<Select.Trigger id="time-format" class="w-full justify-start font-normal sm:w-64">
							<Clock3 />
							<Select.Value />
						</Select.Trigger>
						<Select.Content>
							{#each timeFormats as format (format.value)}
								<Select.Item value={format.value} label={format.label}>
									{format.label}
								</Select.Item>
							{/each}
						</Select.Content>
					</Select.Root>
				</Field.Field>
			</Card.Content>
		</Card.Root>
	{/key}
</fieldset>
