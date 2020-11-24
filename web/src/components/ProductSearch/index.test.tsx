import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import ProductSearch from ".";

it("renders 'Product Search' heading", () => {
  render(
    <MockedProvider mocks={[]}>
      <ProductSearch />
    </MockedProvider>
  );
  const heading = screen.getByText(/product search/i);
  expect(heading).toBeInTheDocument();
});
