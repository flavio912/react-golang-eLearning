import * as React from 'react';
import Card, { PaddingOptions } from '../../core/Card';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import Button from "../../core/Button";
import Icon from "../../core/Icon";
import { Theme } from 'helpers/theme';
import FooterIcon from './FooterIcon';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
  },
  noBorder: {
    borderRadius: `0 ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px`,
  },
  mainContainer: {
    borderRadius: `0 ${theme.primaryBorderRadius}px 0 0`,
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between'
  },
  heading: {
    alignSelf: 'flex-start',
    fontSize: theme.fontSizes.default,
    fontWeight: '700',
    color: theme.colors.primaryWhite,
    borderRadius: `0 0 ${theme.primaryBorderRadius}px 0`,
    padding: '10px 30px'
  },
  icon: {
    alignSelf: 'flex-start',
    margin: '10px'
  },
  price: {
    color: theme.colors.primaryWhite,
    fontSize:  theme.fontSizes.large,
    margin: '30px 30px 10px 30px',
    fontWeight: '800'
  },
  title: {
    color: theme.colors.primaryWhite,
    margin: '0 30px 50px 30px',
    fontSize:  theme.fontSizes.heading,
    fontWeight: '900'
  },
  description: {
    flex: 3,
    lineHeight: '1.5em',
    height: '4.5em',
    overflow: 'hidden',
    fontWeight: '300',
    fontSize:  theme.fontSizes.small,
    margin: '30px 20px 30px 20px'
  },
  button: {
    flex: 1,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  footer: {
    padding: '0 15px',
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid ${theme.colors.borderGrey}`,
    borderWidth: '1px 0 0 0',
    borderRadius: `0 0 ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px`,
  },
  filler: {
    flex: 0.5
  },
  small: {
    width: 340,
    cursor: 'pointer'
  },
  large: {
    width: 680
  }
}));

export type SizeOptions = "small" | "large";
export interface Course {
  type: string;
  url: string;
  title: string;
  price: number;
  description: string;
  assigned: number;
  expiring: number;
  date: string;
  location: string;
}

type Props = {
  course: Course;
  color: string;
  size?: SizeOptions;
  padding?: PaddingOptions;
  className?: string;
};

function CourseCard({ course, color, size = 'small', className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const sizeLink = {
    small: classes.small,
    large: classes.large,
  };

  const backgroundColor = { backgroundColor: color };
  const backgroundImage = { backgroundImage: `linear-gradient(${color}4D, ${color}4D), url(${course.url})` };

  return (
    <Card className={classNames(classes.root, classes.noBorder, sizeLink[size], className)}>
      <div className={classNames(classes.mainContainer)} style={backgroundImage}>

        <div className={classNames(classes.row)}>
          <div className={classNames(classes.heading)} style={backgroundColor}>
            {course.type}
          </div>
          <div className={classNames(classes.icon)}>
            <Icon name="Card_SecondaryActon_Dots" size={18} />
          </div>
        </div>

      <div className={classNames(classes.price)}>£{course.price.toFixed(2)}</div>
      <div className={classNames(classes.title)}>{course.title}</div>

    </ div>

      <div className={classNames(classes.row)}>
        <div className={classNames(classes.description)}>
          {course.description}
        </div>
        {size === 'large' && (
          <div className={classNames(classes.button)}>
            <Button archetype="submit">
                Book Now
            </Button>
          </div>
          )}
      </div>

    {size === 'small' ? (
      <div className={classNames(classes.row, classes.footer)}>
          <FooterIcon name="LeftNav_Icon_Delegates" size={30} value={course.assigned} />
          <FooterIcon name="CourseExpiringSoon" size={25} value={course.expiring} />
          <div className={classNames(classes.filler)} />
      </div>
      ) : (
        <div className={classNames(classes.row, classes.footer)}>
          <FooterIcon name="Course_Calendar" size={25} text={course.date} />
          <FooterIcon name="Location_Pin" size={25} text={course.location} />
          <div className={classNames(classes.filler)} />
      </div>
      )}
    </Card>
  );
}

export default CourseCard;
