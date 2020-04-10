import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  table: {
    width: "100%",
    borderCollapse: "collapse",
  },
  blankRow: {
    height: 10,
  },
  row: {
    background: "white",
    border: `1px solid ${theme.colors.borderGrey}`,
    padding: [10, 0],
    transition: "box-shadow 0.1s ease",
    "&:hover": {
      boxShadow: theme.shadows.primary,
    },
  },
  headerCell: {
    color: theme.colors.textGrey,
    textTransform: "uppercase",
    fontSize: theme.fontSizes.extraSmall,
    padding: [0, 10],
  },
  cell: {
    padding: [0, 10],
    "&:first-child": {
      borderRadius: [5, 0, 0, 5],
    },
    "&:last-child": {
      borderRadius: [0, 5, 5, 0],
    },
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
            <>
              <tr className={classes.blankRow} />
              <tr key={key} className={classes.row}>
                {cells.map(({ component: Cell, sort }) => (
                  <td className={classes.cell}>
                    {typeof Cell === "string" ? (
                      <p key={`${key}-${sort}`}>{Cell}</p>
                    ) : (
                      <Cell key={`${key}-${sort}`} />
                    )}
                  </td>
                ))}
              </tr>
            </>
          );
        })}
      </tbody>
    </table>
  );
}

export default Table;
