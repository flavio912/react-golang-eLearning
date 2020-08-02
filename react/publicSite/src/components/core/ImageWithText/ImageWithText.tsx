import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { useRouter } from 'found';
import { Theme } from 'helpers/theme';
import Icon, { IconNames } from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'space-between',
    flexWrap: 'wrap',
    '@media (max-width: 700px)': {
      flexDirection: 'column',
      justifyContent: 'flex-start',
    },
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'flex-start',
    alignItems: 'center',
    marginTop: '15px',
  },
  container: {
    display: 'flex',
    flex: 1,
    flexDirection: 'column',
    justifyContent: 'space-around',
    '@media (max-width: 700px)': {
      alignSelf: 'center',
    },
  },
  title: {
    fontSize: 23,
    fontWeight: 400,
    color: theme.colors.textGrey,
    marginBottom: '15px',
  },
  subtitle: {
    fontSize: 32,
    fontWeight: '800',
    marginBottom: '32px',
  },
  description: {
    fontSize: theme.fontSizes.xSmallHeading,
    fontWeight: 300,
    color: theme.colors.textGrey,
    marginBottom: '25px',
    '@media (max-width: 700px)': {
      textAlign: 'center',
    },
  },
  link: {
    cursor: 'pointer',
    fontSize: theme.fontSizes.xSmallHeading,
    fontWeight: 600,
    color: theme.colors.navyBlue,
    marginBottom: '15px',
  },
  image: {
    height: '457px',
    width: '100%',
  },
  text: {
    fontSize: theme.fontSizes.extraLarge,
    fontWeight: '500',
    color: theme.colors.textGrey,
    marginTop: '10px',
  },
  center: {
    alignItems: 'center',
  },
  maxWidth: {
    maxWidth: '400px',
    marginRight: '100px',
    '@media (max-width: 700px)': {
      marginRight: '0',
      alignItems: 'center',
    },
  },
  rowReverse: {
    flexDirection: 'row-reverse',
    '@media (max-width: 700px)': {
      flexDirection: 'column',
    },
  },
  marginLeft: {
    margin: '0 0 0 100px',
    '@media (max-width: 700px)': {
      margin: 0,
    },
  },
}));

export type Link = {
  title: string;
  link: string;
};

export type Row = {
  iconName: IconNames;
  text: string;
  link: Link;
};

type Props = {
  title?: string;
  subtitle?: string;
  description?: string;
  link?: Link;
  image?: string;
  stack?: Row[];
  textRight?: boolean;
  className?: string;
};

function ImageWithText({
  title,
  subtitle,
  description,
  link,
  image,
  stack,
  textRight,
  className,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const { router } = useRouter();
  const onClick = (link?: string) => {
    link && router.push(link);
  };

  return (
    <div
      className={classNames(
        classes.root,
        textRight && classes.rowReverse,
        className,
      )}
    >
      {stack ? (
        <div
          className={classNames(
            classes.container,
            classes.maxWidth,
            textRight && classes.marginLeft,
          )}
        >
          {stack.map((row: Row) => (
            <div className={classes.row}>
              <Icon
                name={row.iconName}
                size={68}
                style={{ marginRight: '35px' }}
              />
              <div onClick={() => onClick(row.link?.link)}>
                <div className={classes.link}>{row.link?.title}</div>
                <div className={classes.text}>{row.text}</div>
              </div>
            </div>
          ))}
        </div>
      ) : (
        <div
          className={classNames(
            classes.container,
            classes.maxWidth,
            textRight && classes.marginLeft,
          )}
        >
          <div style={{ flex: 0.1 }} />
          <div className={classes.title}>{title}</div>
          <div className={classes.subtitle}>{subtitle}</div>
          <div className={classes.description}>{description}</div>
          <div className={classes.link} onClick={() => onClick(link?.link)}>
            {link?.title}
            {link && (
              <Icon
                name="ArrowRightNavyBlue"
                size={10}
                style={{ margin: '0 0 1px 5px' }}
              />
            )}
          </div>
          <div style={{ flex: 0.1 }} />
        </div>
      )}
      <div className={classNames(classes.container, classes.center)}>
        <img className={classes.image} src={image} />
      </div>
    </div>
  );
}

export default ImageWithText;
