import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import CustomerSearch from ".";

it("renders 'Customer Search' heading", () => {
  render(
    <MockedProvider mocks={[]}>
      <CustomerSearch />
    </MockedProvider>
  );
  const heading = screen.getByText(/customer search/i);
  expect(heading).toBeInTheDocument();
});
