FROM alpine:latest
# copy over the binary from the first stage
RUN ls -la

COPY chatgpt-adapter /chatgpt-adapter/chatgpt-adapter

#ADD config.yaml /chatgpt-adapter
WORKDIR "/chatgpt-adapter"
ENTRYPOINT [ "/chatgpt-adapter/chatgpt-adapter" ]
