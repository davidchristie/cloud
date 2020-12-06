import { gql, useQuery } from "@apollo/client";
import React from "react";
import { Link } from "react-router-dom";
import { getOrderDetailPageUrl } from "../../utilities";

interface OrderListQuery {
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

interface Props {
  customerId?: string
}

const ORDER_LIST_QUERY = gql`
  query OrderList($customerID: String) {
    orders(customerID: $customerID) {
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

export default function OrderList({ customerId }: Props) {
  const { data, loading, error } = useQuery<OrderListQuery>(ORDER_LIST_QUERY, {
    variables: {
      customerID: customerId,
    }
  });
  return (
    <div className="OrderList" data-testid="OrderList">
      {error && <div>Error</div>}
      {loading && <div>Loading...</div>}
      {data && data.orders.map((order) => (
        <div key={order.id}>
          <Link to={getOrderDetailPageUrl(order.id)}>
            {order.createdAt}
          </Link>
        </div>
      ))}
      {!loading && data && data.orders.length === 0 && (
        <div>No results</div>
      )}
    </div>
  );
}
