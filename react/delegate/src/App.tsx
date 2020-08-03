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
import FinaliseLogin from 'views/FinaliseLogin';
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
import { SideModalProvider } from 'views/SideModalProvider';
import RecoverPassword from 'views/RecoverPassword/RecoverPassword';
import Lesson from 'views/Lesson';

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
      <Route path="/password" Component={RecoverPassword} />
      <Route path="/finalise/:token" Component={FinaliseLogin} />
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
        <Route
          path="/"
          Component={TrainingZone}
          query={graphql`
            query App_TrainingZone_Query {
              user {
                ...TrainingZone_user
              }
            }
          `}
        />
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
            return (
              <ErrorBoundary>
                <OnlineCourses {...args.props} />
              </ErrorBoundary>
            );
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
            if (!args.props) {
              <div></div>;
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
        <Route
          path="/progress"
          Component={TrainingProgress}
          query={graphql`
            query App_Progress_Query($offset: Int, $limit: Int) {
              user {
                activity(page: { offset: $offset, limit: $limit }) {
                  ...TrainingProgress_activity
                }
                ...TrainingProgress_user
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { offset, limit } = location.query;
            return {
              offset: offset,
              limit: limit
            };
          }}
          render={(args: any) => {
            if (args.error) {
              console.log(args.error);
            }
            if (!args.props) {
              return <div></div>;
            }
            return (
              <ErrorBoundary>
                <TrainingProgress
                  {...args.props}
                  activity={args.props?.user?.activity}
                  myCourses={args.props?.user?.myCourses}
                />
              </ErrorBoundary>
            );
          }}
        />
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
            console.log('args', args);
            if (args.error) {
              args.match.router.push('/app');
            }
            if (!args.props) {
              return <div></div>;
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
        <Route
          path="/courses/:courseID/lesson/:lessonUUID"
          Component={Lesson}
          query={graphql`
            query App_Lesson_Query($id: Int!, $uuid: UUID!) {
              user {
                myActiveCourse(id: $id) {
                  ...Lesson_myActiveCourse
                }
              }
              lesson(uuid: $uuid) {
                ...Lesson_lesson
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { courseID, lessonUUID } = params;
            return {
              id: parseInt(courseID),
              uuid: lessonUUID
            };
          }}
          render={(args: any) => {
            console.log('args', args);
            if (args.error) {
              args.match.router.push('/app');
            }
            if (!args.props) {
              return <div></div>;
            }
            return (
              <ErrorBoundary>
                <Lesson
                  {...args.props}
                  myActiveCourse={args.props?.user?.myActiveCourse}
                />
              </ErrorBoundary>
            );
          }}
        />
      </Route>
      <Route
        path="/cert-generator"
        Component={CertGenerator}
        query={graphql`
          query App_Certificate_Query($token: String!) {
            certificateInfo(token: $token) {
              ...CertGenerator_certificateInfo
            }
          }
        `}
        render={(args: any) => {
          console.log('argts', args);
          if (!args.props) {
            return <div></div>;
          }
          return (
            <ErrorBoundary>
              <CertGenerator {...args.props} />
            </ErrorBoundary>
          );
        }}
        prepareVariables={(params: any, { location }: any) => {
          const { token } = location.query;
          return {
            token
          };
        }}
      />
    </Route>
  ),
  render: createRender({})
});

const App = () => (
  <ThemeProvider theme={theme}>
    <SideModalProvider>
      <Router resolver={new Resolver(environment)} />
    </SideModalProvider>
  </ThemeProvider>
);

export default App;
