FROM golang:1.20-alpine as build
WORKDIR /usr/src/app
COPY ./code .
RUN go build main.go

FROM scratch
COPY --from=build /usr/src/app/main .
CMD ["./main"]