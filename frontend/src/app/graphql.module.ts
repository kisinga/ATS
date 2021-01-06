import {NgModule} from '@angular/core';
import {APOLLO_OPTIONS} from 'apollo-angular';
import {ApolloClientOptions, InMemoryCache,ApolloLink} from '@apollo/client/core';
import {HttpLink} from 'apollo-angular/http';
import { setContext } from '@apollo/client/link/context';

const uri = 'http://localhost:4242/api'; // <-- add the URL of the GraphQL server here
export function createApollo(httpLink: HttpLink): ApolloClientOptions<any> {
  const basic = setContext((operation, context) => ({
    headers: {
      Accept: 'charset=utf-8'
    }
  }));

  const auth = setContext((operation, context) => {
    const token = window.localStorage.getItem('jwtToken');

    if (token === null) {
      return {};
    } else {
      return {
        headers: {
          Authorization: `Bearer ${token}`
        }
      };
    }
  });
  const link = ApolloLink.from([basic, auth, httpLink.create({ uri })]);
  const cache = new InMemoryCache();
  return {
    link,
    cache
  }

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
