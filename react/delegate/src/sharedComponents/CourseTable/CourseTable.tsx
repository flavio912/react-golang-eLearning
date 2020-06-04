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
  key: -1,
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
    {
      component: () => <Attempt attempt={attempt} />
    },
    { component: () => <Status isComplete={status} expires={expires} /> }
  ]
});

const defaultFilterCourseOptions: DropdownOption[] = [];

type Props = {
  EmptyComponent: React.ReactElement;
  bookCourseHandler?: Function;
  className?: string;
  rowClicked: () => void | undefined;
};

const CourseTable = ({
  EmptyComponent,
  bookCourseHandler,
  className,
  rowClicked
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [filterCourse, setFilterCourse] = React.useState<DropdownOption>();

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
          'ATTEMPT',
          'STATUS'
        ]}
        rows={[
          courseRow(
            '1',
            'Dangerous Goods by Road Awareness',
            rowClicked,
            'DANGEROUS GOODS(ROAD)',
            80,
            32,
            '1',
            false,
            '',
            classes
          ),
          courseRow(
            '2',
            'Dangerous Goods by Road Awareness',
            rowClicked,
            'DANGEROUS GOODS(ROAD)',
            80,
            32,
            '1',
            true,
            '20/02/2022',
            classes
          )
        ]}
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
