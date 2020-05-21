import * as React from "react";
import classNames from "classnames";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import Button from "sharedComponents/core/Button";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    row: {
        display: "flex",
        flexDirection: "row",
        flexWrap: 'wrap'
      },
      pageNumber: {
        width: 40, marginLeft: 10,
      },
      currentPage: {
        borderColor: theme.colors.primaryBlue,
        boxShadow: '0 2px 4px 0 rgba(0,0,0,0.12)'
      },
}));

type Props = {
    currentPage: number;
    setCurrentPage: Function;
    range: number;
    numberOfPages: number;
};

function PageNumbers({ currentPage, setCurrentPage, range = 4, numberOfPages }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const [pageRange, setPageRange] = React.useState<number>(range);
    const changeRange = (change: number) => {
        if (pageRange + change < range) {
            setPageRange(range);
        } else if (pageRange + change > numberOfPages) {
            setPageRange(numberOfPages);
        } else {
            setPageRange(pageRange + change);
        }
    }

    const pages: number[] = [];
    for (let pageNum = pageRange - range; pageNum < pageRange; pageNum++) {
        pages.push(pageNum + 1)
    }

    return (
      <div className={classes.row}>
        <Button
          style={{ width: 40, marginLeft: 25 }}
          onClick={() => changeRange(-range)}
          small={true}
        >
          <Icon name="ArrowLeft" size={15} />
        </Button>
        {pages.map((pageNum: number) => (
          <Button
            className={classNames(pageNum === currentPage
              ? [classes.pageNumber, classes.currentPage]
              : classes.pageNumber)}
            onClick={() => setCurrentPage(pageNum)}
            small={true}
          >
            {pageNum}
          </Button>
        ))}
        <Button
          style={{ width: 40, marginLeft: 10 }}
          onClick={() => changeRange(range)}
          small={true}
        >
          <Icon name="ArrowRight" size={15} />
        </Button>
      </div>
  );
}

export default PageNumbers;