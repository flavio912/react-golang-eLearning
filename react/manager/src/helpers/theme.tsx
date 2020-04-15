const spacing = {
  0: 5,
  1: 10,
  2: 20,
  3: 50,
};

export type Theme = {
  spacing: (amount: number) => number;
  primaryBorderRadius: number;
  secondaryBorderRadius: number;
  buttonBorderRadius: number;
  primaryGradient: string;
  colors: {
    primaryBlack: string;
    primaryBlue: string;
    primaryGreen: string;
    secondaryGreen: string;
    hoverGreen: string;
    primaryRed: string;
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
    body: string;
  };
  fontSizes: {
    heading: number;
    smallHeading: number;
    extraLarge: number;
    large: number;
    default: number;
    small: number;
    tiny: number;
  };
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  secondaryBorderRadius: 10,
  buttonBorderRadius: 4,
  primaryGradient: "linear-gradient(50deg, #0b57ff 0%,#16C225 100%)",
  colors: {
    primaryBlack: "#0C152E",
    primaryBlue: "#0b57ff",
    primaryGreen: "#10b73b",
    secondaryGreen: "#15C324",
    hoverGreen: "#E7F8E6",
    primaryRed: "#CB463A",
    secondaryGrey: "#9ea2ad",
    primaryWhite: "#FFFFFF",
    borderGrey: "#ededed",
    backgroundGrey: "#f7f9fb",
    searchHoverGrey: "#F5FAFC",
    textBlue: "#1081AA",
    textGrey: "#737988",
  },
  shadows: {
    primary: "0px 2px 10px #0000001f",
    body: "inset 0px -2px 10px 0px #0000001f",
  },
  fontSizes: {
    heading: 25,
    smallHeading: 24,
    extraLarge: 18,
    large: 16,
    default: 14,
    small: 13,
    tiny: 11,
  },
};

export default theme;
