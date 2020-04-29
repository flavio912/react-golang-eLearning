import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import SingleUserSearch from "./SingleUserSearch";
import Button from "components/core/Button";
import { ResultItem } from "../UserSearch";
import converter from "number-to-words";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: "flex",
    flexDirection: "column",
    alignItems: "stretch",
  },
  button: {
    alignSelf: "flex-end",
  },
  section: {
    marginBottom: 15,
  },
  nthdelegate: {
    textTransform: "capitalize",
    fontSize: 15,
    margin: [theme.spacing(1), 0],
  },
  remove: {
    fontSize: theme.fontSizes.small,
    color: theme.colors.textGrey,
    marginLeft: theme.spacing(1),
    cursor: "pointer",
    "&:hover": {
      textDecoration: "underline",
    },
  },
}));

type Props = {
  searchFunction: (query: string) => Promise<ResultItem[]>;
  users: (ResultItem | undefined)[];
  setUsers: (users: (ResultItem | undefined)[]) => void;
  style?: React.CSSProperties;
};

function MultiUserSearch({ searchFunction, users, setUsers, style }: Props) {
  const classes = useStyles();

  const setUserByIndex = (index: number, value: ResultItem | undefined) =>
    setUsers(
      users.reduce((arr, curr, i) => [...arr, index === i ? value : curr], [])
    );

  const removeUserByIndex = (index: number) =>
    setUsers(
      users.reduce((arr, curr, i) => (index === i ? arr : [...arr, curr]), [])
    );

  return (
    <div className={classes.container} style={style}>
      {users.map((user, index) => (
        <div className={classes.section}>
          <p className={classes.nthdelegate}>
            {converter.toWordsOrdinal(index + 1)} Delegate
            {index !== 0 && (
              <span
                className={classes.remove}
                onClick={() => removeUserByIndex(index)}
              >
                Remove
              </span>
            )}
          </p>
          <SingleUserSearch
            searchFunction={searchFunction}
            selection={user}
            setSelection={(user) => setUserByIndex(index, user)}
          />
        </div>
      ))}
      {users.slice(-1)[0] !== undefined && (
        <Button
          archetype="grey"
          icon={{ left: "AddDelegateRepeater" }}
          onClick={() => setUsers([...users, undefined])}
          className={classes.button}
          bold
          small
        >
          Add Delegate
        </Button>
      )}
    </div>
  );
}

export default MultiUserSearch;
