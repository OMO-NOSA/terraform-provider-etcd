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
  key = "Nosa"
  value = "Male"
  provider = etcd
    
}

resource "user_resource" "user"{
  username = "Mario"
  password = "Lugi"
  provider = etcd
}

resource "role_resource" "role" {
  name = "Security Engineer"
  provider = etcd

}

resource "grant_user_role_resource" "gmt" {
  role_name = "Security Engineer"
  username = "Mario"
  provider = etcd
}

# resource "grant_role_permission" "perm" {
#   role_name = "Security Engineer"
#   key = "checking"
#   permission = "READ"
#   range = "test"
#   provider = etcd

# }

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

output "user" {
  value = user_resource.user
}

output "key" {
  value = key_value_resource.edu
}

output "user_data" {
  value = data.users_data_source.edu
}

# output "role" {
#   value = role_resource.role
# }

output "val" {
  value = data.key_value_data_source.edu
}

# output "role_perms" {
#   value = grant_role_permission.perm
# }

output "cluster_data" {
  value = data.cluster_data_source.edu
}

output "users_role" {
  value = grant_user_role_resource.gmt
}