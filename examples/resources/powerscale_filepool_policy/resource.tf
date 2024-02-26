/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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
# After `terraform apply` of this example file it will create a new file pool policy with the name set in `name` attribute on the PowerScale.

# PowerScale File Pool Policy can identify logical groups of files and specify storage operations for these files.
resource "powerscale_filepool_policy" "example_filepool_policy" {
  # Required name for creating and updating.
  # A unique name for this policy. If the policy is default policy, its name should be "Default policy". (Update Supported)
  name = "filePoolPolicySample"

  # Optional is_default_policy for creating. Specifies if the policy is default policy. 
  # Default policy applies to all files not selected by higher-priority policies.
  is_default_policy = false

  # Optional params for creating and updating.
  # Specifies the file matching rules for determining which files will be managed by this policy. (Update Supported)
  file_matching_pattern = {
    or_criteria = [
      {
        and_criteria = [
          # {
          #   # The file attribute to be compared to a given value. 
          #   # Acceptable values are: name, path, link_count, accessed_time, birth_time, changed_time, metadata_changed_time, size, file_type, custom_attribute.
          #   type = "custom_attribute"
          #   # The comparison operator to use while comparing an attribute with its value.
          #   # Acceptable values are: ==, !=, >, >=, <, <=, !.
          #   operator = "=="
          #   # The value to be compared against a file attribute
          #   value = "high"
          #   # Whether time units refer to a calendar date and time (e.g., Jun 3, 2009) or a relative duration (e.g., 2 weeks) (valid only with 'type' in {accessed_time, birth_time, changed_time or metadata_changed_time}
          #   use_relative_time = true
          #   # Indicates whether the existence of an attribute indicates a match (valid only with 'type' = 'custom_attribute')
          #   attribute_exists = true
          #   # True to match the path exactly, False to match any subtree. (valid only with 'type' = 'path')
          #   begins_with = true
          #   # True to indicate case sensitivity when comparing file attributes (valid only with 'type' = 'name' or 'type' = 'path')
          #   case_sensitive = true
          #   # File attribute field name to be compared in a custom comparison (valid only with 'type' = 'custom_attribute')
          #   field = "importance"
          #   # Size unit value. One of 'B','KB','MB','GB','TB','PB','EB' (valid only with 'type' = 'size')
          #   units = "GB"
          # },
          {
            operator = ">"
            type     = "size"
            units    = "B"
            value    = "1073741824"
          },
          {
            operator          = ">"
            type              = "birth_time"
            use_relative_time = true
            value             = "20"
          },
          {
            operator          = ">"
            type              = "metadata_changed_time"
            use_relative_time = false
            value             = "1704742200"
          },
          {
            operator          = "<"
            type              = "accessed_time"
            use_relative_time = true
            value             = "20"
          }
        ]
      },
      {
        and_criteria = [
          {
            operator          = "<"
            type              = "changed_time"
            use_relative_time = false
            value             = "1704820500"
          },
          {
            attribute_exists = false
            field            = "test"
            type             = "custom_attribute"
            value            = ""
          },
          {
            operator = "!="
            type     = "file_type"
            value    = "directory"
          },
          {
            begins_with    = false
            case_sensitive = true
            operator       = "!="
            type           = "path"
            value          = "test"
          },
          {
            case_sensitive = true
            operator       = "!="
            type           = "name"
            value          = "test"
          }
        ]
      }
    ]
  }
  # A list of actions to be taken for matching files. (Update Supported)
  actions = [
    # {
    #   # Action for set_data_access_pattern type. Set data access pattern optimization. Acceptable values: random, concurrency, streaming.
    #   data_access_pattern_action = "concurrency"
    #   # Action for apply_data_storage_policy.
    #   data_storage_policy_action = {
    #     # Specifies the SSD strategy. Acceptable values: metadata, metadata-write, data, avoid.
    #     ssd_strategy ="metadata"
    #     # Specifies the storage target.
    #     storagepool = "anywhere"
    #   }
    #   # Action for apply_snapshot_storage_policy.
    #   snapshot_storage_policy_action = {
    #     # Specifies the SSD strategy. Acceptable values: metadata, metadata-write, data, avoid.
    #     ssd_strategy ="metadata"
    #     # Specifies the snapshot storage target.
    #     storagepool = "anywhere"
    #   }
    #   # Action for enable_coalescer type. Set write performance optimization. True to enable SmartCache action.
    #   enable_coalescer_action = true
    #   # Action for enable_packing type. True to enable enable_packing action.
    #   enable_packing_action = true
    #   # Action for set_requested_protection type. Acceptable values: default, +1n, +2d:1n, +2n, +3d:1n, +3d:1n1d, +3n, +4d:1n, +4d:2n, +4n, 2x, 3x, 4x, 5x, 6x, 7x, 8x.
    #   requested_protection_action = "default"
    #   # Action for set_cloudpool_policy type.
    #   cloudpool_policy_action = {
    #     # Specifies if files with snapshots should be archived.
    #     archive_snapshot_files = true
    #     # Specifies default cloudpool cache settings for new filepool policies.
    #     cache = {
    #       # Specifies cache expiration.
    #       expiration = 86400
    #       # Specifies cache read ahead type. Acceptable values: partial, full.
    #       read_ahead = "partial"
    #       # Specifies cache type. Acceptable values: cached, no-cache.
    #       type = "cached"
    #     }
    #     # Specifies if files should be compressed.
    #     compression = true
    #     # Specifies the minimum amount of time archived data will be retained in the cloud after deletion.
    #     data_retention = 604800
    #     # Specifies if files should be encrypted.
    #     encryption = true
    #     # The minimum amount of time cloud files will be retained after the creation of a full NDMP backup. (Used with NDMP backups only.  Not applicable to SyncIQ.)
    #     full_backup_retention = 145152000
    #     # The minimum amount of time cloud files will be retained after the creation of a SyncIQ backup or an incremental NDMP backup. (Used with SyncIQ and NDMP backups.)
    #     incremental_backup_retention = 145152000
    #     pool = "cloudPool_policy"
    #     # The minimum amount of time to wait before updating cloud data with local changes.
    #     writeback_frequency = 32400
    #   },
    #   # action_type Acceptable values: set_requested_protection, set_data_access_pattern, enable_coalescer, apply_data_storage_policy, apply_snapshot_storage_policy, set_cloudpool_policy, enable_packing.
    #   action_type = "action_type_acceptable_value"
    # },
    {
      data_access_pattern_action = "concurrency"
      action_type                = "set_data_access_pattern"
    },
    {
      data_storage_policy_action = {
        ssd_strategy = "metadata"
        storagepool  = "anywhere"
      }
      action_type = "apply_data_storage_policy"
    },
    {
      snapshot_storage_policy_action = {
        ssd_strategy = "metadata"
        storagepool  = "anywhere"
      }
      action_type = "apply_snapshot_storage_policy"
    },
    {
      requested_protection_action = "default"
      action_type                 = "set_requested_protection"
    },
    {
      enable_coalescer_action = true
      action_type             = "enable_coalescer"
    },
    {
      enable_packing_action = true,
      action_type           = "enable_packing"
    },
    {
      action_type = "set_cloudpool_policy"
      cloudpool_policy_action = {
        archive_snapshot_files = true
        cache = {
          expiration = 86400
          read_ahead = "partial"
          type       = "cached"
        }
        compression                  = true
        data_retention               = 604800
        encryption                   = true
        full_backup_retention        = 145152000
        incremental_backup_retention = 145152000
        pool                         = "cloudPool_policy"
        writeback_frequency          = 32400
      }
    }
  ]
  # A description for this policy. (Update Supported)
  description = "filePoolPolicySample description"
  # The order in which this policy should be applied (relative to other policies). (Update Supported)
  apply_order = 1
}
# After the execution of above resource block, file pool policy would have been created on the PowerScale array. 
# For more information, Please check the terraform state file. 


