import * as React from 'react';
//@ts-ignore
import { BrowserProtocol, queryMiddleware } from 'farce';
import {
  createFarceRouter,
  createRender,
  makeRouteConfig,
  Route,
  RouteRenderArgs,
  RenderErrorArgs,
  RedirectException,
  Redirect
} from 'found';
//@ts-ignore
import { Resolver } from 'found-relay';
import environment from './api/environment';
import { graphql, createFragmentContainer } from 'react-relay';
import LoginPage from 'views/Login';
import { ThemeProvider } from 'react-jss';
import theme from './helpers/theme';
import AppHolder from 'views/AppHolder';
import OnlineCoursePage from 'views/OnlineCourse';
import TrainingZone from 'views/TrainingZone/TrainingZone';
import OnlineCourses from 'views/OnlineCourses';
import CertGenerator from 'views/CertGenerator';
import TrainingProgress from 'views/TrainingProgress';
import ErrorBoundary from 'components/ErrorBoundarys/PageBoundary';
import Module from 'views/Module';
import Test from 'views/Test/Test';

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
            user {
              ...AppHolder_user
            }
          }
        `}
        render={protectedRenderer(AppHolder)}
      >
        <Route path="/" Component={TrainingZone} />
        <Route
          path="/courses"
          Component={OnlineCourses}
          query={graphql`
            query App_Courses_Query {
              user {
                ...OnlineCourses_user
              }
            }
          `}
          render={(args: any) => {
            return <OnlineCourses {...args.props} />;
          }}
        />
        <Route
          path="/courses/:id"
          Component={OnlineCoursePage}
          query={graphql`
            query App_Course_Query($ident: Int!) {
              user {
                myActiveCourse(id: $ident) {
                  ...OnlineCourse_myActiveCourse
                }
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { id } = params;
            return {
              ident: parseInt(id)
            };
          }}
          render={(args: any) => {
            return (
              <ErrorBoundary>
                <OnlineCoursePage
                  {...args.props}
                  myActiveCourse={args.props?.user?.myActiveCourse ?? null}
                />
              </ErrorBoundary>
            );
          }}
        />
        <Route
          path="/courses/:courseID/module/:moduleUUID"
          Component={Module}
          query={graphql`
            query App_Module_Query($id: Int!, $uuid: UUID!) {
              user {
                myActiveCourse(id: $id) {
                  ...Module_myActiveCourse
                }
              }
              module(uuid: $uuid) {
                ...Module_module
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { courseID, moduleUUID } = params;
            return {
              id: parseInt(courseID),
              uuid: moduleUUID
            };
          }}
          render={(args: any) => {
            console.log('args', args);
            if (args.error) {
              args.match.router.push('/app');
            }
            return (
              <ErrorBoundary>
                <Module
                  {...args.props}
                  myActiveCourse={args.props?.user?.myActiveCourse}
                />
              </ErrorBoundary>
            );
          }}
        />
        <Route path="/progress" Component={TrainingProgress} />
        <Route
          path="/courses/:courseID/test/:testUUID"
          Component={Test}
          query={graphql`
            query App_Test_Query($id: Int!, $uuid: UUID!) {
              user {
                myActiveCourse(id: $id) {
                  ...Test_myActiveCourse
                }
              }
              test(uuid: $uuid) {
                ...Test_test
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { courseID, testUUID } = params;
            return {
              id: parseInt(courseID),
              uuid: testUUID
            };
          }}
          render={(args: any) => {
            if (args.error) {
              args.match.router.push('/app');
            }
            return (
              <ErrorBoundary>
                <Test
                  {...args.props}
                  myActiveCourse={args.props?.user?.myActiveCourse}
                />
              </ErrorBoundary>
            );
          }}
        />
      </Route>
      <Route path="/cert-generator" Component={CertGenerator} />
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
