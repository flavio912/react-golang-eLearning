import * as React from 'react';
//@ts-ignore
import { BrowserProtocol, queryMiddleware } from 'farce';
import {
  createFarceRouter,
  createRender,
  makeRouteConfig,
  Route,
  useRouter,
  RouteRenderArgs,
  RouterRenderArgs
} from 'found';
//@ts-ignore
import { Resolver } from 'found-relay';
import environment, { FetchError } from './api/environment';
import { graphql } from 'react-relay';
import LoginPage from 'views/Login';
import RecoverPassword from 'views/RecoverPassword';
import { ThemeProvider } from 'react-jss';
import theme from './helpers/theme';
import AppHolder from 'views/AppHolder';
import OrgOverview from 'views/OrgOverview';
import DelegatesPage from 'views/DelegatesPage';
import CoursesPage from 'views/CoursesPage';
import DelegateProfilePage from 'views/DelegateProfilePage';
import DelegateProfilePageFrag from 'views/DelegateProfilePage';
import ErrorBoundary from 'components/ErrorBoundarys/PageBoundary';

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
      <Route path="/(password)?" Component={RecoverPassword} />
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
            query App_DelegatesPage_Query($offset: Int, $limit: Int) {
              delegates(page: { offset: $offset, limit: $limit }) {
                ...DelegatesPage_delegates
              }
              manager {
                ...DelegatesPage_manager
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { offset, limit } = location.query;
            return {
              ...params,
              page: {
                offset: offset,
                limit: limit,
              }
            }
          }}
        />
        <Route
          path="/delegates/:uuid"
          Component={DelegateProfilePage}
          query={graphql`
            query App_DelegatesProfile_Query($uuid: UUID!, $offset: Int, $limit: Int) {
              delegate(uuid: $uuid) {
                activity(page: { offset: $offset, limit: $limit }) {
                  ...DelegateProfilePage_activity
                }
                ...DelegateProfilePage_delegate
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            console.log(params);
            console.log(location);
            const { uuid } = params;
            const { offset, limit } = location.query;
            return {
              uuid,
              offset,
              limit
            };
          }}
          render={(args: any) => {
            if (args.error && args.error != null) {
              if (args.error.type == 'ErrUnauthorized'){
                args.match.router.push('/login');
              }else {
                args.match.router.push('/app/delegates');
              }
              return;
            }
            if (!args.props) {
              return <div></div>;
            }

            return (
              <ErrorBoundary>
                <DelegateProfilePageFrag
                  {...args.props}
                  activity={args.props?.delegate.activity}
                  delegate={args.props?.delegate}
                />
              </ErrorBoundary>
            )
          }}
        />
        <Route
          path="/courses"
          Component={CoursesPage}
          query={graphql`
            query App_Courses_Query($page: Page) {
              courses(page: $page) {
                ...CoursesPage_courses
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            console.log(params);
            console.log(location);
            const { offset, limit } = location.query;
            return {
              ...params,
              page: {
                offset: offset,
                limit: limit
              }
            };
          }}
        />
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
