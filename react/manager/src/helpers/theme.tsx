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
  loginBackgroundGradient: string;
  paymentSuccessBackgroundGradient: string;
  paymentButtonBackgroundGradient: string;
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
    borderBlack: string;
    backgroundGrey: string;
    textBlue: string;
    textGrey: string;
    searchHoverGrey: string;
    progressGrey: string;
    secondaryDanger: string;
    hoverDanger: string;
    textNavyBlue: string;
    textSolitude: string;
    borderGreyBold: string;
    textIron: string;
    textNavyBlue2: string;
    textGrey2: string;
    textGrey3: string;
    lightGreen: string;
  };
  shadows: {
    primary: string;
    body: string;
  };
  fontSizes: {
    heading: number;
    smallHeading: number;
    tinyHeading: number;
    extraLarge: number;
    large: number;
    smallLarge: number;
    default: number;
    small: number;
    xSmall: number;
    tiny: number;
  };
};

const theme: Theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  primaryBorderRadius: 5,
  secondaryBorderRadius: 10,
  buttonBorderRadius: 4,
  primaryGradient: 'linear-gradient(50deg, #0b57ff 0%,#16C225 100%)',
  loginBackgroundGradient: 'linear-gradient(50deg, #0f6fcc 0%,#16a858 100%)',
  paymentSuccessBackgroundGradient:
    'linear-gradient(200.08deg, #0E5AF9 0%, #0D57FF 100%)',
  paymentButtonBackgroundGradient:
    'linear-gradient(221.01deg, #16C225 0%, #0D57FF 100%)',
  colors: {
    primaryBlack: '#0C152E',
    primaryBlue: '#0b57ff',
    primaryGreen: '#10b73b',
    secondaryBlack: '#34373A',
    secondaryGreen: '#15C324',
    secondaryDanger: '#DB5C5D',
    hoverGreen: '#E7F8E6',
    hoverDanger: '#FBC7C5',
    primaryRed: '#CB463A',
    secondaryGrey: '#9ea2ad',
    primaryWhite: '#FFFFFF',
    borderGrey: '#ededed',
    borderGreyBold: '#E9EBEB',
    backgroundGrey: '#f7f9fb',
    searchHoverGrey: '#F5FAFC',
    textBlue: '#1081AA',
    textGrey: '#737988',
    progressGrey: '#d2d6db',
    textNavyBlue: '#6BAAE7',
    textSolitude: '#DFEEFD',
    textIron: '#CCCDCD',
    textNavyBlue2: '#0E63E8',
    textGrey2: '#5C7487',
    textGrey3: '#E3E3E3',
    borderBlack: '#08080814',
    lightGreen: '#81BE86'
  },
  shadows: {
    primary: '2px 2px 10px rgba(0,0,0,0.07)',
    body: 'inset 0px -2px 10px 0px #0000001f'
  },
  fontSizes: {
    heading: 25,
    smallHeading: 24,
    tinyHeading: 20,
    extraLarge: 18,
    large: 16,
    smallLarge: 15, // NOT A HELPFUL NAME!
    default: 14,
    small: 13,
    xSmall: 12,
    tiny: 11
  }
};

export default theme;
