import {defineConfig} from 'vite';
import artefak from './resources/js/Plugins/vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
    plugins: [
        artefak({
            input: './resources/js/app.js',
            refresh: true,
        }),
        vue({
            template: {
                transformAssetUrls: {
                    base: null,
                    includeAbsolute: false,
                },
            },
            reactivityTransform: true,
        }),
    ],
});
