FROM node:latest

COPY ./client/ /app/
WORKDIR /app/
RUN npm i -g
RUN npm i react-scripts

CMD ["npm", "start"]
