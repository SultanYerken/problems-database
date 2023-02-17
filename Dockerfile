FROM golang

ENV GOPATH=/

COPY ./ ./

RUN go build -o problems-database ./cmd/main.go

CMD ["./problems-database"]