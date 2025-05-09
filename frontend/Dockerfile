# syntax=docker/dockerfile:1

# Build image
FROM node:22 AS build

WORKDIR /usr/src/app

COPY package.json package-lock.json ./
RUN npm install

COPY . .
RUN npm run build --configuration="production" --base-href="./"

# Runtime image
FROM nginx:stable-bullseye

# Copy Angular build output to nginx html folder
COPY --from=build /usr/src/app/dist/frontend/browser/ /usr/share/nginx/html

RUN printf 'server {\n\
    root /usr/share/nginx/html;\n\
    index index.html;\n\
\n\
    include /etc/nginx/mime.types;\n\
    default_type application/octet-stream;\n\
\n\
    location / {\n\
        try_files $uri $uri/ /index.html;\n\
    }\n\
}\n' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
