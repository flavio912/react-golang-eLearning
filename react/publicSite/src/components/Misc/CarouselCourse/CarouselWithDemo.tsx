import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from 'classnames';
import { Theme } from "helpers/theme";
import CarouselCourse from ".";
import { Course } from 'sharedComponents/Overview/CourseCard';
import Spacer from "sharedComponents/core/Spacers/Spacer";
import Button from "sharedComponents/core/Input/Button";

const useStyles = createUseStyles((theme: Theme) => ({
    carouselRoot: {
        width: '100%'
    },
    heading: {
        fontSize: 32,
        color: theme.colors.primaryBlack,
        fontWeight: 800,
        padding: '60px 0px',
        textAlign: 'center'
    },
    exploreCont: {
        maxWidth: '100%',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center'
    },
    exploreText: {
        textAlign: 'center',
        fontSize: theme.fontSizes.extraLarge,
        maxWidth: 500
    },
    buttonHolder: {
        display: 'grid',
        gridTemplateColumns: '1fr 75px 1fr',
        gridSpace: '50px',
        '@media (max-width: 575px)': {
            gridTemplateColumns: '1fr',
            gridTemplateRows: '1fr 1fr 1fr',
        }
    },
    text: {
        alignSelf: 'center',
        textAlign: 'center',
        fontSize: theme.fontSizes.extraLarge,
        color: theme.colors.primaryBlack
    },
    button: {
        height: 52,
        fontSize: 18,
        fontWeight: 800,
        boxShadow: '0px 2px 9px #00000014',
        padding: '0px 36px'
    }
}));

type Props = {
    heading: string;
    description: string;
    courses: Course[];
    className?: string;
};

function CarouselWithDemo({ heading, description, courses, className}: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });
  return (
    <div className={classNames(classes.carouselRoot, className)}>
        <div className={classes.heading}>{heading}</div>
        <CarouselCourse
            courses={courses}
        />
        <Spacer vertical spacing={3} />
        <div className={classes.exploreCont}>
        <p className={classes.exploreText}>
            {description}
        </p>
        <Spacer vertical spacing={3} />
        <div className={classes.buttonHolder}>
            <Button archetype="gradient" className={classes.button}>
            Register your Team
            </Button>
            <div className={classes.text}>OR</div>
            <Button archetype="default" className={classes.button}>
            Request Demo
            </Button>
        </div>
        </div>
        <Spacer vertical spacing={3} />
  </div>
  );
}

export default CarouselWithDemo;