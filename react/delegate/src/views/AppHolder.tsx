import * as React from 'react';
import HeaderMenu from 'components/Menu/HeaderMenu';
import SideMenu from 'components/Menu/SideMenu';
import { Tab } from 'components/Menu/SideMenu/SideMenu';
import { createUseStyles, useTheme } from 'react-jss';
import { useRouter } from 'found';

type Props = {
  children?: React.ReactChildren;
};

const useStyles = createUseStyles(() => ({
  appHolder: {
    display: 'flex',
    padding: '42px 60px',
    justifyContent: 'center'
  },
  appHolderRoot: {
    display: 'grid',
    height: '100vh',
    gridTemplateColumns: 'auto 1fr',
    gridTemplateRows: '82px auto'
  }
}));

export const AppHolder = ({ children }: Props) => {
  const classes = useStyles();
  const { match, router } = useRouter();
  const tabs: Tab[] = [
    { id: 0, icon: 'LeftNav_Icon_Dashboard', title: 'Dashboard' },
    { id: 1, icon: 'LeftNav_Icon_Courses', title: 'Online Courses', size: 23 }
  ];

  const selected = () => {
    const { routes } = match;
    const currentRouter = routes[routes.length - 1];
    console.log(currentRouter.path);
    switch (currentRouter.path) {
      case '/':
        return tabs[0];
      case '/courses':
        return tabs[1];
      case '/courses/:id':
        return tabs[1];
      default:
        return tabs[0];
    }
  };
  return (
    <div className={classes.appHolderRoot}>
      <SideMenu
        tabs={tabs}
        selected={selected()}
        logo={require('../assets/logo/ttc-logo.svg')}
        onClick={(tab) => {
          switch (tab.id) {
            case 0:
              router.push('/app');
              break;
            case 1:
              router.push('/app/courses');
              break;
            default:
              break;
          }
        }}
      />
      <HeaderMenu user={{ name: 'James Smith', url: '' }} />
      <div className={classes.appHolder}>{children}</div>
    </div>
  );
};
