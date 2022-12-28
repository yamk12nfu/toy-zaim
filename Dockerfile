FROM python:3.10-slim

RUN pip install --upgrade pip
WORKDIR /var/toy-zaim

COPY requirements.txt /var/toy-zaim
