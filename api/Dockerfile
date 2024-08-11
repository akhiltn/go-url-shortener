# Multistage Docker build. This reduces the size of the final Docker image.
# Stage 1: Build the app
FROM golang:alpine as builder

RUN mkdir /build

COPY . /build/

WORKDIR /build

RUN go build -o main .

# Stage 2: Deploy the app built in stage 1
FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]
