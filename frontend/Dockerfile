FROM node:14 as build

WORKDIR /frontend

COPY . .

ARG GATSBY_DOMAIN
ENV GATSBY_DOMAIN $GATSBY_DOMAIN

ARG GATSBY_PORT
ENV GATSBY_PORT $GATSBY_PORT

RUN npm ci

RUN npm run build

FROM nginx

COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /frontend/public /usr/share/nginx/html