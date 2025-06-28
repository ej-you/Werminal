# --------
# FRONTEND
# --------

FROM node:alpine3.22 AS frontend

WORKDIR /frontend

COPY ./website/package.json .

RUN npm install

# ----------
# WEB-SERVER
# ----------

FROM nginx:alpine3.22 AS web-server

WORKDIR /frontend

# frontend files
COPY --from=frontend /frontend/node_modules ./node_modules
COPY ./website/index.html ./index.html
COPY ./website/static ./static

WORKDIR /etc/nginx

# nginx configs
COPY ./nginx/nginx.conf ./nginx.conf
COPY ./nginx/site.conf ./conf.d/site.conf
