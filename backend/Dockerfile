FROM golang:latest as builder
COPY go.mod go.sum /go/src/github.com/dunkbing/sfw-checker-viet/backend/
WORKDIR /go/src/github.com/dunkbing/sfw-checker-viet/backend
RUN go mod download
COPY . /go/src/github.com/dunkbing/sfw-checker-viet/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/checker-backend github.com/dunkbing/sfw-checker-viet/backend

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/dunkbing/sfw-checker-viet/backend/build/checker-backend /usr/bin/checker-backend
EXPOSE 8080 8080
ENTRYPOINT [ "/usr/bin/checker-backend" ]
