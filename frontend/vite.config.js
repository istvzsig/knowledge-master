import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import dotenv, { process } from "dotenv";
import path from "path";

dotenv.config({ path: path.resolve(path.__filename, ".env") });

const BACKEND_URL = process.env.VITE_BACKEND_URL;
const BACKEND_PORT = process.env.VITE_BACKEND_PORT;

export default defineConfig({
  plugins: [react()],
  server: {
    port: process.env.VITE_FRONTEND_PORT,
    proxy: {
      "/faqs": {
        target: `${BACKEND_URL}:${BACKEND_PORT}`,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/faqs/, "/faqs"),
      },
    },
  },
});
