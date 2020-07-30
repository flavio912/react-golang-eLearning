/* eslint-disable react/no-multi-comp */
/* eslint-disable react/display-name */
import React, { lazy } from 'react';
import { Redirect } from 'react-router-dom';
import AuthLayout from './layouts/Auth';
import ErrorLayout from './layouts/Error';
import DashboardLayout from './layouts/Dashboard';
import TutorsList from './views/TutorsList';
import CreateTutor from './views/Tutor/CreateTutor';
import UpdateTutor from './views/Tutor/UpdateTutor';
import OverviewView from './views/Overview';
import CoursesView from './views/CoursesView';
import CreateCourse from './views/CreateCourse';
import ModulesList from './views/ModulesList';
import QuestionsList from './views/QuestionsList';
import CreateQuestion from './views/Question/CreateQuestion';
import UpdateQuestion from './views/Question/UpdateQuestion';
import CreateModule from './views/Module/CreateModule';
import UpdateModule from './views/Module/UpdateModule';
import TestsList from './views/TestsList';
import CreateTest from './views/Test/CreateTest';
import UpdateTest from './views/Test/UpdateTest';
import LessonsList from './views/LessonsList';
import CreateLesson from './views/Lesson/CreateLesson';
import UpdateLesson from './views/Lesson/UpdateLesson';
import CertificateTypes from './views/CertificateTypesList';
import CreateCertificateType from './views/CertificateType/CreateCertificateType';
import UpdateCertificateType from './views/CertificateType/UpdateCertificateType';

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
        path: '/approve-companies',
        exact: true,
        component: lazy(() => import('src/views/UnapprovedCompanies'))
      },
      {
        path: '/management/orders/:id',
        exact: true,
        component: lazy(() => import('src/views/OrderManagementDetails'))
      },
      {
        path: '/categories',
        exact: true,
        component: lazy(() => import('src/views/CategoriesList'))
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
        path: '/individuals',
        exact: true,
        component: lazy(() => import('src/views/IndividualsList'))
      },
      {
        path: '/individuals/:id',
        exact: true,
        component: lazy(() => import('src/views/IndividualDetails'))
      },
      {
        path: '/individuals/:id/:tab',
        exact: true,
        component: lazy(() => import('src/views/IndividualDetails'))
      },
      {
        path: '/delegates',
        exact: true,
        component: lazy(() => import('src/views/DelegatesList'))
      },
      {
        path: '/delegates/:id',
        exact: true,
        component: lazy(() => import('src/views/DelegateDetails'))
      },
      {
        path: '/delegates/:id/:tab',
        exact: true,
        component: lazy(() => import('src/views/DelegateDetails'))
      },
      {
        path: '/tutors',
        exact: true,
        component: TutorsList
      },
      {
        path: '/tutor/create/:tab',
        exact: true,
        component: CreateTutor
      },
      {
        path: '/tutor/:ident/:tab',
        exact: true,
        component: UpdateTutor
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
        path: '/tests',
        exact: true,
        component: TestsList
      },
      {
        path: '/test/create/:tab',
        exact: true,
        component: CreateTest
      },
      {
        path: '/test/:ident/:tab',
        exact: true,
        component: UpdateTest
      },
      {
        path: '/modules',
        exact: true,
        component: ModulesList
      },
      {
        path: '/modules/create/:tab',
        exact: true,
        component: CreateModule
      },
      {
        path: '/modules/:ident/:tab',
        exact: true,
        component: UpdateModule
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
        path: '/lessons',
        exact: true,
        component: LessonsList
      },
      {
        path: '/lessons/create/:tab',
        exact: true,
        component: CreateLesson
      },
      {
        path: '/lesson/:ident/:tab',
        exact: true,
        component: UpdateLesson
      },
      {
        path: '/certificateTypes',
        exact: true,
        component: CertificateTypes
      },
      {
        path: '/certificateTypes/create/:tab',
        exact: true,
        component: CreateCertificateType
      },
      {
        path: '/certificateTypes/:ident/:tab',
        exact: true,
        component: UpdateCertificateType
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
