export function getCreateCustomerPageUrl() {
  return '/create/customer'
}

export function getCreateProductPageUrl() {
  return '/create/product'
}

export function getCustomerDetailPageUrl(customerId: string) {
  return `/customers/${customerId}`;
}

export function getCustomerListPageUrl() {
  return '/customers'
}

export function getHomePageUrl() {
  return '/'
}

export function getOrderDetailPageUrl(orderId: string) {
  return `/order/${orderId}`
}

export function getOrderListPageUrl() {
  return '/orders'
}

export function getProductDetailPageUrl(productId: string) {
  return `/product/${productId}`;
}

export function getProductListPageUrl() {
  return '/products'
}
