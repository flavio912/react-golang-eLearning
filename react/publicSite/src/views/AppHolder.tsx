import * as React from 'react';
import HeaderMenu, { Tab } from 'components/Menu/HeaderMenu';
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
    { id: 0, title: 'Dashboard' },
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
      />
      <div className={classes.appHolder}>{children}</div>
    </div>
  );
};
