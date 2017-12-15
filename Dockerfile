FROM stepsaway/baseimage
LABEL maintainer="JD Trout <trout.jd@gmail.com>"

# Set correct environment variables and packages
ENV HOME=/root GOPATH=/go PATH=/go/bin:/usr/local/go/bin:${PATH}

# Make sure ruby matches RBENV and perform cores setup
RUN curl -sL https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz | tar -xzf - -C /usr/local && \
    mkdir -p /go/src/github.com/StepsAway/svcwr && \
    rm -rf /tmp/* && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    echo /root > /etc/container_environment/HOME # ensure home is set when command is run

# Set working directory to /go and copy in source
WORKDIR /go/src/github.com/iDiogenes/exctracker
COPY . /go/src/github.com/iDiogenes/exctracker


# Build and clean source
RUN go get && \
    go build && \
    rm -rf /go/src

WORKDIR /go

# Use baseimage-docker's init process.
CMD ["/sbin/my_init"]
