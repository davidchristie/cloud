import { gql, useQuery } from "@apollo/client";
import React, { useState } from "react";
import { Link } from "react-router-dom";
import { getProductDetailPageUrl } from "../../utilities";

interface ProductSearchQuery {
  products: Array<{
    description: string;
    id: string;
    name: string;
  }>;
}

const PRODUCT_SEARCH_QUERY = gql`
  query ProductSearch($query: String!) {
    products(query: $query) {
      description
      id
      name
    }
  }
`;

export default function ProductSearch() {
  const [query, setQuery] = useState("");
  const { data, loading, error } = useQuery<ProductSearchQuery>(
    PRODUCT_SEARCH_QUERY,
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
    <div className="ProductSearch" data-testid="ProductSearch">
      <h2>Product Search</h2>
      <input onChange={handleQueryChange} value={query} />
      {error && <div>Error</div>}
      {loading && <div>Loading...</div>}
      {query &&
        data &&
        data.products.map((product) => (
          <div key={product.id}>
            <Link to={getProductDetailPageUrl(product.id)}>{product.name}</Link>
          </div>
        ))}
      {query && !loading && data && data.products.length === 0 && (
        <div>No results</div>
      )}
    </div>
  );
}
