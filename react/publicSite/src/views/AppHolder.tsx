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
    paddingTop: 83
  },
  appHolderRoot: {}
}));

const defaultLink = {
  id: 0,
  name: 'Feature',
  link: ''
};

const defaultNew = {
  id: 1,
  name: 'Feature',
  link: '',
  alert: {
    type: 'new',
    value: 'NEW'
  }
};

const defaultIncrease = {
  id: 2,
  name: 'Feature',
  link: '',
  alert: {
    type: 'increase',
    value: '99.9%'
  }
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
      defaultLink
    ]
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
      defaultLink
    ]
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
      defaultLink
    ]
  },
  {
    id: 3,
    header: 'Get Help',
    links: [defaultLink, defaultLink, defaultLink, defaultLink, defaultIncrease]
  }
];

export const AppHolder = ({ children }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const { match, router } = useRouter();
  const tabs: Tab[] = [
    { id: 0, title: 'Features', options: ['Some', 'Different', 'Options'] },
    { id: 1, title: 'Courses', options: ['Some', 'Different', 'Options'] },
    { id: 2, title: 'Resources', options: ['Some', 'Different', 'Options'] },
    { id: 3, title: 'Consultancy' }
  ];

  const selected = () => {
    const { routes } = match;
    const currentRouter = routes[routes.length - 1];
    switch (currentRouter.path) {
      case '/':
        return tabs[0];
      default:
        return tabs[0];
    }
  };

  return (
    <div className={classes.appHolderRoot}>
      <HeaderMenu
        tabs={tabs}
        selected={selected()}
        onClick={(tab) => {
          switch (tab.id) {
            case 0:
              router.push('/app');
              break;
            default:
              break;
          }
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
