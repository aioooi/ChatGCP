FROM golang:latest as build-backend

WORKDIR /backend

COPY ./backend .
# Enable static go builds:
ENV CGO_ENABLED=0
RUN go build -o main .

FROM scratch

COPY --from=build-backend /backend/main .
EXPOSE 80

CMD ["./main"]
