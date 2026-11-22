import prettier from 'eslint-config-prettier';
import path from 'node:path';
import js from '@eslint/js';
import svelte from 'eslint-plugin-svelte';
import { defineConfig, globalIgnores, includeIgnoreFile } from 'eslint/config';
import globals from 'globals';
import ts from 'typescript-eslint';

const gitignorePath = path.resolve(import.meta.dirname, '.gitignore');

export default defineConfig(
	includeIgnoreFile(gitignorePath),
	globalIgnores(['src/lib/api/generated/']),
	js.configs.recommended,
	ts.configs.strict,
	{
		files: ['src/**/*.ts', 'src/**/*.svelte'],
		ignores: ['src/lib/components/ui/**'],
		languageOptions: {
			parserOptions: { projectService: true }
		},
		rules: {
			'@typescript-eslint/await-thenable': 'error',
			'@typescript-eslint/no-deprecated': 'error',
			'@typescript-eslint/no-floating-promises': 'error',
			'@typescript-eslint/no-misused-promises': 'error',
			'@typescript-eslint/only-throw-error': 'error',
			complexity: ['error', { max: 16, variant: 'modified' }],
			'max-depth': ['error', 3],
			eqeqeq: ['error', 'always', { null: 'ignore' }]
		}
	},
	{
		files: ['src/**/*.ts'],
		ignores: ['src/lib/components/ui/**'],
		rules: {
			'@typescript-eslint/no-unsafe-assignment': 'error',
			'@typescript-eslint/no-unsafe-return': 'error'
		}
	},
	svelte.configs.recommended,
	prettier,
	svelte.configs.prettier,
	{
		languageOptions: {
			globals: { ...globals.browser, ...globals.node },
			parserOptions: {
				extraFileExtensions: ['.svelte'],
				tsconfigRootDir: import.meta.dirname
			}
		},
		linterOptions: { reportUnusedDisableDirectives: 'error' }
	},
	{
		files: ['**/*.ts', '**/*.svelte'],
		rules: {
			// TypeScript checks for undefined names.
			'no-undef': 'off'
		}
	},
	{
		files: ['**/*.svelte', '**/*.svelte.ts', '**/*.svelte.js'],
		languageOptions: {
			parserOptions: {
				parser: ts.parser
			}
		}
	}
);
