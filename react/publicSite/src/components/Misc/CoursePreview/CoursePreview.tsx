import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Button from "sharedComponents/core/Input/Button";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        height: '674.5px',
        width: '390.5px',
        display: 'flex',
        flexDirection: 'column',
        border: '1px solid #E9EBEB',
        borderRadius: '5px',
        backgroundColor: theme.colors.primaryWhite,
        boxShadow: '0 2px 10px 0 rgba(0,0,0,0.15), 4px 2px 10px -2px rgba(0,0,0,0.06)',
    },
    price: {
        fontSize: theme.fontSizes.tinyHeading,
        fontWeight: '800',
        margin: '18px 0 15px 33px'
    },
    vat: {
        fontSize: theme.fontSizes.large,
        fontWeight: '500',
    },
    button: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        height: '52px',
        width:' 327px',
        marginTop: '10px',
        alignSelf: 'center',
        boxShadow: '0 1px 4px 0 rgba(0,0,0,0.09)'
    },
    list: {
        fontSize: theme.fontSizes.default,
        fontWeight: '800',
        margin: '22px 0 0 33px'
    },
    bullet: {
        height: '8px',
        width: '8px',
        borderRadius: '8px',
        margin: '7px 8px 7px 0',
        backgroundColor: theme.colors.navyBlue
    },
    detail: {
        fontSize: theme.fontSizes.small,
        fontWeight: '500',
        color: theme.colors.textGrey
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center'
    }
}));

type Props = {
    price: string;
    details: string[];
    onBasket?: () => void;
    onBuy?: () => void;
};

function CoursePreview({ price, details, onBasket, onBuy }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const videoStyle = { height: '238px', width: '382px', margin: '4px',  borderRadius: '5px 5px 0 0' };

  return (
      <div className={classes.root}>
          <div>
                <Icon name="SampleImage_ClassroomCoursesDetail_Feat" style={videoStyle} />
          </div>

          <div className={classes.price}>{price} <span className={classes.vat}>+VAT</span></div>

          <Button
            archetype="submit"
            className={classes.button}
            onClick={() => onBasket && onBasket()}
          >
            Add to Basket
          </Button>

          <Button
            className={classes.button}
            onClick={() => onBuy && onBuy()}
          >
            Buy Now
          </Button>

          <div className={classes.list}>
              This course includes:
              {details.map((detail: string) => (
                <div className={classes.row}>
                    <div className={classes.bullet}/>
                    <div className={classes.detail}>{detail}</div>
                </div>
              ))}
          </div>
      </div>
  );
}

export default CoursePreview;