import { render } from "@testing-library/react";
import { ReactNode } from "react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import RegisterView from "./RegisterView";

const Router = (children: ReactNode) => {
  return (
    <MemoryRouter initialEntries={["/register"]}>
      <Routes>
        <Route path="/register" element={children} />
      </Routes>
    </MemoryRouter>
  );
};

describe("RegisterView", () => {
  it("renders the register form", () => {
    const { getByTestId } = render(Router(<RegisterView />));
    expect(getByTestId("register-form")).toBeInTheDocument();
  });

  describe("Register form renders correctly", () => {
    it("renders the username field", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("username-field")).toBeInTheDocument();
    });

    it("renders the name field", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("name-field")).toBeInTheDocument();
    });

    it("renders the email field", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("email-field")).toBeInTheDocument();
    });

    it("renders the password field", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("password-field")).toBeInTheDocument();
    });

    it("renders the confirm password field", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("confirm-password-field")).toBeInTheDocument();
    });

    it("renders the login button", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("login-button")).toBeInTheDocument();
    });

    it("renders the register button", () => {
      const { getByTestId } = render(Router(<RegisterView />));
      expect(getByTestId("register-button")).toBeInTheDocument();
    });
  });
});
