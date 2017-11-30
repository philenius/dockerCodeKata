FROM node:slim

WORKDIR /app

COPY . /app

RUN npm install

ENV PATH=/app/node_modules/http-server/bin:$PATH

ENTRYPOINT [ "http-server" ]