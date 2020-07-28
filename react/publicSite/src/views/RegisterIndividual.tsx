import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import RegisterInd from 'components/Overview/Registration/RegisterIndividual';
import { Router } from 'found';
import RegistrationCarousel from 'components/Overview/Registration/RegistrationCarousel';
import { graphql, commitMutation } from 'react-relay';
import { Image } from 'components/Misc/CarouselImage';
import environment from 'api/environment';
import { GraphError } from 'types/general';
import {
  RegisterIndividual_CreateIndividualMutationVariables,
  RegisterIndividual_CreateIndividualMutationResponse,
} from '__generated__/RegisterIndividual_CreateIndividualMutation.graphql';
import { delegateLogin } from 'api/config';

const useStyles = createUseStyles((theme: Theme) => ({
  registerRoot: {
    display: 'grid',
    minHeight: '100%',
    background: 'white',
    gridTemplateColumns: '600px 1fr',
    '@media (max-width: 1000px)': {
      gridTemplateColumns: '1fr',
    },
  },
  fancyBackground: {
    background: theme.loginBackgroundGradient,
  },
  picker: {
    padding: 48,
    background: 'white',
  },
}));

const defaultImage: Image = {
  url: require('assets/carouselImage.svg'),
  alt: 'Image',
};

const mutation = graphql`
  mutation RegisterIndividual_CreateIndividualMutation(
    $firstName: String!
    $lastName: String!
    $email: String!
    $password: String!
  ) {
    createIndividual(
      input: {
        firstName: $firstName
        lastName: $lastName
        email: $email
        password: $password
      }
    ) {
      user {
        email
      }
    }
  }
`;

const submitForm = (
  variables: RegisterIndividual_CreateIndividualMutationVariables,
) => {
  return new Promise<RegisterIndividual_CreateIndividualMutationResponse>(
    (resolve, reject) => {
      commitMutation(environment, {
        mutation,
        variables,
        onCompleted: (
          response: RegisterIndividual_CreateIndividualMutationResponse,
          errors: GraphError[],
        ) => {
          if (errors) {
            // Display error
            reject(`${errors[0]?.extensions?.message}`);
            return;
          }
          resolve(response);
        },
        onError: (err) => {
          reject(err);
        },
      });
    },
  );
};

type Props = {
  router: Router;
};

function RegisterIndividual({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const images = [1, 2, 3].map((item) => ({
    ...defaultImage,
    alt: `${defaultImage.alt} ${item}`,
  }));

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <RegisterInd
          onSubmit={async (fname, lname, email, password, telephone) => {
            try {
              const resp = await submitForm({
                firstName: fname,
                lastName: lname,
                email: email,
                password: password,
              });

              window.location.href = delegateLogin;
            } catch (err) {
              alert(err);
            }
          }}
          onLogoClick={() => {
            router.push('/');
          }}
        />
      </div>
      <div className={classes.fancyBackground}>
        <RegistrationCarousel
          onBook={() => console.log('Book')}
          images={images}
        />
      </div>
    </div>
  );
}

export default RegisterIndividual;
