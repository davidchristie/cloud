import { gql, useQuery } from '@apollo/client';
import React from 'react'
import { useParams } from "react-router-dom";
import OrderList from '../../components/OrderList';
import Page from '../../components/Page';
import PageHeading from '../../components/PageHeading'
import NotFoundPage from '../NotFoundPage';

interface Params {
  customerId?: string;
}

interface CustomerQuery {
  customer: {
    firstName: string
    id: string
    lastName: string
  }
}

const CUSTOMER_QUERY = gql`
  query Customer($id: String!) {
    customer(id: $id) {
      firstName
      id
      lastName
    }
  }
`

export default function CustomerDetailPage() {
  const { customerId = '' } = useParams<Params>();
  const { data, error, loading } = useQuery<CustomerQuery>(CUSTOMER_QUERY, {
    variables: {
      id: customerId,
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
          <PageHeading>{data.customer.firstName} {data.customer.lastName}</PageHeading>
          <h2>Orders</h2>
          <OrderList customerId={data.customer.id} />
        </>
      )}
    </Page>
  );
}
