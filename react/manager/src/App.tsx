import * as React from 'react';
import { createUseStyles } from 'react-jss';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect
} from 'react-router-dom';
import { connect } from 'react-redux';

type Props = {
  classes: any;
};

type State = {};

function App() {
  return (
    <div>
      <Router>
        <Switch>
          <Route path='/' exact>
            <Redirect to='/home' />
          </Route>
          <Route path='/login'>
            <p>Login</p>
          </Route>
          <Route>
            <Switch>
              <Route path='/home' exact>
                <p>test</p>
              </Route>
            </Switch>
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default connect(
  () => ({}),
  (dispatch: any) => ({})
)(App);
