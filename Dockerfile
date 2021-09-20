FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -o /go/bin/app main.go

FROM scratch
COPY --from=build /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]