import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Button from "sharedComponents/core/Input/Button";
import Icon from "sharedComponents/core/Icon";
import VideoPlayer from "components/core/VideoPlayer";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        width: '390.5px',
        display: 'flex',
        flexDirection: 'column',
        border: '1px solid #E9EBEB',
        borderRadius: '5px',
        backgroundColor: theme.colors.primaryWhite,
        boxShadow: '0 2px 10px 0 rgba(0,0,0,0.15), 4px 2px 10px -2px rgba(0,0,0,0.06)',
    },
    video: {
        margin: '4px',
        borderRadius: '5px 5px 0 0',
        overflow: 'hidden'
    },
    price: {
        fontSize: theme.fontSizes.smallHeading,
        fontWeight: '600',
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
        fontSize: theme.fontSizes.small,
        fontWeight: '600',
        margin: '22px 0 0 33px',
        marginBottom: '30px'
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
    column: {
        alignSelf: 'flex-end',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
    },
    preview: {
        fontSize: theme.fontSizes.xSmall,
        fontWeight: 'bold',
        color: theme.colors.primaryWhite,
        margin: '45px 0 10px 0'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center'
    },
    // Thumbnail styles
    playCircle: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        width: '76px',
        height: '76px',
        borderRadius: '76px',
        backgroundColor: theme.colors.navyBlue,
        opacity: 0.8
    },
    playTriangle: {
        width: 0, 
        height: 0,
        marginLeft: 6,
        borderTop: "13.5px solid transparent",
        borderBottom: "13.5px solid transparent",
        borderLeft: ["27px", "solid", theme.colors.primaryWhite]
    }
}));

type Props = {
    price: string;
    details: string[];
    video: string;
    onBasket?: () => void;
    onBuy?: () => void;
};

function CoursePreview({ price, details, video, onBasket, onBuy }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
      <div className={classes.root}>

          <VideoPlayer
            source={video}
            width={384} height={216}
            className={classes.video}
            thumbnail={
                <div className={classes.column}>
                    <div />
                    <div className={classes.playCircle}>
                            <div className={classes.playTriangle}/>
                    </div>
                    <div className={classes.preview}>Preview this course</div>
                </div>
            }
        />

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