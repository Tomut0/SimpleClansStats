import { defineEnvVars } from '@sveltejs/kit/env';

export const variables = defineEnvVars({
	APP_NAME: {
		public: true,
		default: 'SCStats'
	},
	API_URL: {
		public: true,
		default: 'http://127.0.0.1:4010'
	},
	PAGE_SEPARATOR: {
		public: true,
		default: ' - '
	}
});
