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

  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },

  build: {
    outDir: "../bin/dist",
    emptyOutDir: true,
    rollupOptions: {
      output: {
        manualChunks: {
          react: ["react", "react/jsx-runtime"],
          "react-dom": ["react-dom", "react-dom/client"],
          "tailwind-merge": ["tailwind-merge"],
          "react-router": ["react-router-dom"],
          icons: ["@radix-ui/react-icons", "lucide-react"],
          "ui/avatar": ["./src/components/ui/avatar.tsx"],
          "ui/button": ["./src/components/ui/button.tsx"],
          "ui/form": ["./src/components/ui/form.tsx"],
          "ui/input": ["./src/components/ui/input.tsx"],
          "ui/label": ["./src/components/ui/label.tsx"],
          "ui/menubar": ["./src/components/ui/menubar.tsx"],
          "ui/textarea": ["./src/components/ui/textarea.tsx"],
          "views/home": ["./src/views/HomeView.tsx"],
          "views/login": ["./src/views/LoginView.tsx"],
          "views/register": ["./src/views/RegisterView.tsx"],
          "views/user-home": ["./src/views/UserHomeView.tsx"],
          "views/pastebin": ["./src/views/PasteBinView.tsx"],
          zustand: ["zustand", "zustand/middleware"],
        },
      },
    },
  },

  server: {
    proxy: {
      "/auth": "http://localhost:8080",
      "/avatar": "http://localhost:8080",
      "/avatars": "http://localhost:8080",
      "/pastebin": "http://localhost:8080",
      "/pastebin/stats": "http://localhost:8080",
    },
  },
});
