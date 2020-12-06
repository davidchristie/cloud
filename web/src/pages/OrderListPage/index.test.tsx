import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import { MemoryRouter } from "react-router-dom";
import OrderListPage from ".";

function renderOrderListPage() {
  render(
    <MemoryRouter>
      <MockedProvider mocks={[]}>
        <OrderListPage />
      </MockedProvider>
    </MemoryRouter>
  );
}

it("renders 'Order List' page heading", () => {
  renderOrderListPage();
  const pageHeading = screen.getByText(/order list/i, {
    selector: 'h1',
  });
  expect(pageHeading).toBeInTheDocument();
});

it("renders order search", () => {
  renderOrderListPage();
  const orderSearch = screen.getByTestId("OrderList");
  expect(orderSearch).toBeInTheDocument();
});
