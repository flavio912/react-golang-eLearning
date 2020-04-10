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
  primaryGradient: string;
  colors: {
    primaryBlack: string;
    primaryBlue: string;
    primaryGreen: string;
    secondaryGrey: string;
    primaryWhite: string;
    borderGrey: string;
    backgroundGrey: string;
    textBlue: string;
    textGrey: string;
    searchHoverGrey: string;
  };
  shadows: {
    primary: string;
  };
  fontSizes: {
    heading: number;
    large: number;
    default: number;
    small: number;
    extraSmall: number;
  };
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  buttonBorderRadius: 4,
  primaryGradient: "linear-gradient(to right, #0b57ff 0%,#10b73b 100%)",
  colors: {
    primaryBlack: "#0C152E",
    primaryBlue: "#0b57ff",
    primaryGreen: "#10b73b",
    secondaryGrey: "#9ea2ad",
    primaryWhite: "#FFFFFF",
    borderGrey: "#ededed",
    backgroundGrey: "#f7f9fb",
    searchHoverGrey: "#F5FAFC",
    textBlue: "#1B759E",
    textGrey: "#616575",
  },
  shadows: {
    primary: "0px 2px 10px #0000001f",
  },
  fontSizes: {
    heading: 26,
    large: 16,
    default: 14,
    small: 13,
    extraSmall: 11,
  },
};

export default theme;
