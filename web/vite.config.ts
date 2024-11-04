import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    {
      name: "inject-env-js",
      transformIndexHtml(html) {
        // Inject /env.js only in production
        if (process.env.NODE_ENV === "production") {
          return html.replace("</head>", `<script src="/env.js"></script>\n</head>`);
        }
        return html;
      },
    },
  ],

  server: {
    port: 5173,
    strictPort: true,
  },
});
