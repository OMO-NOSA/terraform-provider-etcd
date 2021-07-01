output "user" {
  value = etcd_user.user
  sensitive = true
}

output "key" {
  value = etcd_key_value.edu
}

output "user_data" {
  value = data.etcd_users.edu
}

output "role" {
  value = etcd_role.role
}

output "val" {
  value = data.etcd_key_value.edu
}

output "role_perms" {
  value = etcd_grant_role_permission.perm
}

output "cluster_data" {
  value = data.etcd_cluster.edu
}

output "users_role" {
  value = etcd_grant_user_role.gmt
}

output "auth" {
  value = etcd_auth.auth
}
