export function baseUrl(): string {
  const viteApiUrl = import.meta.env.VITE_API_URL;
  const prodApiUrl = (window as any).env?.API_URL;

  if (!viteApiUrl && !prodApiUrl) {
    throw new Error("API_URL is not defined");
  }

  if (import.meta.env.DEV) {
    return viteApiUrl || "http://localhost:3000";
  }

  return prodApiUrl;
}
