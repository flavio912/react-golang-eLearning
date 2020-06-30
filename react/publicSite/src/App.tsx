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
import { ThemeProvider } from 'react-jss';
import theme from './helpers/theme';
import { AppHolder } from 'views/AppHolder';
import { Redirect } from 'react-router-dom';
import Home from 'views/Home';
import AboutUs from 'views/AboutUs';
import Register from 'views/Register';

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
      <Route
        Component={AppHolder}
        //query={ExamplePageQuery}
        render={({ props, error }: any) => {
          // Check if user is logged in, if not redirect to login
          // if (props?.manager) return <AppHolder {...props} />;
          // if (error) {
          //   throw new RedirectException("/login");
          // }
          // return undefined;
          return <AppHolder {...props} />;
        }}
      >
        <Route path="/" Component={Home} />
        <Route path="/aboutus" Component={AboutUs} />
        <Route path="/register" Component={Register} />
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
