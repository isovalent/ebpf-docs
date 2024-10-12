FROM python:3.9.16-slim

LABEL org.opencontainers.image.source=https://github.com/isovalent/ebpf-docs

COPY requirements.txt /root/requirements.txt
RUN apt update && \
    apt install -y git aspell libcairo2-dev libfreetype6-dev libffi-dev \
        libjpeg-dev libpng-dev libz-dev && \
    pip install -r /root/requirements.txt && \
    rm -rf /var/cache/apt/archives /var/lib/apt/lists/*

SHELL ["/bin/bash", "-c"]

RUN echo $'#!/bin/bash \n\
\n\
if [[ -n "${AS_USER}" ]]; then \n\
    AS_GROUP=${AS_GROUP:-${AS_USER}} \n\
    \n\
    groupadd -g $AS_GROUP docs || true \n\
    useradd -u $AS_USER -g $AS_GROUP -m -s /bin/bash docs || true \n\
    \n\
fi \n\
if [ "$#" -ne 0 ]; then \n\
    exec "runuser" "-u" $(id -nu ${AS_USER}) "--" "/bin/bash" "-c" "$@"\n\
else \n\
    exec "runuser" "-u" $(id -nu ${AS_USER}) "--" "/bin/bash" \n\
fi \n\
' >> /bin/entrypoint.sh && \
chmod +x /bin/entrypoint.sh

ENTRYPOINT ["/bin/entrypoint.sh"]
