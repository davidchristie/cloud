import React from "react";
import "./index.css";

interface Props {
  id: string;
  label: string;
  name: string;
  onChange: (value: string) => void;
  value: string;
}

export default function TextInput({ id, label, name, onChange, value }: Props) {
  return (
    <div className="TextInput">
      <label htmlFor={id}>{label}</label>
      <input
        id={id}
        name={name}
        onChange={(event) => {
          onChange(event.target.value);
        }}
        value={value}
      />
    </div>
  );
}
