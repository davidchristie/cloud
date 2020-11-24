import React from "react";
import CreateProductButton from "../../components/CreateProductButton";
import ProductSearch from "../../components/ProductSearch";

function HomePage() {
  return (
    <div className="HomePage" data-testid="HomePage">
      <h1>Home</h1>
      <CreateProductButton />
      <ProductSearch />
    </div>
  );
}

export default HomePage;
