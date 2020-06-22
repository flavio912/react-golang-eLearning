import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  searchResultItemRoot: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    margin: [30, 0]
  },
  courseImage: {
    height: 120,
    width: 144,
    marginRight: 37,
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    backgroundPosition: 'center',
    position: 'relative',
    cursor: 'pointer',
    borderRadius: 5
  },
  courseInfo: {
    flex: 1
  },
  courseTitle: {
    lineHeight: '37px',
    letterSpacing: -0.63,
    fontWeight: 800,
    fontSize: theme.fontSizes.heading,
    color: theme.colors.primaryBlack,
    marginBottom: 10,
    margin: 0
  },
  courseDescription: {
    lineHeight: '24px',
    letterSpacing: -0.38,
    fontWeight: 500,
    fontSize: theme.fontSizes.xLarge,
    color: theme.colors.textGrey,
    margin: 0,
    display: '-webkit-box',
    '-webkit-line-clamp': 2,
    '-webkit-box-orient': 'vertical',
    overflow: 'hidden',
    textOverflow: 'ellipsis'
  }
}));
export type Course = {
  id: string | number;
  image: string;
  title: string;
  description: string;
};
type Props = {
  className?: string;
  course: Course;
  onClick: Function;
};
function SearchResultItem({ className, onClick, course }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const backgroundImage = `url(${course.image})`;
  return (
    <div
      className={classNames(classes.searchResultItemRoot, className)}
      onClick={() => onClick()}
    >
      <div
        className={classes.courseImage}
        style={{
          backgroundImage
        }}
      />
      <div className={classes.courseInfo}>
        <h2 className={classes.courseTitle}>{course.title}</h2>
        <p className={classes.courseDescription}>{course.description}</p>
      </div>
    </div>
  );
}

export default SearchResultItem;
