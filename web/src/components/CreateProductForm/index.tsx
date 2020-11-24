import { gql, useMutation } from "@apollo/client";
import React, { FormEventHandler } from "react";
import { useHistory } from "react-router-dom";
import FormControl from "../FormControl";
import SaveButton from "../SaveButton";
import TextArea from "../TextArea";
import TextInput from "../TextInput";

interface CreateProductMutation {
  createProduct: {
    description: string;
    id: string;
    name: string;
  };
}

const CREATE_PRODUCT_MUTATION = gql`
  mutation CreateProduct($input: CreateProductInput!) {
    createProduct(input: $input) {
      description
      id
      name
    }
  }
`;

export default function CreateProductForm() {
  const [name, setName] = React.useState("");
  const [description, setDescription] = React.useState("");
  const valid = name.length > 0;
  const [
    createProduct,
    { data, error, loading },
  ] = useMutation<CreateProductMutation>(CREATE_PRODUCT_MUTATION);
  const history = useHistory();
  React.useEffect(() => {
    if (data && !error && !loading) {
      history.push("/");
    }
  }, [data, error, history, loading]);
  const handleSubmit: FormEventHandler = (event) => {
    event.preventDefault();
    if (valid) {
      createProduct({
        variables: {
          input: {
            description,
            name,
          },
        },
      });
    }
  };
  return (
    <div className="CreateProductForm">
      <form onSubmit={handleSubmit}>
        <FormControl>
          <TextInput
            id="name"
            label="Name"
            name="name"
            onChange={setName}
            value={name}
          />
        </FormControl>
        <FormControl>
          <TextArea
            id="description"
            label="Description"
            name="description"
            onChange={setDescription}
            value={description}
          />
        </FormControl>
        <SaveButton loading={loading} valid={valid} />
      </form>
    </div>
  );
}
