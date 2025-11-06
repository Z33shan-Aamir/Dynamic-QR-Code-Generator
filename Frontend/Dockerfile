FROM node:lts as build
WORKDIR /app
COPY ./dynamic-qr-code-frontend/package*.json ./
COPY ./dynamic-qr-code-frontend .
RUN npm install 
RUN npm run build

# Production
FROM nginx:stable-alpine AS production
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]