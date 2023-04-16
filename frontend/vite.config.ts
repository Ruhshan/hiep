import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api':{
        target: 'http://0.0.0.0:8989',
        changeOrigin: true,
        secure: false
      }
    }
  }
})
