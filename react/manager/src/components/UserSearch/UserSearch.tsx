import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";

import { Theme } from "helpers/theme";
import Icon from "components/core/Icon";
import CoreInput from "components/core/CoreInput";
import UserResults from "./UserResults";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    border: [1, "solid", theme.colors.borderGrey],
    borderRadius: theme.primaryBorderRadius,
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    background: "white",
  },
  rootFocused: {
    extend: "root",
    boxShadow: theme.shadows.primary,
  },
  search: {
    display: "flex",
    alignItems: "center",
    flex: 1,
    padding: theme.spacing(1),
  },
  searchInput: {
    fontSize: theme.fontSizes.default,
    fontWeight: 200,
    paddingLeft: theme.spacing(1),
  },
  openDelegate: {
    background: theme.colors.primaryBlue,
    color: "white",
    alignItems: "center",
    padding: [0, 38],
    display: "flex",
    fontWeight: 600,
    fontSize: theme.fontSizes.default,
    borderRadius: [0, theme.primaryBorderRadius, theme.primaryBorderRadius, 0],
    cursor: "pointer",
  },
  buttonCont: {
    display: "flex",
    justifyContent: "space-between",
  },
  ttcID: {},
}));

export type ResultItem = {
  value: string;
  key: string;
};

type Props = {
  companyName: string;
  searchFunction: (query: string) => Promise<ResultItem[]>;
  debounceTime?: number;
};

function UserSearch({
  companyName,
  searchFunction,
  debounceTime = 600,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [selection, setSelection]: [
    ResultItem | undefined,
    any
  ] = React.useState();
  const [loading, setLoading]: [boolean, any] = React.useState(false);
  const [focus, setFocus]: [boolean, any] = React.useState(false);
  const [results, setResults]: [ResultItem[], any] = React.useState([]);
  const [input, setInput]: [string, any] = React.useState("");
  const [debouncer, setDebouncer]: [number | undefined, any] = React.useState();

  const onFocus = () => {
    setFocus(true);
    onChange("");

    // Reset the input when focusing again
    if (selection !== undefined) {
      setInput("");
      setSelection(undefined);
    }
  };
  const onBlur = () => {
    setFocus(false);
  };
  const onChange = (text: string) => {
    setLoading(true);
    clearTimeout(debouncer);
    const timeout = setTimeout(async () => {
      console.log("CALLED");
      if (text.length === 0 && results.length > 0) {
        return;
      }
      const res = await searchFunction(text);
      setLoading(false);
      setResults(res);
    }, debounceTime);
    setDebouncer(timeout);
  };
  const onSelect = (result: ResultItem) => {
    setFocus(false);
    setSelection(result);
    setInput(result.key);
  };

  return (
    <div className={focus ? classes.rootFocused : classes.root}>
      <div className={classes.buttonCont}>
        <div className={classes.search}>
          <Icon name={"SearchGlass"} size={15} />
          <CoreInput
            type="search"
            placeholder={`Search within ${companyName}`}
            onChange={onChange}
            className={classes.searchInput}
            onFocus={onFocus}
            onBlur={onBlur}
            value={input}
            setValue={setInput}
          />
        </div>
        {selection && <div className={classes.openDelegate}>Show Delegate</div>}
      </div>
      {focus && (
        <UserResults loading={loading} results={results} onSelect={onSelect} />
      )}
    </div>
  );
}

export default UserSearch;
