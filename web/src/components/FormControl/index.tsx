import React from "react";
import "./index.css";

interface Props {
  children: React.ReactNode;
}

export default function FormControl({ children }: Props) {
  return <div className="FormControl">{children}</div>;
}
