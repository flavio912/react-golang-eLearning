import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import TrustedCard from 'components/core/Cards/TrustedCard';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    height: '100vh',
    alignItems: 'center',
    justifyContent: 'space-evenly'
  },
  title: {
    fontSize: '37px',
    fontWeight: '800'
  },
  subtitle: {
    fontSize: theme.fontSizes.tinyHeading,
    color: theme.colors.textGrey
  },
  image: {
    height: '315px',
    width: '354px'
  },
  imageTitle: {
    fontSize: '28px',
    fontWeight: '800',
    color: theme.colors.navyBlue
  },
  imageSubtitle: {
    fontSize: theme.fontSizes.large,
    color: theme.colors.textGrey
  },
  center: {
    textAlign: 'center',
    maxWidth: '450px'
  }
}));

type Props = {
  title: string;
  subtitle: string;
  imageURL: string;
  imageTitle: string;
  imageSubtitle: string;
  className?: string;
};

function InfoCard({
  title,
  subtitle,
  imageURL,
  imageTitle,
  imageSubtitle,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classNames(classes.root, className)}>
      <Icon name="TTC_Logo" style={{ height: '44px', width: '130px' }} />
      <div className={classNames(classes.title, classes.center)}>{title}</div>
      <div className={classNames(classes.subtitle, classes.center)}>
        {subtitle}
      </div>
      <img className={classes.image} src={imageURL} />
      <div className={classNames(classes.imageTitle, classes.center)}>
        {imageTitle}
      </div>
      <div className={classNames(classes.imageSubtitle, classes.center)}>
        {imageSubtitle}
      </div>
      <TrustedCard padding="none" noShadow className={classes.center} />
    </div>
  );
}

export default InfoCard;
