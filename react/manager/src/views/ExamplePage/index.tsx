import React from "react";
import { Route, RouteRenderArgs } from 'found';
import { graphql } from "react-relay";

import ExamplePage from "./ExamplePage";

const ExamplePageQuery = graphql`
  query ExamplePage_Query {
    ...ExampleComponent_info
  }
`

const ExamplePageRoute = () => (
  <Route
    path="/"
    Component={ExamplePage}
    query={ExamplePageQuery}
    render={({ props }: RouteRenderArgs) => {
      console.log(props)
    return (
      <ExamplePage data={props} />
    )}}
  />
);

export default ExamplePageRoute;
