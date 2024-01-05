FROM golang:1.20 as builder

WORKDIR /app

COPY ../go.mod go.sum ./

COPY .. .

RUN go mod download

RUN go mod tidy

RUN go build -o main .


FROM busybox

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x main

EXPOSE 5555

CMD ["./main"]