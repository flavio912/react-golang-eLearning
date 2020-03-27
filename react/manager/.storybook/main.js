const path = require('path');
module.exports = {
  stories: ['../src/**/*.stories.(ts|tsx|js|jsx|mdx)'],
  addons: [
    '@storybook/preset-typescript',
    '@storybook/addon-knobs/register',
    {
      name: '@storybook/addon-docs',
      options: {
        configureJSX: true
      }
    }
  ]
};
