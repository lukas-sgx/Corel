FROM golang:alpine AS builder
WORKDIR /app/
COPY . /app/
RUN go build -o corel .

FROM scratch
COPY --from=builder /app/corel /app/corel
WORKDIR /app
CMD [ "corel" ]
