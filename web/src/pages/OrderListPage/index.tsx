import React from "react";
import Page from "../../components/Page";
import PageHeading from "../../components/PageHeading";
import OrderList from "../../components/OrderList";

export default function OrderListPage() {
  return (
    <Page>
      <PageHeading>Order List</PageHeading>
      <OrderList />
    </Page>
  );
}
