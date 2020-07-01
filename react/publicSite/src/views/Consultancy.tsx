import * as React from "react";
import { createUseStyles } from "react-jss";
import classNames from 'classnames';
import { Theme } from "helpers/theme";
import PageHeader from "components/core/PageHeader";
import FourPanel, { Panel } from "components/core/FourPanel";
import Curve from "components/core/Curve";
import PeopleCurve from "components/core/Curve/PeopleCurve";
import VideoPlayer from "components/core/VideoPlayer";
import { Row } from "components/core/ImageWithText/ImageWithText";
import BrandCard, { Author } from "components/core/Cards/BrandCard";

const useStyles = createUseStyles((theme: Theme) => ({
    consultancyRoot: {
        width: '100%',
        backgroundColor: theme.colors.primaryWhite
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
    textCentered: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start',
    },
    heading: {
        fontSize: theme.fontSizes.heading,
        color: theme.colors.primaryBlack,
        fontWeight: 800,
        margin: '100px 0px',
        textAlign: 'center'
    },
    subheading: {
        fontSize: theme.fontSizes.heading,
        color: theme.colors.primaryBlack,
        fontWeight: 800,
        margin: '100px 0 25px 0',
        textAlign: 'left'
    },
    text: {
        fontSize: theme.fontSizes.tinyHeading,
        color: '#34373A',
        fontWeight: 400,
        textAlign: 'left',
        maxWidth: '750px'
    },
    greyBackground: {
        backgroundColor: theme.colors.backgroundGrey,
        marginTop: '125px',
        paddingBottom: '100px'
    },
    card: {
        flex: 1,
        margin: '15px',
        maxWidth: '400px',
        '@media (max-width: 1150px)': {
            minWidth: '40%',
            margin: '5px'
        }
    },
    lessMargin: {
        margin: '75px 0'
    },
    row: {
        display: 'flex',
        justifyContent: 'center',
        flexWrap: 'wrap'
    },
    //  Thumbnail styles
    video: {
        flex: 2,
        margin: '15px',
        alignItems: 'flex-end',
        boxShadow: '4px 2px 10px -2px rgba(0,0,0,0.06)',
        '@media (max-width: 1000px)': {
            width: '100vw',
            margin: '0 0 15px 0'
        },
        '@media (max-width: 1150px)': {
            boxShadow: 'none'
        }
    },
    thumbnailRoot: {
        alignSelf: 'flex-end',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        width: '100%',
        '@media (max-width: 500px)': {
            height: 0,
            width: 0,
            opacity: 0
        }
    },
    playCircle: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        width: '76px',
        height: '76px',
        borderRadius: '76px',
        backgroundColor: 'white',
        opacity: 0.8
    },
    playTriangle: {
        width: 0,
        height: 0,
        marginLeft: 6,
        borderTop: '13.5px solid transparent',
        borderBottom: '13.5px solid transparent',
        borderLeft: '27px solid #0E66E0'
    }
}));

const defaultIconPanels: Panel[] = [
    {
        title: "Addressing annual reports", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore", iconName: "TTC_Logo_Icon"
    },
    {
        title: "Cutting compliance costs", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore", iconName: "TTC_Logo_Icon"
    },
    {
        title: "Supporting security managers", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore", iconName: "TTC_Logo_Icon"
    },
    {
        title: "Continually evolving processes", desciption: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqu Lorem ipsum dolor sit amet, consectetur adipisci Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore", iconName: "TTC_Logo_Icon"
    }
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
];

const defaultAuthor: Author = {
    name: "James Mchale",
    title: "Head of Compliance",
    quote: "What TTC deliver is world class training at such a pace this industry have never seen before. Highly recommended!"
}

type Props = {};

function Consultancy(props: Props) {
  const classes = useStyles();
  return (
      <div className={classes.consultancyRoot}>
        <PageHeader
            title="Consultancy"
            description="Take advantage of our full dangerous goods safety adviser service."
        />
        <div className={classes.centerer}>
            <div className={classes.centered}>
                <div className={classes.heading}>How we work</div>
                <FourPanel panels={defaultIconPanels} />
            </div>
        </div>
        <div className={classNames(classes.centerer, classes.greyBackground)}>
            <div className={classes.centered}>
                <div className={classNames(classes.heading, classes.lessMargin)}>Don't just take our word for it</div>
                <div className={classes.row}>
                    <VideoPlayer
                        width={853}
                        className={classes.video}
                        source={require('assets/Stock_Video.mp4')}
                        thumbnail={
                            <div className={classes.thumbnailRoot}>
                                <div className={classes.playCircle}>
                                    <div className={classes.playTriangle} />
                                </div>
                                <Curve
                                    width={853}
                                    logo="Dhl"
                                    description="“ TTC Lorem ipsum dolor sit amet, consectetur adipiscing elit, se oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusm.”"
                                    link={{
                                        title: 'Read & watch their story',
                                        link: '/'
                                    }}
                                    blue
                                />
                            </div>
                        }
                    />
                    <BrandCard
                        className={classes.card}
                        logoURL={require("assets/maersk-logo.svg")}
                        link="/"
                        author={defaultAuthor}
                    />
                
                    <BrandCard
                        className={classes.card}
                        logoURL={require("assets/heathrow-logo.png")}
                        link="/"
                        text="TTC Lorem ipsum dolor sit amet, consectetur adipiscing elitse oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consecte"
                    />
                    <BrandCard
                        className={classes.card}
                        logoURL={require("assets/ups-logo.svg")}
                        link="/"
                        text="TTC Lorem ipsum dolor sit amet, consectetur adipiscing elitse oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consecte"
                    />
                    <BrandCard
                        className={classes.card}
                        logoURL={require("assets/nippon-express-logo.png")}
                        link="/"
                        text="TTC Lorem ipsum dolor sit amet, consectetur adipiscing elitse oddo eiusmod tempor incididunt ut labore Lorem ipsum dolor sit amet, consecte"
                    />
                </div>
            </div>
        </div>
        <div className={classes.centerer}>
            <div className={classes.textCentered}>
                <div className={classes.subheading}>DGSA Services</div>
                <div className={classes.text}>TTC offers a full dangerous goods safety adviser service. This includes an audit, telephone and email support as well as completing the annual report, which is a legal requirement.<br/><br/>We also offer a DGSA support service for those companies that need support for their in- house DGSA or who do not require the full DGSA service. We offer competitive rates and work with you to ensure you are compliant whilst also helping you to save money with efficient and effective processes.</div>
                <div className={classes.subheading}>Dangerous Goods Consultancy Services</div>
                <div className={classes.text}>Through our compliance and consultancy services, TTC can help you to ensure that your dangerous goods processes operate in the most cost-effective way as well as being fully compliant with all the relevant dangerous goods transport regulations.<br/><br/>We can undertake a compliance audit of your processes and offer honest and cost-effective solutions, if and when these are required.<br/><br/>We are happy to discuss any consultancy services you may need such as producing standard operating procedures or helping companies to navigate the maze of requirements for the transport of dangerous goods.</div>
                <div className={classes.subheading}>Aviation Security Consultancy Services</div>
                <div className={classes.text}>TTC offers a full security manager service ensuring that your security compliance is taken care of to the highest standard leaving you free to run your business. We can also take a more advisory role supporting your security manager allowing them to manage day to day operations with the safety net of an expert at the end of a phone.</div>
            </div>
        </div>
        <PeopleCurve stack={defaultStack} />
      </div>
  );
}

export default Consultancy;