import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import { Row } from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import PageHeader, { ButtonLink } from 'components/core/PageHeader';
import CourseSearch, { Tab } from 'components/CourseSearch';
import { CourseProps } from 'components/CourseSearch/CourseItem';
import TrustedCard from 'components/core/Cards/TrustedCard';
import FourPanel from 'components/core/FourPanel';
import { Panel } from 'components/core/FourPanel/FourPanel';
import PeopleCurve from 'components/core/Curve/PeopleCurve';
import ExpandableListView, {
  ExpandItemType,
} from 'components/Misc/ExpandableListView';
import PageMargin from 'components/core/PageMargin';
import { createFragmentContainer, graphql, fetchQuery } from 'react-relay';
import { Courses_courses } from './__generated__/Courses_courses.graphql';
import { Courses_category } from './__generated__/Courses_category.graphql';
import environment from 'api/environment';
import { CoursesQueryResponse } from './__generated__/CoursesQuery.graphql';
import { useRouter } from 'found';

const useStyles = createUseStyles((theme: Theme) => ({
  courseRoot: {
    width: '100%',
    backgroundColor: theme.colors.primaryWhite,
  },
  heading: {
    fontSize: 32,
    color: theme.colors.primaryBlack,
    fontWeight: 800,
    padding: '20px 0px',
    textAlign: 'center',
  },
  text: {
    fontSize: theme.fontSizes.large,
    fontWeight: 500,
    textAlign: 'center',
    maxWidth: '750px',
  },
  courseSearch: {
    paddingBottom: '30px',
    backgroundColor: '#F7F9FB',
  },
  margin: {
    margin: '80px 0',
  },
  smallMargin: {
    margin: '40px 0',
  },
  marginBottom: {
    marginBottom: '40px',
  },
}));

const defaultButtons: ButtonLink[] = [
  { title: 'Regulated Agents', link: '/' },
  { title: 'Known Consignor', link: '/' },
  { title: 'GSAT', link: '/' },
];

const defaultTabs: Tab[] = [
  {
    name: 'All Courses',
    value: '',
  },
  {
    name: 'Regulated Agents',
    value: 'Regulated Agent',
  },
  {
    name: 'Known Consignor',
    value: 'Known Consignor',
  },
  {
    name: 'GSAT',
    value: 'GSAT',
  },
];

const defaultImagePanels: Panel[] = [
  {
    title: 'Start with this thing',
    desciption:
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.',
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  },
  {
    title: 'Then do this thing',
    desciption:
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.',
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  },
  {
    title: 'Then give this a shot',
    desciption:
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.',
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  },
  {
    title: 'Finish with some of that',
    desciption:
      'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.',
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
  },
];

const defaultStack: Row[] = [
  {
    iconName: 'CourseCertificates',
    text:
      'All of our friendly and knowledgable team are available via email and live chat.',
    link: { title: 'World Class 24x7 Support', link: '/' },
  },
  {
    iconName: 'CourseCertificates',
    text:
      'Stay tuned for regular webinars and live QA sessions with the TTC team.',
    link: { title: 'Webinars and Live Sessions', link: '/' },
  },
  {
    iconName: 'CourseCertificates',
    text:
      'Got a question that needs an immediate answer? Try our knowledge base.',
    link: { title: 'Knowledge Base', link: '/' },
  },
];
const FAQOne: ExpandItemType = {
  id: 0,
  title: 'What do I need to know about Aviation Security?',
  description:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum iplorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut.',
  isExpanded: false,
};
const FAQTwo: ExpandItemType = {
  id: 1,
  title: 'How do I enroll on a Aviation Security Course?',
  description:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum iplorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut.',
  isExpanded: false,
};
const FAQThree: ExpandItemType = {
  id: 2,
  title: 'Are these Courses free?',
  description:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum iplorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut.',
  isExpanded: false,
};
const FAQFour: ExpandItemType = {
  id: 3,
  title: 'How long is each Course?',
  description:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum iplorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut.',
  isExpanded: false,
};
const FAQFive: ExpandItemType = {
  id: 4,
  title: 'What is the meaning of life?',
  description:
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Ex ea commodo consequat.Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum iplorem ipsum Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut.',
  isExpanded: false,
};

