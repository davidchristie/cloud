import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import { MemoryRouter } from "react-router-dom";
import NotFoundPage from ".";

function renderNotFoundPage() {
  render(
    <MemoryRouter>
      <MockedProvider mocks={[]}>
        <NotFoundPage />
      </MockedProvider>
    </MemoryRouter>
  );
}

it("renders 'Not Found' page heading", () => {
  renderNotFoundPage();
  const pageHeading = screen.getByText(/not found/i, {
    selector: 'h1',
  });
  expect(pageHeading).toBeInTheDocument();
});
