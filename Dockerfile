FROM centos:7

LABEL maintainer="shubham@linux.com"

COPY bin/ /bin/

CMD ["thedarnapi"]
