import React from "react";
import { Route, Switch } from "react-router-dom";
import CreateProductPage from "./pages/CreateProductPage";
import HomePage from "./pages/HomePage";
import NotFoundPage from "./pages/NotFoundPage";
import ProductDetailPage from "./pages/ProductDetailPage";
import ProductListPage from "./pages/ProductListPage";
import { getCreateProductPageUrl, getHomePageUrl, getProductDetailPageUrl, getProductListPageUrl } from "./utilities";

export default function Routes() {
  return (
    <Switch>
      <Route exact path={getHomePageUrl()}>
        <HomePage />
      </Route>
      <Route exact path={getCreateProductPageUrl()}>
        <CreateProductPage />
      </Route>
      <Route exact path={getProductDetailPageUrl(':productId')}>
        <ProductDetailPage />
      </Route>
      <Route exact path={getProductListPageUrl()}>
        <ProductListPage />
      </Route>
      <Route path="*">
        <NotFoundPage />
      </Route>
    </Switch>
  );
}
