import React from "react";
import CustomerSearch from "../../components/CustomerSearch";
import Page from "../../components/Page";
import PageHeading from '../../components/PageHeading'

export default function CustomerListPage() {
  return (
    <Page>
      <PageHeading>Customer List</PageHeading>
      <CustomerSearch />
    </Page>
  );
}
