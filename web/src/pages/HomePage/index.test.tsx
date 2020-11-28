import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import { MemoryRouter } from "react-router-dom";
import HomePage from ".";

function renderHomePage() {
  render(
    <MemoryRouter>
      <MockedProvider mocks={[]}>
        <HomePage />
      </MockedProvider>
    </MemoryRouter>
  );
}

it("renders 'Home' page heading", () => {
  renderHomePage();
  const pageHeading = screen.getByText(/home/i, {
    selector: 'h1',
  });
  expect(pageHeading).toBeInTheDocument();
});
