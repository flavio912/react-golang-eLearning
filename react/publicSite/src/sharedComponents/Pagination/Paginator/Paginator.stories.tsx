import * as React from "react";
import Paginator from ".";
import { withKnobs } from "@storybook/addon-knobs";

export default {
  title: "/~/Desktop/GitHub/admin-react/react/delegate/src/sharedComponents/Paginator",
  decorators: [withKnobs],
};

const Default = () => {
    const [currentPage, setCurrentPage] = React.useState<number>(1);
  
    return (
        <Paginator
            currentPage={currentPage}
            numPages={8}
            itemsPerPage={10}
            updatePage={(page: number) => setCurrentPage(page)}
        />
    )
}
  
export const normal = () => {
    return <Default />;
};