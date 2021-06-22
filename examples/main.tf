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



resource "key_value_resource" "edu" {
    key = "Nosa"
    value = "Male"
    provider = etcd
    
}

resource "user_resource" "user"{
  username = "Alan123"
  password = "1456"
  provider = etcd
}

resource "role_resource" "role" {
  name = "Security Engineer"
  provider = etcd

}

resource "grant_role_permission" "perm" {
  name = "Security Engineer"
  key = "checking"
  permission = "READ"
  range = "test"
  provider = etcd

}

data "users_data_source" "edu" {
    provider = etcd
}

output "user" {
  value = user_resource.user
}

output "key" {
    value = key_value_resource.edu
}

output "user_data" {
  value = data.users_data_source.edu
}

output "role" {
  value = role_resource.role
}