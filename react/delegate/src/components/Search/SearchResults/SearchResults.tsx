import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import SearchResultItem from 'components/Search/SearchResultItem';
import SearchInput from 'components/Search/SearchInput';
import Paginator from 'sharedComponents/Pagination/Paginator';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import { Record } from 'relay-runtime';

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
export type PageInfo = {
  totalPages: number;
  offset: number;
  limit: number;
  totalItems: number;
};
export type Result = {
  resultItems: ResultItem[];
  pageInfo: PageInfo;
};

type Props = {
  searchFunction: (query: string) => Promise<Result>
  debounceTime?: number;
};

function SearchResults({ searchFunction, debounceTime = 400 }: Props) {
  const theme = useTheme();
  const classes = useStyles({
    theme
  });
  const [searchText, setSearchText] = React.useState<string>('');

  const [results, setResults]: [ResultItem[], any] = React.useState([]);
  const [pageInfo, setPageInfo]: [PageInfo | undefined, any] = React.useState();
  const [debouncer, setDebouncer]: [number | undefined, any] = React.useState();
  
  const onChange = (text: string) => {
    clearTimeout(debouncer);
    const timeout = setTimeout(async () => {
      setSearchText(text);
      if (text.length === 0 && results.length > 0){
        return;
      }
      const res = await searchFunction(text);
      
      setResults(res.resultItems);
      setPageInfo(res.pageInfo);
    }, debounceTime);
    setDebouncer(timeout);
  };

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
            {results.map((item, index) => (
              <SearchResultItem course={item} key={index} onClick={() => {}} />
            ))}
          </div>
          <Spacer vertical spacing={2} />
          <Paginator
            currentPage={pageInfo?.offset ?? 0}
            updatePage={() => {}}
            numPages={pageInfo?.totalPages ?? 10}
            itemsPerPage={10}
          />
        </>
      )}
    </div>
  );
}

export default SearchResults;
