import * as React from "react";
import { withKnobs } from "@storybook/addon-knobs";
import RegistrationCarousel from "./RegistrationCarousel";
import { Image } from "components/Misc/CarouselImage";

export default {
  title: "Overview/Registration/RegistrationCarousel",
  decorators: [withKnobs],
};

const background = {
    background: 'linear-gradient(222.02deg, #16BB33 0%, #0E69DA 100%)',
}

const defaultImage: Image = {
    url: require('assets/carouselImage.svg'),
    alt: 'Image'
  };

export const normal = () => {
  const images = [1, 2, 3].map((item) => ({
    ...defaultImage,
    alt: `${defaultImage.alt} ${item}`
    }));

  return <div style={background}>
        <RegistrationCarousel  
            images={images}
            onBook={() => console.log('Book')}
        />;
      </div>
};