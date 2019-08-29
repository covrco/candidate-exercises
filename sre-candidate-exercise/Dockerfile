FROM golang:1.12
ADD . /go/src/github.com/covrco/sre-candidate-exercise
RUN go get github.com/golang/protobuf/proto
RUN go get google.golang.org/grpc
CMD go run /go/src/github.com/covrco/sre-candidate-exercise/cmd/word_splitter/main.go
EXPOSE 9001