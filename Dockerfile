FROM alpine:3.15.0
# copy over the binary from the first stage
COPY chatgpt-adapter /chatgpt-adapter/chatgpt-adapter

WORKDIR "/chatgpt-adapter"
ENTRYPOINT [ "/chatgpt-adapter/chatgpt-adapter" ]
