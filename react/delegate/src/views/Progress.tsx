import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import { useRouter } from 'found';
import Heading from 'components/core/Heading';
import CourseTable from 'sharedComponents/CourseTable';

const useStyles = createUseStyles((theme: Theme) => ({
  progressRoot: {
    display: 'grid',
    gridGap: theme.spacing(3),
    gridTemplateAreas: `
      "header header  .       .     "
      "active active  active  active"
    `,
    gridTemplateRows: 'repeat(4, auto)',
    gridTemplateColumns: 'repeat(4, 1fr)'
  },
  header: {
    gridArea: 'header'
  },
  activeTable: {
    gridArea: 'active'
  }
}));

type Props = {};

function Progress({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();

  const userName = 'James';
  return (
    <div className={classes.progressRoot}>
      <div className={classes.header}>
        <Heading text="Training Progress" size={'large'} />
        <Heading
          text={`${userName}, here you can see your training progress, and keep up to date with your daily activity.`}
          size={'medium'}
        />
      </div>
      <CourseTable
        EmptyComponent={<div>No Courses to show</div>}
        className={classes.activeTable}
        rowClicked={() => {
          router.push('/app/courses/1');
        }}
      />
    </div>
  );
}

export default Progress;
