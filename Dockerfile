FROM golang

WORKDIR /src

COPY cmd/kubia/main.go .

RUN go build main.go

FROM ubuntu

COPY --from=0 /src/main .

CMD ["./main"]
