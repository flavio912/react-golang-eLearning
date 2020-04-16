import * as React from "react";
import { createUseStyles } from "react-jss";
import ThemeObject, { Theme } from "helpers/theme";
import classNames from "classnames";
import Icon from "../Icon";
import { ScaleLoader } from "react-spinners";
import CoreInput from "../CoreInput";

const useStyles = createUseStyles((theme: Theme) => ({
  tagContainer: {
    width: "100%",
  },
  tag: {
    height: 40,
    display: "inline-flex",
    alignItems: "center",
    flexDirection: "row",
    backgroundColor: theme.colors.backgroundGrey,
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: theme.buttonBorderRadius,
    fontSize: theme.fontSizes.default,
    padding: [0, 15],
    margin: [0, 15, 15, 0],
  },
  removeTag: {
    margin: [0, 0, 5, 15],
    height: 20,
    width: 12,
  },
  container: {
    width: "100%",
    height: 40,
    position: "relative",
    display: "flex",
    flexDirection: "column",
    alignItems: "strech",
  },
  clickableBox: {
    height: 40,
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    padding: [0, 15],
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: theme.buttonBorderRadius,
    cursor: "pointer",
  },
  dropdown: {
    position: "absolute",
    top: 50,
    width: "100%",
    maxHeight: 450, // don't know how this should be calculated
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: theme.primaryBorderRadius,
    display: "flex",
    flexDirection: "column",
  },
  searchContainer: {
    padding: 10,
    display: "flex",
    flexDirection: "column",
    alignItems: "strech",
    position: "relative",
  },
  searchIcon: {
    position: "absolute",
    top: 21,
    right: 21,
  },
  searchBox: {
    height: 40,
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: 3,
    padding: [0, 10],
    fontSize: 14,
    flex: "auto",
    "&:focus": {
      borderColor: theme.colors.primaryBlue,
    },
  },
  tempResults: {
    height: 100,
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    borderTop: `1px solid ${theme.colors.borderGrey}`,
  },
  searchResults: {
    overflowY: "auto",
  },
  option: {
    height: 40,
    cursor: "pointer",
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    padding: [0, 15],
    fontSize: theme.fontSizes.default,
    borderBottom: `1px solid ${theme.colors.borderGrey}`,
    "&:first-child, &:last-child": {
      borderBottom: "none",
    },
  },
  header: {
    backgroundColor: theme.colors.primaryBlue,
    cursor: "default",
  },
  title: {
    textTransform: "uppercase",
    color: "white",
    fontWeight: 900,
  },
  pill: {
    textTransform: "uppercase",
    color: "white",
    fontSize: 9,
    border: "1px solid white",
    borderRadius: 10,
    padding: [2, 10],
  },
  checkbox: {
    marginRight: 15,
  },
  // CSS triangle hack :)
  // https://css-tricks.com/snippets/css/css-triangle/
  triangleDown: {
    width: 0,
    height: 0,
    borderLeft: "5px solid transparent",
    borderRight: "5px solid transparent",
    borderTop: ` ${5 * 2 * 0.866}px solid ${theme.colors.primaryBlue}`,
  },
  triangleUp: {
    width: 0,
    height: 0,
    borderLeft: "5px solid transparent",
    borderRight: "5px solid transparent",
    borderBottom: `${5 * 2 * 0.866}px solid ${theme.colors.primaryBlue}`,
  },
}));

export type CourseCategory = {
  title: string;
  courses: Course[];
};

export type Course = {
  id: number | string; // not sure yet
  name: string;
  price: number;
};

type Props = {
  multiselect?: boolean;
  selected: Course[];
  setSelected: (selected: ((current: Course[]) => Course[]) | Course[]) => void;
  searchQuery: (query: string) => Promise<CourseCategory[]>;
  debounceTime?: number;
};

