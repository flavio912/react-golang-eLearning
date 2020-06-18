import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import CourseItem, { CourseProps } from './CourseItem';
import Dropdown, { DropdownOption } from 'sharedComponents/core/Input/Dropdown';
import Button from 'sharedComponents/core/Input/Button';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'column'
  },
  tabBar: {
    borderBottom: ['1px', 'solid', theme.colors.borderGrey]
  },
  tab: {
    cursor: 'pointer',
    fontSize: theme.fontSizes.large,
    color: theme.colors.primaryBlack,
    fontWeight: 400,
    textAlign: 'center',
    marginRight: 65,
    padding: '14px 20px',
    opacity: 0.4,
    transition: 'opacity 0.5s linear'
  },
  selected: {
    borderBottom: ['3.5px', 'solid', theme.colors.primaryGreen],
    opacity: 1,
    transition: 'opacity 0.5s linear'
  },
  noMargin: {
    marginRight: 0
  },
  list: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    backgroundColor: '#F7F9FB',
    padding: '0 15px'
  },
  listOptions: {
    maxWidth: '1003px',
    width: '100%',
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    margin: '57px 0 52px 0'
  },
  dropDown: {
    marginLeft: '25px'
  },
  listItem: {
    maxWidth: '1083px',
    marginBottom: '30px'
  },
  searchText: {
    display: 'flex',
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.smallHeading,
    fontWeight: '800',
    marginLeft: '10px'
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center'
  },
  button: {
    height: '52px',
    width: '182px',
    margin: '0 20px',
    boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)',
    fontSize: theme.fontSizes.large,
    fontWeight: '800'
  },
  line: {
    borderBottom: ['1.5px', 'solid', theme.colors.borderGrey],
    width: '400px'
  }
}));

type TabNames = 'All Courses' | 'Regulated Agents' | 'Known Consignor' | 'GSAT';

export type Tab = {
  name: TabNames;
  value: string;
};

type Props = {
  tabs: Tab[];
  courses: CourseProps[];
  className?: string;
};

function CourseSearch({ tabs, courses, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [selectedTab, setSelectedTab] = React.useState(tabs[0]);
  const [showLimit, setShowLimit] = React.useState(5);
  const [showFilter, setShowFilter] = React.useState({
    id: 0,
    title: 'Show All'
  });
  const [priceFilter, setPriceFilter] = React.useState({
    id: 0,
    title: 'Inital'
  });

  // TODO: Find out list search filters
  const showOptions: DropdownOption[] = [];
  const priceOptions: DropdownOption[] = [];

  const filteredCourses =
    selectedTab.name !== 'All Courses'
      ? courses.filter((x) => x.type === selectedTab.name)
      : courses;

  return (
    <div className={classNames(classes.root, className)}>
      <div className={classNames(classes.row, classes.tabBar)}>
        {tabs.map((tab: Tab, index: number) => (
          <div
            className={classNames(
              classes.tab,
              index === tabs.length - 1 && classes.noMargin,
              selectedTab === tab && classes.selected
            )}
            onClick={() => setSelectedTab(tab)}
          >
            {tab.name}
          </div>
        ))}
      </div>
      <div className={classes.list}>
        <div className={classes.listOptions}>
          <div
            className={classes.searchText}
          >{`${filteredCourses.length} ${selectedTab.value} Courses Available`}</div>
          <div className={classes.row}>
            <Dropdown
              placeholder="Show All"
              selected={showFilter}
              setSelected={(selected: DropdownOption) =>
                setShowFilter(selected)
              }
              options={showOptions}
            />
            <Dropdown
              className={classes.dropDown}
              placeholder="Initial"
              selected={priceFilter}
              setSelected={(selected: DropdownOption) =>
                setPriceFilter(selected)
              }
              options={priceOptions}
            />
          </div>
        </div>
        {filteredCourses.slice(0, showLimit).map((courseItem: CourseProps) => (
          <div className={classes.listItem}>
            <CourseItem
              title={courseItem.title}
              description={courseItem.description}
              price={courseItem.price}
              type={courseItem.type}
              colour={courseItem.colour}
              imageURL={courseItem.imageURL}
              className={courseItem?.className}
              viewCourse={courseItem.viewCourse}
              addToBasket={courseItem.addToBasket}
            />
          </div>
        ))}
        {showLimit < filteredCourses.length && (
          <div className={classes.row}>
            <div className={classes.line} />
            <Button
              className={classes.button}
              small
              onClick={() => setShowLimit(showLimit + 5)}
            >
              {`Show More (${filteredCourses.length - showLimit})`}
            </Button>
            <div className={classes.line} />
          </div>
        )}
      </div>
    </div>
  );
}

export default CourseSearch;
