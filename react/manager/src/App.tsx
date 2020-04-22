import * as React from "react";
//@ts-ignore
import { BrowserProtocol, queryMiddleware } from "farce";
import {
  createFarceRouter,
  createRender,
  makeRouteConfig,
  Route,
  RouteRenderArgs,
} from "found";
//@ts-ignore
import { Resolver } from "found-relay";
import environment from "./api/environment";
import { graphql, createFragmentContainer } from "react-relay";
import LoginPage from "views/Login";
import { ThemeProvider } from "react-jss";
import theme from "./helpers/theme";
import { Redirect } from "react-router-dom";

const ExamplePageQuery = graphql`
  query App_Query {
    info
  }
`;

const Router = createFarceRouter({
  historyProtocol: new BrowserProtocol(),
  historyMiddlewares: [queryMiddleware],
  routeConfig: makeRouteConfig(
    <Route>
      <Route
        path="/(login)?"
        Component={LoginPage}
        query={ExamplePageQuery}
        render={({ props }: RouteRenderArgs) => {
          console.log(props);
          return <LoginPage data={props} />;
        }}
      />
    </Route>
  ),

  render: createRender({}),
});

const App = () => (
  <ThemeProvider theme={theme}>
    <Router resolver={new Resolver(environment)} />)
  </ThemeProvider>
);

export default App;
