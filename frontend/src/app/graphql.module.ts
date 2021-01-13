import { APOLLO_OPTIONS } from "apollo-angular";
import { HttpLink } from "apollo-angular/http";
import {
  split,
  ApolloClientOptions,
  ApolloLink,
  DefaultOptions,
  InMemoryCache,
} from "@apollo/client/core";
import { WebSocketLink } from "@apollo/client/link/ws";
import { getMainDefinition } from "@apollo/client/utilities";
import { NgModule } from "@angular/core";
import { setContext } from "@apollo/client/link/context";

let local: boolean;
let uri = "atske.herokuapp.com/api"; // <-- add the URL of the GraphQL server here
let wsSchema = "wss://";
let httpSchema = "https://";

const defaultOptions: DefaultOptions = {
  watchQuery: {
    fetchPolicy: "no-cache",
    errorPolicy: "ignore",
  },
  query: {
    fetchPolicy: "no-cache",
    errorPolicy: "all",
  },
};

export function createApollo(httpLink: HttpLink): ApolloClientOptions<any> {
  if (location.hostname === "localhost" || location.hostname === "127.0.0.1") {
    local = true;
  }
  if (local) {
    wsSchema = "ws://";
    uri = "localhost:4242/api";
    httpSchema = "http://";
  }
  const basic = setContext((operation, context) => ({
    headers: {
      Accept: "charset=utf-8",
    },
  }));
  const token = window.localStorage.getItem("token");

  const auth = setContext((operation, context) => {
    if (token === null) {
      return {};
    } else {
      return {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      };
    }
  });
  const httpUrlLInk = httpLink.create({ uri: httpSchema + uri });

  const webSocketLink = new WebSocketLink({
    uri: wsSchema + uri,
    options: {
      reconnect: true,
      lazy: true,
      connectionParams: {
        Authorization: `Bearer ${token}`,
      },
    },
  });

  const enhanced_link = split(
    // split based on operation type
    ({ query }) => {
      const dif = getMainDefinition(query);
      return (
        dif.kind === "OperationDefinition" && dif.operation === "subscription"
      );
    },
    webSocketLink,
    httpUrlLInk
  );

  // const link = ApolloLink.from([basic, auth, httpUrlLInk]);
  const link = ApolloLink.from([basic, auth, enhanced_link]);

  const cache = new InMemoryCache();

  return {
    link,
    cache,
    defaultOptions,
  };
}

@NgModule({
  providers: [
    {
      provide: APOLLO_OPTIONS,
      useFactory: createApollo,
      deps: [HttpLink],
    },
  ],
})
export class GraphQLModule {}
