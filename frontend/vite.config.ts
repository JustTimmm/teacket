import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'
import { tanstackRouter } from "@tanstack/router-plugin/vite";

// https://vite.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/rpc': {
        target: 'http://localhost:8080',
        rewrite: (path) => path.replace(/^\/rpc/, ''),
      }
    }
  },
  plugins: [
    tanstackRouter({ target: 'react', autoCodeSplitting: true }),
    react(),
  ],
})
