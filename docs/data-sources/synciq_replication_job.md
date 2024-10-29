---
# Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "powerscale_synciq_replication_job data source"
linkTitle: "powerscale_synciq_replication_job"
page_title: "powerscale_synciq_replication_job Data Source - terraform-provider-powerscale"
subcategory: ""
description: |-
  This datasource is used to query the SyncIQ replication jobs from PowerScale array. The information fetched from this datasource can be used for getting the details.
---

# powerscale_synciq_replication_job (Data Source)

This datasource is used to query the SyncIQ replication jobs from PowerScale array. The information fetched from this datasource can be used for getting the details.

## Example Usage

```terraform
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

# PowerScale SyncIQ Replication Job datasource allows you to get a list of SyncIQ replication jobs.

# Returns a list of PowerScale SyncIQ Replication Jobs 
data "powerscale_synciq_replication_job" "all_jobs" {
}

# Returns a list of PowerScale SyncIQ Replication Jobs with running state
data "powerscale_synciq_replication_job" "jobByState" {
  filter {
    state = "running"
  }
}

# Output value of above block by executing 'terraform output' command.
# The user can use the fetched information by the variable data.powerscale_synciq_replication_job.all_jobs.synciq_jobs
output "powerscale_synciq_all_replication_jobs" {
  value = data.powerscale_synciq_replication_job.all_jobs.synciq_jobs
}

# The user can use the fetched information by the variable data.powerscale_synciq_replication_job.all_jobs.synciq_jobs
output "replicationJobByState" {
  value = data.powerscale_synciq_replication_job.jobByState.synciq_jobs
}

# After the successful execution of above said block, We can see the output value by executing 'terraform output' command.
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (Block, Optional) (see [below for nested schema](#nestedblock--filter))

### Read-Only

- `id` (String) Identifier.
- `synciq_jobs` (Attributes List) List of user groups. (see [below for nested schema](#nestedatt--synciq_jobs))

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Optional:

- `dir` (String) The direction of the sort.
- `limit` (Number) Return no more than this many results.
- `sort` (String) The field that will be used for sorting.
- `state` (String) Only list SyncIQ replication jobs matching this state.


<a id="nestedatt--synciq_jobs"></a>
### Nested Schema for `synciq_jobs`

Read-Only:

- `action` (String) The action to be taken by this job.
- `ads_streams_replicated` (Number) The number of ads streams replicated by this job.
- `block_specs_replicated` (Number) The number of block specs replicated by this job.
- `bytes_recoverable` (Number) The number of bytes recoverable by this job.
- `bytes_transferred` (Number) The number of bytes that have been transferred by this job.
- `char_specs_replicated` (Number) The number of char specs replicated by this job.
- `committed_files` (Number) The number of WORM committed files.
- `corrected_lins` (Number) The number of LINs corrected by this job.
- `dead_node` (Boolean) This field is true if the node running this job is dead.
- `directories_replicated` (Number) The number of directories replicated.
- `dirs_changed` (Number) The number of directories changed by this job.
- `dirs_deleted` (Number) The number of directories deleted by this job.
- `dirs_moved` (Number) The number of directories moved by this job.
- `dirs_new` (Number) The number of directories created by this job.
- `duration` (Number) The amount of time in seconds between when the job was started and when it ended.  If the job has not yet ended, this is the amount of time since the job started.  This field is null if the job has not yet started.
- `encrypted` (Boolean) If true, syncs will be encrypted.
- `end_time` (Number) The time the job ended in unix epoch seconds. The field is null if the job hasn't ended.
- `error` (String) The primary error message for this job.
- `error_checksum_files_skipped` (Number) The number of files with checksum errors skipped by this job.
- `error_io_files_skipped` (Number) The number of files with io errors skipped by this job.
- `error_net_files_skipped` (Number) The number of files with network errors skipped by this job.
- `errors` (List of String) A list of error messages for this job.
- `failed_chunks` (Number) They number of data chunks that failed transmission.
- `fifos_replicated` (Number) The number of fifos replicated by this job.
- `file_data_bytes` (Number) The number of bytes transferred that belong to files.
- `files_changed` (Number) The number of files changed by this job.
- `files_linked` (Number) The number of files linked by this job.
- `files_new` (Number) The number of files created by this job.
- `files_selected` (Number) The number of files selected by this job.
- `files_transferred` (Number) The number of files transferred by this job.
- `files_unlinked` (Number) The number of files unlinked by this job.
- `files_with_ads_replicated` (Number) The number of files with ads replicated by this job.
- `flipped_lins` (Number) The number of LINs flipped by this job.
- `hard_links_replicated` (Number) The number of hard links replicated by this job.
- `hash_exceptions_fixed` (Number) The number of hash exceptions fixed by this job.
- `hash_exceptions_found` (Number) The number of hash exceptions found by this job.
- `id` (String) A unique identifier for this object.
- `job_id` (Number) The ID of the job.
- `lins_total` (Number) The number of LINs transferred by this job.
- `network_bytes_to_source` (Number) The total number of bytes sent to the source by this job.
- `network_bytes_to_target` (Number) The total number of bytes sent to the target by this job.
- `new_files_replicated` (Number) The number of new files replicated by this job.
- `num_retransmitted_files` (Number) The number of files that have been retransmitted by this job.
- `phases` (Attributes List) Data for each phase of this job. (see [below for nested schema](#nestedatt--synciq_jobs--phases))
- `policy` (Attributes) The policy associated with this job, or null if there is currently no policy associated with this job (this can happen if the job is newly created and not yet fully populated in the underlying database). (see [below for nested schema](#nestedatt--synciq_jobs--policy))
- `policy_action` (String) This is the action the policy is configured to perform.
- `policy_id` (String) The ID of the policy.
- `policy_name` (String) The name of the policy.
- `quotas_deleted` (Number) The number of quotas removed from the target.
- `regular_files_replicated` (Number) The number of regular files replicated by this job.
- `resynced_lins` (Number) The number of LINs resynched by this job.
- `retransmitted_files` (List of String) The files that have been retransmitted by this job.
- `retry` (Number) The number of times the job has been retried.
- `running_chunks` (Number) The number of data chunks currently being transmitted.
- `service_report` (Attributes List) Data for each component exported as part of service replication. (see [below for nested schema](#nestedatt--synciq_jobs--service_report))
- `sockets_replicated` (Number) The number of sockets replicated by this job.
- `source_bytes_recovered` (Number) The number of bytes recovered on the source.
- `source_directories_created` (Number) The number of directories created on the source.
- `source_directories_deleted` (Number) The number of directories deleted on the source.
- `source_directories_linked` (Number) The number of directories linked on the source.
- `source_directories_unlinked` (Number) The number of directories unlinked on the source.
- `source_directories_visited` (Number) The number of directories visited on the source.
- `source_files_deleted` (Number) The number of files deleted on the source.
- `source_files_linked` (Number) The number of files linked on the source.
- `source_files_unlinked` (Number) The number of files unlinked on the source.
- `sparse_data_bytes` (Number) The number of sparse data bytes transferred by this job.
- `start_time` (Number) The time the job started in unix epoch seconds. The field is null if the job hasn't started.
- `state` (String) The state of the job.
- `succeeded_chunks` (Number) The number of data chunks that have been transmitted successfully.
- `symlinks_replicated` (Number) The number of symlinks replicated by this job.
- `sync_type` (String) The type of sync being performed by this job.
- `target_bytes_recovered` (Number) The number of bytes recovered on the target.
- `target_directories_created` (Number) The number of directories created on the target.
- `target_directories_deleted` (Number) The number of directories deleted on the target.
- `target_directories_linked` (Number) The number of directories linked on the target.
- `target_directories_unlinked` (Number) The number of directories unlinked on the target.
- `target_files_deleted` (Number) The number of files deleted on the target.
- `target_files_linked` (Number) The number of files linked on the target.
- `target_files_unlinked` (Number) The number of files unlinked on the target.
- `target_snapshots` (List of String) The target snapshots created by this job.
- `total_chunks` (Number) The total number of data chunks transmitted by this job.
- `total_data_bytes` (Number) The total number of bytes transferred by this job.
- `total_exported_services` (Number) The total number of components exported as part of service replication.
- `total_files` (Number) The number of files affected by this job.
- `total_network_bytes` (Number) The total number of bytes sent over the network by this job.
- `total_phases` (Number) The total number of phases for this job.
- `unchanged_data_bytes` (Number) The number of bytes unchanged by this job.
- `up_to_date_files_skipped` (Number) The number of up-to-date files skipped by this job.
- `updated_files_replicated` (Number) The number of updated files replicated by this job.
- `user_conflict_files_skipped` (Number) The number of files with user conflicts skipped by this job.
- `warnings` (List of String) A list of warning messages for this job.
- `workers` (Attributes List) A list of workers for this job. (see [below for nested schema](#nestedatt--synciq_jobs--workers))
- `worm_committed_file_conflicts` (Number) The number of WORM committed files which needed to be reverted. Since WORM committed files cannot be reverted, this is the number of files that were preserved in the compliance store.

<a id="nestedatt--synciq_jobs--phases"></a>
### Nested Schema for `synciq_jobs.phases`

Read-Only:

- `end_time` (Number) The time the job ended this phase.
- `phase` (String) The phase that the job was in.
- `start_time` (Number) The time the job began this phase.


<a id="nestedatt--synciq_jobs--policy"></a>
### Nested Schema for `synciq_jobs.policy`

Read-Only:

- `action` (String) The action to be taken by the job.
- `file_matching_pattern` (Attributes) A file matching pattern, organized as an OR'ed set of AND'ed file criteria, for example ((a AND b) OR (x AND y)) used to define a set of files with specific properties.  Policies of type 'sync' cannot use 'path' or time criteria in their matching patterns, but policies of type 'copy' can use all listed criteria. (see [below for nested schema](#nestedatt--synciq_jobs--policy--file_matching_pattern))
- `name` (String) User-assigned name of this sync policy.
- `source_exclude_directories` (List of String) Directories that will be excluded from the sync.  Modifying this field will result in a full synchronization of all data.
- `source_include_directories` (List of String) Directories that will be included in the sync.  Modifying this field will result in a full synchronization of all data.
- `source_root_path` (String) The root directory on the source cluster the files will be synced from.  Modifying this field will result in a full synchronization of all data.
- `target_host` (String) Hostname or IP address of sync target cluster.  Modifying the target cluster host can result in the policy being unrunnable if the new target does not match the current target association.
- `target_path` (String) Absolute filesystem path on the target cluster for the sync destination.

<a id="nestedatt--synciq_jobs--policy--file_matching_pattern"></a>
### Nested Schema for `synciq_jobs.policy.file_matching_pattern`

Read-Only:

- `or_criteria` (Attributes List) An array containing objects with "and_criteria" properties, each set of and_criteria will be logically OR'ed together to create the full file matching pattern. (see [below for nested schema](#nestedatt--synciq_jobs--policy--file_matching_pattern--or_criteria))

<a id="nestedatt--synciq_jobs--policy--file_matching_pattern--or_criteria"></a>
### Nested Schema for `synciq_jobs.policy.file_matching_pattern.or_criteria`

Read-Only:

- `and_criteria` (Attributes List) An array containing individual file criterion objects each describing one criterion.  These are logically AND'ed together to form a set of criteria. (see [below for nested schema](#nestedatt--synciq_jobs--policy--file_matching_pattern--or_criteria--and_criteria))

<a id="nestedatt--synciq_jobs--policy--file_matching_pattern--or_criteria--and_criteria"></a>
### Nested Schema for `synciq_jobs.policy.file_matching_pattern.or_criteria.and_criteria`

Read-Only:

- `attribute_exists` (Boolean) For "custom_attribute" type criteria.  The file will match as long as the attribute named by "field" exists.  Default is true.
- `case_sensitive` (Boolean) If true, the value comparison will be case sensitive.  Default is true.
- `field` (String) The name of the file attribute to match on (only required if this is a custom_attribute type criterion).  Default is an empty string "".
- `operator` (String) How to compare the specified attribute of each file to the specified value.
- `type` (String) The type of this criterion, that is, which file attribute to match on.
- `value` (String) The value to compare the specified attribute of each file to.
- `whole_word` (Boolean) If true, the attribute must match the entire word.  Default is true.





<a id="nestedatt--synciq_jobs--service_report"></a>
### Nested Schema for `synciq_jobs.service_report`

Read-Only:

- `component` (String) The component that was processed.
- `directory` (String) The directory of the service export.
- `end_time` (Number) The time the job ended this component.
- `error_msg` (List of String) A list of error messages generated while exporting components.
- `filter` (List of String) A list of path-based filters for exporting components.
- `handlers_failed` (Number) The number of handlers failed during export.
- `handlers_skipped` (Number) The number of handlers skipped during export.
- `handlers_transferred` (Number) The number of handlers exported.
- `records_failed` (Number) The number of records failed during export.
- `records_skipped` (Number) The number of records skipped during export.
- `records_transferred` (Number) The number of records exported.
- `start_time` (Number) The time the job began this component.
- `status` (String) The current status of export for this component.


<a id="nestedatt--synciq_jobs--workers"></a>
### Nested Schema for `synciq_jobs.workers`

Read-Only:

- `connected` (Boolean) Whether there is a connection between the source and target.
- `last_split` (Number) The last time a network split occurred.
- `last_work` (Number) The last time the worker performed work.
- `lin` (Number) The LIN being worked on.
- `lnn` (Number) The lnn the worker is assigned to run on.
- `process_id` (Number) The process ID of the worker.
- `source_host` (String) The source host for this worker.
- `target_host` (String) The target host for this worker.
- `worker_id` (Number) The ID of the worker.