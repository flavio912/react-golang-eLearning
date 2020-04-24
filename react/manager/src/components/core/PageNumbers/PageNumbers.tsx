import * as React from "react";
import classNames from "classnames";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import Button from "components/core/Button";
import Icon from "components/core/Icon";

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
    pages: number[];
    currentPage: number;
    pageRange: number[];
    setPageRange: Function;
    setCurrentPage: Function;
};

function PageNumbers({ pages, currentPage, pageRange, setPageRange, setCurrentPage }: Props) {
    const classes = useStyles();

    return (
      <div className={classes.row}>
          <Button
          style={{ width: 40, marginLeft: 25 }}
          onClick={() => setPageRange(pages.slice(0,4))}
          small={true}
        >
          <Icon name="ArrowLeft" size={15} />
        </Button>
        {pageRange.map((pageNum: number) => (
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
          onClick={() => setPageRange(pages.slice(4,8))}
          small={true}
        >
          <Icon name="ArrowRight" size={15} />
        </Button>
      </div>
  );
}

export default PageNumbers;