import viteConfig from "./vite.config";
import { defineConfig, mergeConfig } from "vitest/config";

export default defineConfig(
  mergeConfig(viteConfig, {
    test: {
      environment: "jsdom",
      setupFiles: ["./src/tests/setup.ts"],
      include: ["./src/**/*.{test,spec}.{js,mjs,cjs,ts,mts,cts,jsx,tsx}"],
      globals: true,
    },
  })
);
