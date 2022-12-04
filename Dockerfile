FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY ../tinyurl-api ./
ADD . /app
RUN echo $(ls -l ./)
#RUN go get /app
#RUN go install
RUN go build -o tinyl_api
Expose 9080
CMD ["./tinyl_api"]