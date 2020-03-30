const spacing = {
  0: 5,
  1: 10,
  2: 20,
  3: 50,
};

export type Theme = {
  spacing: (amount: number) => number;
  primaryBorderRadius: number;
  colors: {
    primaryBlack: string;
    primaryBlue: string;
    primaryGreen: string;
    secondaryGrey: string;
  };
  shadows: {
    primary: string;
  };
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  colors: {
    primaryBlack: "#0c152f",
    primaryBlue: "#0b57ff",
    primaryGreen: "#10b73b",
    secondaryGrey: "#9ea2ad",
  },
  shadows: {
    primary: "0px 2px 10px #0000001f",
  },
};

export default theme;
