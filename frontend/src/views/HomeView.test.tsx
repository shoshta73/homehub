import { ReactNode } from "react";
import { MemoryRouter, Route, Routes } from "react-router-dom";
import { render } from "@testing-library/react";
import HomeView from "./HomeView";

const Router = (children: ReactNode) => {
  return (
    <MemoryRouter initialEntries={["/"]}>
      <Routes>
        <Route path="/" element={children} />
      </Routes>
    </MemoryRouter>
  );
};

describe("HomeView", () => {
  it("renders the menubar", () => {
    const { getByTestId } = render(Router(<HomeView />));
    expect(getByTestId("menubar")).toBeInTheDocument();
  });

  describe("menubar renders correctly", () => {
    it("renders the register button", () => {
      const { getByTestId } = render(Router(<HomeView />));
      expect(getByTestId("register-button")).toBeInTheDocument();
    });

    it("renders the login button", () => {
      const { getByTestId } = render(Router(<HomeView />));
      expect(getByTestId("login-button")).toBeInTheDocument();
    });

    it("renders the mode button", () => {
      const { getByTestId } = render(Router(<HomeView />));
      expect(getByTestId("mode-button")).toBeInTheDocument();
    });
  });

  describe("Buttons render correctly", () => {
    it("renders the register correctly", () => {
      const { getByTestId } = render(Router(<HomeView />));
      expect(getByTestId("register-icon")).toBeInTheDocument();
      expect(getByTestId("register-text")).toBeInTheDocument();
    });

    it("renders the login correctly", () => {
      const { getByTestId } = render(Router(<HomeView />));
      expect(getByTestId("login-icon")).toBeInTheDocument();
      expect(getByTestId("login-text")).toBeInTheDocument();
    });

    it("renders the mode correctly", () => {
      const { getByTestId } = render(Router(<HomeView />));
      expect(getByTestId("mode-icon")).toBeInTheDocument();
    });
  });
});
