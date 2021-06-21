terraform {
  required_providers {
    etcd = {
      version = "0.1"
      source  = "hashicorp.com/passbase/etcd"
    }
  }
}

provider "etcd" {}


resource "kv_resource" "edu" {
    key = "Nosa"
    value = "Male"
}

output "Kv" {
    value = kv_resource.edu
}