/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

# Available actions: Create, Update, Delete and Import.
# After `terraform apply` of this example file it will create a new LDAP provider with the name set in `name` attribute on the PowerScale.

# PowerScale LDAP provider enables you to define, query, and modify directory services and resources.
resource "powerscale_ldap_provider" "example_ldap_provider" {
  # Required params for creating and updating.
  # Specifies the name of the LDAP provider.
  name = "ldap_provider_test"
  # Specifies the root of the tree in which to search identities.
  base_dn = "dc=tthe,dc=testLdap,dc=com"
  # Specifies the server URIs. Begin URIs with ldap:// or ldaps://
  server_uris = ["ldap://10.225.108.54"]

  # Optional ignore_unresolvable_server_urls for creating and updating. If true, ignore unresolvable server URIs.
  ignore_unresolvable_server_urls = false

  # Optional groupnet for creating. Specifies the groupnet identifier.
  groupnet = "groupnet0"

  # Optional params for creating and updating.
  # Specifies the attribute name used when searching for alternate security identities.
  alternate_security_identities_attribute = "altSecurityIdentities"
  # If true, enables authentication and identity management through the authentication provider.
  authentication = true
  # If true, connects the provider to a random server.
  balance_servers = true
  # Specifies the distinguished name for binding to the LDAP server.
  bind_dn = ""
  # Specifies which bind mechanism to use when connecting to an LDAP server. The only supported option is the 'simple' value.
  bind_mechanism = "simple"
  # Specifies the timeout in seconds when binding to an LDAP server. Value should between 1 - 3600.
  bind_timeout = 10
  # Specifies the path to the root certificates file.
  certificate_authority_file = ""
  # Specifies the time in seconds between provider online checks. Value should between 0 - 3600.
  check_online_interval = 100
  # Specifies the canonical name.
  cn_attribute = "cn"
  # Automatically create the home directory on the first login.
  create_home_directory = false
  # Specifies the hashed password value.
  crypt_password_attribute = ""
  # Specifies the LDAP Email attribute.
  email_attribute = "mail"
  # If true, enables the LDAP provider.
  enabled = true
  # If true, allows the provider to enumerate groups.
  enumerate_groups = true
  # If true, allows the provider to enumerate users.
  enumerate_users = true
  # Specifies the list of groups that can be resolved.
  findable_groups = []
  # Specifies the list of users that can be resolved.
  findable_users = []
  # Specifies the LDAP GECOS attribute.
  gecos_attribute = "gecos"
  # Specifies the LDAP GID attribute.
  gid_attribute = "gidNumber"
  # Specifies the distinguished name of the entry where LDAP searches for groups are started.
  group_base_dn = ""
  # Specifies the domain for this provider through which groups are qualified.
  group_domain = "LDAP_GROUPS"
  # Specifies the LDAP filter for group objects.
  group_filter = "(objectClass=posixGroup)"
  # Specifies the LDAP Group Members attribute.
  group_members_attribute = "memberUid"
  # Specifies the depth from the base DN to perform LDAP searches. 
  # Acceptable values: default, base, onelevel, subtree, children.
  group_search_scope = "default"
  # Specifies the path to the home directory template.
  home_directory_template = ""
  # Specifies the LDAP Homedir attribute.
  homedir_attribute = "homeDirectory"
  # If true, continues over secure connections even if identity checks fail.
  ignore_tls_errors = false
  # Specifies the groups that can be viewed in the provider.
  listable_groups = []
  # Specifies the users that can be viewed in the provider.
  listable_users = []
  # Specifies the login shell path.
  login_shell = "/bin/bash"
  # Sets the method by which group member lookups are performed. Use caution when changing this option directly.
  # Acceptable values: default, rfc2307bis.
  member_lookup_method = "default"
  # Specifies the LDAP Query Member Of attribute, which performs reverse membership queries.
  member_of_attribute = ""
  # Specifies the LDAP UID attribute, which is used as the login name.
  name_attribute = "uid"
  # Specifies the distinguished name of the entry where LDAP searches for netgroups are started.
  netgroup_base_dn = ""
  # Specifies the LDAP filter for netgroup objects.
  netgroup_filter = "(objectClass=nisNetgroup)"
  # Specifies the LDAP Netgroup Members attribute.
  netgroup_members_attribute = "memberNisNetgroup"
  # Specifies the depth from the base DN to perform LDAP searches.
  # Acceptable values: default, base, onelevel, subtree, children.
  netgroup_search_scope = "default"
  # Specifies the LDAP Netgroup Triple attribute.
  netgroup_triple_attribute = "nisNetgroupTriple"
  # Normalizes group names to lowercase before look up.
  normalize_groups = false
  # Normalizes user names to lowercase before look up.
  normalize_users = false
  # Specifies the LDAP NT Password attribute.
  nt_password_attribute = ""
  # Specifies which NTLM versions to support for users with NTLM-compatible credentials.
  # Acceptable values: all, v2only, none.
  ntlm_support = "all"
  # Specifies the provider domain.
  provider_domain = ""
  # Determines whether to continue over a non-TLS connection.
  require_secure_connection = false
  # If true, checks the provider for filtered lists of findable and unfindable users and groups.
  restrict_findable = true
  # If true, checks the provider for filtered lists of listable and unlistable users and groups.
  restrict_listable = false
  # Specifies the default depth from the base DN to perform LDAP searches.
  # Acceptable values: base, onelevel, subtree, children.
  search_scope = "subtree"
  # Specifies the search timeout period in seconds. Value should between 10 - 3600.
  search_timeout = 100
  # Sets the attribute name that indicates the absolute date to expire the account.
  shadow_expire_attribute = "shadowExpire"
  # Sets the attribute name that indicates the section of the shadow map that is used to store the flag value.
  shadow_flag_attribute = "shadowFlag"
  # Sets the attribute name that indicates the number of days of inactivity that is allowed for the user.
  shadow_inactive_attribute = "shadowInactive"
  # Sets the attribute name that indicates the last change of the shadow information.
  shadow_last_change_attribute = "shadowLastChange"
  # Sets the attribute name that indicates the maximum number of days a password can be valid.
  shadow_max_attribute = "shadowMax"
  # Sets the attribute name that indicates the minimum number of days between shadow changes.
  shadow_min_attribute = "shadowMin"
  # Sets LDAP filter for shadow user objects.
  shadow_user_filter = "(objectClass=shadowAccount)"
  # Sets the attribute name that indicates the number of days before the password expires to warn the user.
  shadow_warning_attribute = "shadowWarning"
  # Specifies the LDAP Shell attribute.
  shell_attribute = "loginShell"
  # Sets the attribute name that indicates the SSH Public Key for the user.
  ssh_public_key_attribute = "sshPublicKey"
  # Specifies the status of the provider.
  status = "online"
  # If true, indicates that this provider instance was created by OneFS and cannot be removed.
  system = false
  # Specifies the minimum TLS protocol version.
  tls_protocol_min = "1.2"
  # Specifies the LDAP UID Number attribute.
  uid_attribute = "uidNumber"
  # Specifies the groups that cannot be resolved by the provider.
  unfindable_groups = ["wheel", "0", "insightiq", "15", "isdmgmt", "16"]
  # Specifies users that cannot be resolved by the provider.
  unfindable_users = ["root", "0", "insightiq", "15", "isdmgmt", "16"]
  # Sets the LDAP Unique Group Members attribute.
  unique_group_members_attribute = ""
  # Specifies a group that cannot be listed by the provider.
  unlistable_groups = []
  # Specifies a user that cannot be listed by the provider.
  unlistable_users = []
  # Specifies the distinguished name of the entry at which to start LDAP searches for users.
  user_base_dn = ""
  # Specifies the domain for this provider through which users are qualified.
  user_domain = "LDAP_USERS"
  # Specifies the LDAP filter for user objects.
  user_filter = "(objectClass=posixAccount)"
  # Specifies the depth from the base DN to perform LDAP searches.
  # Acceptable values: default, base, onelevel, subtree, children.
  user_search_scope = "default"

  # Optional params for creating and updating -  Only available for PowerScale 9.5 and above.
  # This setting controls the behavior of the certificate revocation checking algorithm when the LDAP provider is presented with a digital certificate by an LDAP server. 
  # Acceptable values: none, allowNoData, allowNoSrc, strict.
  # tls_revocation_check_level = "none"
  # Specifies the OCSP server URIs. Begin URIs with http://
  # ocsp_server_uris = []
}

# After the execution of above resource block, LDAP provider would have been created on the PowerScale array. 
# For more information, Please check the terraform state file. 