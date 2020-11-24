import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import React from "react";
import NotFoundPage from ".";

function renderNotFoundPage() {
  render(
    <MockedProvider mocks={[]}>
      <NotFoundPage />
    </MockedProvider>
  );
}

it("renders 'Not Found' page heading", () => {
  renderNotFoundPage();
  const pageHeading = screen.getByText(/not found/i);
  expect(pageHeading).toBeInTheDocument();
});
