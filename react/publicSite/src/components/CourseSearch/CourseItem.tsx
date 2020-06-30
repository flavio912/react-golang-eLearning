import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Button from 'sharedComponents/core/Input/Button';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    backgroundColor: theme.colors.primaryWhite,
    boxShadow: '4px 2px 10px -2px rgba(0,0,0,0.06)',
    border: ['1px', 'solid', theme.colors.borderGrey],
    borderRadius: '13px',
    padding: '24px',
    '@media (max-width: 700px)': {
      flexDirection: 'column',
      maxWidth: '351px'
    }
  },
  title: {
    fontSize: theme.fontSizes.heading,
    fontWeight: '800',
    color: theme.colors.primaryBlack,
    '@media (max-width: 700px)': {
      marginBottom: '10px'
    }
  },
  description: {
    fontSize: theme.fontSizes.default,
    fontWeight: '500',
    color: theme.colors.textGrey,
    maxWidth: '600px',
    marginRight: '20px',
    '@media (max-width: 700px)': {
      margin: '0 0 20px 0'
    }
  },
  price: {
    fontSize: theme.fontSizes.heading,
    fontWeight: '800',
    color: theme.colors.primaryBlue,
    '@media (max-width: 700px)': {
      marginBottom: '20px'
    }
  },
  button: {
    height: '52px',
    width: '182px',
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)',
    marginRight: '20px',
    fontSize: theme.fontSizes.large,
    fontWeight: '800',
    '@media (max-width: 700px)': {
      width: '100%',
      margin: '0 0 20px 0'
    }
  },
  extraShadow: {
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)'
  },
  image: {
    order: 1,
    display: 'flex',
    height: '260px',
    width: '300px',
    borderRadius: theme.secondaryBorderRadius,
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover'
  },
  type: {
    alignSelf: 'flex-start',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    fontSize: theme.fontSizes.xSmall,
    fontWeight: '700',
    color: theme.colors.primaryWhite,
    borderRadius: `${theme.secondaryBorderRadius}px 0 ${theme.secondaryBorderRadius}px 0`,
    padding: `${theme.spacing(1)}px ${theme.spacing(2)}px`
  },
  column: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    '@media (max-width: 700px)': {
      order: 2,
      marginTop: '25px'
    }
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'wrap',
    justifyContent: 'center',
    alignItems: 'center',
    '@media (max-width: 700px)': {
      alignSelf: 'center'
    }
  }
}));

export type CourseProps = {
  title: string;
  description: string;
  price: string;
  type: string;
  colour: string;
  imageURL: string;
  viewCourse: () => void;
  addToBasket: () => void;
  className?: string;
};

function CourseItem({
  title,
  description,
  price,
  type,
  colour,
  imageURL,
  viewCourse,
  addToBasket,
  className
}: CourseProps) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const backgroundColor = { backgroundColor: colour };
  const backgroundImage = {};
  if (imageURL)
    backgroundImage[
      'backgroundImage'
    ] = `linear-gradient(#00000040, #00000040), url(${imageURL})`;

  return (
    <div className={classNames(classes.root, className)}>
      <div className={classes.column}>
        <div className={classes.title}>{title}</div>
        <div className={classes.description}>{description}</div>
        <div className={classes.price}>{price}</div>
        <div className={classes.row}>
          <Button
            archetype="submit"
            onClick={viewCourse}
            className={classNames(classes.button, classes.extraShadow)}
          >
            View Course
          </Button>
          <Button onClick={addToBasket} className={classes.button}>
            Add to Basket
          </Button>
        </div>
      </div>
      <div className={classes.image} style={backgroundImage}>
        <div className={classes.type} style={backgroundColor}>
          {type.toUpperCase()}
        </div>
      </div>
    </div>
  );
}

export default CourseItem;
