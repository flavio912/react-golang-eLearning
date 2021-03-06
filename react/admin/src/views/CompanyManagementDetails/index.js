import React from 'react';
import { Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import { Container, Tabs, Tab, Divider, colors } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery, useMutation } from '@apollo/react-hooks';

import Page from 'src/components/Page';
import Header from './Header';
import Summary from './Summary';
import Invoices from './Invoices';
import Managers from './Managers';
import Delegates from './Delegates';
import Logs from './Logs';
import ErrorModal from 'src/components/ErrorModal';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  tabs: {
    marginTop: theme.spacing(3)
  },
  divider: {
    backgroundColor: colors.grey[300]
  },
  content: {
    marginTop: theme.spacing(3)
  }
}));

const GET_COMPANY = gql`
  query GetCompany($id: UUID!) {
    company(uuid: $id) {
      uuid
      name
      approved
      contactEmail
      isContract
      logoURL
      managers {
        edges {
          uuid
          email
          firstName
          lastName
          profileImageUrl
          createdAt
          lastLogin
        }
        pageInfo {
          total
        }
      }
      delegates {
        edges {
          uuid
          email
          firstName
          lastName
          lastLogin
          profileImageUrl
          createdAt
        }
        pageInfo {
          total
        }
      }
      address {
        addressLine1
        addressLine2
        county
        postCode
        country
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

function CompanyManagementDetails({ match, history }) {
  const classes = useStyles();
  const { id, tab: currentTab } = match.params;

  const { loading, error, data, refetch } = useQuery(GET_COMPANY, {
    variables: {
      id: id
    },
    fetchPolicy: 'cache-and-network',
    skip: !id
  });
  const [approveCompany, { mutationErr }] = useMutation(APPROVE_COMPANY);

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;
  const company = data?.company;

  let tabs = [
    { value: 'summary', label: 'Summary' },
    { value: 'Courses', label: 'Courses' },
    { value: 'managers', label: 'Managers' },
    { value: 'delegates', label: 'Delegates' }
  ];

  if (!company.approved) {
    tabs = [
      { value: 'summary', label: 'Summary' },
      { value: 'managers', label: 'Managers' }
    ];
  }

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  if (!currentTab) {
    return <Redirect to={`/companies/${id}/summary`} />;
  }

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const handleUpdatedCompany = () => {
    refetch();
  };

  return (
    <Page className={classes.root} title="Company Management Details">
      <ErrorModal error={mutationErr} />
      <Container maxWidth={false}>
        <Header
          companyName={data.company.name}
          approved={data.company.approved}
          onApprove={async () => {
            approveCompany({
              variables: { uuid: data.company.uuid }
            });
          }}
        />
        <Tabs
          className={classes.tabs}
          onChange={handleTabsChange}
          scrollButtons="auto"
          value={currentTab}
          variant="scrollable"
        >
          {tabs.map(tab => (
            <Tab key={tab.value} label={tab.label} value={tab.value} />
          ))}
        </Tabs>
        <Divider className={classes.divider} />
        <div className={classes.content}>
          {currentTab === 'summary' && (
            <Summary
              company={data.company}
              onUpdate={() => {
                refetch();
              }}
            />
          )}
          {currentTab === 'invoices' && <Invoices />}
          {currentTab === 'managers' && (
            <Managers
              company={company}
              onUpdateCompany={handleUpdatedCompany}
            />
          )}
          {currentTab === 'delegates' && (
            <Delegates
              company={company}
              onUpdateCompany={handleUpdatedCompany}
            />
          )}
          {currentTab === 'logs' && <Logs />}
        </div>
      </Container>
    </Page>
  );
}

CompanyManagementDetails.propTypes = {
  history: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired
};

export default CompanyManagementDetails;
