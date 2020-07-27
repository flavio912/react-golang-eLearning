import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import { createFragmentContainer, graphql, commitMutation } from 'react-relay';
import RegisterComp from 'components/Overview/Registration/RegisterCompany';
import { Router } from 'found';
import RegistrationCarousel from 'components/Overview/Registration/RegistrationCarousel';
import { Image } from 'components/Misc/CarouselImage';
import { GraphError } from 'types/general';
import {
  RegisterCompany_CompanyRequestMutationVariables,
  RegisterCompany_CompanyRequestMutationResponse,
} from '__generated__/RegisterCompany_CompanyRequestMutation.graphql';
import environment from 'api/environment';

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
  mutation RegisterCompany_CompanyRequestMutation(
    $companyName: String!
    $addr1: String!
    $addr2: String!
    $county: String!
    $postcode: String!
    $country: String!
    $firstName: String!
    $lastName: String!
    $email: String!
    $password: String!
  ) {
    createCompanyRequest(
      company: {
        companyName: $companyName
        addressLine1: $addr1
        addressLine2: $addr2
        county: $county
        postCode: $postcode
        country: $country
      }
      manager: {
        firstName: $firstName
        lastName: $lastName
        email: $email
        jobTitle: ""
        telephone: ""
        password: $password
      }
      recaptcha: ""
    )
  }
`;

const submitForm = (
  variables: RegisterCompany_CompanyRequestMutationVariables,
) => {
  return new Promise<RegisterCompany_CompanyRequestMutationResponse>(
    (resolve, reject) => {
      commitMutation(environment, {
        mutation,
        variables,
        onCompleted: (
          response: RegisterCompany_CompanyRequestMutationResponse,
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

function RegisterCompany({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const images = [1, 2, 3].map((item) => ({
    ...defaultImage,
    alt: `${defaultImage.alt} ${item}`,
  }));

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <RegisterComp
          onSubmit={async (
            firstName,
            lastName,
            email,
            password,
            telephone,
            companyName,
          ) => {
            try {
              await submitForm({
                companyName,
                addr1: '',
                addr2: '',
                county: '',
                country: '',
                postcode: '',
                firstName,
                lastName,
                email,
                password,
              });
              // Show success popup
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

export default RegisterCompany;
