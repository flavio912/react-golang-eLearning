import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Carousel from 'react-multi-carousel';
import 'react-multi-carousel/lib/styles.css';
const useStyles = createUseStyles((theme: Theme) => ({
  rootCarouselImage: {
    padding: theme.spacing(2),
    background: theme.carouselImageBackgroundGradient
  },
  carousel: {},
  containerClass: {
    paddingBottom: 29
  },
  imageSlide: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center'
  },
  imageWrapper: {
    border: `1px solid ${theme.colors.approxZircon}`,
    borderRadius: 6,
    backgroundColor: theme.colors.primaryWhite,
    boxShadow: `0 2px 8px 0 rgba(0,0,0,0.17)`,
    padding: 6.5
  },
  dot: {
    position: 'relative',
    width: 14,
    height: 14,
    border: `1px solid ${theme.colors.primaryWhite}`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 0,
    borderRadius: 50,
    margin: [0, 7.5],
    cursor: 'pointer',
    background: 'transparent',
    '&:focus': {
      outline: 'none'
    }
  },
  dotActive: {
    '&:after': {
      content: "''",
      display: 'block',
      width: 8,
      height: 8,
      borderRadius: 50,
      backgroundColor: theme.colors.primaryWhite
    }
  },
  image: {
    maxWidth: '100%'
  }
}));
export interface Image {
  url: string;
  alt?: string;
}
type Props = {
  images: Image[];
};
const responsive = {
  superLargeDesktop: {
    breakpoint: { max: 4000, min: 3000 },
    items: 1
  },
  desktop: {
    breakpoint: { max: 3000, min: 1024 },
    items: 1
  },
  tablet: {
    breakpoint: { max: 1024, min: 464 },
    items: 1
  },
  mobile: {
    breakpoint: { max: 464, min: 0 },
    items: 1
  }
};
const CustomDot = ({ onClick, active, classes }: any) => {
  return (
    <li>
      <button
        className={classNames({ [classes.dotActive]: active }, classes.dot)}
        onClick={() => onClick()}
      />
    </li>
  );
};
function CarouselImage({ images }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.rootCarouselImage}>
      <div className={classes.carousel}>
        <Carousel
          additionalTransfrom={0}
          arrows={false}
          autoPlaySpeed={3000}
          centerMode={false}
          className=""
          containerClass={classes.containerClass}
          draggable
          focusOnSelect={false}
          infinite={true}
          itemClass=""
          keyBoardControl
          minimumTouchDrag={80}
          renderButtonGroupOutside={false}
          renderDotsOutside={false}
          responsive={responsive}
          showDots={true}
          sliderClass=""
          slidesToSlide={1}
          customDot={<CustomDot classes={classes} />}
          swipeable
        >
          {images.map((image, index) => (
            <div key={index} className={classes.imageSlide}>
              <div className={classes.imageWrapper}>
                <img
                  src={image.url}
                  alt={image.alt ?? ''}
                  className={classes.image}
                />
              </div>
            </div>
          ))}
        </Carousel>
      </div>
    </div>
  );
}

export default CarouselImage;
