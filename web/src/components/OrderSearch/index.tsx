import { gql, useQuery } from "@apollo/client";
import React from "react";

interface OrderSearchQuery {
  orders: Array<{
    customer: {
      id: string;
    }
    createdAt: string
    id: string;
    lineItems: Array<{
      product: {
        id: string
      }
      quantity: number
    }>
  }>;
}

const ORDER_SEARCH_QUERY = gql`
  query OrderSearch {
    orders {
      customer {
        id
      }
      createdAt
      id
      lineItems {
        product {
          id
        }
        quantity
      }
    }
  }
`;

export default function OrderSearch() {
  const { data, loading, error } = useQuery<OrderSearchQuery>(ORDER_SEARCH_QUERY);
  return (
    <div className="OrderSearch" data-testid="OrderSearch">
      {error && <div>Error</div>}
      {loading && <div>Loading...</div>}
      {data && data.orders.map((order) => (
        <div key={order.id}>
          {order.createdAt}
        </div>
      ))}
      {!loading && data && data.orders.length === 0 && (
        <div>No results</div>
      )}
    </div>
  );
}
