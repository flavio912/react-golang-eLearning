import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Button from "sharedComponents/core/Input/Button";
import Attempt from "components/Delegate/Attempt";
import ActiveCoursesEmpty from "components/Delegate/ActiveCoursesEmpty";
import theme, { Theme } from "helpers/theme";
import Table from "components/core/Table";
import Text from "components/core/Table/Text/Text";
import Status from "components/core/Table/Status";
import Action from "components/core/Table/Action";
import CourseCompletion from "sharedComponents/core/CourseCompletion";
import Dropdown, { DropdownOption } from "sharedComponents/core/Input/Dropdown";
// import CheckboxSingle from "components/core/CheckboxSingle";

type Props = {};
const useStyles = createUseStyles((theme: Theme) => ({
  root: {},
  sectionTitleWrapper: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-start",
    "& h2": {
      fontSize: 15,
      color: theme.colors.primaryBlack,
      fontWeight: 300,
    },
    marginBottom: 16,
  },
  courseDropdown: {
    marginLeft: 49,
    background: theme.colors.primaryWhite,
  },
  courseButton: {
    display: "flex",
    justifyContent: "flex-end",
    marginTop: 27,
    "& button": {
      width: 214,
    },
  },
}));
const courseRowEmpty = () => ({
  key: -1,
  cells: [
    {
      component: () => (
        <ActiveCoursesEmpty title="Book John on their first Course" />
      ),
      colspan: 5,
    },
    {
      component: () => <Action />,
    },
  ],
});
const courseRow = (
  key: string | number,
  title: string,
  category: string,
  totalProcess: number,
  totalCompleted: number,
  attempt: string,
  status: boolean,
  expires?: string,
  classes?: any
): any => ({
  key,
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
      component: () => (
        <Text text={title} color={theme.colors.secondaryBlack} />
      ),
    },
    {
      component: () => (
        <Text text={category} color={theme.colors.secondaryBlack} />
      ),
    },
    {
      component: () => (
        <CourseCompletion total={totalProcess} complete={totalCompleted} />
      ),
    },
    {
      component: () => <Attempt attempt={attempt} />,
    },
    { component: () => <Status isComplete={status} expires={expires} /> },
    {
      component: () => <Action />,
    },
  ],
});
const defaultFilterCourseOptions: DropdownOption[] = [
  {
    id: 1,
    title: "Show Historical Courses 1",
    component: <div>Show Historical Courses 1</div>,
  },
  {
    id: 2,
    title: "Show Historical Courses 2",
    component: <div>Show Historical Courses 2</div>,
  },
  {
    id: 3,
    title: "Show Historical Courses 3",
    component: <div>Show Historical Courses 3</div>,
  },
];
const CourseTable = (props: any) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [filterCourse, setFilterCourse] = React.useState<DropdownOption>();

  return (
    <div className={classes.root}>
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
          "COURSE TITLE",
          "CATEGORY",
          "PROGRESS",
          "ATTEMPT",
          "STATUS",
          "ACTIONS",
        ]}
        rows={[
          courseRow(
            1,
            "Dangerous Goods by Road Awareness",
            "DANGEROUS GOODS(ROAD)",
            80,
            32,
            "1",
            false,
            "",
            classes
          ),
          courseRow(
            2,
            "Dangerous Goods by Road Awareness",
            "DANGEROUS GOODS(ROAD)",
            80,
            32,
            "1",
            true,
            "20/02/2022",
            classes
          ),
          courseRowEmpty(),
        ]}
      />
      <div className={classes.courseButton}>
        <Button bold archetype="submit">
          Book on new Course
        </Button>
      </div>
    </div>
  );
};

export default CourseTable;
