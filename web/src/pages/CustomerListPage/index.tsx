import React from "react";
import CustomerSearch from "../../components/CustomerSearch";
import Page from "../../components/Page";

export default function CustomerListPage() {
  return (
    <Page>
      <h1>Customer List</h1>
      <CustomerSearch />
    </Page>
  );
}
