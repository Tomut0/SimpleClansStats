<script lang="ts">
	import { API_URL } from '$app/env/public';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import {
		FieldGroup,
		Field,
		FieldLabel,
		FieldDescription
	} from '$lib/components/ui/field/index.js';
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
				body: JSON.stringify({ email: email, password })
			});

			if (!response.ok) {
				errorMessage = response.status === 401 ? 'Invalid email or password' : 'Login failed';
				return;
			}

			window.location.assign('/');
		} catch {
			errorMessage = 'Could not reach the server. Please try again.';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<Card.Root class="mx-auto w-full max-w-sm">
	<Card.Header>
		<Card.Title class="text-2xl">{m['auth.login']()}</Card.Title>
		<Card.Description>{m['auth.loginForm.description']()}</Card.Description>
	</Card.Header>
	<Card.Content>
		<form onsubmit={submit}>
			<FieldGroup>
				<Field>
					<FieldLabel for="email-{id}">{m['auth.loginForm.email']()}</FieldLabel>
					<Input
						id="email-{id}"
						name="email"
						type="email"
						placeholder={m['auth.loginForm.emailPlaceholder']()}
						required
					/>
				</Field>
				<Field>
					<div class="flex items-center">
						<FieldLabel for="password-{id}">{m['auth.loginForm.password']()}</FieldLabel>
						<a href="##" class="ms-auto inline-block text-sm underline">
							{m['auth.loginForm.forgotPassword']()}
						</a>
					</div>
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
					<Button variant="outline" class="w-full">
						<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
							<path
								d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
								fill="currentColor"
							/>
						</svg>
						{m['auth.loginForm.loginWithGoogle']()}
					</Button>
					<FieldDescription class="text-center">
						{m['auth.loginForm.noAccount']()} <a href="##">{m['auth.loginForm.signUp']()}</a>
					</FieldDescription>
				</Field>
			</FieldGroup>
		</form>
	</Card.Content>
</Card.Root>
