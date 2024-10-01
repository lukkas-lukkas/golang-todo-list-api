FROM golang:1.23

WORKDIR /usr/src/app

COPY . .

ENTRYPOINT [ "go", "run" ]
CMD [ "main.go" ]