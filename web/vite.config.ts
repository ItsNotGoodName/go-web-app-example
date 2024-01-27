import { defineConfig } from "vite"
import FullReload from 'vite-plugin-full-reload'

export default defineConfig({
  plugins: [
    FullReload("reload-vite.local")
  ],
  build: {
    // generate manifest.json in outDir
    manifest: true,
    rollupOptions: {
      // overwrite default .html entry
      input: 'src/main.ts',
    },
  },
  server: {
    // ./web_dev.go assumes the 5173 port
    port: 5173,
    strictPort: true,
    // Disable HMR to prevent page reload when templ files change
    hmr: false,
  }
})
