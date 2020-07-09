import * as React from "react";
import classNames from 'classnames';
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import CarouselImage, { Image } from "components/Misc/CarouselImage";
import Button from "sharedComponents/core/Input/Button";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        height: '100vh',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'space-around',
        padding: '70px 0'
    },
    carousel: {
        width: '40vw',
        maxWidth: '668px'
    },
    image: {
        height: '68px',
        width: '178px',
        margin: '25px 0 10px 0'
    },
    text: {
        fontSize: '28px',
        fontWeight: '800',
        color: theme.colors.primaryWhite,
        textAlign: 'center',
        maxWidth: '525px'
    },
    button: {
        marginTop: '20px',
        height: '62px',
        width: '188px',
        fontSize: theme.fontSizes.large,
        fontWeight: '800',
        color: theme.colors.primaryBlack,
        boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)'
    }
}));

type Props = {
    images: Image[];
    onBook: () => void;
    className?: string;
};

function RegistrationCarousel({ images, onBook, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
      <div className={classNames(classes.root, className)}>
        <CarouselImage className={classes.carousel} images={images} />
        <img
            className={classes.image}
            src={require('assets/CustomerChampions.svg')}
        />
        <div className={classes.text}>Still not sure? Our customer success team are here to help</div>
        <Button
            className={classes.button}
            onClick={onBook}
        >
            Book a Demo
        </Button>
      </div>
  );
}

export default RegistrationCarousel;