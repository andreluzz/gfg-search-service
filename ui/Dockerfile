# docker build --build-arg VUE_APP_API_URL=http://localhost:8080 -t andreluzz/go-gfg-search-ui .
# docker run -it andreluzz/go-gfg-search-ui
# OBS.: Assuming rest api server is running on the host on port 8080

# Build phase
FROM node:alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm install -g @vue/cli
COPY . .
ARG VUE_APP_API_URL
ENV VUE_APP_API_URL $VUE_APP_API_URL
RUN npm run build

# Deploy phase
FROM nginx:stable-alpine
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]