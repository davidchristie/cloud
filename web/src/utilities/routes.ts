export function getCreateProductPageUrl() {
  return '/create/product'
}

export function getHomePageUrl() {
  return '/'
}

export function getProductDetailPageUrl(productId: string) {
  return `/product/${productId}`;
}

export function getProductListPageUrl() {
  return '/products'
}
