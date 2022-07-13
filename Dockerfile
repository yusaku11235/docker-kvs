FROM  golang:1.13.4-alpine3.10 as builder
WORKDIR /src
COPY ./index.go /src
RUN go build -o yusaku_appserver index.go

FROM alpine:3.10.3
COPY --from=builder /src/yusaku_appserver /bin/yusaku_appserver
CMD ["/bin/yusaku_appserver"]
