# Build the Go Api
FROM golang:latest AS builder
# ADD . /app
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Build the React application
FROM node:alpine AS node_builder
COPY --from=builder /app/frontend/package*.json ./
ENV NODE_ENV=production
RUN npm install --production
COPY --from=builder /app/frontend/ ./
RUN npm run build

# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
COPY --from=node_builder /build ./frontend
RUN chmod +x ./main
EXPOSE 8080
CMD ./main