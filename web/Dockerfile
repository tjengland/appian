FROM node:alpine3.13 as base

ADD . .
RUN npm i
ENV PORT="3000"
EXPOSE ${PORT}