# PowerScale Default File Pool Policy applies to all files not selected by higher-priority and can specify storage operations for these files.
resource "powerscale_filepool_policy" "example_default_policy" {
  # Required name for creating and updating.
  # Default Policy name should be "Default policy". (Update Supported)
  name = "Default policy"

  # Optional is_default_policy for creating. Specifies if the policy is default policy. 
  # Default policy applies to all files not selected by higher-priority policies.
  is_default_policy = true

  # Optional actions when creating and updating.
  # A list of actions to be taken for matching files. (Update Supported)
  actions = [
    # {
    #   # Action for set_data_access_pattern type. Set data access pattern optimization. Acceptable values: random, concurrency, streaming.
    #   data_access_pattern_action = "concurrency"
    #   # Action for apply_data_storage_policy.
    #   data_storage_policy_action = {
    #     # Specifies the SSD strategy. Acceptable values: metadata, metadata-write, data, avoid.
    #     ssd_strategy ="metadata"
    #     # Specifies the storage target.
    #     storagepool = "anywhere"
    #   }
    #   # Action for apply_snapshot_storage_policy.
    #   snapshot_storage_policy_action = {
    #     # Specifies the SSD strategy. Acceptable values: metadata, metadata-write, data, avoid.
    #     ssd_strategy ="metadata"
    #     # Specifies the snapshot storage target.
    #     storagepool = "anywhere"
    #   }
    #   # Action for enable_coalescer type. Set write performance optimization. True to enable SmartCache action.
    #   enable_coalescer_action = true
    #   # Action for enable_packing type. True to enable enable_packing action.
    #   enable_packing_action = true
    #   # Action for set_requested_protection type. Acceptable values: default, +1n, +2d:1n, +2n, +3d:1n, +3d:1n1d, +3n, +4d:1n, +4d:2n, +4n, 2x, 3x, 4x, 5x, 6x, 7x, 8x.
    #   requested_protection_action = "default"
    #   # action_type Acceptable values: set_requested_protection, set_data_access_pattern, enable_coalescer, apply_data_storage_policy, apply_snapshot_storage_policy, enable_packing.
    #   action_type = "action_type_acceptable_value"
    # },
    {
      data_access_pattern_action = "concurrency"
      action_type                = "set_data_access_pattern"
    },
    {
      data_storage_policy_action = {
        ssd_strategy = "metadata"
        storagepool  = "anywhere"
      }
      action_type = "apply_data_storage_policy"
    },
    {
      snapshot_storage_policy_action = {
        ssd_strategy = "metadata"
        storagepool  = "anywhere"
      }
      action_type = "apply_snapshot_storage_policy"
    },
    {
      requested_protection_action = "default"
      action_type                 = "set_requested_protection"
    },
    {
      enable_coalescer_action = true
      action_type             = "enable_coalescer"
    },
    {
      enable_packing_action = true,
      action_type           = "enable_packing"
    }
  ]
}
# After the execution of above resource block, Default File Pool Policy would have been cached in terraform state file, or
# Default File Pool Policy would have been updated on PowerScale.
# For more information, Please check the terraform state file.