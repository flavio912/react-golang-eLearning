import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        padding: '40px',
        borderRadius: '14px',
        boxShadow: '0 1px 7px 3px rgba(0,0,0,0.11)'
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
        flex: 1
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center'
    },
    header: {
        fontSize: theme.fontSizes.heading,
        fontWeight: '900',
        marginBottom: '50px'
    },
    tableHeader: {
        textAlign: 'center',
        fontSize: theme.fontSizes.default,
        color: theme.colors.textGrey
    },
    name: {
        fontSize: theme.fontSizes.large,
        fontWeight: '700'
    },
    id: {
        fontSize: theme.fontSizes.large,
        color: theme.colors.textGrey,
        marginTop: '5px'
    },
    image: {
        height: '50px',
        width: '54px',
        borderRadius: '5px',
        marginRight: '15px'
    },
    detail: {
        textAlign: 'center',
        fontSize: theme.fontSizes.large,
        color: theme.colors.textGrey
    },
    padding: {
        padding: '20px 0 20px 0'
    },
    border: {
        paddingBottom: '20px',
        borderBottom: ['2px', 'solid', theme.colors.borderGrey]
    },
    left: {
        textAlign: 'left',
    }
}));

export type OrderItem = {
    id: string | number;
    name: string;
    imageURL: string;
    quantity: number;
    price: number;
}

type Props = {
    orderItems: OrderItem[];
    className?: string;
};

function OrderCard({ orderItems, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}> 
        <div className={classes.header}>Order Details</div>

        <div className={classNames(classes.row, classes.border)}>
            <div className={classNames(classes.tableHeader, classes.left)} style={{ flex: 3 }}>PRODUCT</div>
            <div className={classes.tableHeader} style={{ flex: 1 }}>QUANTITY</div>
            <div className={classes.tableHeader} style={{ flex: 1 }}>PRICE</div>
            <div className={classes.tableHeader} style={{ flex: 1 }}>TOTAL</div>
        </div>

        {orderItems && orderItems.map((item: OrderItem) => (
            <div className={classNames(classes.row, classes.padding)} key={item.id}>
                <div className={classes.row} style={{ flex: 3 }}>
                    <img className={classes.image} src={item.imageURL} />
                    <div className={classes.column}>
                        <div className={classes.name}>{item.name}</div>
                        <div className={classes.id}>SKU {item.id}</div>
                    </div>
                </div>
                <div className={classes.detail} style={{ flex: 1 }}>{item.quantity}</div>
                <div className={classes.detail} style={{ flex: 1 }}>£{item.price}</div>
                <div className={classes.detail} style={{ flex: 1 }}>£{(item.quantity * item.price).toFixed(2)}</div>
        </div>
        ))}
    </div> 
  );
}

export default OrderCard;