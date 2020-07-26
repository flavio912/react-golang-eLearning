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
import { createFragmentContainer, graphql } from 'react-relay';
import { Router } from 'found';
import { DelegateProfilePage_delegate } from './__generated__/DelegateProfilePage_delegate.graphql';
import DelegateSlideIn from 'components/Delegate/DelegateSlideIn';

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

type Props = {
  delegate: DelegateProfilePage_delegate;
  router: Router;
};

const DelegateProfilePage = ({ delegate, router }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [action, setAction] = React.useState<DropdownOption>();
  
  const [openDelegateSlideIn, setOpenDelegateSlideIn] = React.useState(false);
  
  const headerActionOptions: DropdownOption[] = [
    {
      id: 1,
      title: 'Edit',
      component: <div onClick={() => setOpenDelegateSlideIn(true)}>Edit</div>
    }
  ];

  const inpDelegate = {
    uuid: delegate.uuid,
    firstName: delegate.firstName,
    lastName: delegate.lastName,
    jobTitle: delegate.jobTitle,
    email: delegate.email ?? '',
    phone: delegate.telephone ?? '',
    ttcId: delegate.TTC_ID,
    generatePassword: false,
    generatedPassword: '',
  };

  const myCourses =
    delegate?.myCourses?.map((myCourse) => ({
      title: myCourse.course.name,
      categoryName: myCourse.course.category?.name ?? '',
      progress: {
        total: 100,
        completed: myCourse.status === 'complete' ? 100 : 0
      },
      attempt: 1,
      status: {
        isComplete: myCourse.status === 'complete'
      },
      onClick: () => {}
    })) ?? [];
  return (
    <div className={classes.root}>
      <div className={classes.top}>
        <PageHeader
          showCreateButtons={false}
          title={`${delegate.firstName} ${delegate.lastName}`}
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
          title={`${delegate.firstName} ${delegate.lastName}'s summary`}
          className={classes.quickOverview}
        >
          <Summary
            numActiveCourses={3}
            numLastActive={4}
            numCertificates={2}
            numExpiringSoon={1}
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
        courses={myCourses}
        EmptyComponent={
          <ActiveCoursesEmpty
            title={`Book ${delegate.firstName} on their first Course`}
          />
        }
        rowClicked={() => {}}
      />
      <Spacer vertical spacing={3} />
      <ActivityTable />
      <DelegateSlideIn 
        isOpen={openDelegateSlideIn}
        onClose={() => setOpenDelegateSlideIn(false)}
        delegate={inpDelegate}
      />
    </div>
  );
};

const DelegateProfilePageFrag = createFragmentContainer(DelegateProfilePage, {
  delegate: graphql`
    fragment DelegateProfilePage_delegate on Delegate {
      uuid
      firstName
      lastName
      email
      jobTitle
      telephone
      TTC_ID
      myCourses {
        status
        course {
          name
          category {
            name
          }
        }
      }
    }
  `
});

export default DelegateProfilePageFrag;
