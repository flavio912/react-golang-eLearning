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
} from 'found';
//@ts-ignore
import { Resolver } from 'found-relay';
import environment from './api/environment';
import { graphql, createFragmentContainer } from 'react-relay';
import { ThemeProvider } from 'react-jss';
import theme from './helpers/theme';
import AppHolder from 'views/AppHolder';
import { Redirect } from 'react-router-dom';
import Home from 'views/Home';
import AboutUs from 'views/AboutUs';
import PaymentPage from 'views/PaymentPage';
import RegisterStart from 'views/RegisterStart';
import RegisterIndividual from 'views/RegisterIndividual';
import RegisterCompany from 'views/RegisterCompany';
import Courses from 'views/Courses';
import Consultancy from 'views/Consultancy';
import CourseDetailsPage from 'views/CourseDetailsPage';
import PrivacyPolicy from 'views/PrivacyPolicy';
import ContactUs from 'views/ContactUs';
import RegisterCalendar from 'views/RegisterCalendar';
import ArticleLandingPage from 'views/ArticleLandingPage';
import Article from 'views/Article';

const Router = createFarceRouter({
  historyProtocol: new BrowserProtocol(),
  historyMiddlewares: [queryMiddleware],
  routeConfig: makeRouteConfig(
    <Route>
      <Route
        Component={AppHolder}
        //query={ExamplePageQuery}
        query={graphql`
          query App_AppHolder_Query {
            categories {
              ...AppHolder_categories
            }
          }
        `}
        render={({ props, error }: any) => {
          window.scrollTo(0, 0);
          if (!props) {
            return <div></div>;
          }
          // Check if user is logged in, if not redirect to login
          // if (props?.manager) return <AppHolder {...props} />;
          // if (error) {
          //   throw new RedirectException("/login");
          // }
          // return undefined;
          return <AppHolder {...props} />;
        }}
      >
        <Route path="/" Component={Home} />
        <Route path="/aboutus" Component={AboutUs} />
        <Route path="/payment" Component={PaymentPage} />
        <Route
          path="/course/:id"
          Component={CourseDetailsPage}
          query={graphql`
            query App_CourseDetailsPage_Query($id: Int!) {
              course(id: $id) {
                ...CourseDetailsPage_course
              }
            }
          `}
          prepareVariables={(params: any) => {
            const { id } = params;
            return {
              id: parseInt(id),
            };
          }}
        />
        <Route
          path="/courses/:categoryUUID"
          Component={Courses}
          query={graphql`
            query App_Courses_Query(
              $page: Page
              $filter: CoursesFilter
              $categoryUUID: UUID!
            ) {
              courses(page: $page, filter: $filter) {
                ...Courses_courses
              }
              category(uuid: $categoryUUID) {
                ...Courses_category
              }
            }
          `}
          prepareVariables={(params: any, { location }: any) => {
            const { categoryUUID } = params;
            return {
              ...params,
              page: {
                offset: 0,
                limit: 5,
              },
              filter: {
                categoryUUID,
              },
              categoryUUID,
            };
          }}
        />
        <Route path="/consultancy" Component={Consultancy} />
        <Route path="/contact" Component={ContactUs} />
        <Route path="/privacypolicy" Component={PrivacyPolicy} />
        <Route path="/articles" Component={ArticleLandingPage} />
        <Route path="/article" Component={Article} />
      </Route>
      <Route path="/register" Component={RegisterStart} />
      <Route path="/register/individual" Component={RegisterIndividual} />
      <Route path="/register/company" Component={RegisterCompany} />
      <Route path="/register/calendar" Component={RegisterCalendar} />
    </Route>,
  ),
  render: createRender({}),
});

const App = () => (
  <ThemeProvider theme={theme}>
    <Router resolver={new Resolver(environment)} />
  </ThemeProvider>
);

export default App;
