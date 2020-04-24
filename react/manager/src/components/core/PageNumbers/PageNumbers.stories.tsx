import * as React from "react";
import PageNumbers from "./PageNumbers";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "core/PageNumbers",
  decorators: [withKnobs],
};

const Default = () => {
  const pages: number[] = [1,2,3,4,5,6,7,8];
  const [currentPage, setCurrentPage] = React.useState<number>(1);
  const [pageRange, setPageRange] = React.useState<number[]>([1,2,3,4]);

    return (
        <PageNumbers
            pages={pages}
            currentPage={currentPage}
            pageRange={pageRange}
            setCurrentPage={(page: number) => setCurrentPage(page)}
            setPageRange={(range: number[]) => setPageRange(range)}
        />
    )
}

export const normal = () => {
  return <Default />;
};