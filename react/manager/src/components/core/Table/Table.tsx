import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  table: {
    width: "100%",
    borderCollapse: "separate",
    borderSpacing: '0 10px',
  },
  row: {
    background: "white",
    height: 70,
    transition: "box-shadow 0.1s ease",
    "&:hover": {
      boxShadow: theme.shadows.primary,
      "& $cell": {"&:first-child": {
        borderRadius: 0,
      },
      "&:last-child": {
        borderRadius: 0,
      },}
    },
  },
  headerCell: {
    color: theme.colors.textGrey,
    textTransform: "uppercase",
    fontSize: theme.fontSizes.extraSmall,
    padding: [0, 10],
    textAlign: "left",
  },
  cell: {
    padding: [0, 10],
    borderTop: `1px solid ${theme.colors.borderGrey}`,
    borderRight: '0',
    borderBottom: `1px solid ${theme.colors.borderGrey}`,
    borderLeft: '0',
    transition: "border-radius 0.1s ease",
    "&:first-child": {
      borderRadius: [5, 0, 0, 5],
      borderLeft: `1px solid ${theme.colors.borderGrey}`,
    },
    "&:last-child": {
      borderRadius: [0, 5, 5, 0],
      borderRight: `1px solid ${theme.colors.borderGrey}`,
    },
  },
  defaultCell: {
    fontSize: theme.fontSizes.default,
  },
}));

type Props = {
  header: string[];
  rows: TableRow[];
  sort: Sort;
};

type TableRow = {
  key: string;
  cells: TableCell[];
};

type TableCell = {
  component: (() => JSX.Element) | string;
  sort?: string | number;
};

type Sort = {
  by: string;
  dir: "ASC" | "DESC";
};

function Table({ header, rows, sort }: Props) {
  const classes = useStyles();

  const sorter = (a: TableRow, b: TableRow): number => {
    const column: number = header.findIndex((col) => col === sort.by);
    if (column > 0) {
      if (sort.dir === "DESC") {
        const t = b;
        b = a;
        a = t;
      }

      const aCell = a[column];
      const bCell = b[column];

      if (!(aCell && bCell && aCell.sort && bCell.sort)) {
        console.warn(`Table column ${sort.by} is not able to be sorted`);
      } else {
        if (aCell.sort < bCell.sort) return 1;
        if (aCell.sort > bCell.sort) return -1;
      }
    } else {
      console.warn(`Table column ${sort.by} is not present in the table`);
    }
    return 0;
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
        {rows.sort(sorter).map(({ key, cells }) => {
          return (
            <tr key={key} className={classes.row}>
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
