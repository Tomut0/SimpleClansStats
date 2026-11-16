import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { resolve } from '$app/paths';

export const load: PageLoad = async ({ parent }) => {
	const { servers } = await parent();

	if (servers?.length == 0) {
		throw redirect(303, resolve('/(app)/settings'));
	}
};
