FROM alpine:3.15.0
# copy over the binary from the first stage
COPY chatgpt-adapter /helloworld/helloworld

WORKDIR "/helloworld"
ENTRYPOINT [ "/helloworld/helloworld" ]