const getMore = async (categoryUUID: string, offset: number) => {
  const query = graphql`
    query CoursesQuery($categoryUUID: UUID, $offset: Int!) {
      courses(
        filter: { categoryUUID: $categoryUUID }
        page: { limit: 5, offset: $offset }
      ) {
        edges {
          ident: id
          name
          price
          excerpt
          introduction
          bannerImageURL
        }
        pageInfo {
          total
          offset
          limit
          given
        }
      }
    }
  `;

  const variables = {
    offset,
    categoryUUID,
  };

  const data = (await fetchQuery(environment, query, variables)) as {
    courses: Courses_courses;
  };

  if (!data || !data.courses || !data.courses.edges || !data.courses.pageInfo) {
    console.error('Could not get data', data);
    return undefined;
  }

  return data.courses;
};

type Props = {
  courses: Courses_courses;
  category: Courses_category;
};

function Courses({ courses, category }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const { router } = useRouter();
  const [currentCourses, setCourses] = React.useState<Courses_courses['edges']>(
    courses.edges,
  );
  const [selectedTab, setSelectedTab] = React.useState(defaultTabs[0]);

  React.useEffect(() => {
    setCourses(courses.edges);
  }, [courses]);

  const courseItems = (currentCourses || []).map((course) => ({
    title: course?.name ?? '',
    description: course?.excerpt ?? '',
    price: `£${course?.price}`,
    type: category.name,
    colour: category.color,
    imageURL: course?.bannerImageURL ?? '',
    viewCourse: () => {
      router.push('/course/' + course?.ident);
    },
    addToBasket: () => {},
  }));

  return (
    <div className={classes.courseRoot}>
      <PageHeader
        title={`${category.name} Courses`}
        description="Aviation security is a combination of human and material resources to safeguard civil aviation against unlawful interference. Unlawful interference could be acts of terrorism, sabotage, threat to life and property, communication of false threat, bombing, etc."
        archetype="buttons"
        buttons={defaultButtons}
        history={['Courses', category.name]}
      />
      <Spacer spacing={4} vertical />
      <PageMargin>
        <div className={classNames(classes.heading, classes.smallMargin)}>
          Our four-step training process
        </div>
        <FourPanel
          className={classes.marginBottom}
          panels={defaultImagePanels}
          noBorders
        />
        <div className={classes.margin}>
          <div className={classes.heading}>Explore {category.name} Courses</div>
          <div className={classes.text}>
            We have over 100 {category.name} Courses lorem ipsum dolor sit amet,
            consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
            labore et dolore magna aliqua
          </div>
        </div>
      </PageMargin>
      <CourseSearch
        className={classes.courseSearch}
        tabs={defaultTabs}
        courses={courseItems}
        selectedTab={selectedTab}
        onChangeTab={(tab: Tab) => setSelectedTab(tab)}
        moreToShow={courses.pageInfo?.total != currentCourses?.length}
        totalCourses={courses.pageInfo?.total}
        onMore={async () => {
          const more = await getMore(
            category.uuid ?? '',
            (currentCourses ?? []).length,
          );
          if (more) {
            if (!currentCourses || !more.edges) return;
            setCourses([...currentCourses, ...more.edges]);
          }
        }}
      />
      <PageMargin>
        <div className={classNames(classes.heading, classes.margin)}>
          Frequently Asked Questions about {category.name}
        </div>
        <ExpandableListView
          data={[FAQOne, FAQTwo, FAQThree, FAQFour, FAQFive]}
        />
      </PageMargin>
      <TrustedCard
        text="Trusted by more than 1,000 businesses in 120 countries."
        className={classes.margin}
        noShadow
      />
      <PageMargin>
        <div className={classes.margin}>
          <FloatingVideo
            height={352}
            width={628}
            source={require('assets/Stock_Video.mp4')}
            author={{
              name: 'Kristian Durhuus',
              title: 'Chief Executive Officer',
              quote:
                'TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore .',
            }}
          />
        </div>
      </PageMargin>
      <PeopleCurve stack={defaultStack} />
    </div>
  );
}

export default createFragmentContainer(Courses, {
  courses: graphql`
    fragment Courses_courses on CoursePage {
      edges {
        ident: id
        name
        price
        excerpt
        introduction
        bannerImageURL
      }
      pageInfo {
        total
        offset
        limit
        given
      }
    }
  `,
  category: graphql`
    fragment Courses_category on Category {
      uuid
      name
      color
    }
  `,
});