function SearchableDropdown({
  multiselect,
  selected = [],
  setSelected,
  searchQuery,
  debounceTime = 600,
}: Props) {
  const classes = useStyles();
  const [isOpen, setOpen] = React.useState<boolean>(false);
  const [isLoading, setLoading] = React.useState<boolean>(false);
  const [search, setSearch] = React.useState<string>("");
  const [categories, setCategories] = React.useState<CourseCategory[]>([]);
  const searchRef = React.useRef<HTMLInputElement>(null);
  const debounceRef = React.useRef<number | undefined>();

  const placeholder = multiselect
    ? "Please select one or more Courses"
    : selected.length > 0
    ? selected[0].name
    : "Please select the Course you wish to book";

  React.useEffect(() => {
    if (isOpen && searchRef && searchRef.current) {
      searchRef.current.focus();
    }
  }, [isOpen]);

  React.useEffect(() => {
    if (search !== "") {
      setLoading(true);
      clearTimeout(debounceRef.current);
      debounceRef.current = window.setTimeout(async () => {
        searchQuery(search).then((results) => {
          setCategories(results);
          setLoading(false);
        });
      }, debounceTime);
    } else {
      setLoading(false);
    }
    return () => {
      clearTimeout(debounceRef.current);
    };
  }, [search, debounceTime, searchQuery]);

  return (
    <>
      {multiselect && selected.length > 0 && (
        <div className={classes.tagContainer}>
          {selected.map(({ name, id }: Course) => (
            <div className={classes.tag} key={id}>
              {name}
              <Icon
                pointer
                name="RemoveSelectedCourse_X"
                className={classes.removeTag}
                size={null} // set in class
                onClick={() => setSelected((s) => s.filter((f) => f.id !== id))}
              />
            </div>
          ))}
        </div>
      )}
      <div className={classes.container}>
        <div
          className={classes.clickableBox}
          onClick={() => setOpen((o) => !o)}
        >
          <p>{placeholder}</p>
          <div className={isOpen ? classes.triangleUp : classes.triangleDown} />
        </div>
        {isOpen && (
          <div className={classes.dropdown}>
            <div className={classes.searchContainer}>
              <Icon
                name="SearchGlass"
                className={classes.searchIcon}
                size={20}
              />
              <CoreInput
                ref={searchRef}
                placeholder="Search"
                type="text"
                onChange={(text) => setSearch(text)}
                value={search}
                className={classes.searchBox}
              />
            </div>
            {isLoading ? (
              <div className={classes.tempResults}>
                <ScaleLoader
                  color={ThemeObject.colors.borderGrey}
                  height={20}
                  loading={isLoading}
                />
              </div>
            ) : search === "" ? (
              <div className={classes.tempResults}>
                <p>Start typing to see courses</p>
              </div>
            ) : categories.length === 0 ? (
              <div className={classes.tempResults}>
                <p>No Courses found</p>
              </div>
            ) : (
              <div className={classes.searchResults}>
                {categories.map(({ title, courses }: CourseCategory) => (
                  <div key={title}>
                    <div className={classNames(classes.header, classes.option)}>
                      <span className={classes.title}>{title}</span>
                      <span className={classes.pill}>Category</span>
                    </div>
                    {courses.map(({ name, id, price }: Course) => (
                      <div
                        key={id}
                        className={classes.option}
                        onClick={() => {
                          if (multiselect) {
                            if (
                              selected.map((course) => course.id).includes(id)
                            ) {
                              setSelected((s) => s.filter((f) => f.id !== id));
                            } else {
                              setSelected((s) => [...s, { name, id, price }]);
                            }
                          } else {
                            setSelected([{ name, id, price }]);
                            setSearch("");
                            setOpen(false);
                          }
                        }}
                      >
                        <span>
                          {multiselect &&
                            (selected
                              .map((course) => course.id)
                              .includes(id) ? (
                              <Icon
                                name="FormCheckbox_Checked"
                                className={classes.checkbox}
                                size={17}
                              />
                            ) : (
                              <Icon
                                name="FormCheckbox_Unchecked"
                                className={classes.checkbox}
                                size={17}
                              />
                            ))}
                          {name}
                        </span>
                        <span>£{price.toFixed(2)}</span>
                      </div>
                    ))}
                  </div>
                ))}
              </div>
            )}
          </div>
        )}
      </div>
    </>
  );
}

export default SearchableDropdown;
