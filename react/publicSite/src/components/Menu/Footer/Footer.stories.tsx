import * as React from "react";
import Footer, { Column, Link } from "./Footer";
import { withKnobs, object } from "@storybook/addon-knobs";

export default {
  title: "Menu/Footer",
  decorators: [withKnobs],
};

const defaultLink: Link = {
    name: "Feature",
    link: ""
}

const defaultNew: Link = {
    name: "Feature",
    link: "",
    alert: {
        type: "new",
        value: "NEW"
    }
}

const defaultIncrease: Link = {
    name: "Feature",
    link: "",
    alert: {
        type: "increase",
        value: "99.9%"
    }
}

const defaultColumns: Column[] = [
    {
        header: "Features", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultNew, defaultLink, defaultLink, defaultLink, defaultLink
        ]
    },
    {
        header: "Learn More", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultLink, defaultLink
        ]
    },
    {
        header: "Company", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultLink, defaultLink
        ]
    },
    {
        header: "Get Help", links: [
            defaultLink, defaultLink, defaultLink, defaultLink, defaultIncrease
        ]
    }
]

export const normal = () => {
    const columns: Column[] = object("Columns", defaultColumns);
  return <Footer columns={columns} />;
};