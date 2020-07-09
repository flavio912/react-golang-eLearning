const spacing = {
  0: 5,
  1: 10,
  2: 20,
  3: 50,
  4: 70,
  5: 100,
};

export type Theme = {
  spacing: (amount: number) => number;
  primaryBorderRadius: number;
  secondaryBorderRadius: number;
  buttonBorderRadius: number;
  primaryGradient: string;
  loginBackgroundGradient: string;
  carouselImageBackgroundGradient: string;
  centerColumnWidth: number;
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
    lightBlue: string;
    navyBlue: string;
    footerBlue: string;
    footerGrey: string;
    blueRibbon: string;
    approxIron: string;
    approxZircon: string;
    secondaryDanger: string;
    navyBlue2: string;
    silver: string;
  };
  shadows: {
    primary: string;
    body: string;
  };
  fontSizes: {
    extraLargeHeading: number;
    heading: number;
    smallHeading: number;
    xSmallHeading: number;
    tinyHeading: number;
    extraLarge: number;
    large: number;
    default: number;
    small: number;
    xSmall: number;
    tiny: number;
    xTiny: number;
  };
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  secondaryBorderRadius: 10,
  buttonBorderRadius: 4,
  primaryGradient: 'linear-gradient(50deg, #0b57ff 0%,#16C225 100%)',
  loginBackgroundGradient: 'linear-gradient(50deg, #0f6fcc 0%,#16a858 100%)',
  carouselImageBackgroundGradient:
    'linear-gradient(222.02deg, #16BB33 0%, #0E69DA 100%)',
  centerColumnWidth: 1200,
  colors: {
    primaryBlack: '#0C152E',
    primaryBlue: '#0b57ff',
    primaryGreen: '#10b73b',
    secondaryBlack: '#34373A',
    secondaryGreen: '#15C324',
    hoverGreen: '#E7F8E6',
    primaryRed: '#CB463A',
    secondaryGrey: '#9ea2ad',
    primaryWhite: '#FFFFFF',
    borderGrey: '#ededed',
    backgroundGrey: '#f7f9fb',
    searchHoverGrey: '#F5FAFC',
    textBlue: '#1081AA',
    textGrey: '#737988',
    progressGrey: '#d2d6db',
    lightBlue: '#E2E9F8',
    navyBlue: '#0E66E0',
    footerBlue: '#1D2B35',
    footerGrey: '#93A1B0',
    blueRibbon: '#0E5AF9',
    approxIron: '#CCCDCD',
    approxZircon: '#E9EBEB',
    secondaryDanger: '#DB5C5D',
    navyBlue2: '#0E63E8',
    silver: '#BFBFBF',
  },
  shadows: {
    primary: '2px 2px 10px rgba(0,0,0,0.07)',
    body: 'inset 0px -2px 10px 0px #0000001f',
  },
  fontSizes: {
    extraLargeHeading: 34,
    heading: 25,
    smallHeading: 24,
    xSmallHeading: 21,
    tinyHeading: 18,
    extraLarge: 18,
    large: 16,
    default: 14,
    small: 13,
    xSmall: 12,
    tiny: 11,
    xTiny: 9,
  },
};

export default theme;
