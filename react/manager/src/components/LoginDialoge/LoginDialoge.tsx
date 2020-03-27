import * as React from 'react';
import Card from '../core/Card';

type Props = {

}

function LoginDialoge(props: Props) {
  return (
    <Card padding="medium">
      <h1>Login to TTC Hub</h1>
      <p>Glad to have you back, please enter your login details to proceed</p>
    </Card>
  )
}

export default LoginDialoge


