FROM golang:latest as build-stage

# enable static go builds
ENV CGO_ENABLED=0

WORKDIR /app
COPY ./backend .
RUN go build -o main .


FROM scratch

COPY --from=build-stage /app/main .
EXPOSE 8080

CMD ["./main"]
