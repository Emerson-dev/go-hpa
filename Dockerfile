FROM golang:1.11 as builder

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix "static" go-hpa

FROM scratch

USER 1000

COPY --from=builder /go/bin/go-hpa /bin/go-hpa

EXPOSE 8000

ENTRYPOINT ["/bin/go-hpa"]
