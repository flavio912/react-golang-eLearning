import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import Heading from 'components/Heading';
import SearchResultItem from 'components/SearchResultItem';
import Paginator from 'sharedComponents/Paginator';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

const useStyles = createUseStyles((theme: Theme) => ({
  searchRoot: {},
  searchList: {
    borderTop: `1px solid ${theme.colors.borderGrey}`,
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

type Props = {};

function SearchResults({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const results = [
    {
      id: 1,
      title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
      image: 'https://www.gstatic.com/webp/gallery/1.jpg',
      description:
        'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    },
    {
      id: 2,
      title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
      image: 'https://www.gstatic.com/webp/gallery/1.jpg',
      description:
        'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    },
    {
      id: 3,
      title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
      image: 'https://www.gstatic.com/webp/gallery/1.jpg',
      description:
        'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    },
    {
      id: 4,
      title: 'Cargo Manager (CM) – VC, HS, XRY, EDS',
      image: 'https://www.gstatic.com/webp/gallery/1.jpg',
      description:
        'This course is for those who screen air cargo and mail, to provide them with the knowledge and skills needed to deliver effective screening in visual check, hand search…'
    }
  ];
  return (
    <div className={classes.searchRoot}>
      <div className={classes.searchText}>
        <Heading text="Carg" size={'medium'} />
      </div>
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
    </div>
  );
}

export default SearchResults;
