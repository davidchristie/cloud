import React from "react";
import "./index.css";

interface Props {
  id: string;
  label: string;
  name: string;
  onChange: (value: string) => void;
  value: string;
}

export default function TextArea({ id, label, name, onChange, value }: Props) {
  return (
    <div className="TextArea">
      <label htmlFor={id}>{label}</label>
      <textarea
        cols={30}
        id={id}
        name={name}
        onChange={(event) => {
          onChange(event.target.value);
        }}
        rows={10}
        value={value}
      />
    </div>
  );
}
