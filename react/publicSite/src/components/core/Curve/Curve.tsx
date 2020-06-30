import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { useRouter } from 'found';
import { Theme } from "helpers/theme";
import blueCurve from '../../../assets/blue-curve.svg';
import curve from '../../../assets/curve.svg';
import Icon, { IconNames } from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    backgroundColor: 'transparent'
  },
  curve: {
    backgroundPosition: 'center',
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    width: '100%'
  },
  children: {
    display: 'flex',
    justifyContent: 'center',
    marginTop: '125px'
  },
  contents: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'flex-start',
    margin: '50px 60px 20px 60px',
    '@media (max-width: 1000px)': {
      margin: '5% 2% 2% 6%',
    }
  },
  row: {
    marginTop: '35px',
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    width: '100%'
  },
  description: {
    fontSize: theme.fontSizes.large,
    fontWeight: '600',
    color: theme.colors.primaryWhite,
    maxWidth: '455px'
  },
  link: {
    cursor: 'pointer',
    fontSize: theme.fontSizes.default,
    fontWeight: 'bold',
    color: theme.colors.primaryWhite,
    textDecorationLine: 'underline'
  },
  black: {
    color: theme.colors.primaryBlack
  }
}));

export type Link ={
  title: string;
  link: string;
}

type Props = {
  height?: string | number;
  width?: string | number;
  logo?: IconNames;
  description?: string;
  link?: Link;
  blue?: boolean;
  children?: React.ReactNode;
  className?: string;
};

function Curve({ height, width, logo, description, link, blue, children, className }: Props) {
  const theme = useTheme();
  const classes = useStyles( theme );

  const { router } = useRouter();
  const onClick = (link?: string) => {
      link && router.push(link);
  }

  return (
      <div
        className={classNames(classes.root, classes.curve, className)}
        style={{ backgroundImage: `url(${blue ? blueCurve : curve})`, height, maxWidth: width }}
      >
          {children ? (
            <div className={classes.children}>
              {children}
            </div>
          ) : (
            <div className={classes.contents}>
              {logo && <Icon name={logo} style={{ width: '155px' }} />}
              <div className={classes.row}>
                <div className={classNames(classes.description, !blue && classes.black)}>{description}</div>
                <div
                  className={classNames(classes.link, !blue && classes.black)}
                  onClick={() => onClick(link?.link)}
                >
                  {link?.title}
                </div>
              </div>
            </div>
          )}
      </div>
  );
}

export default Curve;