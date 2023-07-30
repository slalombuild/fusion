FROM golang:1.18-alpine

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -v -o /bin/fusion cmd/fusion/main.go

LABEL author="Slalom Build"
LABEL github="https://github.com/slalombuild/fusion"

ENTRYPOINT ["/bin/fusion"]

CMD [ "/bin/fusion" ]