/* eslint-disable react/no-multi-comp */
/* eslint-disable react/display-name */
import React, { lazy } from 'react';
import { Redirect } from 'react-router-dom';
import AuthLayout from './layouts/Auth';
import ErrorLayout from './layouts/Error';
import DashboardLayout from './layouts/Dashboard';
import TutorsView from './views/TutorsList';
import OverviewView from './views/Overview';
import CoursesView from './views/CoursesView';
import CreateCourse from './views/CreateCourse';
import ModulesList from './views/ModulesList';
import QuestionsList from './views/QuestionsList';
import CreateQuestion from './views/Question/CreateQuestion';
import UpdateQuestion from './views/Question/UpdateQuestion';

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
        path: '/admins',
        exact: true,
        component: lazy(() => import('src/views/AdminsList'))
      },
      {
        path: '/admins/:id',
        exact: true,
        component: lazy(() => import('src/views/AdminDetails'))
      },
      {
        path: '/admins/:id/:tab',
        exact: true,
        component: lazy(() => import('src/views/AdminDetails'))
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
        path: '/courses',
        exact: true,
        component: CoursesView
      },
      {
        path: '/course/:ident/:tab',
        exact: true,
        component: CreateCourse
      },
      {
        path: '/courses/create',
        exact: true,
        component: CreateCourse
      },
      {
        path: '/courses/create/:tab',
        exact: true,
        component: CreateCourse
      },
      {
        path: '/modules',
        exact: true,
        component: ModulesList
      },
      {
        path: '/questions',
        exact: true,
        component: QuestionsList
      },
      {
        path: '/questions/create/:tab',
        exact: true,
        component: CreateQuestion
      },
      {
        path: '/question/:ident/:tab',
        exact: true,
        component: UpdateQuestion
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
