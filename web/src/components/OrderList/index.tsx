import { gql, useQuery } from "@apollo/client";
import React from "react";
import { Link } from "react-router-dom";
import { getOrderDetailPageUrl } from "../../utilities";
import DateFormat from '../DateFormat'

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

const LIMIT = 25

const ORDER_LIST_QUERY = gql`
  query OrderList($customerID: String, $limit: Int, $skip: Int) {
    orders(customerID: $customerID, limit: $limit, skip: $skip) {
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
  const { data, fetchMore, loading, error } = useQuery<OrderListQuery>(ORDER_LIST_QUERY, {
    variables: {
      customerID: customerId,
      limit: LIMIT,
    }
  });
  return (
    <div className="OrderList" data-testid="OrderList">
      {error && <div>Error</div>}
      {loading && <div>Loading...</div>}
      {data && (
        <div>
          {data.orders.map((order) => (
            <div key={order.id}>
              <Link to={getOrderDetailPageUrl(order.id)}>
                <DateFormat value={order.createdAt} />
              </Link>
            </div>
          ))}
          <button
            onClick={() => {
              fetchMore({
                updateQuery: (previousResult, { fetchMoreResult }) => {
                  return {
                    ...previousResult,
                    orders: [
                      ...previousResult.orders,
                      ...(fetchMoreResult?.orders || []),
                    ],
                  };
                },
                variables: {
                  customerID: customerId,
                  limit: LIMIT,
                  skip: data.orders.length,
                }
              })
            }}
          >Load more</button>
        </div>
      )}
      {!loading && data && data.orders.length === 0 && (
        <div>No results</div>
      )}
    </div>
  );
}
