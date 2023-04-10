FROM golang:latest as build-stage

WORKDIR /app

COPY ./backend .
# Enable static go builds:
ENV CGO_ENABLED=0
RUN go build -o main .


FROM scratch

COPY --from=build-stage /app/main .
COPY ./frontend ./public

EXPOSE 80

CMD ["./main"]
