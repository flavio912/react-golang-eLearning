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
  results: ResultItem[];
};

function SearchResults({ results }: Props) {
  const theme = useTheme();
  const classes = useStyles({
    theme
  });
  const [searchText, setSearchText] = React.useState<string>('');
  return (
    <div
      className={classes.searchRoot}
      onClick={(event: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
        event?.stopPropagation();
      }}
    >
      <SearchInput
        placeholder="Search for Courses..."
        onChange={setSearchText}
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
