import { defineConfig } from "vite";
import { visualizer } from "rollup-plugin-visualizer";
import react from "@vitejs/plugin-react-swc";
import path from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    visualizer({
      filename: "./stats/treemap.html",
      template: "treemap",
    }),
    visualizer({
      filename: "./stats/flamegraph.html",
      template: "flamegraph",
    }),
    visualizer({
      filename: "./stats/network.html",
      template: "network",
    }),
    visualizer({
      filename: "./stats/sunburst.html",
      template: "sunburst",
    }),
  ],

  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          react: ["react", "react/jsx-runtime"],
          "react-dom": ["react-dom", "react-dom/client"],
          "tailwind-merge": ["tailwind-merge"],
        },
      },
    },
  },

  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
