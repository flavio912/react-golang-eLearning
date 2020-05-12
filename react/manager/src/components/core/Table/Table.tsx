import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  table: {
    width: "100%",
    borderCollapse: "separate",
    borderSpacing: "0 10px",
  },
  row: {
    background: "white",
    height: 70,
    borderRadius: theme.primaryBorderRadius,
    transition: "box-shadow 0.1s ease",
    "&:hover": {
      boxShadow: theme.shadows.primary,
      // "& $cell": {
      //   "&:first-child": {
      //     borderRadius: 0,
      //   },
      //   "&:last-child": {
      //     borderRadius: 0,
      //   },
      // },
    },
  },
  headerCell: {
    color: theme.colors.secondaryBlack,
    textTransform: "uppercase",
    fontSize: theme.fontSizes.tiny,
    fontWeight: 400,
    padding: [0, 10],
    textAlign: "left",
  },
  cell: {
    padding: [0, 10],
    borderTop: `1px solid ${theme.colors.borderGrey}`,
    borderRight: "0",
    borderBottom: `1px solid ${theme.colors.borderGrey}`,
    borderLeft: "0",
    transition: "border-radius 0.1s ease",
    "&:first-child": {
      borderRadius: [
        theme.primaryBorderRadius,
        0,
        0,
        theme.primaryBorderRadius,
      ],
      borderLeft: `1px solid ${theme.colors.borderGrey}`,
    },
    "&:last-child": {
      borderRadius: [
        0,
        theme.primaryBorderRadius,
        theme.primaryBorderRadius,
        0,
      ],
      borderRight: `1px solid ${theme.colors.borderGrey}`,
    },
  },
  defaultCell: {
    fontSize: theme.fontSizes.default,
  },
}));

type Props = {
  header: Array<string | JSX.Element>;
  rows: TableRow[];
  filter?: Filter;
  sort?: Sort;
};

type TableRow = {
  key: string;
  cells: TableCell[];
  onClick?: () => void;
};

type TableCell = {
  component: (() => JSX.Element) | string;
  sort?: string | boolean | number;
};

type Sort = {
  column: number;
  dir: "ASC" | "DESC";
};

type Filter = {
  column: number;
  filterFunc: (value: string | number | boolean) => boolean;
};

function Table({ header, rows, sort, filter }: Props) {
  const classes = useStyles();

  const sorter = (a: TableRow, b: TableRow): number => {
    if (sort === undefined) return 0;
    if (sort.column < header.length) {
      if (sort.dir === "DESC") {
        const t = b;
        b = a;
        a = t;
      }

      const aCell = a.cells[sort.column];
      const bCell = b.cells[sort.column];

      if (
        !(
          "sort" in aCell &&
          aCell.sort !== undefined &&
          "sort" in bCell &&
          bCell.sort !== undefined
        )
      ) {
        console.warn(`Table column ${sort.column} is not able to be sorted`);
      } else {
        if (aCell.sort < bCell.sort) return 1;
        if (aCell.sort > bCell.sort) return -1;
      }
    } else {
      console.warn(`Table column ${sort.column} is not present in the table`);
    }
    return 0;
  };

  const filterer = (row: TableRow) => {
    if (filter === undefined) return true;
    if (filter.column < row.cells.length) {
      const value = row.cells[filter.column];

      if ("sort" in value && value.sort !== undefined) {
        return filter.filterFunc(value.sort);
      } else {
        console.warn(
          `Table column ${filter.column} is not able to be filtered`
        );
        return true;
      }
    } else {
      console.warn(`Table column ${filter.column} is not present in the table`);
      return true;
    }
  };

  return (
    <table className={classes.table}>
      <thead>
        <tr>
          {header.map((title) => (
            <th className={classes.headerCell}>{title}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {rows
          .filter(filterer)
          .sort(sorter)
          .map(({ key, cells, onClick }) => {
            return (
              <tr key={key} className={classes.row} onClick={onClick}>
                {cells.map(({ component: Cell, sort }) => (
                  <td className={classes.cell}>
                    {typeof Cell === "string" ? (
                      <p className={classes.defaultCell} key={`${key}-${sort}`}>
                        {Cell}
                      </p>
                    ) : (
                      <Cell key={`${key}-${sort}`} />
                    )}
                  </td>
                ))}
              </tr>
            );
          })}
      </tbody>
    </table>
  );
}

export default Table;
