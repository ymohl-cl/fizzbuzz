# Image base to get deps
FROM golang:1.11.4 AS deps

# Force the go compiler to use modules
ENV GO111MODULE=on
 
WORKDIR /go/src/app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get -u golang.org/x/lint/golint

# Image ci to test the validate code
FROM deps AS ci

COPY . .
RUN golint ./...
RUN go test ./...
RUN go build

# Image runtime to execute the binary
FROM ci AS runtime

ENTRYPOINT ["./fizzbuzz"]