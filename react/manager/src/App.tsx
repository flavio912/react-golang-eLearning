import * as React from "react";
import { BrowserProtocol, queryMiddleware } from "farce";
import { createFarceRouter, createRender, makeRouteConfig, Route } from "found";
import { Resolver } from "found-relay";
import environment from "./api/environment";
import ExamplePageRoute from "views/ExamplePage";

const Router = createFarceRouter({
  historyProtocol: new BrowserProtocol(),
  historyMiddlewares: [queryMiddleware],
  routeConfig: makeRouteConfig(
    <Route path="/" Component={ExamplePageRoute}></Route>
  ),

  render: createRender({}),
});

const App = () => <Router resolver={new Resolver(environment)} />;

export default App;
