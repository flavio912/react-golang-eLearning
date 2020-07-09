import * as React from 'react';
import HeaderMenu, { Tab } from 'components/Menu/HeaderMenu';
import { createUseStyles, useTheme } from 'react-jss';
import { useRouter } from 'found';
import { Theme } from 'helpers/theme';
import Footer from 'components/Menu/Footer';

type Props = {
  children?: React.ReactChildren;
};

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

export const AppHolder = ({ children }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const { match, router } = useRouter();
  const tabs: Array<Tab> = [
    {
      id: 0,
      title: 'Features',
      options: [
        {
          title: 'For Companies',
          text:
            "We're training the finest minds in air, road and sea - are you on the list?",
          link: '/register/company',
        },
        {
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
          title: 'Aviation Security',
          text:
            'Training courses specifically designed for those who work in Aviation Security',
          link: '/courses',
        },
        {
          title: 'Dangerous Goods',
          text:
            'Courses for both air and road, all in accordance with CAA Regulations',
          link: '/',
        },
        {
          title: 'Health & Safety',
          text:
            'All our courses can be taken online in conjunction withyour internal policies',
          link: '/',
        },
        {
          title: 'Classroom Courses',
          text:
            'All classroom courses are delivered in London at our purpose built facility',
          link: '/',
        },
      ],
    },
    {
      id: 2,
      title: 'Resources',
      options: [
        {
          title: 'News & Blog',
          text:
            'Stay in touch with all the industry news from the team at TTC Hub',
          link: '/articles',
        },
        {
          title: 'Book A Demo',
          text:
            'Looking for more information on our platform? Let us show you the ropes',
          link: '/',
        },
        {
          title: 'About Us',
          text:
            "We're a growing team of industry experts with 40+ years of experience",
          link: '/aboutus',
        },
        {
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
      />
      <div className={classes.appHolder}>{children}</div>
      <Footer columns={footerColumns} />
    </div>
  );
};
