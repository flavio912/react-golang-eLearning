import * as React from "react";
import PageNumbers from "./PageNumbers";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "Pagination/PageNumbers",
  decorators: [withKnobs],
};

const Default = () => {
  const [currentPage, setCurrentPage] = React.useState<number>(1);

    return (
        <PageNumbers
            numberOfPages={8}
            range={4}
            currentPage={currentPage}
            setCurrentPage={(page: number) => setCurrentPage(page)}
        />
    )
}

export const normal = () => {
  return <Default />;
};