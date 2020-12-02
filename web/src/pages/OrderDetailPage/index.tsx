import { gql, useQuery } from '@apollo/client';
import React from 'react'
import { Link, useParams } from "react-router-dom";
import Page from '../../components/Page';
import PageHeading from "../../components/PageHeading";
import { getCustomerDetailPageUrl, getProductDetailPageUrl } from '../../utilities';
import NotFoundPage from '../NotFoundPage';

interface Params {
  orderId?: string;
}

interface OrderQuery {
  order: {
    customer: {
      firstName: string
      id: string
      lastName: string
    } | null
    createdAt: string
    id: string
    lineItems: Array<{
      product: {
        id: string
        name: string
      } | null
      quantity: number
    }>
  }
}

const ORDER_QUERY = gql`
  query Order($id: String!) {
    order(id: $id) {
      customer {
        firstName
        id
        lastName
      }
      createdAt
      id
      lineItems {
        product {
          id
          name
        }
        quantity
      }
    }
  }
`

export default function OrderDetailPage() {
  const { orderId = '' } = useParams<Params>();
  const { data, error, loading } = useQuery<OrderQuery>(ORDER_QUERY, {
    variables: {
      id: orderId,
    }
  })
  if (error) {
    return <NotFoundPage />
  }
  return (
    <Page>
      {loading && (
        <div>Loading...</div>
      )}
      {data && (
        <>
          <PageHeading>{data.order.createdAt}</PageHeading>
          <p>Customer: {data.order.customer ? (
            <Link to={getCustomerDetailPageUrl(data.order.customer.id)}>
              {data.order.customer.firstName} {data.order.customer.lastName}
            </Link>
          ) : 'None'}</p>
          {data.order.lineItems.map((lineItem, index) => (
            <div key={index}>
              {lineItem.quantity} | {lineItem.product ? (
                <Link to={getProductDetailPageUrl(lineItem.product.id)}>
                  {lineItem.product.name}
                </Link>
              ) : 'None'}
            </div>
          ))}
        </>
      )}
    </Page>
  );
}
