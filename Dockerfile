FROM golang:alpine3.16

COPY . /app/
WORKDIR /app
# run the app
RUN go mod tidy; go build -o app
RUN chmod +x app
CMD [ "./app" ]