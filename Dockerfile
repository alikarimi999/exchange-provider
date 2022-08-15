FROM golang:alpine3.16

RUN apk add gcompat

COPY . /app/
WORKDIR /app


RUN chmod +x order_service
CMD ["./order_service" ]