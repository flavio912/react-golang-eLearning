import React, { useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import Page from 'src/components/Page';
import SearchBar from 'src/components/SearchBar';
import Header from './Header';
import Results from './Results';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  }
}));

const GET_LESSONS = gql`
  query GetLessons($page: Page, $filter: LessonFilter, $orderBy: OrderBy) {
    lessons(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        name
        text
      }
      pageInfo {
        given
        total
      }
    }
  }
`;

function LessonsList({ match, history }) {
  const classes = useStyles();

  const handleFilter = () => {};
  const handleSearch = () => {};

  const { loading, error, data, refetch } = useQuery(GET_LESSONS, {
    variables: {
      page: {
        offset: 0,
        limit: 100
      },
      filter: {}
    }
  });

  useEffect(() => {
    if (!data || loading || error) return;

    refetch();
  }, [data]);

  return (
    <Page className={classes.root} title="Lessons">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/lessons/create/overview');
          }}
        />
        <SearchBar onFilter={handleFilter} onSearch={handleSearch} />
        {data?.lessons && (
          <Results className={classes.results} lessons={data.lessons.edges} />
        )}
      </Container>
    </Page>
  );
}

export default LessonsList;
