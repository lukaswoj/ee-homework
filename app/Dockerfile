FROM golang:1.15.7-alpine AS build
RUN apk add --no-cache git
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go get github.com/gorilla/mux
RUN go build -o /app src/main.go

FROM scratch AS bin
COPY --from=build /app /
