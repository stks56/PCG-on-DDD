FROM golang:1.17 AS builder

RUN mkdir /workspace
WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/pcg server.go


FROM scratch

COPY --from=builder /bin/pcg /bin/pcg

EXPOSE 1323
CMD ["/bin/pcg"]
