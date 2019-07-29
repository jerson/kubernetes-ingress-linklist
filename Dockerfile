FROM alpine
MAINTAINER Julian Kleinhans <julian.kleinhans@aoe.com>
RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN mkdir -p /usr/bin
ADD bin/kubernetes-ingress-linklist /usr/bin/
RUN chmod +x /usr/bin/kubernetes-ingress-linklist
ENV PATH $PATH:/usr/bin

EXPOSE 8080

CMD ["/usr/bin/kubernetes-ingress-linklist"]