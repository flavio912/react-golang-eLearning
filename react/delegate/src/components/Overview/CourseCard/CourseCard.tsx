import * as React from 'react';
import Card, { PaddingOptions } from '../../../sharedComponents/core/Card';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import Button from "../../../sharedComponents/core/Button";
import Icon from "../../../sharedComponents/core/Icon";
import { Theme } from 'helpers/theme';
import FooterIcon from './FooterIcon';
import CourseCompletion from 'sharedComponents/core/Table/CourseCompletion';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    '&:hover': {
      boxShadow: '0 2px 12px 0 rgba(0,0,0,0.18)'
    }
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
    fontSize: theme.fontSizes.small,
    fontWeight: '700',
    color: theme.colors.primaryWhite,
    borderRadius: `0 0 ${theme.secondaryBorderRadius}px 0`,
    padding: '10px 20px'
  },
  icon: {
    alignSelf: 'flex-start',
    margin: '10px'
  },
  price: {
    color: theme.colors.primaryWhite,
    fontSize:  theme.fontSizes.large,
    margin: '20px 20px 5px 20px',
    fontWeight: '800'
  },
  title: {
    color: theme.colors.primaryWhite,
    margin: '0 20px 30px 20px',
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
    margin: '20px'
  },
  progress: {
    color: theme.colors.primaryBlack,
    margin: '15px 0',
    fontSize:  theme.fontSizes.tiny,
    fontWeight: '400'
  },
  button: {
    flex: 1,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center'
  },
  footer: {
    padding: '0 20px',
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid ${theme.colors.borderGrey}`,
    borderWidth: '1px 0 0 0',
    borderRadius: `0 0 ${theme.primaryBorderRadius}px ${theme.primaryBorderRadius}px`,
  },
  filler: {
    flex: 0.5
  },
  small: {
    width: '298px',
    cursor: 'pointer'
  },
  large: {
    width: '610px'
  }
}));

export type SizeOptions = "small" | "large";
export type Completion = {
  total: number;
  complete: number;
}
export interface Course {
  type: string;
  colour: string;
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
  filterColour: string;
  size?: SizeOptions;
  progress?: Completion;
  onClick: Function,
  padding?: PaddingOptions;
  className?: string;
};

function CourseCard({ course, filterColour, onClick, size = 'small', progress, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const backgroundColor = { backgroundColor: course.colour };
  const backgroundImage = { backgroundImage: `linear-gradient(${filterColour}, ${filterColour}), url(${course.url})` };

  return (
    <Card className={classNames(classes.root, classes.noBorder, classes[size], className)}>
      <div className={classNames(classes.mainContainer)} style={backgroundImage}>

        <div className={classNames(classes.row)}>
          <div className={classNames(classes.heading)} style={backgroundColor}>
            {course.type}
          </div>
          <Icon className={classNames(classes.icon)} name="Card_SecondaryActon_Dots" size={18} />
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
            <Button archetype="submit" onClick={() => onClick()}>
                Book Now
            </Button>
          </div>
          )}
      </div>

    {progress && progress.complete && progress.total ? 
      <div className={classNames(classes.row, classes.footer)}>
        <div className={classes.progress}>PROGRESS</div>
        <CourseCompletion complete={progress.complete} total={progress.total} width={125} fraction={false} />
      </div>
    : size === 'small' ? (
      <div className={classNames(classes.row, classes.footer)}>
          <FooterIcon name="Icon_Delegates" size={20} value={course.assigned} />
          <FooterIcon name="CourseExpiringSoon" size={20} value={course.expiring} />
          <div className={classNames(classes.filler)} />
      </div>
      ) : (
        <div className={classNames(classes.row, classes.footer)}>
          <FooterIcon name="Course_Calendar" size={20} text={course.date} />
          <FooterIcon name="Location_Pin" size={20} text={course.location} />
          <div className={classNames(classes.filler)} />
      </div>
      )}
    </Card>
  );
}

export default CourseCard;
