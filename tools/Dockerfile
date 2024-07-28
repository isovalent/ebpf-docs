FROM python:3.9.16-slim

ARG UID=1000
ARG GID=1000

COPY requirements.txt /root/requirements.txt
RUN apt update && \
    apt install -y git aspell && \
    pip install -r /root/requirements.txt && \
    groupadd -g $GID docs && \
    useradd -u $UID -g $GID -m -s /bin/bash docs
ENTRYPOINT bash
