import React from "react";
import { Link } from "react-router-dom";

export default function CreateProductButton() {
  return (
    <Link to="/create/product">
      <button>Create Product</button>
    </Link>
  );
}
