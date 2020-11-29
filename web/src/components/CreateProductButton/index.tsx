import React from "react";
import { Link } from "react-router-dom";
import { getCreateProductPageUrl } from "../../utilities";

export default function CreateProductButton() {
  return (
    <Link to={getCreateProductPageUrl()}>
      <button>Create Product</button>
    </Link>
  );
}
