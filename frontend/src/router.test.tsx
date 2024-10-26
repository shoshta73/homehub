import { render } from "@testing-library/react";
import { RouterProvider } from "react-router-dom";
import router from "./router";

describe("Router", () => {
  it("renders with wrapper", () => {
    const { getByTestId } = render(<RouterProvider router={router} />);

    expect(getByTestId("wrapper")).toBeInTheDocument();

    expect(getByTestId("wrapper")).toHaveClass("flex");
    expect(getByTestId("wrapper")).toHaveClass("flex-col");
    expect(getByTestId("wrapper")).toHaveClass("min-h-screen");
    expect(getByTestId("wrapper")).toHaveClass("h-screen");
    expect(getByTestId("wrapper")).toHaveClass("max-h-screen");
  });
});
