import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import ImageWithText, { Row } from 'components/core/ImageWithText';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    width: '100%'
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center'
  },
  centered: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    width: theme.centerColumnWidth,
  },
  heading: {
    fontSize: theme.fontSizes.extraLargeHeading,
    color: theme.colors.primaryBlack,
    fontWeight: 800,
    margin: '75px 0 20px 0',
    textAlign: 'center'
  },
  text: {
    fontSize: theme.fontSizes.extraLarge,
    color: theme.colors.textGrey,
    fontWeight: 500,
    textAlign: 'center',
    maxWidth: '750px'
  },
  margin: {
    margin: '80px 0',
  },
  curve: {
    height: '100px',
    maxWidth: '100vw',
    marginBottom: '-4px' // Weird gap otherwise
  },
  greyBackground: {
      backgroundColor: theme.colors.backgroundGrey
  }
}));

type Props = {
    stack: Row[];
    className?: string;
};

function PeopleCurve({ stack, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
      <div className={classNames(classes.root, className)}>
        <img className={classes.curve} src={require('assets/grey-curve.svg')} />
        <div className={classNames(classes.centerer, classes.greyBackground)}>
            <div className={classes.centered}>   
                <div className={classes.heading}>Our people make the difference.</div>
                <div className={classes.text}>Not only do we offer incredible training, but our customer service is world-class too</div>
                <ImageWithText
                    className={classes.margin}
                    image={require("assets/StockUKTeam.svg")}
                    stack={stack}
                />
            </div>
        </div>
      </div>
  );
}

export default PeopleCurve;