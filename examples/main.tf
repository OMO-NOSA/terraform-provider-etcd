terraform {
  required_providers {
    etcd = {
      version = "0.1"
      source  = "hashicorp.com/passbase/etcd"
    }
  }
}

provider "etcd" {
  endpoints = [ "localhost:2379" ]
}

resource "key_value_resource" "edu" {
  key = "Passbase"
  value = "Awesome"
  provider = etcd
    
}

resource "user_resource" "user"{
  username = "Passbase"
  password = var.user_password
  provider = etcd
}

resource "role_resource" "role" {
  name = "developer"
  provider = etcd

}

resource "grant_role_permission" "perm" {
  role_name = "developer"
  key = "checking"
  permission = "WRITE"
  range = "test"
  provider = etcd

}

resource "grant_user_role_resource" "gmt" {
  role_name = "developer"
  username = "passbase"
  provider = etcd
  
  depends_on = [
    user_resource.user,
    role_resource.role,
  ]
}

data "users_data_source" "edu" {
  provider = etcd
}

data "key_value_data_source" "edu" {
   provider = etcd
   key = "Nosa" 
}
 
data "cluster_data_source" "edu" {
   provider = etcd
}