const spacing = {
  0: 5,
  1: 10,
  2: 20,
  3: 50
};

export type Theme = {
  spacing: (amount: number) => number;
  primaryBorderRadius: number;
  secondaryBorderRadius: number;
  buttonBorderRadius: number;
  primaryGradient: string;
  verticalGradient: string;
  loginBackgroundGradient: string;
  searchBackground: string;
  colors: {
    primaryBlack: string;
    primaryBlue: string;
    primaryGreen: string;
    secondaryGreen: string;
    secondaryBlack: string;
    hoverGreen: string;
    primaryRed: string;
    secondaryGrey: string;
    primaryWhite: string;
    borderGrey: string;
    backgroundGrey: string;
    textBlue: string;
    textGrey: string;
    searchHoverGrey: string;
    progressGrey: string;
    navyBlue: string;
    blueRibbon: string;
    approxIron: string;
    approxZircon: string;
    secondaryDanger: string;
    certBackgroundGrey: string;
  };
  shadows: {
    primary: string;
    body: string;
  };
  fontSizes: {
    heading: number;
    smallHeading: number;
    xSmallHeading: number;
    tinyHeading: number;
    extraLarge: number;
    large: number;
    xLarge: number;
    default: number;
    small: number;
    xSmall: number;
    tiny: number;
  };
  paperSizes: {
    A4: {
      width: number,
      height: number
    }
  }
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  secondaryBorderRadius: 10,
  buttonBorderRadius: 4,
  primaryGradient: 'linear-gradient(50deg, #0b57ff 0%,#16C225 100%)',
  verticalGradient: 'linear-gradient(to bottom left, #16C225 0%,#0b57ff 100%)',
  loginBackgroundGradient: 'linear-gradient(50deg, #0f6fcc 0%,#16a858 100%)',
  searchBackground: 'rgba(7,67,121,0.75)',
  colors: {
    primaryBlack: '#0C152E',
    primaryBlue: '#0b57ff',
    primaryGreen: '#10b73b',
    secondaryBlack: '#34373A',
    secondaryGreen: '#15C324',
    hoverGreen: 'rgb(243, 251, 242)',
    primaryRed: '#CB463A',
    secondaryGrey: '#9ea2ad',
    primaryWhite: '#FFFFFF',
    borderGrey: '#ededed',
    backgroundGrey: '#f7f9fb',
    searchHoverGrey: '#F5FAFC',
    textBlue: '#1081AA',
    textGrey: '#737988',
    progressGrey: '#d2d6db',
    navyBlue: '#0E66E0',
    blueRibbon: '#0E5AF9',
    approxIron: '#CCCDCD',
    approxZircon: '#E9EBEB',
    secondaryDanger: '#DB5C5D',
    certBackgroundGrey: '#F5F5FE'
  },
  shadows: {
    primary: '2px 2px 10px rgba(0,0,0,0.07)',
    body: 'inset 0px -2px 10px 0px #0000001f'
  },
  fontSizes: {
    heading: 25,
    smallHeading: 24,
    xSmallHeading: 21,
    tinyHeading: 20,
    extraLarge: 18,
    large: 16,
    xLarge: 15,
    default: 14,
    small: 13,
    xSmall: 12,
    tiny: 11
  },
  paperSizes: {
    A4: {
      width: 1240,
      height: 1754
    }
  }
};

export default theme;
