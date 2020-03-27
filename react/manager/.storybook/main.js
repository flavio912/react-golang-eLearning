const path = require('path');
module.exports = {
  stories: ['../src/**/*.stories.[tj]s[x]'],
  addons: ['@storybook/preset-typescript', '@storybook/addon-knobs/register']
};
