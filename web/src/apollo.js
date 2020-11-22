import { ApolloClient, InMemoryCache } from "@apollo/client";

export function createApolloClient() {
  return new ApolloClient({
    cache: new InMemoryCache(),
    uri: "/query",
  });
}
