import { createBrowserRouter } from "react-router-dom";
import App from "./App";
import RegisterView from "./views/RegisterView";
import LoginView from "./views/LoginView";

const router = createBrowserRouter([
  {
    path: "/",
    children: [
      {
        index: true,
        element: <App />,
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
        path: "*",
        element: <h1>Not Implented Yet</h1>,
      },
    ],
  },
]);

export default router;
