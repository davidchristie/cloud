import React from "react";
import CreateProductButton from "../../components/CreateProductButton";
import Page from "../../components/Page";
import ProductSearch from "../../components/ProductSearch";

export default function ProductListPage() {
  return (
    <Page>
      <h1>Product List</h1>
      <CreateProductButton />
      <ProductSearch />
    </Page>
  );
}
