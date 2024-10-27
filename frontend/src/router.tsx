import { createBrowserRouter, Outlet } from "react-router-dom";
import HomeView from "./views/HomeView";
import { lazy, Suspense } from "react";

const RegisterView = lazy(() => import("./views/RegisterView"));
const LoginView = lazy(() => import("./views/LoginView"));
const UserHomeView = lazy(() => import("./views/UserHomeView"));

const Wrapper = () => (
  <div className="flex flex-col min-h-screen h-screen max-h-screen" data-testid="wrapper">
    <Suspense fallback={<div>Loading...</div>}>
      <Outlet />
    </Suspense>
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
      {
        path: "register",
        element: <RegisterView />,
      },
      {
        path: "login",
        element: <LoginView />,
      },
      {
        path: "home",
        element: <UserHomeView />,
      },
    ],
  },
]);

export default router;
