import React from "react";
import { Route, Switch } from "react-router-dom";
import CreateCustomerPage from "./pages/CreateCustomerPage";
import CreateProductPage from "./pages/CreateProductPage";
import CustomerDetailPage from "./pages/CustomerDetailPage";
import CustomerListPage from "./pages/CustomerListPage";
import HomePage from "./pages/HomePage";
import NotFoundPage from "./pages/NotFoundPage";
import OrderDetailPage from "./pages/OrderDetailPage";
import OrderListPage from "./pages/OrderListPage";
import ProductDetailPage from "./pages/ProductDetailPage";
import ProductListPage from "./pages/ProductListPage";
import { getCreateCustomerPageUrl, getCreateProductPageUrl, getCustomerDetailPageUrl, getCustomerListPageUrl, getHomePageUrl, getOrderDetailPageUrl, getOrderListPageUrl, getProductDetailPageUrl, getProductListPageUrl } from "./utilities";

const routes = [
  {
    Component: HomePage,
    path: getHomePageUrl(),
  },
  {
    Component: CreateCustomerPage,
    path: getCreateCustomerPageUrl(),
  },
  {
    Component: CustomerListPage,
    path: getCustomerListPageUrl(),
  },
  {
    Component: CustomerDetailPage,
    path: getCustomerDetailPageUrl(':customerId'),
  },
  {
    Component: OrderListPage,
    path: getOrderListPageUrl(),
  },
  {
    Component: OrderDetailPage,
    path: getOrderDetailPageUrl(':orderId'),
  },
  {
    Component: CreateProductPage,
    path: getCreateProductPageUrl(),
  },
  {
    Component: ProductDetailPage,
    path: getProductDetailPageUrl(':productId'),
  },
  {
    Component: ProductListPage,
    path: getProductListPageUrl(),
  }
]

export default function Routes() {
  return (
    <Switch>
      {routes.map(({ Component, path }, index) => (
        <Route exact key={index} path={path}>
          <Component />
        </Route>
      ))}
      <Route path="*">
        <NotFoundPage />
      </Route>
    </Switch>
  );
}
