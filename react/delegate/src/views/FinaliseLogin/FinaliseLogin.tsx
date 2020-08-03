import React from 'react';

import FinaliseDialogue from 'components/LoginDialogue/FinaliseDialogue';
import { Theme } from 'helpers/theme';
import environment from 'api/environment';
import { createUseStyles, useTheme } from 'react-jss';
import { commitMutation, graphql } from 'react-relay';
import { GraphError } from 'types/general';
import { Router, useRouter } from 'found';
import jwt from 'jsonwebtoken';

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
  mutation FinaliseLogin_FinaliseMutation($token: String!, $password: String!) {
    finaliseDelegate(input: { token: $token, password: $password }) {
      token
    }
  }
`;

const AttemptLogin = (router: Router, token: String) => (
  password: string,
  passwordRepeat: string,
  errorCallback: (err: string) => void
) => {
  const variables = {
    token,
    password
  };

  if (!password || !passwordRepeat) {
    return errorCallback('Please enter password twice');
  }
  if (password !== passwordRepeat) {
    return errorCallback('Passwords are not the same');
  }

  commitMutation(environment, {
    mutation,
    variables,
    onCompleted: (
      response: { delegateLogin: { token: string } },
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

const FinaliseLogin = ({ data }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router, match } = useRouter();
  const { token } = match?.params;
  const decoded = jwt.decode(token, { json: true });
  const { TTC_ID } = decoded ? decoded.claims : { TTC_ID: 'Invalid TTC ID' };

  return (
    <>
      {/* <RedirectRequest/> */}
      <div className={classes.root}>
        <FinaliseDialogue
          TTC_ID={TTC_ID}
          onSubmit={AttemptLogin(router, token)}
        />
      </div>
    </>
  );
};

export default FinaliseLogin;
