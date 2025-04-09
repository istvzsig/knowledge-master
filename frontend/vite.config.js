import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

const BACKEND_URL = process.env.BACKEND_URL || "http://127.0.0.1";
const BACKEND_PORT = process.env.BACKEND_PORT || 5555;
const FRONTEND_PORT = process.env.FRONTEND_PORT || 5173;

export default defineConfig({
  plugins: [react()],
  server: {
    port: FRONTEND_PORT,
    proxy: {
      "/faqs": {
        target: `${BACKEND_URL}:${BACKEND_PORT}`,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/faqs/, "/faqs"),
      },
    },
  },
});
