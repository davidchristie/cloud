import { MockedProvider } from "@apollo/client/testing";
import { screen, render } from "@testing-library/react";
import HomePage from ".";

function renderHomePage() {
  render(
    <MockedProvider mocks={[]}>
      <HomePage />
    </MockedProvider>
  );
}

it("renders 'Home' page heading", () => {
  renderHomePage();
  const pageHeading = screen.getByText(/home/i);
  expect(pageHeading).toBeInTheDocument();
});

it("renders product search", () => {
  renderHomePage();
  const productSearch = screen.getByTestId("ProductSearch");
  expect(productSearch).toBeInTheDocument();
});
