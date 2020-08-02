import * as React from 'react';
import HeaderMenu, { Tab } from 'components/Menu/HeaderMenu';
import { createUseStyles, useTheme } from 'react-jss';
import { useRouter } from 'found';
import { Theme } from 'helpers/theme';
import Footer from 'components/Menu/Footer';
import { delegateLogin } from 'api/config';
import { createFragmentContainer, graphql } from 'react-relay';
import { AppHolder_categories } from './__generated__/AppHolder_categories.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  appHolder: {
    display: 'flex',
    justifyContent: 'center',
    color: theme.colors.primaryBlack,
    paddingTop: 83,
  },
  appHolderRoot: {},
}));

const defaultLink = {
  id: 0,
  name: 'Feature',
  link: '',
};

const defaultNew = {
  id: 1,
  name: 'Feature',
  link: '',
  alert: {
    type: 'new',
    value: 'NEW',
  },
};

const defaultIncrease = {
  id: 2,
  name: 'Feature',
  link: '',
  alert: {
    type: 'increase',
    value: '99.9%',
  },
};

const footerColumns = [
  {
    id: 0,
    header: 'Features',
    links: [
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultNew,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
    ],
  },
  {
    id: 1,
    header: 'Learn More',
    links: [
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
    ],
  },
  {
    id: 2,
    header: 'Company',
    links: [
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
    ],
  },
  {
    id: 3,
    header: 'Get Help',
    links: [
      defaultLink,
      defaultLink,
      defaultLink,
      defaultLink,
      defaultIncrease,
    ],
  },
];

type Props = {
  children?: React.ReactChildren;
  categories: AppHolder_categories;
};

const AppHolder = ({ children, categories }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const { match, router } = useRouter();

  const categoryOptions = (categories.edges ?? []).map((category) => ({
    id: category?.uuid ?? '',
    title: category?.name ?? '',
    text:
      'Training courses specifically designed for those who work in Aviation Security',
    link: `/courses/${category?.uuid}`,
  }));

  const tabs: Array<Tab> = [
    {
      id: 0,
      title: 'Features',
      options: [
        {
          id: 0,
          title: 'For Companies',
          text:
            "We're training the finest minds in air, road and sea - are you on the list?",
          link: '/register/company',
        },
        {
          id: 1,
          title: 'For Individuals',
          text:
            'Do you need trainingsolutions for your next role in Transport Compliance?',
          link: '/register/individual',
        },
      ],
    },
    {
      id: 1,
      title: 'Courses',
      options: [
        {
          id: 0,
          title: 'Online Courses',
          text:
            'Training courses specifically designed for those who work in Aviation Security',
          options: categoryOptions,
        },
        {
          id: 1,
          title: 'Classroom Courses',
          text:
            'All classroom courses are delivered in London at our purpose built facility',
          options: categoryOptions,
        },
      ],
    },
    {
      id: 2,
      title: 'Resources',
      options: [
        {
          id: 0,
          title: 'News & Blog',
          text:
            'Stay in touch with all the industry news from the team at TTC Hub',
          link: '/articles',
        },
        {
          id: 1,
          title: 'Book A Demo',
          text:
            'Looking for more information on our platform? Let us show you the ropes',
          link: '/',
        },
        {
          id: 2,
          title: 'About Us',
          text:
            "We're a growing team of industry experts with 40+ years of experience",
          link: '/aboutus',
        },
        {
          id: 3,
          title: 'Contact Us',
          text:
            "We're a growing team of industry experts with 100+ years of experience",
          link: '/contact',
        },
      ],
    },
    { id: 3, title: 'Consultancy', link: '/consultancy' },
  ];

  const selectedRoute = () => {
    const { routes } = match;
    const currentRouter = routes[routes.length - 1];
    switch (currentRouter.path) {
      case '/':
        return tabs[0];
      default:
        return tabs[0];
    }
  };

  const [selected, setSelected] = React.useState<Tab | undefined>();

  return (
    <div className={classes.appHolderRoot}>
      <HeaderMenu
        tabs={tabs}
        selected={selected}
        setSelected={(tab?: Tab) => setSelected(tab)}
        onClick={(link) => {
          setSelected(undefined);
          router.push(link);
        }}
        onRegisterClick={() => {
          router.push('/register');
        }}
        onLogoClick={() => {
          router.push('/');
        }}
        onCheckout={() => console.log('Checkout')}
        onLoginClick={() => {
          window.location.href = delegateLogin;
        }}
      />
      <div className={classes.appHolder}>{children}</div>
      <Footer columns={footerColumns} />
    </div>
  );
};

export default createFragmentContainer(AppHolder, {
  categories: graphql`
    fragment AppHolder_categories on CategoryPage {
      edges {
        uuid
        name
        color
      }
      pageInfo {
        total
      }
    }
  `,
});
