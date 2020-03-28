const spacing = {
  0: 5,
  1: 10,
  2: 20,
  3: 50,
};

const theme = {
  spacing: (amount: 0 | 1 | 2 | 3) => spacing[amount],
  borderRadius: 5,
  colors: {
    primaryBlack: "#061729",
  },
};

export default theme;
