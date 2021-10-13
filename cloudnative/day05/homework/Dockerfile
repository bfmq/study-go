FROM busybox

EXPOSE 8000
ENV VERSION=123456
COPY homework /root
RUN chmod +x /root/homework

ENTRYPOINT ["/root/homework"]
