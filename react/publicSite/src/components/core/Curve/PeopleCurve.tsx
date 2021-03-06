import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import ImageWithText, { Row } from 'components/core/ImageWithText';
import PageMargin from '../PageMargin';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    width: '100%',
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center',
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
    textAlign: 'center',
  },
  text: {
    fontSize: theme.fontSizes.extraLarge,
    color: theme.colors.textGrey,
    fontWeight: 500,
    textAlign: 'center',
    maxWidth: '750px',
  },
  margin: {
    margin: '80px 0',
  },
  curve: {
    height: '100px',
    maxWidth: '100vw',
    marginBottom: '-16px', // Weird gap otherwise
    '@media (max-width: 800px) and (min-width: 550px)': {
      marginBottom: '-5%',
    },
    '@media (max-width: 550px)': {
      marginBottom: '-40px',
    },
  },
  greyBackground: {
    backgroundColor: theme.colors.backgroundGrey,
  },
}));

const defaultStack: Row[] = [
  {
    iconName: 'CourseCertificates',
    text:
      'All of our friendly and knowledgable team are available via email and live chat.',
    link: { title: 'World Class 24x7 Support', link: '/' },
  },
  {
    iconName: 'CourseCertificates',
    text:
      'Stay tuned for regular webinars and live QA sessions with the TTC team.',
    link: { title: 'Webinars and Live Sessions', link: '/' },
  },
  {
    iconName: 'CourseCertificates',
    text:
      'Got a question that needs an immediate answer? Try our knowledge base.',
    link: { title: 'Knowledge Base', link: '/' },
  },
];

type Props = {
  stack?: Row[];
  className?: string;
};

function PeopleCurve({ stack = defaultStack, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
      <img className={classes.curve} src={require('assets/grey-curve.svg')} />
      <PageMargin
        centererStyle={classNames(classes.centerer, classes.greyBackground)}
        centeredStyle={classes.centered}
      >
        <div className={classes.heading}>Our people make the difference.</div>
        <div className={classes.text}>
          Not only do we offer incredible training, but our customer service is
          world-class too
        </div>
        <ImageWithText
          className={classes.margin}
          image={require('assets/StockUKTeam.svg')}
          stack={stack}
        />
      </PageMargin>
    </div>
  );
}

export default PeopleCurve;
