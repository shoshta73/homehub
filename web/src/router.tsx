import { createBrowserRouter } from "react-router-dom";
import App from "./App";
import RegisterView from "./views/RegisterView";
import LoginView from "./views/LoginView";
import Home from "./views/user/Home";
import Root from "./views/user/pastebin/Root.tsx";
import Create from "./views/user/pastebin/Create.tsx";
import Paste from "./views/user/pastebin/Paste.tsx";

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
        path: "/home",
        element: <Home />,
        children: [
          {
            path: "pastebin",
            children: [
              {
                index: true,
                element: <Root />,
              },
              {
                path: "create",
                element: <Create />,
              },
              {
                path: "paste/:id",
                element: <Paste />,
              },
            ],
          },
        ],
      },
      {
        path: "*",
        element: <h1>Not Implented Yet</h1>,
      },
    ],
  },
]);

export default router;
