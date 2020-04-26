#!/bin/bash

mkdir cert && cd cert

# Generate a CA key
openssl genrsa -out ca.key 2048

# Generate a CA certificate
openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 10000 -out ca.crt


function genkeypair {
  # Generate a $1 key
  openssl genrsa -out $1.key 2048

  # Generate a certificate signing request
  # (interactive mode) openssl req -new -key $1.key -out $1.csr
  openssl req -new -sha256 -key $1.key -subj "/C=US/ST=CA/O=MyOrg, Inc./CN=localhost" -out $1.csr

  # Generate a $1 certificate
   openssl x509 -req -in $1.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out $1.crt -days 10000 -sha256
}

genkeypair "server"
genkeypair "client"
