import { gql, useQuery } from '@apollo/client';
import React from 'react'
import { useParams } from "react-router-dom";
import Page from '../../components/Page';
import PageHeading from "../../components/PageHeading";
import NotFoundPage from '../NotFoundPage';

interface Params {
  productId?: string;
}

interface ProductQuery {
  product: {
    description: string
    id: string
    name: string
  }
}

const PRODUCT_QUERY = gql`
  query Product($id: String!) {
    product(id: $id) {
      description
      id
      name
    }
  }
`

export default function ProductDetailPage() {
  const { productId = '' } = useParams<Params>();
  const { data, error, loading } = useQuery<ProductQuery>(PRODUCT_QUERY, {
    variables: {
      id: productId,
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
          <PageHeading>{data.product.name}</PageHeading>
          <p>{data.product.description}</p>
        </>
      )}
    </Page>
  );
}
