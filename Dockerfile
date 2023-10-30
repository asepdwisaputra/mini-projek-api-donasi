FROM golang:alpine3.18 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
#RUN go clean --modcache #kalo udah pernah build/udah download
RUN go build -o main
# EXPOSE 8000
# CMD ["/app/main"]

# stage 2
FROM alpine:3.18
WORKDIR /root/
# COPY --from=builder /app/config.json .
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]