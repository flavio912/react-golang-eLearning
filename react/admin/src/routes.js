/* eslint-disable react/no-multi-comp */
/* eslint-disable react/display-name */
import React, { lazy } from 'react';
import { Redirect } from 'react-router-dom';
import AuthLayout from './layouts/Auth';
import ErrorLayout from './layouts/Error';
import DashboardLayout from './layouts/Dashboard';
import TutorsView from './views/TutorsList';
import OverviewView from './views/Overview';

export default [
  {
    path: '/',
    exact: true,
    component: () => <Redirect to="/dashboard" />
  },
  {
    path: '/auth',
    component: AuthLayout,
    routes: [
      {
        path: '/auth/login',
        exact: true,
        component: lazy(() => import('src/views/Login'))
      },
      {
        component: () => <Redirect to="/errors/error-404" />
      }
    ]
  },
  {
    path: '/errors',
    component: ErrorLayout,
    routes: [
      {
        path: '/errors/error-401',
        exact: true,
        component: lazy(() => import('src/views/Error401'))
      },
      {
        path: '/errors/error-404',
        exact: true,
        component: lazy(() => import('src/views/Error404'))
      },
      {
        path: '/errors/error-500',
        exact: true,
        component: lazy(() => import('src/views/Error500'))
      },
      {
        component: () => <Redirect to="/errors/error-404" />
      }
    ]
  },
  {
    route: '*',
    component: DashboardLayout,
    routes: [
      {
        path: '/dashboard',
        exact: true,
        component: lazy(() => import('src/views/Calendar'))
      },
      {
        path: '/companies',
        exact: true,
        component: lazy(() => import('src/views/CompaniesManagementList'))
      },
      {
        path: '/companies/:id',
        exact: true,
        component: lazy(() => import('src/views/CompanyManagementDetails'))
      },
      {
        path: '/users/:id',
        exact: true,
        component: lazy(() => import('src/views/UserPage'))
      },
      {
        path: '/users/:id/:tab',
        exact: true,
        component: lazy(() => import('src/views/UserPage'))
      },
      {
        path: '/companies/:id/:tab',
        exact: true,
        component: lazy(() => import('src/views/CompanyManagementDetails'))
      },
      {
        path: '/companies/orders',
        exact: true,
        component: lazy(() => import('src/views/OrderManagementList'))
      },
      {
        path: '/management/orders/:id',
        exact: true,
        component: lazy(() => import('src/views/OrderManagementDetails'))
      },
      {
        path: '/tutors',
        exact: true,
        component: TutorsView
      },
      {
        path: '/overview',
        exact: true,
        component: OverviewView
      },
      {
        path: '/settings',
        exact: true,
        component: lazy(() => import('src/views/Settings'))
      },
      {
        path: '/settings/:tab',
        exact: true,
        component: lazy(() => import('src/views/Settings'))
      },
      {
        path: '/social-feed',
        exact: true,
        component: lazy(() => import('src/views/SocialFeed'))
      },
      {
        path: '/getting-started',
        exact: true,
        component: lazy(() => import('src/views/GettingStarted'))
      },
      {
        component: () => <Redirect to="/errors/error-404" />
      }
    ]
  }
];
