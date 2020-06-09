import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import Button, { ButtonProps } from 'components/core/Input/Button';
import FlatCard from 'components/core/Cards/FlatCard';
import classnames from 'classnames';

const useStyles = createUseStyles((theme: Theme) => ({
  cardItemRoot: {
    border: `1px solid ${theme.colors.approxZircon}`,
    flexDirection: 'column',
    position: 'relative',
    '&:after': {
      content: "''",
      display: 'block',
      border: `1px solid ${theme.colors.approxZircon}`,
      position: 'absolute',
      width: 'calc(100% - 10px)',
      height: 5,
      borderRadius: 6,
      bottom: -5,
      left: 5,
      zIndex: -1
    },
    '&:before': {
      content: "''",
      display: 'block',
      border: `1px solid ${theme.colors.approxZircon}`,
      position: 'absolute',
      width: 'calc(100% - 20px)',
      height: 3,
      borderRadius: 6,
      bottom: -9,
      left: 10,
      zIndex: -1
    }
  },
  cardItemTitle: {
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.heading,
    letterSpacing: -0.63,
    fontWeight: 800,
    margin: [0, 0, 23]
  },
  cardItemDescription: {
    color: theme.colors.secondaryBlack,
    fontSize: theme.fontSizes.extraLarge,
    letterSpacing: -0.45,
    marginBottom: 31,
    fontWeight: 300,
    marginTop: 0,
    lineHeight: `30px`
  },
  buttonWrapper: {
    display: 'inline-block'
  },
  button: {
    fontSize: theme.fontSizes.extraLarge,
    letterSpacing: -0.45,
    fontWeight: 'bold'
  }
}));

type Props = {
  title: string;
  description?: string;
  buttonProps: ButtonProps;
  className?: string;
};

function CardItem({ title, description, buttonProps, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <FlatCard
      shadow={false}
      className={classnames(classes.cardItemRoot, className)}
      padding={'large'}
    >
      <h6 className={classes.cardItemTitle}>{title}</h6>
      <p className={classes.cardItemDescription}>{description}</p>
      <div className={classes.buttonWrapper}>
        <Button
          {...buttonProps}
          padding={'medium'}
          className={classes.button}
        />
      </div>
    </FlatCard>
  );
}

export default CardItem;
