import * as React from 'react';
//@ts-ignore
import { BrowserProtocol, queryMiddleware } from 'farce';
import {
  createFarceRouter,
  createRender,
  makeRouteConfig,
  Route,
  useRouter,
  RouteRenderArgs
} from 'found';
//@ts-ignore
import { Resolver } from 'found-relay';
import environment, { FetchError } from './api/environment';
import { graphql } from 'react-relay';
import LoginPage from 'views/Login';
import { ThemeProvider } from 'react-jss';
import theme from './helpers/theme';
import AppHolder from 'views/AppHolder';
import OrgOverview from 'views/OrgOverview';
import DelegatesPage from 'views/DelegatesPage';
import { OnlineCoursesPage } from 'views/CoursesPage';
import DelegateProfilePage from 'views/DelegateProfilePage';

import type { AppHolder_manager } from './views/__generated__/AppHolder_manager.graphql';
import { Redirect } from 'react-router-dom';

const protectedRenderer = (Comp: React.ReactNode) => (
  args: RouteRenderArgs
) => {
  // Sadly found-relay has no types...
  //@ts-ignore
  if (args?.error && args?.error.type == 'ErrUnauthorized') {
    args.match.router.push('/login');
    return;
  }
  //@ts-ignore
  return <Comp {...args.props} />;
};

const Router = createFarceRouter({
  historyProtocol: new BrowserProtocol(),
  historyMiddlewares: [queryMiddleware],
  routeConfig: makeRouteConfig(
    <Route>
      <Route path="/(login)?" Component={LoginPage} />
      <Route
        path="/app"
        Component={AppHolder}
        query={graphql`
          query App_Holder_Query {
            manager {
              ...AppHolder_manager
            }
          }
        `}
        render={protectedRenderer(AppHolder)}
      >
        <Route
          path="/"
          Component={OrgOverview}
          query={graphql`
            query App_Org_Query {
              manager {
                ...OrgOverview_manager
              }
            }
          `}
        />
        <Route
          path="/delegates"
          Component={DelegatesPage}
          query={graphql`
            query App_DelegatesPage_Query {
              delegates {
                ...DelegatesPage_delegates
              }
              manager {
                ...DelegatesPage_manager
              }
            }
          `}
        />
        <Route
          path="/delegates/:uuid"
          Component={DelegateProfilePage}
          query={graphql`
            query App_DelegatesProfile_Query($uuid: UUID!) {
              delegate(uuid: $uuid) {
                ...DelegateProfilePage_delegate
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            console.log(params);
            console.log(location);
            const { uuid } = params;
            return {
              uuid
            };
          }}
        />
        <Route
          path="/courses"
          Component={OnlineCoursesPage}
          query={graphql`
            query App_OnlineCourses_Query($page: Page) {
              onlineCourses(page: $page) {
                ...CoursesPage_onlineCourses
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            console.log(params);
            console.log(location);
            const { pageNum, type } = location.query;
            return {
              ...params,
              page: {
                offset: pageNum,
                limit: 10
              }
            };
          }}
        />
        {/* <Route
          path="/courses/classroom"
          Component={CoursesPage}
          query={graphql`
            query App_ClassroomCourses_Query {

            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            console.log(params);
            console.log(location);
            const { page, type } = location.query;
            return {
              page,

            };
          }}
        /> */}
      </Route>
    </Route>
  ),
  render: createRender({})
});

const App = () => (
  <ThemeProvider theme={theme}>
    <Router resolver={new Resolver(environment)} />
  </ThemeProvider>
);

export default App;
