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
 
COPY --from=build /usr/src/app/dist/frontend/browser/ /usr/share/nginx/html