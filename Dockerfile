FROM golang:latest as build-backend

WORKDIR /backend

COPY ./backend .
# Enable static go builds:
ENV CGO_ENABLED=0
RUN go build -o main .


FROM node:18-alpine as build-frontend

WORKDIR /frontend

COPY ./frontend .
RUN npm install
RUN npm run build
RUN echo `ls -l build`


FROM scratch

COPY --from=build-backend /backend/main .
COPY --from=build-frontend /frontend/build ./public

EXPOSE 80

CMD ["./main"]
