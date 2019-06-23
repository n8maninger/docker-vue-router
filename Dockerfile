FROM nginx:alpine

COPY ./vue-router.conf /etc/nginx/conf.d/default.conf