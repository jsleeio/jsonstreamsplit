# Start by building the application.
FROM golang:1.11-alpine3.8 AS build
WORKDIR /jsonstreamsplit
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

# Now copy it into our base image.
FROM scratch
COPY --from=build /jsonstreamsplit/jsonstreamsplit /jsonstreamsplit
USER 1000
ENTRYPOINT ["/jsonstreamsplit"]
