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
import ExamplePage from "views/ExamplePage";

const ExamplePageQuery = graphql`
  query ExamplePage_Query {
    ...ExampleComponent_info
  }
`;

const Router = createFarceRouter({
  historyProtocol: new BrowserProtocol(),
  historyMiddlewares: [queryMiddleware],
  routeConfig: makeRouteConfig(
    <Route
      path="/"
      Component={ExamplePage}
      query={ExamplePageQuery}
      render={({ props }: RouteRenderArgs) => {
        console.log(props);
        return <ExamplePage data={props} />;
      }}
    />
  ),

  render: createRender({}),
});

const App = () => <Router resolver={new Resolver(environment)} />;

export default App;
