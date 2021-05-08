FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN apk -U --no-cache add upx git gcc make \
	&& make static \
	&& upx ./dist/server

FROM scratch

COPY --from=build /app/dist/server /

EXPOSE 80/tcp

VOLUME [ "/www" ]

CMD ["/server", "/www"]