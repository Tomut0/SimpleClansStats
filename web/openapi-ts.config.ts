import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
	input: '../shared/openapi.yaml',
	output: 'src/lib/api/generated',
	plugins: ['@hey-api/typescript']
});
