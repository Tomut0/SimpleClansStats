<script lang="ts">
	import { API_URL } from '$app/env/public';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { FieldGroup, Field, FieldLabel } from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { m } from '$lib/paraglide/messages.js';

	const id = $props.id();
	let errorMessage = $state('');
	let isSubmitting = $state(false);

	async function submit(event: SubmitEvent) {
		event.preventDefault();

		const form = event.currentTarget as HTMLFormElement;
		const formData = new FormData(form);
		const email = String(formData.get('email') ?? '');
		const password = String(formData.get('password') ?? '');

		errorMessage = '';
		isSubmitting = true;

		try {
			const response = await fetch(`${API_URL}/auth/login`, {
				method: 'POST',
				credentials: 'include',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ email, password })
			});

			if (!response.ok) {
				errorMessage = response.status === 401 ? 'Invalid email or password' : 'Login failed';
				return;
			}

			window.location.assign('/');
		} catch {
			errorMessage = m['errors.http_error']();
		} finally {
			isSubmitting = false;
		}
	}
</script>

<Card.Root class="mx-auto w-full max-w-sm">
	<Card.Header>
		<Card.Title class="text-2xl">{m['auth.login']()}</Card.Title>
		<Card.Description>{m['pages.login.form.description']()}</Card.Description>
	</Card.Header>
	<Card.Content>
		<form onsubmit={submit}>
			<FieldGroup>
				<Field>
					<FieldLabel for="email-{id}">{m['pages.login.form.email']()}</FieldLabel>
					<Input
						id="email-{id}"
						name="email"
						type="email"
						placeholder={m['pages.login.form.emailPlaceholder']()}
						required
					/>
				</Field>
				<Field>
					<FieldLabel for="password-{id}">{m['pages.login.form.password']()}</FieldLabel>
					<Input
						id="password-{id}"
						name="password"
						type="password"
						autocomplete="current-password"
						required
					/>
				</Field>
				<Field>
					<Button type="submit" class="w-full" disabled={isSubmitting}>
						{isSubmitting ? 'Signing in…' : m['auth.login']()}
					</Button>
					{#if errorMessage}
						<p class="text-sm text-destructive" role="alert">{errorMessage}</p>
					{/if}
				</Field>
			</FieldGroup>
		</form>
	</Card.Content>
</Card.Root>
