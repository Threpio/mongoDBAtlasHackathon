FROM golang:1.16-alpine AS build
WORKDIR /platform
RUN apk update && apk add git make
COPY . .
RUN go mod download
RUN go build -o /dist/backend .


FROM alpine
WORKDIR /app/
COPY --from=build /dist/backend /app/backend
EXPOSE 8080
ENTRYPOINT ["/app/core", "start"]