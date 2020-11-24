import React from "react";

interface Props {
  loading: boolean;
  valid: boolean;
}

export default function SaveButton({ loading, valid }: Props) {
  return (
    <button disabled={loading || !valid} type="submit">
      {loading ? "Saving" : "Save"}
    </button>
  );
}
