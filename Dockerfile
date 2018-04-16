FROM ysitd/dep

RUN dep ensure -vendor-only && \
    go build main.go

FROM alpine:3.6

COPY --from=builder /go/src/code.ysitd.cloud/component/aviation/sky/sky /

CMD ["/sky"]