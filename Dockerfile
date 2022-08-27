FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN apk -U --no-cache add git gcc make
RUN CGO_ENABLED=0 go build -o dist/ -a -trimpath -ldflags="-s -w" -tags='netgo timetzdata' ./cmd/routerd

FROM scratch

COPY --from=build /app/dist/routerd /usr/local/bin

EXPOSE 80/tcp

CMD ["routerd", "/www"]