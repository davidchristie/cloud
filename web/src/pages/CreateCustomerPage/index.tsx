import React from "react";
import CreateCustomerForm from "../../components/CreateCustomerForm";
import Page from "../../components/Page";
import PageHeading from '../../components/PageHeading'

export default function CreateCustomerPage() {
  return (
    <Page>
      <PageHeading>Create Customer</PageHeading>
      <CreateCustomerForm />
    </Page>
  );
}
