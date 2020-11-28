import React from "react";
import CreateProductForm from "../../components/CreateProductForm";
import Page from "../../components/Page";

export default function CreateProductPage() {
  return (
    <Page>
      <h1>Create Product</h1>
      <CreateProductForm />
    </Page>
  );
}
