import { defineConfig } from 'vite';
import dotenv from 'dotenv';
import path from 'path';
// import viteRestart from 'vite-plugin-restart';
import symfonyPlugin from "vite-plugin-symfony";
import tailwindcss from '@tailwindcss/vite'

dotenv.config()
const siteUrl = process.env.DEFAULT_URI;

export default defineConfig(({ command }) => ({
//  css:{postcss},
  plugins: [
    symfonyPlugin({
        stimulus: true,
    }),
    tailwindcss(),
    // viteRestart({
    //   reload: ['templates/**/*.twig', 'translations/**/*.yaml'],
    // }),
  ],

  resolve: {
    alias: {
      '@': path.resolve(__dirname, './assets'),
    },
  },

  base: command === 'serve' ? '' : '/build/',

  build: {
    target: 'esnext',
    outDir: './public/build',
    rollupOptions: {
      input: {
        app: './assets/app.js',
      },
      output: {
        assetFileNames: (assetInfo) => {
         // const extType = assetInfo.name.split('.').pop().match(/css/) ? 'css' : 'assets';
         // return `${extType}/[name].[hash][extname]`;
          // Place asset in each corresponding folder "build/{img,font,etc.}/*".
          const info = assetInfo.name.split('.');
          let extType = info[info.length - 1];
          if (/png|jpe?g|svg|gif|tiff|bmp|ico|avif|webp/i.test(extType)) {
              extType = 'img';
          } else if (/woff2?|otf|ttf|eot/.test(extType)) {
              extType = 'fonts';
          }
          return `${extType}/[name].[hash][extname]`;
        },
        chunkFileNames: 'js/[name].[hash].js',
        entryFileNames: 'js/[name].[hash].js',
      },
    },
    manifest: true,
  },

   server: {
    cors: true,
    hmr: { host: 'localhost' },
    host: '0.0.0.0',
    port: 5173,
    strictPort: true,
    watch: { interval: 100, usePolling: true },
  },
}));
