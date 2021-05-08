A minimal static web server to serve Vue applications with history-mode

Example:
```dockerfile
# build
FROM node:12 AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM n8maninger/vue-router

COPY --from=build /app/dist /www
```