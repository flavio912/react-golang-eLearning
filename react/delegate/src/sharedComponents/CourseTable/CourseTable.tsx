import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Button from 'sharedComponents/core/Input/Button';
import Attempt from './Attempt';
import theme, { Theme } from 'helpers/theme';
import Table, { TableRow } from 'sharedComponents/core/Table';
import Text from 'sharedComponents/core/Table/Text/Text';
import Status from 'sharedComponents/core/Table/Status';
import classnames from 'classnames';
import CourseCompletion from 'sharedComponents/core/CourseCompletion';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
// import CheckboxSingle from "components/core/CheckboxSingle";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {},
  sectionTitleWrapper: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-start',
    '& h2': {
      fontSize: 15,
      color: theme.colors.primaryBlack,
      fontWeight: 300
    },
    marginBottom: 16
  },
  courseDropdown: {
    marginLeft: 49,
    background: theme.colors.primaryWhite
  },
  courseButton: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginTop: 27,
    '& button': {
      width: 214
    }
  }
}));

const courseRowEmpty = (EmptyComponent: React.ReactElement) => ({
  key: '-1',
  cells: [
    {
      component: () => EmptyComponent,
      colspan: 5
    }
  ]
});

const courseRow = (
  key: string,
  title: string,
  onClick: () => void | undefined,
  category: string,
  totalProcess: number,
  totalCompleted: number,
  attempt: string,
  status: boolean,
  expires?: string,
  classes?: { [key: string]: string }
): TableRow => ({
  key,
  onClick,
  cells: [
    // {
    //   component: () => (
    //     <CheckboxSingle
    //       box={{
    //         label: "",
    //         checked: false,
    //       }}
    //       setBox={() => {}}
    //     />
    //   ),
    // },
    {
      component: () => <Text text={title} color={theme.colors.secondaryBlack} />
    },
    {
      component: () => (
        <Text text={category} color={theme.colors.secondaryBlack} />
      )
    },
    {
      component: () => (
        <CourseCompletion total={totalProcess} complete={totalCompleted} />
      )
    },
    { component: () => <Status isComplete={status} expires={expires} /> }
  ]
});

const defaultFilterCourseOptions: DropdownOption[] = [];

type Course = {
  title: string;
  categoryName: string;
  progress: {
    total: number;
    completed: number;
  };
  attempt: number;
  status: {
    isComplete: boolean;
    expires?: string;
  };
  onClick: () => void;
};

type Props = {
  EmptyComponent: React.ReactElement;
  bookCourseHandler?: Function;
  className?: string;
  rowClicked: () => void | undefined;
  courses: Course[];
};

const CourseTable = ({
  EmptyComponent,
  bookCourseHandler,
  className,
  courses,
  rowClicked
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [filterCourse, setFilterCourse] = React.useState<DropdownOption>();

  let courseRows: TableRow[] = [];
  if (courses.length > 0) {
    courseRows = courses.map((course, index) =>
      courseRow(
        String(index),
        course.title,
        course.onClick,
        course.categoryName,
        course.progress.total,
        course.progress.completed,
        String(course.attempt),
        course.status.isComplete,
        course.status.expires,
        classes
      )
    );
  } else {
    courseRows = [courseRowEmpty(EmptyComponent)];
  }
  return (
    <div className={classnames(classes.root, className)}>
      <div className={classes.sectionTitleWrapper}>
        <h2>Active Courses</h2>
        <div className={classes.courseDropdown}>
          <Dropdown
            placeholder="Show Historical Courses"
            options={defaultFilterCourseOptions}
            selected={filterCourse}
            setSelected={setFilterCourse}
          />
        </div>
      </div>
      <Table
        header={[
          //   <CheckboxSingle
          //     box={{ label: "", checked: false }}
          //     setBox={() => {}}
          //   />,
          'COURSE TITLE',
          'CATEGORY',
          'PROGRESS',
          'STATUS'
        ]}
        rows={courseRows}
      />
      {bookCourseHandler && (
        <div
          className={classes.courseButton}
          onClick={() => bookCourseHandler()}
        >
          <Button bold archetype="submit">
            Book on new Course
          </Button>
        </div>
      )}
    </div>
  );
};

export default CourseTable;
