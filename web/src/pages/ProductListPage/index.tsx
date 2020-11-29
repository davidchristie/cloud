import React from "react";
import CreateProductButton from "../../components/CreateProductButton";
import Page from "../../components/Page";
import PageHeading from "../../components/PageHeading";
import ProductSearch from "../../components/ProductSearch";

export default function ProductListPage() {
  return (
    <Page>
      <PageHeading>Product List</PageHeading>
      <CreateProductButton />
      <ProductSearch />
    </Page>
  );
}
