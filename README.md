A minimal static web server to serve Vue applications with history-mode

Example:
```dockerfile
# build
FROM node:lts AS build

WORKDIR /app

COPY . .
RUN npm install && npm run build

FROM n8maninger/vue-router

COPY --from=build /app/dist /www
```