import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import { MemoryRouter } from "react-router-dom";
import ProductListPage from ".";

function renderProductListPage() {
  render(
    <MemoryRouter>
      <MockedProvider mocks={[]}>
        <ProductListPage />
      </MockedProvider>
    </MemoryRouter>
  );
}

it("renders 'Product List' page heading", () => {
  renderProductListPage();
  const pageHeading = screen.getByText(/product list/i, {
    selector: 'h1',
  });
  expect(pageHeading).toBeInTheDocument();
});

it("renders product search", () => {
  renderProductListPage();
  const productSearch = screen.getByTestId("ProductSearch");
  expect(productSearch).toBeInTheDocument();
});
