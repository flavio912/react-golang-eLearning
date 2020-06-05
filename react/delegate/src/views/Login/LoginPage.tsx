import React from 'react';

import LoginDialogue from 'components/LoginDialogue';
import { Theme } from 'helpers/theme';
import environment from 'api/environment';
import { createUseStyles, useTheme } from 'react-jss';
import { commitMutation, graphql } from 'react-relay';
import { GraphError } from 'types/general';
import { Redirect, RedirectException } from 'found';

type Props = {
  data: any;
};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    background: theme.colors.backgroundGrey,
    width: '100%',
    height: '100%',
    position: 'absolute',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  }
}));

const mutation = graphql`
  mutation LoginPage_LoginMutation($email: String!, $password: String!) {
    managerLogin(input: { email: $email, password: $password, noResp: true }) {
      token
    }
  }
`;

const AttemptLogin = (
  email: string,
  password: string,
  errorCallback: (err: string) => void
) => {
  const variables = {
    email,
    password
  };

  commitMutation(environment, {
    mutation,
    variables,
    onCompleted: (
      response: { managerLogin: { token: string } },
      errors: GraphError[]
    ) => {
      if (errors) {
        // Display error
        errorCallback(`${errors[0]?.extensions?.message}`);
        return;
      }
      console.log('Response received from server.', response, errors);
      window.location.href = '/app';
    },
    onError: (err) => console.error(err)
  });
};

const LoginPage = ({ data }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <>
      {/* <RedirectRequest/> */}
      <div className={classes.root}>
        <LoginDialogue onSubmit={AttemptLogin} />
      </div>
    </>
  );
};

export default LoginPage;
