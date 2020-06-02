FROM node:10

ENV HOME /app

WORKDIR /app

ENTRYPOINT ["yarn", "run", "storybook"]
EXPOSE 4001