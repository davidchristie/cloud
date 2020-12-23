import { gql, useMutation } from "@apollo/client";
import React from 'react'
import { useHistory, useRouteMatch } from "react-router-dom";
import { getProductDetailPageUrl, getProductListPageUrl } from "../../utilities";

interface DeleteProductMutation {
  deleteProduct: {
    id: string;
  };
}

interface Props {
  productId: string
}

const DELETE_PRODUCT_MUTATION = gql`
  mutation CreateProduct($id: String!) {
    deleteProduct(id: $id) {
      id
    }
  }
`;

export default function DeleteProductButton({ productId }: Props) {
  const match = useRouteMatch()
  const history = useHistory()
  const [
    deleteProduct,
    { data, loading },
  ] = useMutation<DeleteProductMutation>(DELETE_PRODUCT_MUTATION, {
    variables: {
      id: productId
    }
  });
  console.log({ data, loading, match })
  React.useEffect(() => {
    if (data && match.url === getProductDetailPageUrl(data.deleteProduct.id)) {
      history.push(getProductListPageUrl())
    }
  }, [data, match, history])
  const handleClick = () => {
    deleteProduct()
  }
  return (
    <button disabled={loading} onClick={handleClick}>
      {loading ? 'Deleting...' : 'Delete'}
    </button>
  )
}
