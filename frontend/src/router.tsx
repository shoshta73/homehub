import { createBrowserRouter, Outlet } from "react-router-dom";
import { Button } from "./components/ui/button";

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
        element: (
          <>
            <h1 className="text-3xl font-bold underline">Hello from HomeHub!</h1>
            <Button>Click me</Button>
          </>
        ),
      },
    ],
  },
]);

export default router;
