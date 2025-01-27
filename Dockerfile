# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/Saoirse/api

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/justinas/alice
RUN go get github.com/gorilla/context
RUN go get github.com/julienschmidt/httprouter
RUN go install Saoirse/api

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/api

EXPOSE 5000
