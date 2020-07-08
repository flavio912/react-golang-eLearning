import * as React from "react";
import GenreHeader from "./GenreHeader";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Overview/Article/GenreHeader",
  decorators: [withKnobs],
};

export const normal = React.createElement(() => {
    const genres: string[] = ['Latest Articles', 'Popular', 'Dangerous Goods', 'Aviation', 'Health & Safety', 'Product']
    const [selected, setSelected] = React.useState(genres[0]);
  return <GenreHeader
    genres={genres}
    selected={selected}
    setSelected={setSelected}
    />;
});