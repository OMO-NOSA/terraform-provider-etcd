# Terraform Provider Etcd

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)

- ![Build Action Status](https://github.com/passbase/terraform-provider-etcd/actions/workflows/test.yml/badge.svg)
- ![Release Action Status](https://github.com/passbase/terraform-provider-etcd/actions/workflows/release.yml/badge.svg)


This repository contains Etcd [Terraform](https://www.terraform.io) provider for our internal tooling. It contains:

 - A resource, and a data source (`internal/provider/`),
 - Examples (`examples/`) and generated documentation (`docs/`),
 - Miscellaneous meta files.
 

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```
### Using Go Build For Linux
```sh
$ go build -o terraform-provider-etcd
```
```sh
$ export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
```
```sh
$ mkdir -p ~/.terraform.d/plugins/hashicorp.com/passbase/etcd/0.1/$OS_ARCH
```
```sh
$ mv terraform-provider-etcd ~/.terraform.d/plugins/hashicorp.com/passbase/etcd/0.1/$OS_ARCH
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/passbase/terraform-provider-etcd
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
