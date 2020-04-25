import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import CourseCard, { SizeOptions } from "components/Overview/CourseCard";
import Dropdown, { DropdownOption } from "components/core/Dropdown";
import PageHeader from "components/PageHeader";
import SelectButton from "components/core/SelectButton";
import CircleBorder from "components/core/CircleBorder";
import PageNumbers from "components/core/PageNumbers";

const defaultCourse = {
  type: "DANGEROUS GOODS AIR",
  colour: "#8C1CB4",
  url: "https://www.stickpng.com/assets/images/58428e7da6515b1e0ad75ab5.png",
  title: "Dangerous goods by air category 7",
  price: 60,
  description: "This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage, This course is for those involved in the handling, storage and loading of cargo or mail and baggage",
  assigned: 40,
  expiring: 9,
  date: "MAR 3rd 2020",
  location: "TTC at Hilton T4"
};

const filterColour = "#AAAAAA90";
const courses = [defaultCourse, defaultCourse, defaultCourse, defaultCourse, defaultCourse, defaultCourse, defaultCourse, defaultCourse];

const defaultComponent = () => (
  <div>
    PlaceHolder Name
  </div>
)

const defaultOption: DropdownOption = {
  id: 1,
  title: "Test Option",
  component: defaultComponent()
}

const defaultOptions = [defaultOption, defaultOption, defaultOption];

type Props = {};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    flexGrow: 1,
    maxWidth: 1275,
  },
  course: {
    margin: '15px 10px',
  },
  dropdown: {
    width: 177,
    marginLeft: 15,
    backgroundColor: theme.colors.primaryWhite
  },
  pageDropdown: {
    width: 117,
    backgroundColor: theme.colors.primaryWhite
  },
  row: {
    display: "flex",
    flexDirection: "row",
    flexWrap: 'wrap',
  },
  spaceEvenly: {
    justifyContent: 'space-evenly'
  },
  filterRow: {
    margin: '30px 0px',
    justifyContent: 'space-between'
  },
}));

const CoursesPage = () => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  // Card size state
  const sizeOptions: SizeOptions[] = ["small", "large"];
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
        {courses && courses.map((course) => (
          <CourseCard
            className={classes.course}
            course={course}
            filterColour={filterColour}
            onClick={() => console.log('Pressed')}
            size={sizeOptions[options.indexOf(cardSize)]}
          />
        ))}
      </div>
      <div className={classes.row}>
        <div className={classes.pageDropdown}>
          <Dropdown
            placeholder="Show 10"
            options={defaultOptions}
            selected={itemsPerPage}
            setSelected={setItemsPerPage}
          />
        </div>
        <PageNumbers
          numberOfPages={10}
          range={4}
          currentPage={currentPage}
          setCurrentPage={(page: number) => setCurrentPage(page)}
        />
      </div>
    </div>
  );
};

export default CoursesPage;
