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
} from './__generated__/RegisterCompany_CompanyRequestMutation.graphql';
import environment from 'api/environment';
import RegisterCompanyPart2 from 'components/Overview/Registration/RegisterCompanyPart2';
import RegisterCompanySuccess from 'components/Overview/Registration/RegisterCompanySuccess';

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
  hidden: {
    display: 'none',
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
        contactEmail: $email
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

type CompanyState = {
  fname: string;
  lname: string;
  email: string;
  password: string;
  telephone: string;
  companyName: string;
  address1: string;
  address2: string;
  county: string;
  postcode: string;
  country: string;
};

const initState = {
  fname: '',
  lname: '',
  email: '',
  password: '',
  telephone: '',
  companyName: '',
  address1: '',
  address2: '',
  county: '',
  postcode: '',
  country: '',
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
  const [page, setPage] = React.useState<0 | 1 | 2>(0);
  const [companyInfo, setCompanyInfo] = React.useState<CompanyState>(initState);

  const register = async () => {
    try {
      await submitForm({
        companyName: companyInfo?.companyName,
        addr1: companyInfo.address1,
        addr2: companyInfo.address2,
        county: companyInfo.county,
        country: companyInfo.country,
        postcode: companyInfo.postcode,
        firstName: companyInfo.fname,
        lastName: companyInfo.lname,
        email: companyInfo.email,
        password: companyInfo.password,
      });
      setPage(2);
      // Show success popup
    } catch (err) {
      alert(err);
    }
  };

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <div className={page === 0 ? '' : classes.hidden}>
          <RegisterComp
            onNext={() => setPage(1)}
            onChange={(state) => {
              setCompanyInfo({ ...companyInfo, ...state });
            }}
            onLogoClick={() => {
              router.push('/');
            }}
          />
        </div>

        <div className={page === 1 ? '' : classes.hidden}>
          <RegisterCompanyPart2
            onNext={register}
            onChange={(state) => {
              setCompanyInfo({ ...companyInfo, ...state });
            }}
            onLogoClick={() => {
              router.push('/');
            }}
            onBack={() => setPage(0)}
          />
        </div>

        <div className={page === 2 ? '' : classes.hidden}>
          <RegisterCompanySuccess
            onNext={() => {
              router.push('/');
            }}
            onLogoClick={() => {
              router.push('/');
            }}
          />
        </div>
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
