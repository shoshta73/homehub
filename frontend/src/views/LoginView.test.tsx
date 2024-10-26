import { render } from "@testing-library/react";
import { ReactNode } from "react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import LoginView from "./LoginView";

const Router = (children: ReactNode) => {
  return (
    <MemoryRouter initialEntries={["/login"]}>
      <Routes>
        <Route path="/login" element={children} />
      </Routes>
    </MemoryRouter>
  );
};

describe("LoginView", () => {
  it("renders the login form", () => {
    const { getByTestId } = render(Router(<LoginView />));
    expect(getByTestId("login-form")).toBeInTheDocument();
  });

  describe("Login form renders correctly", () => {
    it("renders the username field", () => {
      const { getByTestId } = render(Router(<LoginView />));
      expect(getByTestId("username-field")).toBeInTheDocument();
    });

    it("renders the email field", () => {
      const { getByTestId } = render(Router(<LoginView />));
      expect(getByTestId("email-field")).toBeInTheDocument();
    });

    it("renders the password field", () => {
      const { getByTestId } = render(Router(<LoginView />));
      expect(getByTestId("password-field")).toBeInTheDocument();
    });

    it("render the register button", () => {
      const { getByTestId } = render(Router(<LoginView />));
      expect(getByTestId("register-button")).toBeInTheDocument();
    });

    it("renders the submit button", () => {
      const { getByTestId } = render(Router(<LoginView />));
      expect(getByTestId("submit-button")).toBeInTheDocument();
    });
  });
});
