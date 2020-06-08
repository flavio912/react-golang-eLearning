import React from 'react';

import LoginDialogue from 'components/LoginDialogue';
import { Theme } from 'helpers/theme';
import environment from 'api/environment';
import { createUseStyles, useTheme } from 'react-jss';
import { commitMutation, graphql } from 'react-relay';
import { GraphError } from 'types/general';
import { useRouter, Router } from 'found';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    background: theme.loginBackgroundGradient,
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
    managerLogin(input: { email: $email, password: $password }) {
      token
    }
  }
`;

const AttemptLogin = (router: Router) => (
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
    onCompleted: async (
      response: { managerLogin: { token: string } },
      errors: GraphError[]
    ) => {
      if (errors) {
        // Display error
        errorCallback(`${errors[0]?.extensions?.message}`);
        return;
      }

      console.log('Response received from server.', response, errors);
      router.push('/app');
    },
    onError: (err) => console.error(err)
  });
};

const LoginPage = () => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  return (
    <>
      {/* <RedirectRequest/> */}
      <div className={classes.root}>
        <LoginDialogue onSubmit={AttemptLogin(router)} />
      </div>
    </>
  );
};

export default LoginPage;
