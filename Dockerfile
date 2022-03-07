FROM alpine as build
RUN apk add go
COPY . /src
WORKDIR /src
RUN go build .

FROM alpine
COPY --from=build /src/goserve /goserve
WORKDIR /
EXPOSE 3000

CMD ["/goserve", "-d", "/http"]