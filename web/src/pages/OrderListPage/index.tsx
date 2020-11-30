import React from "react";
import Page from "../../components/Page";
import PageHeading from "../../components/PageHeading";
import OrderSearch from "../../components/OrderSearch";

export default function OrderListPage() {
  return (
    <Page>
      <PageHeading>Order List</PageHeading>
      <OrderSearch />
    </Page>
  );
}
