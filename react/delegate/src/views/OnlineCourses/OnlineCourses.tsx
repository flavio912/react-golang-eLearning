import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import Heading from 'components/core/Heading';
import SelectButton from 'components/core/Input/SelectButton';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import CourseCard from 'sharedComponents/Overview/CourseCard';
import { Course } from 'sharedComponents/Overview/CourseCard/CourseCard';
import Paginator from 'sharedComponents/Pagination/Paginator';
import { useRouter } from 'found';

const useStyles = createUseStyles((theme: Theme) => ({
  onlineCoursesRoot: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
    maxWidth: 1275
  },
  mainHeading: {
    gridArea: 'headin'
  },
  subHeading: {
    gridArea: 'subhea',
    maxWidth: 466
  },
  selectButton: {
    gridArea: 'select'
  },
  filterRow: {
    marginTop: theme.spacing(3),
    display: 'flex',
    justifyContent: 'space-between'
  },
  courseHolder: {
    display: 'grid',
    gridGap: 26,
    marginTop: theme.spacing(3),
    gridTemplateColumns: 'repeat(auto-fit, 298px)',
    justifyContent: 'space-between'
  },
  page: {
    marginTop: theme.spacing(3)
  }
}));

type Props = {
  className?: string;
};

const defaultComponent = () => <div>PlaceHolder Name</div>;

const defaultOption: DropdownOption = {
  id: 1,
  title: 'Test Option',
  component: defaultComponent()
};

const defaultCourse: Course = {
  type: 'DANGEROUS GOODS AIR',
  colour: '#8C1CB4',
  url: '/static/media/SampleImage_ClassroomCoursesDetail_Feat.d89b5773.png',
  title: 'Dangerous goods by air category 7',
  price: 60,
  description:
    'This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage',
  date: 'MAR 3rd 2020',
  location: 'TTC at Hilton T4',
  modules: 5,
  lessons: 5,
  video_time: 5
};

function OnlineCourses({ className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  const userName = 'James';

  const selectOptions = ['Active Courses', 'Completed Courses'];

  const filterOptions = [defaultOption, defaultOption];

  const courses = [
    defaultCourse,
    defaultCourse,
    defaultCourse,
    defaultCourse,
    defaultCourse,
    defaultCourse
  ];

  return (
    <div className={classes.onlineCoursesRoot}>
      <Heading
        text={'Courses'}
        size={'large'}
        className={classes.mainHeading}
      />
      <Heading
        text={`${userName}, here are all of the courses you’ve currently been enrolled on. If you’ve already passed a course, it’s stored here for safe keeping! `}
        size={'medium'}
        className={classes.subHeading}
      />
      <div className={classes.filterRow}>
        <SelectButton
          options={selectOptions}
          selected={selectOptions[0]}
          onClick={() => {}}
          className={classes.selectButton}
        />
        <Dropdown
          options={filterOptions}
          placeholder={'Show Categories'}
          setSelected={() => {}}
        />
      </div>
      <div className={classes.courseHolder}>
        {courses.map((course, index) => (
          <CourseCard
            key={index}
            course={course}
            onClick={() => {
              router.push('/app/courses/1');
            }}
            progress={30}
          />
        ))}
      </div>
      <div className={classes.page}>
        <Paginator
          currentPage={1}
          numPages={3}
          updatePage={() => {}}
          itemsPerPage={10}
        />
      </div>
    </div>
  );
}
export default OnlineCourses;
