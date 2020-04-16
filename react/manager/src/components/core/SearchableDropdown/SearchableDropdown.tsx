import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";
import Icon from "../Icon";

const useStyles = createUseStyles((theme: Theme) => ({
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
    borderRadius: theme.primaryBorderRadius,
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
    borderRadius: theme.primaryBorderRadius,
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
}));

type Props = {
  placeholder: string;
};

const options = [
  {
    title: "Dangerous Goods - Air",
    content: [
      {
        name: "Cargo Operative Screener (COS) – VC, HS, XRY, ETD",
        price: 200,
      },
      {
        name: "Cargo Operative Screener (COS) Recurrent – VC, HS, XRY, ETD",
        price: 65,
      },
    ],
  },
  {
    title: "Known Consignor",
    content: [
      {
        name: "Known Consignor Responsible Person",
        price: 70,
      },
      {
        name: "Known Consignor (Modules 1-7)",
        price: 55,
      },
      {
        name: "Manual Handling Awareness",
        price: 25,
      },
      {
        name: "Fire Safety Awareness",
        price: 25,
      },
      {
        name: "Noise and Vibration Awareness",
        price: 25,
      },
      {
        name: "Seat Belt Misuse Awareness",
        price: 300,
      },
    ],
  },
];

function SearchableDropdown({ placeholder }: Props) {
  const classes = useStyles();
  const [selected, setSelected] = React.useState<string | null>(null);
  const [isOpen, setOpen] = React.useState<boolean>(false);
  const [search, setSearch] = React.useState<string>("");
  const searchRef = React.useRef<HTMLInputElement>(null);

  React.useEffect(() => {
    if (isOpen && searchRef && searchRef.current) {
      searchRef.current.focus();
    }
  }, [isOpen]);

  return (
    <div className={classes.container}>
      <div className={classes.clickableBox} onClick={() => setOpen((o) => !o)}>
        <p>{selected || placeholder}</p>
        {/* Icon goes here */}
      </div>
      {isOpen && (
        <div className={classes.dropdown}>
          <div className={classes.searchContainer}>
            <Icon name="SearchGlass" className={classes.searchIcon} size={20} />
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
                  name.toLocaleLowerCase().includes(search.toLocaleLowerCase())
              );
              if (filtered.length > 0) {
                return (
                  <>
                    <div className={classNames(classes.header, classes.option)}>
                      <span className={classes.title}>{title}</span>
                      <span className={classes.pill}>Catagory</span>
                    </div>
                    {filtered.map(({ name, price }) => (
                      <div className={classes.option}>
                        <span>{name}</span>
                        <span>£{price.toFixed(2)}</span>
                      </div>
                    ))}
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
  );
}

export default SearchableDropdown;
