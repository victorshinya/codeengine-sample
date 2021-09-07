# Build the source code with Lang container
FROM golang:alpine as build
COPY main.go /
RUN go build -o /main /main.go

# Copy the exec file to a smaller base image
FROM alpine
COPY --from=build /main /main
CMD /main
