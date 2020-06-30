import * as React from "react";
import ArticleHeader from "./ArticleHeader";
import { withKnobs, text, boolean, object } from "@storybook/addon-knobs";

export default {
  title: "Overview/Article/ArticleHeader",
  decorators: [withKnobs],
};

export const normal = () => {
    const showAuthor: boolean = boolean("Show Author", false);
    const title: string = text("Title", "ICAO calls for improved coordination and awareness of COVID-19");
    const date: string = text("Date", "18th April 2020");
    const genre: string = text("Genre", "Health & Safety");
    const featured: string = text("Featured", "Featured Article");
    const author = object("Author", { name: "James Green", url: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")})
  return <ArticleHeader
        title={title}
        date={date}
        genre={genre}
        featured={showAuthor ? undefined : featured }
        author={showAuthor ? author : undefined}
        image={require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")}
  />;
};