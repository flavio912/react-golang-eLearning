import * as React from 'react';
import { createUseStyles } from 'react-jss';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect
} from 'react-router-dom';

import {graphql, QueryRenderer} from 'react-relay';

import environment from './api/environment';

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

const PrivateRoute = ({ children, isAuthenticated, path, exact }: RouteProps) => (
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
            state: { from: location }
          }}
        />
      )
    }
  />
);

type Response = {
  error: any,
  props: any,
};

const App = () => (
  <QueryRenderer
      environment={environment}
      query={graphql`
        query UserQuery {
          user {
            id
          }  
        }
      `}
      variables={{}}
      render={({error, props}: Response) => {
        if (error) {
          return <div>Error!</div>;
        }
        if (!props) {
          return <div>Loading...</div>;
        }
        return (
          <Router>
            <Switch>
              <Route path='/' exact></Route>
              <Route path='/login'>
                <p>Login</p>
              </Route>
              <Route>
              <PrivateRoute isAuthenticated={props.viewer.id === 'TEST'} path='/home' exact>
                <p>test</p>
              </PrivateRoute>
              </Route>
            </Switch>
          </Router>
        );
      }}
    />
);

export default App;
