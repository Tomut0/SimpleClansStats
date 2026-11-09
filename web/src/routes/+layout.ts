import { API_URL } from '$app/env/public';
import { asset } from '$app/paths';
import type { AppSettings } from '$lib/api/generated';
import type { LayoutLoad } from './$types';

export const ssr = false;

export const load: LayoutLoad = async ({ fetch, depends }) => {
  depends('app:session');
  depends('app:settings');
  
  const [meRes, serversRes, settings] = await Promise.all([
    fetch(`${API_URL}/auth/me`, { credentials: 'include' }),
    fetch(`${API_URL}/servers`),
    fetch(`${API_URL}/settings`)
      .then(async (response): Promise<AppSettings | null> =>
        response.ok ? await response.json() : null
      )
      .catch(() => null)
  ]);

  const user = meRes.ok ? await meRes.json() : null;
  const servers = serversRes.ok ? await serversRes.json() : [];
  
  const logoUrl = settings?.logoVersion
    ? `${API_URL}/settings/logo?v=${settings.logoVersion}`
    : asset('/favicon.svg');

  return { user, servers, settings, logoUrl };
};
