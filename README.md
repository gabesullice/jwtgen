# jwtgen

## Usage

```
Usage of ./jwtgen:
  -iss string
    	The token issuer. (default "www.campus.com")
  -sub string
    	The token subject. (default "jack.ryan")
```

```shell
./jwtgen -iss starfleet -sub kirk < ./private-key.pem
```

## Install

Build the program with `make`:

```shell
make
```

Or, if you have [Docker][install-docker] but not [Go][install-go] installed:

```shell
docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app golang:1 make
```

If you do not already have a private key, you can generate a public and private
key pair one with the following commands:

```shell
openssl genrsa -out ./private-key.pem 2048
openssl rsa \
  -in ./private-key.pem \
  -outform PEM \
  -pubout \
  -out ./public-key.pem
```

[install-docker]: https://www.docker.com/get-started/
[install-go]: https://go.dev/doc/install
