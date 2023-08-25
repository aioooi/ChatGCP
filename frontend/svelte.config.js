import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';

const dev = process.argv.includes('dev');

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			// // using "index.html" as a "fallback" here as in
			// // https://kit.svelte.dev/docs/single-page-apps turns it into an SPA:
			// fallback: 'index.html',
			precompress: true,
			strict: true
		}),
		paths: {
			base: dev ? '' : '/ChatGCP'
		}
	}
};

export default config;
