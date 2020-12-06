import React from "react";
import CreateCustomerButton from '../../components/CreateCustomerButton'
import CustomerList from "../../components/CustomerList";
import Page from "../../components/Page";
import PageHeading from '../../components/PageHeading'

export default function CustomerListPage() {
  return (
    <Page>
      <PageHeading>Customer List</PageHeading>
      <CreateCustomerButton />
      <CustomerList />
    </Page>
  );
}
