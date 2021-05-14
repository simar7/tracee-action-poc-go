FROM simar7/trcghaction:latest
COPY entrypoint.sh /
RUN apk --no-cache add bash
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]







# Specify the version of Go to use
#FROM golang:1.16 as builder

# Install upx (upx.github.io) to compress the compiled tracee-action
#RUN apt-get update && apt-get -y install upx

# Turn on Go modules support and disable CGO
#ENV GO111MODULE=on CGO_ENABLED=0 GOPATH=""

# Copy all the files from the host into the container
#COPY . .

# Compile the tracee-action - the added flags instruct Go to produce a
# standalone binary
#RUN go build \
#  -a \
#  -trimpath \
#  -ldflags "-s -w -extldflags '-static'" \
#  -installsuffix cgo \
#  -tags netgo \
#  -o /bin/tracee-action \
#  .

# Strip any symbols - this is not a library
#RUN strip /bin/tracee-action

# Compress the compiled tracee-action
#RUN upx -q -9 /bin/tracee-action


# Step 2

# Use the most basic and empty container - this container has no
# runtime, files, shell, libraries, etc.
#FROM scratch

# Copy over SSL certificates from the first step - this is required
# if our code makes any outbound SSL connections because it contains
# the root CA bundle.
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy over the compiled tracee-action from the first step
#COPY --from=builder /bin/tracee-action /bin/tracee-action

# Specify the container's entrypoint as the tracee-action
#ENTRYPOINT ["/bin/tracee-action"]