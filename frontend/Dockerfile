# syntax=docker/dockerfile:1

# Build image
FROM node:22 AS build

WORKDIR /usr/src/app

COPY package.json package-lock.json ./
RUN npm install

COPY . .
RUN npm run build --configuration=production

# Runtime image
FROM nginx:stable-bullseye

RUN rm -rf /usr/share/nginx/html/*

RUN printf 'server {\n\
    listen 80;\n\
    server_name _;\n\
    root /usr/share/nginx/html;\n\
    location / {\n\
        try_files $uri $uri/ /index.html;\n\
    }\n\
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg)$ {\n\
        try_files $uri =404;\n\
    }\n\
}\n' > /etc/nginx/conf.d/default.conf

COPY --from=build /usr/src/app/dist/frontend/browser/ /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
