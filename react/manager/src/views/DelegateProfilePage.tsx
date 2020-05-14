import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Button from "sharedComponents/core/Button";
import TitleWrapper from "components/Overview/TitleWrapper";
import Summary from "components/Overview/Summary";
import Attempt from "components/Delegate/Attempt";
import ActiveType from "components/Delegate/ActiveType";
import ActiveCoursesEmpty from "components/Delegate/ActiveCoursesEmpty";
import TrainingProgressCard from "components/Overview/TrainingProgressCard";
import theme, { Theme } from "helpers/theme";
import PageHeader from "components/PageHeader";
import Table from "components/core/Table";
import Text from "components/core/Table/Text/Text";
import Status from "components/core/Table/Status";
import Action from "components/core/Table/Action";
import CourseCompletion from "components/core/Table/CourseCompletion";
import Dropdown, { DropdownOption } from "components/core/Dropdown";
import Spacer from "components/core/Spacers/Spacer";
import Paginator from "components/Paginator";
import CheckboxSingle from "components/core/CheckboxSingle";
import TimeSpent from "components/Delegate/TimeSpent";
import ActivityName from "components/Delegate/ActivityName";

type Props = {};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    flexGrow: 1,
    maxWidth: 1275,
  },
  top: {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
  },
  searchAndFilter: {
    display: "flex",
    justifyContent: "space-between",
    height: 40,
  },
  tableOptions: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  dropdown: {
    background: "white",
  },
  divider: {
    width: theme.spacing(1),
  },
  search: {
    flex: 0.4,
  },

  quickOverview: {
    gridArea: "overvw",
  },
  trainingProgress: {
    gridArea: "traini",
  },
  cardFlex: {
    display: "flex",
  },
  grid: {
    marginTop: 19,
    display: "flex",
    justifyContent: "space-between",
  },
  sectionTitleWrapper: {
    display: "flex",
    alignItems: "center",
    justifyContent: "flex-start",
    "& h2": {
      fontSize: 15,
      color: theme.colors.primaryBlack,
      fontWeight: 400,
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
  activeType: {
    display: "flex",
    alignItems: "center",
    "& div": {
      marginLeft: 10,
    },
  },

  headerActions: {
    "& div": {
      background: theme.colors.primaryWhite,
    },
  },
}));
const courseRowEmpty = () => ({
  key: "1",
  cells: [
    {
      component: () => <ActiveCoursesEmpty />,
      colspan: 6,
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
    {
      component: () => (
        <CheckboxSingle
          box={{
            label: "",
            checked: false,
          }}
          setBox={() => {}}
        />
      ),
    },
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
const activityRow = (
  key: string | number,
  activeTime: {
    time: string;
    date: string;
  },
  title: string,
  activeType: {
    icon: any;
    text: string;
  },
  timeSpent:
    | string
    | {
        h: number;
        m: number;
      },
  userName: string,
  classes?: any
): any => ({
  key,
  cells: [
    {
      component: () => (
        <CheckboxSingle
          box={{
            label: "",
            checked: false,
          }}
          setBox={() => {}}
        />
      ),
    },
    {
      component: () => {
        return (
          <Text
            text={`${activeTime.time} on ${activeTime.date}`}
            color={theme.colors.secondaryBlack}
          />
        );
      },
    },
    {
      component: () => <ActivityName userName={userName} title={title} />,
    },
    {
      component: () => (
        <ActiveType icon={activeType.icon} text={activeType.text} />
      ),
    },
    {
      component: () => <TimeSpent timeSpent={timeSpent} />,
    },
    { component: () => null },
    {
      component: () => null,
    },
  ],
  onClick: () => {},
});
const delegateData = {
  userUUID: "asda",
  name: "Bruce Willis",
  email: "bruce.willis@email.com",
  courses: 30,
  certificates: 10,
  lastActive: 30,
  expiringSoon: 30,
};
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
const headerActionOptions: DropdownOption[] = [
  {
    id: 1,
    title: "Action 1",
    component: <div>Action 1</div>,
  },
  {
    id: 2,
    title: "Action 2",
    component: <div>Action 2</div>,
  },
  {
    id: 3,
    title: "Action 3",
    component: <div>Action 3</div>,
  },
];
const DelegateProfilePage = (props: any) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = props;
  const [filterCourse, setFilterCourse] = React.useState<DropdownOption>();
  const [action, setAction] = React.useState<DropdownOption>();

  return (
    <div className={classes.root}>
      <div className={classes.top}>
        <PageHeader
          showCreateButtons={false}
          title={delegateData.name}
          subTitle="Member of Fedex UK Limited"
          backProps={{
            text: "Back to all Delegates",
            onClick: () => router.push("/app/delegates"),
          }}
        />
        <div className={classes.headerActions}>
          <Dropdown
            placeholder="Actions"
            options={headerActionOptions}
            selected={action}
            setSelected={setAction}
          />
        </div>
      </div>
      <div className={classes.grid}>
        <TitleWrapper
          title={`${delegateData.name}'s summary`}
          className={classes.quickOverview}
        >
          <Summary
            numActiveCourses={delegateData.courses}
            numLastActive={delegateData.lastActive}
            numCertificates={delegateData.certificates}
            numExpiringSoon={delegateData.expiringSoon}
          />
        </TitleWrapper>
        <TitleWrapper
          title="Training Progress"
          className={classes.trainingProgress}
        >
          <div className={classes.cardFlex}>
            <TrainingProgressCard
              coursesDone={0}
              courseNewCourseIcon={"CourseNewCourseGrey"}
              courseTimeTrackedIcon={"CourseTimeTrackedGrey"}
              courseTitle="Modules done"
              timeTracked={"n/a"}
              title="Weekly"
            />
            <Spacer spacing={3} horizontal />
            <TrainingProgressCard
              coursesDone={20}
              coursesPercent={300}
              courseNewCourseIcon={"CourseNewCourseGreen"}
              courseTimeTrackedIcon={"CourseTimeTrackedGreen"}
              timeTracked={{ h: 30, m: 10 }}
              timePercent={100}
              title="Monthly"
            />
          </div>
        </TitleWrapper>
      </div>

      <Spacer spacing={3} vertical />
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
          <CheckboxSingle
            box={{ label: "", checked: false }}
            setBox={() => {}}
          />,
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
            1,
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
      <Spacer vertical spacing={3} />
      <div className={classes.sectionTitleWrapper}>
        <h2>Bruce's activity</h2>
        <div className={classes.courseDropdown}>
          <Button
            archetype={"default"}
            icon={{ right: "FilterAdjust" }}
            children={"Filter Activities"}
          />
        </div>
      </div>
      <Table
        header={[
          "",
          "ACTIVITY TIME",
          "NAME",
          "ACTIVE TYPE",
          "TIME SPENT",
          "",
          "",
        ]}
        rows={[
          activityRow(
            1,
            {
              time: "10:29",
              date: "01/02/2020",
            },
            "Bruce failed the Dangerous Goods by Road Awareness Course",
            {
              icon: "CourseFailed",
              text: "Failed Course",
            },
            {
              h: 3,
              m: 15,
            },
            "Bruce Willis",
            classes
          ),
          activityRow(
            1,
            {
              time: "10:29",
              date: "01/02/2020",
            },
            "Bruce failed the Dangerous Goods by Road Awareness Course",
            {
              icon: "CourseNewCourse",
              text: "New Course",
            },
            {
              h: 3,
              m: 15,
            },
            "Bruce Willis",
            classes
          ),
          activityRow(
            1,
            {
              time: "10:29",
              date: "01/02/2020",
            },
            "Bruce failed the Dangerous Goods by Road Awareness Course",
            {
              icon: "CourseCertificates",
              text: "New New certificate!",
            },
            {
              h: 3,
              m: 15,
            },
            "Bruce Willis",
            classes
          ),
          activityRow(
            1,
            {
              time: "10:29",
              date: "01/02/2020",
            },
            "Bruce failed the Dangerous Goods by Road Awareness Course",
            {
              icon: "CourseAccountActivated",
              text: "Account active",
            },
            "n/a",
            "Bruce Willis",
            classes
          ),
        ]}
      />
      <Paginator
        currentPage={1}
        updatePage={() => {}}
        numPages={10}
        itemsPerPage={10}
      />
    </div>
  );
};

export default DelegateProfilePage;
