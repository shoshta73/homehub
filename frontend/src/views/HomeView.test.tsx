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
});
