import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Button } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import Page from 'src/components/Page';
import Results from 'src/components/Results';
import Header from './Header';
import CategorySaveModal from './CategorySaveModal';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  },
  colorBox: {
    width: 60,
    height: 30
  }
}));

const GET_CATEGORIES = gql`
  query GetCategories($name: String, $page: Page) {
    categories(text: $name, page: $page) {
      edges {
        uuid
        name
        color
      }
      pageInfo {
        total
        offset
        limit
        given
      }
    }
  }
`;

function CategoriesList() {
  const classes = useStyles();

  const [modal, setModal] = React.useState({
    open: false,
    categoryUUID: undefined
  });
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data, refetch } = useQuery(GET_CATEGORIES, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage
      }
    },
    fetchPolicy: 'cache-and-network'
  });

  if (error) return <div>{error.message}</div>;

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  // Results table
  const headers = ['Name', 'Color', 'Actions'];
  const cells = [
    {
      field: 'name'
    },
    {
      component: result => (
        <div>
          <div
            style={{ backgroundColor: result.color }}
            className={classes.colorBox}
          ></div>
        </div>
      )
    },
    {
      component: result => (
        <Button
          variant={'outlined'}
          size="small"
          color="primary"
          onClick={() => setModal({ open: true, categoryUUID: result.uuid })}
        >
          Edit
        </Button>
      )
    }
  ];

  return (
    <Page className={classes.root} title="Categories">
      <Container maxWidth={false}>
        <Header onAddCategory={() => setModal({ open: true })} />
        <Results
          className={classes.results}
          results={data?.categories}
          headers={headers}
          cells={cells}
          noPagination
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
        <CategorySaveModal
          onClose={() => setModal({ open: false })}
          onSave={() => {
            setModal({ open: false });
            refetch();
          }}
          categoryUUID={modal.categoryUUID}
          open={modal.open}
        />
      </Container>
    </Page>
  );
}

export default CategoriesList;
