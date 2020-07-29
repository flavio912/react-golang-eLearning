import React from 'react';
import { Router } from 'react-router-dom';
import { renderRoutes } from 'react-router-config';
import { createBrowserHistory } from 'history';
import MomentUtils from '@date-io/moment';
import { Provider as StoreProvider } from 'react-redux';
import { ThemeProvider } from '@material-ui/styles';
import { MuiPickersUtilsProvider } from '@material-ui/pickers';
import 'react-perfect-scrollbar/dist/css/styles.css';
import { theme, themeWithRtl } from './theme';
import { configureStore } from './store';
import routes from './routes';
import GoogleAnalytics from './components/GoogleAnalytics';
import CookiesNotification from './components/CookiesNotification';
import ScrollReset from './components/ScrollReset';
import StylesProvider from './components/StylesProvider';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { ApolloClient } from 'apollo-client';
import { ApolloProvider } from '@apollo/react-hooks';
import { HttpLink } from 'apollo-link-http';
import { ApolloLink, from } from 'apollo-link';
import { onError } from 'apollo-link-error';
import { gql } from 'apollo-boost';

import './mixins/chartjs';
import './mixins/moment';
import './mixins/validate';
import './mixins/prismjs';
import './mock';
import './assets/scss/main.scss';

const history = createBrowserHistory();
const store = configureStore();

const authMiddleware = new ApolloLink((operation, forward) => {
  // add the authorization to the headers
  operation.setContext({
    headers: {
      Authorization: `Bearer ${localStorage.getItem('auth')}`
    }
  });

  return forward(operation);
});

const errorLink = new onError(({ graphQLErrors, networkError }) => {
  if (graphQLErrors)
    graphQLErrors.map(({ message, locations, path }) =>
      console.log(
        `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
      )
    );

  if (networkError) console.log(`[Network error]: ${networkError}`);
});

const link = new HttpLink({
  uri: '//ttc.devserver.london/graphql',
  credentials: 'include' //TODO: Possibly change to same-origin in prod
});

const client = new ApolloClient({
  link: from([errorLink, authMiddleware, link]),
  cache: new InMemoryCache(),
  defaultOptions: {
    query: {
      fetchPolicy: 'cache-and-network',
      errorPolicy: 'all'
    }
  },
  typeDefs: gql`
    enum AccessType {
      restricted
      open
    }

    enum CourseType {
      online
      classroom
    }
  `
});

function App() {
  const direction = 'ltr';

  return (
    <StoreProvider store={store}>
      <ThemeProvider theme={direction === 'rtl' ? themeWithRtl : theme}>
        <StylesProvider direction={direction}>
          <MuiPickersUtilsProvider utils={MomentUtils}>
            <ApolloProvider client={client}>
              <Router history={history}>
                <ScrollReset />
                <GoogleAnalytics />
                <CookiesNotification />
                {renderRoutes(routes)}
              </Router>
            </ApolloProvider>
          </MuiPickersUtilsProvider>
        </StylesProvider>
      </ThemeProvider>
    </StoreProvider>
  );
}

export default App;
