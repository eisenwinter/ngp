FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build ./cmd/server/server.go 

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=build ./app/server ./server
COPY --from=build ./app/views/* ./views
COPY --from=build ./app/templates/* ./templates 
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]