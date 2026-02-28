FROM golang:1.23-bookworm

RUN apt-get update && apt-get install -y \
    nodejs \
    npm \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
