import { version } from '$app/environment';
    
interface Config {
	appName: string;
  pageSeparator: string;
	appVersion: string;
}

const config: Config = {
	appName: import.meta.env.VITE_APP_NAME ?? 'SCStats',
	pageSeparator: import.meta.env.VITE_PAGE_SEPARATOR ?? ' - ',
	appVersion: version ?? 'unknown',
};

export default config;
