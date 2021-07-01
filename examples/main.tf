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
  username = "root"
  password = "root"
  
}

resource "etcd_auth" "auth" {
  is_auth_enabled = true
}

resource "etcd_key_value" "edu" {
  key = "Passbase"
  value = "Awesome"

}

resource "etcd_user" "user"{
  username = "passbase"
  password = var.user_password
 
}

resource "etcd_role" "role" {
  name = "developer"
  
}

resource "etcd_grant_role_permission" "perm" {
  role_name = "developer"
  key = etcd_key_value.edu.key
  permission = "WRITE"
  range = "test"

}

resource "etcd_grant_user_role" "gmt" {
  role_name = etcd_role.role.name
  username = etcd_user.user.username
}

data "etcd_users" "edu" {
  depends_on = [
    etcd_user.user,
  ]
}

data "etcd_key_value" "edu" {
   key = "Nosa"
}
 
data "etcd_cluster" "edu" {
  
}