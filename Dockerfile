FROM centos:7

LABEL maintainer="shubham@linux.com"

COPY bin/ build/entrypoint.sh /bin/

ENTRYPOINT /bin/entrypoint.sh
