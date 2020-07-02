import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Button from 'sharedComponents/core/Input/Button';
import Icon from 'sharedComponents/core/Icon';
import VideoPlayer from '../VideoPlayer';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    width: '100%',
    display: 'flex',
    backgroundColor: theme.colors.lightBlue,
    padding: '57px 0',
    '@media (max-width: 850px)': {
        flexDirection: 'column',
        alignItems: 'center',
        backgroundColor: theme.colors.primaryWhite,
        padding: 0
    }
  },
  courseDetails: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    flex: 8,
    '@media (max-width: 850px)': {
        flex: 1,
        maxWidth: '850px',
        justifyContent: 'center',
    },
  },
  history: {
    display: 'flex',
    width: '100%',
    justifyContent: 'flex-start',
    marginBottom: '50px',
    '@media (max-width: 850px)': {
        justifyContent: 'center',
        backgroundColor: theme.colors.lightBlue,
        padding: '50px 0',
        marginBottom: '0px',
    }
  },
  courseTitle: {
    fontSize: theme.fontSizes.heading,
    color: theme.colors.primaryBlack,
    fontWeight: '800',
    paddingBottom: 30
  },
  courseDesc: {
    fontSize: theme.fontSizes.tinyHeading,
    color: theme.colors.footerBlue,
    fontWeight: '500',
    marginTop: '10px',
    maxWidth: '950px',
  },
  updatedText: {
    marginLeft: '25px'
  },
  extraLarge: {
    fontSize: theme.fontSizes.extraLarge,
    color: theme.colors.primaryBlack,
  },
  bold: {
    fontWeight: 'bold',
    marginLeft: '3px'
  },
  price: {
    fontSize: theme.fontSizes.smallHeading,
    color: theme.colors.primaryBlack,
    fontWeight: '600',
    margin: '30px 25px',
    '@media (min-width: 850px)': {
        display: 'none'
    },
  },
  vat: {
    fontSize: theme.fontSizes.large,
    fontWeight: '850',
  },
  button: {
    fontSize: theme.fontSizes.extraLarge,
    fontWeight: '800',
    margin: '15px',
    height: '52px',
    width: '327px',
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.43)',
    '@media (min-width: 850px)': {
        display: 'none'
    },
  },
  shadow: {
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)'
  },
  component: {
    position: 'absolute',
    top: '21vh',
    right: '10%',
    '@media (min-width: 1250px) and (max-width: 1450px)': {
        right: '5%'
    },
    '@media (min-width: 850px) and (max-width: 1250px)': {
        right: '3%'
    },
    '@media (max-width: 850px)': {
        display: 'none'
    }
  },
  timeRow: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    alignSelf: 'flex-start',
    marginTop: '31px',
    '@media (max-width: 850px)': {
        flexDirection: 'column',
        alignItems: 'flex-start'
    }
  },
  buttonRow: {
    alignSelf: 'center',
    display: 'flex',
    justifyContent: 'center',
    flexWrap: 'wrap'
  },
  spacer: {
    flex: 5,
    '@media (max-width: 1250px)': {
        flex: 8
    }
  },
  mobileMargin: {
    '@media (max-width: 850px)': {
        margin: '0 25px'
    }
  },
  // Video styles
  video: {
    maxHeight: '50vw',
    maxWidth: '100vw',
    marginBottom: '20px',
    borderRadius: '5px 5px 0 0',
    overflow: 'hidden',
    '@media (min-width: 850px)': {
        display: 'none'
    }
  },
  preview: {
    fontSize: theme.fontSizes.xSmall,
    fontWeight: 'bold',
    color: theme.colors.primaryWhite,
    margin: '45px 0 10px 0'
  },
  thumbnail: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center'
  },
  playCircle: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    width: '76px',
    height: '76px',
    borderRadius: '76px',
    backgroundColor: theme.colors.navyBlue,
    opacity: 0.8
},
playTriangle: {
    width: 0, 
    height: 0,
    marginLeft: 6,
    borderTop: "13.5px solid transparent",
    borderBottom: "13.5px solid transparent",
    borderLeft: ["27px", "solid", theme.colors.primaryWhite]
}
}));

export type ButtonLink = {
  title: string;
  link: string;
};

type Props = {
  title: string;
  description: string;
  history: string[];
  estimatedTime: string;
  lastUpdated: string;
  price: string;
  video: string;
  onBuy: () => void;
  onBasket: () => void;
  sideComponent: React.ReactNode;
  className?: string;
};

function CoursePageHeader({
  title,
  description,
  history,
  estimatedTime,
  lastUpdated,
  price,
  video,
  onBuy,
  onBasket,
  sideComponent,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
        <div style={{ flex: 1 }} />
        <div className={classes.courseDetails}>
            {history && (
                <div className={classes.history}>
                    {history.map((page: string, index: number) =>
                        index !== history.length - 1 ? (
                            <div className={classes.extraLarge}>
                                {page} <Icon name="Right_Arrow" size={12} />
                            </div>
                        ) : (
                            <div className={classNames(classes.extraLarge, classes.bold)}>
                                {page}
                            </div>
                        )
                    )}
                </div>
            )}
            <VideoPlayer
                source={video}
                className={classes.video}
                thumbnail={
                    <div className={classes.thumbnail}>
                        <div className={classes.playCircle}>
                            <div className={classes.playTriangle}/>
                        </div>
                        <div className={classes.preview}>Preview this course</div>
                    </div>
                }
            />
            <div className={classNames(classes.courseTitle, classes.mobileMargin)}>
                {title}
            </div>
            <div className={classNames(classes.courseDesc, classes.mobileMargin)}>
                {description}
            </div>
            <div className={classes.timeRow}>
                <div className={classNames(classes.extraLarge, classes.mobileMargin)}>
                    {<strong>Estimated Time:</strong>} {estimatedTime}
                </div>
                <div
                    className={classNames(classes.updatedText, classes.extraLarge, classes.mobileMargin)}
                >
                    {<strong>Last Updated:</strong>} {lastUpdated}
                </div>
            </div>

            <div className={classes.price}>{price} <span className={classes.vat}>+VAT</span></div>

            <div className={classes.buttonRow}>
                <Button
                    archetype="submit"
                    className={classes.button}
                    onClick={() => onBasket && onBasket()}
                >
                    Add to Basket
                </Button>

                <Button
                    className={classNames(classes.button, classes.shadow)}
                    onClick={() => onBuy && onBuy()}
                >
                    Buy Now
                </Button>
            </div>
        </div>
        <div className={classes.spacer} />
        <div className={classes.component}>{sideComponent}</div>
    </div>
  );
}

export default CoursePageHeader;
