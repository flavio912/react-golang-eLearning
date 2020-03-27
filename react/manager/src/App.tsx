import * as React from 'react';
import { createUseStyles } from 'react-jss';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect
} from 'react-router-dom';

type Props = {
  classes: any;
};

type State = {};

function App() {
  return (
    <div>
      {/* <Router>
        <Switch>
          <Route path='/' exact></Route>
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
      </Router> */}
    </div>
  );
}

export default App;
