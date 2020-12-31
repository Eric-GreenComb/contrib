# rsa

## openssl

#### private key

openssl genrsa -out private.pem 1024

#### private 2 pkcs8

openssl pkcs8 -topk8 -inform PEM -in private.pem -outform pem -nocrypt -out pkcs8_private.pem

#### public key

openssl rsa -in private.pem -pubout -out public.pem