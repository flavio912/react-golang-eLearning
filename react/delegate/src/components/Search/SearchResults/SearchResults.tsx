import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import SearchResultItem from 'components/Search/SearchResultItem';
import SearchInput from 'components/Search/SearchInput';
import Paginator from 'sharedComponents/Pagination/Paginator';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

const useStyles = createUseStyles((theme: Theme) => ({
  searchRoot: {
    padding: [70, 60, 59, 91],
    backgroundColor: theme.colors.primaryWhite
  },
  searchList: {
    paddingTop: 7
  },
  searchText: {
    '& h1': {
      lineHeight: `25px`,
      letterSpacing: -0.5,
      color: theme.colors.primaryBlack,
      marginBottom: 30
    }
  }
}));
export type ResultItem = {
  id: string | number;
  title: string;
  image: string;
  description: string;
};
type Props = {
  results?: ResultItem[];
  searchFunction?: (query: string) => Promise<ResultItem[]>
};

function SearchResults({ results, searchFunction }: Props) {
  const theme = useTheme();
  const classes = useStyles({
    theme
  });
  const [searchText, setSearchText] = React.useState<string>('');

  const [showResults, setShowResults]: [ResultItem[], any] = React.useState(results ? results : []);
  const [debouncer, setDebouncer]: [number | undefined, any] = React.useState();

  const onChange = searchFunction ? (text: string) => {
    clearTimeout(debouncer);
    const timeout = setTimeout(async () => {
      if (text.length === 0 && showResults.length > 0){
        return;
      }
      const res = await searchFunction(text);
      
      setSearchText(text);
      setShowResults(res);
    });
    setDebouncer(timeout);
  } : setSearchText;

  return (
    <div
      className={classes.searchRoot}
      onClick={(event: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
        event?.stopPropagation();
      }}
    >
      <SearchInput
        placeholder="Search for Courses..."
        onChange={onChange}
        value={searchText}
      />
      {searchText && (
        <>
          <div className={classes.searchList}>
            {showResults.map((item, index) => (
              <SearchResultItem course={item} key={index} onClick={() => {}} />
            ))}
          </div>
          <Spacer vertical spacing={2} />
          <Paginator
            currentPage={1}
            updatePage={() => {}}
            numPages={10}
            itemsPerPage={10}
          />
        </>
      )}
    </div>
  );
}

export default SearchResults;
