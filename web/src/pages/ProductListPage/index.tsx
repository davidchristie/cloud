import React from "react";
import CreateProductButton from "../../components/CreateProductButton";
import Page from "../../components/Page";
import PageHeading from "../../components/PageHeading";
import ProductList from "../../components/ProductList";

export default function ProductListPage() {
  return (
    <Page>
      <PageHeading>Product List</PageHeading>
      <CreateProductButton />
      <ProductList />
    </Page>
  );
}
