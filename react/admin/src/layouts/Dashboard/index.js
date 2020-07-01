import React, { Suspense, useState } from 'react';
import { renderRoutes } from 'react-router-config';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import { LinearProgress } from '@material-ui/core';
import NavBar from './NavBar';
import TopBar from './TopBar';
import { Redirect } from 'react-router';

const useStyles = makeStyles(theme => ({
  container: {
    minHeight: '100vh',
    display: 'flex',
    '@media all and (-ms-high-contrast:none)': {
      height: 0 // IE11 fix
    }
  },
  content: {
    paddingTop: 64,
    flexGrow: 1,
    maxWidth: '100%',
    overflowX: 'hidden',
    [theme.breakpoints.up('lg')]: {
      paddingLeft: 256
    },
    [theme.breakpoints.down('xs')]: {
      paddingTop: 56
    }
  }
}));

const GET_ADMIN = gql`
  query GetAdmin {
    admin {
      email
    }
  }
`;

function Dashboard({ route }) {
  const classes = useStyles();
  const [openNavBarMobile, setOpenNavBarMobile] = useState(false);

  // Try get login
  const { loading, error, data } = useQuery(GET_ADMIN);
  if (loading) return <div></div>;
  if (error) {
    return <Redirect to={'/auth/login'} />;
  }

  return (
    <>
      <TopBar onOpenNavBarMobile={() => setOpenNavBarMobile(true)} />
      <NavBar
        onMobileClose={() => setOpenNavBarMobile(false)}
        openMobile={openNavBarMobile}
      />
      <div className={classes.container}>
        <div className={classes.content}>
          <Suspense fallback={<LinearProgress />}>
            {renderRoutes(route.routes)}
          </Suspense>
        </div>
      </div>
    </>
  );
}

Dashboard.propTypes = {
  route: PropTypes.object
};

export default Dashboard;
