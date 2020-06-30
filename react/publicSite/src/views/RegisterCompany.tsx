import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import RegisterComp from 'components/Overview/Registration/RegisterCompany';
import { Router } from 'found';
const useStyles = createUseStyles((theme: Theme) => ({
  registerRoot: {
    display: 'grid',
    minHeight: '100%',
    background: 'white',
    gridTemplateColumns: '600px 1fr',
    '@media (max-width: 800px)': {
      gridTemplateColumns: '1fr'
    }
  },
  fancyBackground: {
    background: theme.loginBackgroundGradient
  },
  picker: {
    padding: 48,
    background: 'white'
  }
}));

type Props = {
  router: Router;
};

function RegisterCompany({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <RegisterComp
          onSubmit={() => {}}
          onLogoClick={() => {
            router.push('/');
          }}
        />
      </div>
      <div className={classes.fancyBackground}></div>
    </div>
  );
}

export default RegisterCompany;