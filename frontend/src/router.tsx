import { createBrowserRouter, Outlet } from "react-router-dom";
import HomeView from "./views/HomeView";

const Wrapper = () => (
  <div className="flex flex-col min-h-screen h-screen max-h-screen">
    <Outlet />
  </div>
);

const router = createBrowserRouter([
  {
    path: "/",
    element: <Wrapper />,
    children: [
      {
        index: true,
        element: <HomeView />,
      },
    ],
  },
]);

export default router;
