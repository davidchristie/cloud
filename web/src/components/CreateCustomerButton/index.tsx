import React from "react";
import { Link } from "react-router-dom";
import { getCreateCustomerPageUrl } from '../../utilities'

export default function CreateCustomerButton() {
  return (
    <Link to={getCreateCustomerPageUrl()}>
      <button>Create Customer</button>
    </Link>
  );
}
