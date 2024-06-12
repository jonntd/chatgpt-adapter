FROM alpine:latest
# copy over the binary from the first stage
COPY chatgpt-adapter /chatgpt-adapter/
COPY config.yaml /chatgpt-adapter/
#ADD config.yaml /chatgpt-adapter
WORKDIR "/chatgpt-adapter"
ENTRYPOINT [ "/chatgpt-adapter/chatgpt-adapter" ]
