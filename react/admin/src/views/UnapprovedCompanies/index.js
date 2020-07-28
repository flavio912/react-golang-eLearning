import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import Page from 'src/components/Page';
import SearchBar from 'src/components/SearchBar';
import Header from './Header';
import Results from './Results';
import { gql } from 'apollo-boost';
import { useQuery, useMutation } from '@apollo/react-hooks';
import ErrorModal from 'src/components/ErrorModal';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  }
}));

const GET_COMPANIES = gql`
  query GetCompanies($page: Page, $filter: CompanyFilter) {
    companies(page: $page, filter: $filter) {
      edges {
        uuid
        name
        createdAt
      }
      pageInfo {
        given
        total
      }
    }
  }
`;

const APPROVE_COMPANY = gql`
  mutation ApproveCompany($uuid: UUID!) {
    updateCompany(input: { uuid: $uuid, approved: true }) {
      uuid
    }
  }
`;

function UnapprovedCompanies({ match, history }) {
  const classes = useStyles();

  // const handleFilter = () => {};

  const handleSearch = () => {};

  const { loading, error, data } = useQuery(GET_COMPANIES, {
    variables: {
      page: {
        offset: 0,
        limit: 100
      },
      filter: {
        approved: false
      },
      orderBy: {
        ascending: false,
        field: 'created_at'
      }
    }
  });
  const [approveCompany, { mutationErr }] = useMutation(APPROVE_COMPANY);

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;

  return (
    <Page className={classes.root} title="Company Requests">
      <ErrorModal error={error || mutationErr} />
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/approve-companies/create/overview');
          }}
        />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        {data.companies && (
          <Results
            className={classes.results}
            companies={data.companies.edges}
            onApprove={async uuid => {
              await approveCompany({
                variables: { uuid }
              });
            }}
          />
        )}
      </Container>
    </Page>
  );
}

export default UnapprovedCompanies;
