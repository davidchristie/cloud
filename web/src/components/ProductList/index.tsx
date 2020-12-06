import { gql, useQuery } from "@apollo/client";
import React, { useState } from "react";
import { Link } from "react-router-dom";
import { getProductDetailPageUrl } from "../../utilities";

interface ProductListQuery {
  products: Array<{
    description: string;
    id: string;
    name: string;
  }>;
}

const PRODUCT_LIST_QUERY = gql`
  query ProductList($query: String!) {
    products(query: $query) {
      description
      id
      name
    }
  }
`;

export default function ProductList() {
  const [query, setQuery] = useState("");
  const { data, loading, error } = useQuery<ProductListQuery>(
    PRODUCT_LIST_QUERY,
    {
      variables: {
        query,
      },
    }
  );
  const handleQueryChange: React.ChangeEventHandler = (event) => {
    if (event.target instanceof HTMLInputElement) {
      setQuery(event.target.value);
    }
  };
  return (
    <div className="ProductList" data-testid="ProductList">
      <input onChange={handleQueryChange} value={query} />
      {error && <div>Error</div>}
      {loading && <div>Loading...</div>}
      {data && data.products.map((product) => (
        <div key={product.id}>
          <Link to={getProductDetailPageUrl(product.id)}>{product.name}</Link>
        </div>
      ))}
      {!loading && data && data.products.length === 0 && (
        <div>No results</div>
      )}
    </div>
  );
}
