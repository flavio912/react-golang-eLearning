import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import ImageWithText, { Row } from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import PageHeader, { ButtonLink } from 'components/core/PageHeader';
import CourseSearch, { Tab } from 'components/CourseSearch';
import { CourseProps } from 'components/CourseSearch/CourseItem';
import TrustedCard from 'components/core/Cards/TrustedCard';
import FourPanel from 'components/core/FourPanel';
import { Panel } from 'components/core/FourPanel/FourPanel';

const useStyles = createUseStyles((theme: Theme) => ({
  courseRoot: {
    width: '100%'
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center'
  },
  centered: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    width: theme.centerColumnWidth,
  },
  heading: {
    fontSize: 30,
    color: theme.colors.primaryBlack,
    fontWeight: 800,
    padding: '20px 0px',
    textAlign: 'center'
  },
  text: {
    fontSize: theme.fontSizes.large,
    fontWeight: 500,
    textAlign: 'center',
    maxWidth: '750px'
  },
  margin: {
    margin: '80px 0',
  },
  smallMargin: {
    margin: '40px 0',
  },
  marginBottom: {
      marginBottom: '50px'
  }
}));

const defaultButtons: ButtonLink[] = [
    { title: "Regulated Agents", link: "/"},
    { title: "Known Consignor", link: "/"},
    { title: "GSAT", link: "/"},
  ]

  const defaultTabs: Tab[] = [
    {
        name: 'All Courses', value: ''
    },
    {
        name: 'Regulated Agents', value: 'Regulated Agent'
    },
    {
        name: 'Known Consignor', value: 'Known Consignor'
    },
    {
        name: 'GSAT', value: 'GSAT'
    }
];

const defaultImagePanels: Panel[] = [
    {
        title: "Start with this thing", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    },
    {
        title: "Then do this thing", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    },
    {
        title: "Then give this a shot", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    },
    {
        title: "Finish with some of that", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim.", imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png')
    }
];

const defaultCourseItem: CourseProps = {
    title: "Cargo Manager (CM) – VC, HS, XRY, EDS",
    description: "We have over 100 Aviation Security Courses lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
    price: "£310.00",
    type: "Regulated Agents",
    colour: "#8C1CB4",
    imageURL: require('assets/SampleImage_ClassroomCoursesDetail_Feat.png'),
    viewCourse: () => console.log("View Course"),
    addToBasket: () => console.log("Add to Basket")
};

const defaultCourseItems: CourseProps[] = [
    defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem,defaultCourseItem
];

const defaultStack: Row[] = [
    {
        iconName: "CourseCertificates", text: "All of our friendly and knowledgable team are available via email and live chat.",
        link: { title: "World Class 24x7 Support", link: "/"}
    },
    {
        iconName: "CourseCertificates", text: "Stay tuned for regular webinars and live QA sessions with the TTC team.",
        link: { title: "Webinars and Live Sessions", link: "/"}
    },
    {
        iconName: "CourseCertificates", text: "Got a question that needs an immediate answer? Try our knowledge base.",
        link: { title: "Knowledge Base", link: "/"}
    },
]

type Props = {};

function Courses({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [selectedTab, setSelectedTab] = React.useState(defaultTabs[0]);

  return (
    <div className={classes.courseRoot}>
        <PageHeader
            title="Aviation Security Courses"
            description="Aviation security is a combination of human and material resources to safeguard civil aviation against unlawful interference. Unlawful interference could be acts of terrorism, sabotage, threat to life and property, communication of false threat, bombing, etc."
            archetype="buttons"
            buttons={defaultButtons}
            history={["Courses", "Aviation Security"]}
        />
        <Spacer spacing={4} vertical />
        <div className={classes.centerer}>
            <div className={classes.centered}>
                <div className={classNames(classes.heading, classes.smallMargin)}>Our four-step training process</div>
                <FourPanel
                    className={classes.marginBottom}
                    panels={defaultImagePanels}
                    noBorders
                />
                <div className={classes.margin}>
                    <div className={classes.heading}>Explore Aviation Security Courses</div>
                    <div className={classes.text}>We have over 100 Aviation Security Courses lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua</div>
                </div>
            </div>
        </div>
        <CourseSearch 
            tabs={defaultTabs}
            courses={defaultCourseItems}
            selectedTab={selectedTab}
            onChangeTab={(tab: Tab) => setSelectedTab(tab)}
            moreToShow={true}
            onMore={() => console.log("More")}
        />
        <div className={classes.centerer}>
            <div className={classes.centered}>
                <TrustedCard
                    text="Trusted by more than 1,000 businesses in 120 countries."
                    className={classes.margin}
                    noShadow
                />
                <div className={classes.margin}>
                    <FloatingVideo
                        
                        height={352}
                        width={628}
                        source={require('assets/Stock_Video.mp4')}
                        author={{
                            "name": "Kristian Durhuus",
                            "title": "Chief Executive Officer",
                            "quote": "TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore ."
                        }}
                    />
                </div>
                <div className={classes.heading}>Our people make the difference.</div>
                <div className={classes.text}>Not only do we offer incredible training, but our customer service is world-class too</div>
                <ImageWithText
                    className={classes.margin}
                    image={require("assets/StockUKTeam.svg")}
                    stack={defaultStack}
                />
            </div>
        </div>
    </div>
  );
}

export default Courses;
