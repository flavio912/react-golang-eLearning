const spacing = {
  0: 5,
  1: 10,
  2: 20,
  3: 50,
};

export type Theme = {
  spacing: (amount: number) => number;
  primaryBorderRadius: number;
  buttonBorderRadius: number;
  colors: {
    primaryBlack: string;
    primaryBlue: string;
    primaryGreen: string;
    secondaryGrey: string;
    borderGrey: string;
    backgroundGrey: string;
    textBlue: string;
    textGrey: string;
  };
  shadows: {
    primary: string;
  };
  fontSizes: {
    default: number;
  };
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  buttonBorderRadius: 4,
  colors: {
    primaryBlack: "#0c152f",
    primaryBlue: "#0b57ff",
    primaryGreen: "#10b73b",
    secondaryGrey: "#9ea2ad",
    borderGrey: "#ededed",
    backgroundGrey: "#f7f9fb",
    textBlue: "#1B759E",
    textGrey: "#616575",
  },
  shadows: {
    primary: "0px 2px 10px #0000001f",
  },
  fontSizes: {
    default: 13,
  },
};

export default theme;
