import { gql, useMutation } from "@apollo/client";
import React, { FormEventHandler } from "react";
import { useHistory } from "react-router-dom";
import FormControl from "../FormControl";
import SaveButton from "../SaveButton";
import TextInput from "../TextInput";

interface CreateCustomerMutation {
  createCustomer: {
    firstName: string;
    id: string;
    lastName: string;
  };
}

const CREATE_PRODUCT_MUTATION = gql`
  mutation CreateCustomer($input: CreateCustomerInput!) {
    createCustomer(input: $input) {
      firstName
      id
      lastName
    }
  }
`;

export default function CreateCustomerForm() {
  const [firstName, setFirstName] = React.useState("");
  const [lastName, setLastName] = React.useState("");
  const valid = firstName.length > 0 && lastName.length > 0;
  const [
    createCustomer,
    { data, error, loading },
  ] = useMutation<CreateCustomerMutation>(CREATE_PRODUCT_MUTATION);
  const history = useHistory();
  React.useEffect(() => {
    if (data && !error && !loading) {
      history.push("/");
    }
  }, [data, error, history, loading]);
  const handleSubmit: FormEventHandler = (event) => {
    event.preventDefault();
    if (valid) {
      createCustomer({
        variables: {
          input: {
            firstName,
            lastName,
          },
        },
      });
    }
  };
  return (
    <div className="CreateCustomerForm">
      <form onSubmit={handleSubmit}>
        <FormControl>
          <TextInput
            id="first-name"
            label="First name"
            name="first-name"
            onChange={setFirstName}
            value={firstName}
          />
        </FormControl>
        <FormControl>
          <TextInput
            id="description"
            label="Last name"
            name="last-name"
            onChange={setLastName}
            value={lastName}
          />
        </FormControl>
        <SaveButton loading={loading} valid={valid} />
      </form>
    </div>
  );
}
