import { RouterProvider } from "react-router-dom";
import router from "./router";
import useAppState from "./store/state";

function App() {
  const state = useAppState();

  if (state.mode === "light") {
    if (document.documentElement.classList.contains("dark")) {
      document.documentElement.classList.remove("dark");
      document.documentElement.classList.add("light");
    }

    if (!document.documentElement.classList.contains("light")) {
      document.documentElement.classList.add("light");
    }
  } else {
    if (document.documentElement.classList.contains("light")) {
      document.documentElement.classList.remove("light");
      document.documentElement.classList.add("dark");
    }

    if (!document.documentElement.classList.contains("dark")) {
      document.documentElement.classList.add("dark");
    }
  }

  return <RouterProvider router={router} />;
}

export default App;
