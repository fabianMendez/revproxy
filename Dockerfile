FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -buildvcs=false -o main .

FROM alpine

COPY --from=builder /build/main /app/

ENV PORT=8080
EXPOSE 8080

WORKDIR /app
CMD ["./main"]
