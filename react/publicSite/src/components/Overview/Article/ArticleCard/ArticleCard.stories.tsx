import * as React from "react";
import ArticleCard from "./ArticleCard";
import { withKnobs, text } from "@storybook/addon-knobs";

export default {
  title: "Overview/Article/ArticleCard",
  decorators: [withKnobs],
};

export const normal = () => {
    const type: string = text("Type", "Dangerous Goods");
    const date: string = text("Date", "20th March 2020");
    const description: string = text("Description", "CANSO makes call to action to ensure the stability of the ATM industry");
    const article = { type, date, description, imageURL: require("assets/SampleImage_ClassroomCoursesDetail_Feat.png")}
    return <ArticleCard
        article={article}
        onClick={() => console.log('Clicked')}
  />;
};