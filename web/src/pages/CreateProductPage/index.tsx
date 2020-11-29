import React from "react";
import CreateProductForm from "../../components/CreateProductForm";
import Page from "../../components/Page";
import PageHeading from '../../components/PageHeading'

export default function CreateProductPage() {
  return (
    <Page>
      <PageHeading>Create Product</PageHeading>
      <CreateProductForm />
    </Page>
  );
}
