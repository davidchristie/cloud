import { gql, useQuery } from "@apollo/client";
import React, { useState } from "react";
import { Link } from 'react-router-dom'
import { getCustomerDetailPageUrl } from "../../utilities";

interface CustomerSearchQuery {
  customers: Array<{
    firstName: string;
    id: string;
    lastName: string;
  }>;
}

const CUSTOMER_SEARCH_QUERY = gql`
  query CustomerSearch($query: String!) {
    customers(query: $query) {
      firstName
      id
      lastName
    }
  }
`;

export default function CustomerSearch() {
  const [query, setQuery] = useState("");
  const { data, loading, error } = useQuery<CustomerSearchQuery>(
    CUSTOMER_SEARCH_QUERY,
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
    <div className="CustomerSearch" data-testid="CustomerSearch">
      <input onChange={handleQueryChange} value={query} />
      {error && <div>Error</div>}
      {loading && <div>Loading...</div>}
      {data && data.customers.map((customer) => (
          <div key={customer.id}>
            <Link to={getCustomerDetailPageUrl(customer.id)}>
              {customer.firstName} {customer.lastName}
            </Link>
          </div>
        ))}
      {!loading && data && data.customers.length === 0 && (
        <div>No results</div>
      )}
    </div>
  );
}
