output "user" {
  value = user_resource.user
  sensitive = true
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

output "val" {
  value = data.key_value_data_source.edu
}

output "role_perms" {
  value = grant_role_permission.perm
}

output "cluster_data" {
  value = data.cluster_data_source.edu
}

# output "users_role" {
#   value = grant_user_role_resource.gmt
# }