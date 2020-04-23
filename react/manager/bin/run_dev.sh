#!/bin/bash
yarn
yarn run relay --watch &
yarn run storybook & 
yarn run start