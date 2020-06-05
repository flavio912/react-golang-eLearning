import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import TitleWrapper from 'components/Overview/TitleWrapper';
import Summary from 'components/Overview/Summary';
import TrainingProgressCard from 'components/Overview/TrainingProgressCard';
import { Theme } from 'helpers/theme';
import PageHeader from 'components/PageHeader';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import CourseTable from 'sharedComponents/CourseTable';
import ActivityTable from 'components/Delegate/ActivityTable';
import ActiveCoursesEmpty from 'components/Delegate/ActiveCoursesEmpty';

type Props = {};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
    maxWidth: 1275
  },
  top: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between'
  },
  divider: {
    width: theme.spacing(1)
  },
  quickOverview: {
    gridArea: 'overvw'
  },
  trainingProgress: {
    gridArea: 'traini'
  },
  cardFlex: {
    display: 'flex'
  },
  grid: {
    marginTop: 19,
    display: 'flex',
    justifyContent: 'space-between'
  },
  headerActions: {}
}));
const delegateData = {
  userUUID: 'asda',
  name: 'Bruce Willis',
  email: 'bruce.willis@email.com',
  courses: 30,
  certificates: 10,
  lastActive: 30,
  expiringSoon: 30
};
const headerActionOptions: DropdownOption[] = [
  {
    id: 1,
    title: 'Edit',
    component: <div>Edit</div>
  }
];

const DelegateProfilePage = (props: any) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = props;
  const [action, setAction] = React.useState<DropdownOption>();

  return (
    <div className={classes.root}>
      <div className={classes.top}>
        <PageHeader
          showCreateButtons={false}
          title={delegateData.name}
          subTitle="Member of Fedex UK Limited"
          backProps={{
            text: 'Back to all Delegates',
            onClick: () => router.push('/app/delegates')
          }}
        />
        <div className={classes.headerActions}>
          <Dropdown
            placeholder="Actions"
            options={headerActionOptions}
            selected={action}
            setSelected={() => {}}
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
              courseNewCourseIcon={'CourseNewCourseGrey'}
              courseTimeTrackedIcon={'CourseTimeTrackedGrey'}
              courseTitle="Modules done"
              timeTracked={'n/a'}
              title="Weekly"
            />
            <Spacer spacing={3} horizontal />
            <TrainingProgressCard
              coursesDone={20}
              coursesPercent={300}
              courseNewCourseIcon={'CourseNewCourseGreen'}
              courseTimeTrackedIcon={'CourseTimeTrackedGreen'}
              timeTracked={{ h: 30, m: 10 }}
              timePercent={100}
              title="Monthly"
            />
          </div>
        </TitleWrapper>
      </div>

      <Spacer spacing={3} vertical />
      <CourseTable
        EmptyComponent={
          <ActiveCoursesEmpty title="Book John on their first Course" />
        }
        rowClicked={() => {}}
      />
      <Spacer vertical spacing={3} />
      <ActivityTable />
    </div>
  );
};

export default DelegateProfilePage;
