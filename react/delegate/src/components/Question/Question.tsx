import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {},
  questionItem: {
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: 5,
    backgroundColor: theme.colors.primaryWhite,
  },
  questionImage: {
    height: 189,
    width: 332,
    marginRight: 18,
    marginBottom: 19,
    backgroundSize: "contain",
    backgroundRepeat: "no-repeat",
    backgroundPosition: "center",
    position: "relative",
    cursor: "pointer",
    "& span": {
      right: 15,
      bottom: 14,
    },
  },
  questionText: {
    height: 86,
    width: 682,
    paddingLeft: 20.5,
    paddingRight: 29,
    paddingTop: 24.5,
    paddingBottom: 24.5,
    display: "flex",
    alignItems: "center",
    margin: [24, 0],
    cursor: "pointer",
    position: "relative",
    "& span": {
      right: 29,
    },
  },
  questionTitle: {
    color: theme.colors.secondaryBlack,
    fontSize: 22,
    letterSpacing: -0.55,
    marginBottom: 32.5,
  },
  questionOptions: {
    display: "flex",
  },
  optionCol: {
    flexDirection: "column",
  },
  optionRow: {
    flexDirection: "row",
  },
  optionIndex: {
    color: "#0E5AF9",
    fontSize: 30,
    fontWeight: 900,
    letterSpacing: -0.75,
    marginRight: 21,
  },
  optionTitle: {
    color: theme.colors.secondaryBlack,
    fontSize: 22,
    letterSpacing: -0.55,
    margin: 0,
  },
  optionSelected: {
    border: `1px solid #0E5AF9`,
    backgroundColor: theme.colors.searchHoverGrey,
    boxShadow: `0 2px 4px 0 rgba(0,0,0,0.09)`,
  },
  selectedIcon: {
    borderRadius: 50,
    boxSizing: "border-box",
    backgroundColor: theme.colors.primaryWhite,
    width: 24,
    height: 24,
    border: `1px solid #CCCDCD`,
    position: "absolute",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    "&:after": {
      content: "''",
      display: "block",
      width: 14,
      height: 14,
      backgroundColor: "#0E5AF9",
      borderRadius: 50,
    },
  },
}));
export type Option = {
  id: string | number;
  index?: string | number;
  title?: string;
  image?: string;
};
export type QuestionType = "image" | "text";
export interface Question {
  id: string | number;
  title: string;
  options: Option[];
}
type Props = {
  className?: string;
  question: Question;
  type: QuestionType;
  onSelected: Function;
};
const OptionImageEle = ({
  classes,
  isSelected,
  option,
  onClick,
  index,
}: any) => {
  const backgroundImage = `url(${option.image})`;
  return (
    <div
      className={classNames(classes.questionItem, classes.questionImage, {
        [classes.optionSelected]: isSelected,
      })}
      onClick={onClick}
      key={index}
      style={{
        backgroundImage,
      }}
    >
      {isSelected && <span className={classes.selectedIcon} />}
    </div>
  );
};
const OptionTextEle = ({
  classes,
  isSelected,
  option,
  onClick,
  index,
}: any) => (
  <div
    key={index}
    className={classNames(classes.questionItem, classes.questionText, {
      [classes.questionSelected]: isSelected,
    })}
    onClick={onClick}
  >
    {option.index && (
      <span className={classes.optionIndex}>{option.index}</span>
    )}
    <p className={classes.optionTitle}>{option.title}</p>
    {isSelected && <span className={classes.selectedIcon} />}
  </div>
);
function Question({ className, question, type, onSelected }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [select, setSelect] = React.useState<Option>();
  return (
    <div className={classNames(classes.root, className)}>
      <h2 className={classes.questionTitle}>{question.title}</h2>
      <div
        className={classNames(classes.questionOptions, {
          [classes.optionCol]: type === "text",
          [classes.optionRow]: type === "image",
        })}
      >
        {question.options.map((option: Option, index: string | number) => {
          const isSelected = select && select.id === option.id;
          if (type === "image") {
            return (
              <OptionImageEle
                classes={classes}
                onClick={() => {
                  setSelect(option);
                  onSelected(option);
                }}
                option={option}
                key={index}
                isSelected={isSelected}
              />
            );
          } else {
            return (
              <OptionTextEle
                classes={classes}
                onClick={() => {
                  setSelect(option);
                  onSelected(option);
                }}
                option={option}
                key={index}
                isSelected={isSelected}
              />
            );
          }
        })}
      </div>
    </div>
  );
}

export default Question;
