import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CourseCard, { SizeOptions } from 'sharedComponents/Overview/CourseCard';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import PageHeader from 'components/PageHeader';
import SelectButton from 'components/core/Input/SelectButton';
import CircleBorder from 'sharedComponents/core/CircleBorder';
import Paginator from 'sharedComponents/Pagination/Paginator';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

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
const courses = [
  defaultCourse,
  defaultCourse,
  defaultCourse,
  defaultCourse,
  defaultCourse,
  defaultCourse,
  defaultCourse,
  defaultCourse
];

const defaultComponent = () => <div>PlaceHolder Name</div>;

const defaultOption: DropdownOption = {
  id: 1,
  title: 'Test Option',
  component: defaultComponent()
};

const defaultOptions = [defaultOption, defaultOption, defaultOption];

type Props = {};

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

const CoursesPage = () => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  // Card size state
  const sizeOptions: SizeOptions[] = ['small', 'large'];
  const options = ['Online Courses', 'Classroom Courses'];
  const [cardSize, setCardSize] = React.useState(options[0]);

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
          selected={cardSize}
          options={options}
          onClick={(option: string) => setCardSize(option)}
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
              size={sizeOptions[options.indexOf(cardSize)]}
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

export default CoursesPage;
