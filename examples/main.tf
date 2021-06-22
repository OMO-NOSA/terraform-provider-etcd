terraform {
  required_providers {
    etcd = {
      version = "~>0.1"
      source  = "hashicorp.com/passbase/etcd"
    }
  }
}

provider "etcd" {
   endpoints = [ "localhost:2379" ]
}

data "cluster_data_source" "edu" {
    provider = etcd
}

resource "key_value_resource" "edu" {
    key = "Nosa"
    value = "Male"
    provider = etcd
    
}

output "key" {
    value = key_value_resource.edu
}

output "cluster_data" {
  value = data.cluster_data_source.edu
}