FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -o project-management-hub

EXPOSE 8081

ENTRYPOINT [ "./project-management-hub" ]
