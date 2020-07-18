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
  RedirectException
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
import Questions from 'views/Questions';
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
        <Route path="/progress" Component={TrainingProgress} />
        <Route path="/questions" Component={Questions} />
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
