import * as React from "react";
import { createUseStyles } from "react-jss";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect,
} from "react-router-dom";

import { graphql, QueryRenderer } from "react-relay";

import environment from "./api/environment";
import ExamplePage from "views/ExamplePage";

type Props = {
  classes: any;
};

type State = {};

type RouteProps = {
  children: React.ReactNode;
  isAuthenticated: boolean;
  path: string;
  exact: boolean;
};

const PrivateRoute = ({
  children,
  isAuthenticated,
  path,
  exact,
}: RouteProps) => (
  <Route
    path={path}
    exact={exact}
    render={({ location }) =>
      isAuthenticated ? (
        children
      ) : (
        <Redirect
          to={{
            pathname: "/login",
            state: { from: location },
          }}
        />
      )
    }
  />
);

type Response = {
  error: any;
  props: any;
};

const userID = 1;

const App = () => (
  <Router>
    <Switch>
      <Route path="/" exact>
        <ExamplePage />
      </Route>
      <Route path="/login">
        <p>Login</p>
      </Route>
      <PrivateRoute isAuthenticated={false} path="/home" exact>
        <p>test</p>
      </PrivateRoute>
    </Switch>
  </Router>
);

export default App;
