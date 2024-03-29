import adapter from '@sveltejs/adapter-auto';
import preprocessor from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: preprocessor({
		scss: {
			prependData: `@import './src/static/styles.scss';`,
		}
	}),
	kit: {
		adapter: adapter(),
		vite: {
			// resolve: {
			//   alias: {
			// 	src: path.resolve('./src'),
			//   },
			// },
			server: {
			  fs: {
				allow: ['./static'],
			  },
			},
		  },
	}
};

export default config;


// // import autoAdapter from '@sveltejs/adapter-auto';
// import nodeAdapter from '@sveltejs/adapter-node';
// import preprocess from 'svelte-preprocess';
// import importAssets from 'svelte-preprocess-import-assets';
// import seqPreprocessor from 'svelte-sequential-preprocessor';

// import path from 'path';

// /** @type {import('@sveltejs/kit').Config} */
// const config = {
// 	compilerOptions: {
// 	  enableSourcemap: true,
// 	},
  
// 	// Consult https://github.com/sveltejs/svelte-preprocess
// 	// for more information about preprocessors
// 	preprocess: [seqPreprocessor([preprocess({ sourceMap: true }), importAssets()])],
  
// 	kit: {
// 	  // adapter: autoAdapter({ out: 'build' }),
// 	  adapter: nodeAdapter({ out: 'build' }),
  
// 	  // hydrate the <div id="svelte"> element in src/app.html
// 	  target: '#svelte',
// 	  // Absolute Imports
// 	  vite: {
// 		resolve: {
// 		  alias: {
// 			src: path.resolve('./src'),
// 		  },
// 		},
// 		server: {
// 		  fs: {
// 			allow: ['./assets'],
// 		  },
// 		},
// 	  },
// 	},
//   };
  
//   export default config;