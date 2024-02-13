FROM golang:1.21.4

RUN go version
ENV GOPATH=/

COPY ./ ./

# Build Go app
RUN go mod download
RUN go build -o audit ./cmd/main.go

CMD [ "./audit" ]
