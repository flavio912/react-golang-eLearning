import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import FancyBorder from '../../components/Certificate/FancyBorder';
import CertInfo from '../../components/Certificate/CertInfo';
import { ReactComponent as Logo } from '../../assets/logo/ttc-logo.svg';

const useStyles = createUseStyles((theme: Theme) => ({
  certGeneratorRoot: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
  },
  logo: {
    height: 70
  },
}));

type Props = {
  className?: string;
  certLogo: string;
};

function CertGenerator({ className, certLogo }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.certGeneratorRoot}>
      <FancyBorder>
        <div>

          <Logo className={classes.logo} />          
        </div>
        
        <CertInfo 
          certName="General Security Awareness Training (GSAT)"
          moduleDeliver="Module Delivered: 1-20"
          forEu="For EU 1998/2015: 11.2.2"
          certNo="000000054321"
          trainingDate="DD MMM YYYY"
          expiryDate="DD MMM YYYY"
        />
      </FancyBorder>
    </div>
  )
}

export default CertGenerator;
