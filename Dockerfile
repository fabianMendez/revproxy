FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -o main .

FROM alpine

COPY --from=builder /build/main /app/

ENV PORT=8080
EXPOSE 8080

WORKDIR /app
CMD ["./main"]