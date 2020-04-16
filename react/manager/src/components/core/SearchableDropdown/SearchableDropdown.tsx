import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";
import Icon from "../Icon";

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

    "&:focus": {
      borderColor: theme.colors.primaryBlue,
      boxShadow: "0 0 3px 1px rgba(14,99,232,0.34)",
    },
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

type Props = {
  multiselect?: boolean;
  placeholder: string;
  selected: string[];
  setSelected: (selected: ((current: string[]) => string[]) | string[]) => void;
  options: {
    title: string;
    content: {
      name: string;
      price: number;
    }[];
  }[];
};

function SearchableDropdown({
  multiselect,
  placeholder,
  selected = [],
  setSelected,
  options,
}: Props) {
  const classes = useStyles();
  const [isOpen, setOpen] = React.useState<boolean>(false);
  const [search, setSearch] = React.useState<string>("");
  const searchRef = React.useRef<HTMLInputElement>(null);

  React.useEffect(() => {
    if (isOpen && searchRef && searchRef.current) {
      searchRef.current.focus();
    }
  }, [isOpen]);

  return (
    <>
      {multiselect && selected.length > 0 && (
        <div className={classes.tagContainer}>
          {selected.map((tag) => (
            <div className={classes.tag}>
              {tag}
              <Icon
                pointer
                name="RemoveSelectedCourse_X"
                className={classes.removeTag}
                size={null} // set in class
                onClick={() => setSelected((s) => s.filter((f) => f !== tag))}
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
          <p>{multiselect ? placeholder : selected[0] || placeholder}</p>
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
              <input
                ref={searchRef}
                placeholder="Search"
                type="text"
                onChange={(evt) => setSearch(evt.target.value)}
                value={search}
                className={classes.searchBox}
              />
            </div>
            <div className={classes.searchResults}>
              {options.map(({ title, content }) => {
                const filtered = content.filter(
                  ({ name }) =>
                    !search ||
                    name
                      .toLocaleLowerCase()
                      .includes(search.toLocaleLowerCase())
                );
                if (filtered.length > 0) {
                  return (
                    <>
                      <div
                        className={classNames(classes.header, classes.option)}
                      >
                        <span className={classes.title}>{title}</span>
                        <span className={classes.pill}>Catagory</span>
                      </div>
                      {filtered.map(
                        ({ name, price }: { name: string; price: number }) => (
                          <div
                            className={classes.option}
                            onClick={() => {
                              if (multiselect) {
                                if (selected.includes(name)) {
                                  setSelected((s) =>
                                    s.filter((f) => f !== name)
                                  );
                                } else {
                                  setSelected((s) => [...s, name]);
                                }
                              } else {
                                setSelected([name]);
                                setOpen(false);
                              }
                            }}
                          >
                            <span>
                              {multiselect &&
                                (selected.includes(name) ? (
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
                        )
                      )}
                    </>
                  );
                } else {
                  return null;
                }
              })}
            </div>
          </div>
        )}
      </div>
    </>
  );
}

export default SearchableDropdown;
