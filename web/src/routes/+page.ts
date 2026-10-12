import { API_URL } from '$app/env/public';
import { dev } from '$app/environment';

import { redirect, error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import type { User, GetServersResponse } from '$lib/api/generated';

type MockScenario = 'empty' | 'with_servers';

function mockHeaders(scenario: MockScenario): HeadersInit {
	// В продакшене заголовки не нужны — Prism не стоит перед продом
	if (!dev) return {};
	return { Prefer: `example=${scenario}` };
}

export const load: PageLoad = async ({ fetch }) => {
	const [userResponse, serversResponse] = await Promise.all([
		fetch(`${API_URL}/auth/me`, { credentials: 'include' }),
		fetch(`${API_URL}/servers`, {
			credentials: 'include',
			headers: mockHeaders('empty')
		})
	]);

	if (!serversResponse.ok) {
		throw error(serversResponse.status, 'Failed to load servers');
	}

	const servers: GetServersResponse = await serversResponse.json();

	if (userResponse.status === 401 && servers.length === 0) {
		throw redirect(303, '/login');
	}

	const user: User | undefined = userResponse.ok ? await userResponse.json() : undefined;

	return { user, servers };
};
