import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// https://github.com/vitejs/vite/issues/1960
// https://github.com/anncwb/vite-plugin-compression
import viteCompression from 'vite-plugin-compression'

// https://vitejs.dev/config/
// a costom vite compression
export default defineConfig({
  plugins: [vue(), viteCompression()],
})
