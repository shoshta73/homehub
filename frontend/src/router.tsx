import { createHashRouter, Outlet } from "react-router-dom";
import HomeView from "./views/HomeView";
import { lazy, Suspense } from "react";

const RegisterView = lazy(() => import("./views/RegisterView"));
const LoginView = lazy(() => import("./views/LoginView"));

const Wrapper = () => (
  <div className="flex flex-col min-h-screen h-screen max-h-screen">
    <Suspense fallback={<div>Loading...</div>}>
      <Outlet />
    </Suspense>
  </div>
);

const router = createHashRouter([
  {
    path: "/",
    element: <Wrapper />,
    children: [
      {
        index: true,
        element: <HomeView />,
      },
      {
        path: "register",
        element: <RegisterView />,
      },
      {
        path: "login",
        element: <LoginView />,
      },
    ],
  },
]);

export default router;
