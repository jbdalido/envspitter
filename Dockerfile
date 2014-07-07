FROM mischief/docker-golang
MAINTAINER Jean-Baptiste Dalido <jbdalido@gmail.com>

# Add and build the main.go
ADD main.go /tmp/envspitter.go
RUN go get github.com/dotcloud/docker/pkg/mflag && \
    go get github.com/emicklei/go-restful && \
    go build -o /usr/local/bin/envspitter /tmp/envspitter.go 

# Define envspitter as entrypoint
# ex: docker run -t -i image --listen 127.0.0.1:9098
ENTRYPOINT ["/usr/local/bin/envspitter"]
