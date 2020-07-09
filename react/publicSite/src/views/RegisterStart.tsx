import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import ImageWithText from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import AccountTypeCard from 'components/Overview/Registration/AccountTypeCard';
import { Router } from 'found';
import { Image } from 'components/Misc/CarouselImage';
import RegistrationCarousel from 'components/Overview/Registration/RegistrationCarousel';

const useStyles = createUseStyles((theme: Theme) => ({
  registerRoot: {
    display: 'grid',
    minHeight: '100%',
    background: 'white',
    gridTemplateColumns: '600px 1fr',
    '@media (max-width: 1000px)': {
      gridTemplateColumns: '1fr'
    }
  },
  fancyBackground: {
    background: theme.loginBackgroundGradient
  },
  picker: {
    padding: 48
  },
  carousel: {

  }
}));

const defaultImage: Image = {
  url: require('assets/carouselImage.svg'),
  alt: 'Image'
};

type Props = {
  router: Router;
};

function RegisterStart({ router }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const images = [1, 2, 3].map((item) => ({
    ...defaultImage,
    alt: `${defaultImage.alt} ${item}`
  }));

  return (
    <div className={classes.registerRoot}>
      <div className={classes.picker}>
        <AccountTypeCard
          onIndividual={() => {
            router.push('/register/individual');
          }}
          onCompany={() => {
            router.push('/register/company');
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

export default RegisterStart;
