FROM golang:1.15-alpine3.12 as build
WORKDIR /api
COPY . /api
RUN go build ./service/api

FROM alpine:3.12
COPY --from=build api ./api

EXPOSE 3000

CMD [ "./api" ]
