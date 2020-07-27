import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';
import Page from 'src/components/Page';
import Header from './Header';
import Results from './Results';
import Filter from 'src/components/Filter';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  },
  filter: {
    marginTop: theme.spacing(3)
  }
}));

const GET_MODULES = gql`
  query SearchModules($name: String!, $page: Page!) {
    modules(filter: { name: $name }, page: $page) {
      edges {
        uuid
        name
        syllabus {
          name
          uuid
        }
        tags {
          uuid
          name
          color
        }
      }
      pageInfo {
        total
        limit
        offset
        given
      }
    }
  }
`;

function useSearchQuery(text, page) {
  const { error, data } = useQuery(GET_MODULES, {
    variables: {
      name: text,
      page: page
    }
  });

  if (!data || !data.modules || !data.modules.edges || !data.modules.pageInfo) {
    console.error('Could not get data', data, error);
    return {
      resultItems: [],
      pageInfo: {
        total: 1,
        offset: 0,
        limit: 4,
        given: 0
      }
    };
  }

  console.log(data);

  const resultItems = data.modules.edges.map(module => ({
    uuid: module?.uuid ?? '',
    name: module?.name ?? '',
    syllabus: module?.syllabus ?? ''
  }));

  const pageInfo = {
    total: data.modules.pageInfo?.total,
    offset: data.modules.pageInfo.offset,
    limit: data.modules.pageInfo.limit,
    given: data.modules.pageInfo.given
  };

  return { resultItems, pageInfo };
}

function ModulesList({ match, history }) {
  const classes = useStyles();

  const [searchText, setSearchText] = React.useState('');
  const [page] = React.useState({ limit: 20, offset: 0 });
  const searchResults = useSearchQuery(searchText, page);

  return (
    <Page className={classes.root} title="Modules">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/modules/create/overview');
          }}
        />
        <Filter
          className={classes.filter}
          onChange={text => setSearchText(text)}
        />
        {searchResults && (
          <Results
            className={classes.results}
            modules={searchResults.resultItems}
          />
        )}
      </Container>
    </Page>
  );
}

export default ModulesList;
