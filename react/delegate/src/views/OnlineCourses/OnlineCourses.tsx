import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import Heading from 'components/core/Heading';
import SelectButton from 'components/core/Input/SelectButton';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import CourseCard from 'sharedComponents/Overview/CourseCard';
import { Course } from 'sharedComponents/Overview/CourseCard/CourseCard';
import Paginator from 'sharedComponents/Pagination/Paginator';
import Button from 'sharedComponents/core/Input/Button';
import { useRouter } from 'found';
import { createFragmentContainer, graphql } from 'react-relay';
import { OnlineCourses_user } from './__generated__/OnlineCourses_user.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  onlineCoursesRoot: {
    display: 'flex',
    flexDirection: 'column',
    flexGrow: 1
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
    gridTemplateColumns: 'repeat(auto-fit, 298px)'
  },
  page: {
    marginTop: theme.spacing(3)
  },
  subFilters: {
    display: 'flex'
  },
  filterButton: {
    marginRight: theme.spacing(1)
  }
}));

type Props = {
  className?: string;
  user?: OnlineCourses_user;
};

const defaultComponent = () => <div>Dangerous Goods Air</div>;

const defaultOption: DropdownOption = {
  id: 1,
  title: 'Test Option',
  component: defaultComponent()
};

function OnlineCourses({ className, user }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();

  const selectOptions = ['Active Courses', 'Completed Courses'];
  const [selectedOption, setSelectedOption] = React.useState(selectOptions[0]);

  const filterOptions = [defaultOption];

  return (
    <div className={classes.onlineCoursesRoot}>
      <Heading
        text={'Courses'}
        size={'large'}
        className={classes.mainHeading}
      />
      <Heading
        text={`${user?.firstName}, here are all of the courses you’ve currently been enrolled on. If you’ve already passed a course, it’s stored here for safe keeping! `}
        size={'medium'}
        className={classes.subHeading}
      />
      <div className={classes.filterRow}>
        <SelectButton
          options={selectOptions}
          selected={selectedOption}
          onClick={(option: string) => {
            setSelectedOption(option);
          }}
          className={classes.selectButton}
        />
        <div className={classes.subFilters}>
          <Button
            archetype={'default'}
            icon={{ right: 'FilterAdjust' }}
            children={'Filters'}
            className={classes.filterButton}
          />
          <Dropdown
            options={filterOptions}
            placeholder={'Show Categories'}
            setSelected={() => {
              //
            }}
          />
        </div>
      </div>
      <div className={classes.courseHolder}>
        {user?.myCourses?.map((course, index) => (
          <CourseCard
            key={index}
            course={{
              title: course.course.name,
              id: course.course.ident,
              type: 'Online Course',
              colour: course.course.color ?? '',
              description: course.course.excerpt ?? '',
              url: course.course.bannerImageURL ?? ''
            }}
            onClick={() => {
              router.push(`/app/courses/${course.course.ident}`);
            }}
            progress={30}
          />
        ))}
      </div>
      <div className={classes.page}>
        <Paginator
          currentPage={1}
          numPages={3}
          updatePage={() => {
            //TODO:
          }}
          itemsPerPage={10}
        />
      </div>
    </div>
  );
}

export default createFragmentContainer(OnlineCourses, {
  user: graphql`
    fragment OnlineCourses_user on User {
      firstName
      myCourses {
        course {
          ident: id
          name
          excerpt
          color
          type
          bannerImageURL
        }
        status
      }
    }
  `
});
