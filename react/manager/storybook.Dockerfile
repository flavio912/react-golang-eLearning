FROM node:10

ENV HOME /app

WORKDIR /app

CMD ["yarn"]
ENTRYPOINT ["yarn", "run", "storybook"]
EXPOSE 4001