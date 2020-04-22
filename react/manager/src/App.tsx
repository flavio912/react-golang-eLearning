import * as React from 'react';
//@ts-ignore
import { BrowserProtocol, queryMiddleware } from 'farce';
import {
  createFarceRouter,
  createRender,
  makeRouteConfig,
  Route,
  RouteRenderArgs,
  RenderErrorArgs,
  RedirectException
} from 'found';
//@ts-ignore
import { Resolver } from 'found-relay';
import environment from './api/environment';
import { graphql, createFragmentContainer } from 'react-relay';
import LoginPage from 'views/Login';
import { ThemeProvider } from 'react-jss';
import theme from './helpers/theme';
import { AppHolder } from 'views/AppHolder';
import { Redirect } from 'react-router-dom';
import Card from 'components/core/Card';

const ExamplePageQuery = graphql`
  query App_Query {
    manager {
      uuid
      firstName
      lastName
    }
  }
`;

const Router = createFarceRouter({
  historyProtocol: new BrowserProtocol(),
  historyMiddlewares: [queryMiddleware],
  routeConfig: makeRouteConfig(
    <Route>
      <Route path="/(login)?" Component={LoginPage} query={ExamplePageQuery} />
      <Route
        path="/app"
        Component={AppHolder}
        query={ExamplePageQuery}
        render={({ props, error }: any) => {
          // Check if user is logged in, if not redirect to login
          if (props?.manager) return <AppHolder {...props} />;
          if (error) {
            throw new RedirectException('/login');
          }
          return undefined;
        }}
      >
        {/* Page info goes here */}
        <Route Component={LoginPage} />
        {/* Page info goes here */}
      </Route>
    </Route>
  ),
  render: createRender({})
});

const App = () => (
  <ThemeProvider theme={theme}>
    <Router resolver={new Resolver(environment)} />
  </ThemeProvider>
);

export default App;
