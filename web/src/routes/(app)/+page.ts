import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { resolve } from '$app/paths';

export const load: PageLoad = async ({ parent }) => {
	const { user, servers } = await parent();

	if (!user) {
		throw redirect(303, resolve('login'));
	}

	if (servers?.length == 0) {
		throw redirect(303, resolve('/(app)/settings'));
	}
};
