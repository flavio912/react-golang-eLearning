import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { createFragmentContainer, graphql } from 'react-relay';

import { Theme } from 'helpers/theme';
import CourseCard, {
  SizeOptions,
  Course
} from 'sharedComponents/Overview/CourseCard';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import PageHeader from 'components/PageHeader';
import SelectButton from 'components/core/Input/SelectButton';
import CircleBorder from 'sharedComponents/core/CircleBorder';
import Paginator from 'sharedComponents/Pagination/Paginator';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

import { CoursesPage_onlineCourses } from './__generated__/CoursesPage_onlineCourses.graphql';
import { CoursesPage_classroomCourses } from './__generated__/CoursesPage_classroomCourses.graphql';

const defaultCourse = {
  type: 'DANGEROUS GOODS AIR',
  colour: '#8C1CB4',
  url: '/static/media/SampleImage_ClassroomCoursesDetail_Feat.d89b5773.png',
  title: 'Dangerous goods by air category 7',
  price: 60,
  description:
    'This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage',
  assigned: 40,
  expiring: 9,
  date: 'MAR 3rd 2020',
  location: 'TTC at Hilton T4',
  modules: 5,
  lessons: 5,
  video_time: 5
};

const filterColour = '#AAAAAA90';

const defaultComponent = () => <div>PlaceHolder Name</div>;

const defaultOption: DropdownOption = {
  id: 1,
  title: 'Test Option',
  component: defaultComponent()
};

const defaultOptions = [defaultOption, defaultOption, defaultOption];

type Props = {
  onlineCourses?: CoursesPage_onlineCourses;
  classroomCourses?: CoursesPage_classroomCourses;
};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1,
    maxWidth: 1275
  },
  course: {
    margin: '15px 0px'
  },
  dropdown: {
    width: 177,
    marginLeft: 15,
    backgroundColor: theme.colors.primaryWhite
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'wrap'
  },
  spaceEvenly: {
    justifyContent: 'space-between'
  },
  filterRow: {
    margin: '30px 0px',
    justifyContent: 'space-between'
  }
}));

const CoursesPage = ({ onlineCourses, classroomCourses }: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  // Set card size depending on course type
  const isOnline = true;

  let courses: Course[] = [];

  if (onlineCourses) {
    courses = onlineCourses.edges.map((course: any) => ({
      id: course.info.id,
      type: course.info.category.name,
      colour: course?.info.category.color,
      url: '/static/media/SampleImage_ClassroomCoursesDetail_Feat.d89b5773.png',
      title: course.info.name,
      price: course.info.price,
      description: course.info.excerpt,
      assigned: 40,
      expiring: 9,
      modules: 5,
      lessons: 5,
      video_time: 5
    }));
  }
  if (classroomCourses) {
    // Todo
  }

  // Dropdown states
  const [date, setDate] = React.useState<DropdownOption>();
  const [category, setCategory] = React.useState<DropdownOption>();
  const [itemsPerPage, setItemsPerPage] = React.useState<DropdownOption>();

  // Page number states
  const [currentPage, setCurrentPage] = React.useState<number>(1);

  return (
    <div className={classes.root}>
      <PageHeader
        showCreateButtons
        title="Courses Available"
        subTitle="All active courses within your organisation"
        sideComponent={<CircleBorder text={23} size={30} fontSize={16} />}
      />
      <div className={classNames(classes.row, classes.filterRow)}>
        <SelectButton
          selected={isOnline ? 'Online Courses' : 'Classroom Courses'}
          options={['Online Courses', 'Classroom Courses']}
          onClick={(option: string) => ''}
        />
        <div className={classes.row}>
          <div className={classes.dropdown}>
            <Dropdown
              placeholder="March 2020"
              options={defaultOptions}
              selected={date}
              setSelected={setDate}
            />
          </div>
          <div className={classes.dropdown}>
            <Dropdown
              placeholder="Show Categories"
              options={defaultOptions}
              selected={category}
              setSelected={setCategory}
            />
          </div>
        </div>
      </div>
      <div className={classNames(classes.row, classes.spaceEvenly)}>
        {courses &&
          courses.map((course) => (
            <CourseCard
              className={classes.course}
              course={course}
              onClick={() => console.log('Pressed')}
              size={isOnline ? 'small' : 'large'}
            />
          ))}
      </div>
      <Spacer vertical spacing={3} />
      <div className={classes.row}>
        <Paginator
          currentPage={currentPage}
          updatePage={setCurrentPage}
          numPages={10}
          itemsPerPage={10}
        />
      </div>
      <Spacer vertical spacing={3} />
    </div>
  );
};

export const OnlineCoursesPage = createFragmentContainer(CoursesPage, {
  onlineCourses: graphql`
    fragment CoursesPage_onlineCourses on OnlineCoursePage {
      edges {
        uuid
        info {
          id
          name
          color
          excerpt
          price
          category {
            name
            color
          }
        }
      }
      pageInfo {
        total
        offset
        limit
        given
      }
    }
  `
});

export const ClassroomCoursesPage = createFragmentContainer(CoursesPage, {
  classroomCourses: graphql`
    fragment CoursesPage_classroomCourses on ClassroomCoursePage {
      edges {
        uuid
      }
      pageInfo {
        total
        offset
        limit
        given
      }
    }
  `
});
