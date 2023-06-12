# Table: snyk_projects

This table shows data for Snyk Projects.

https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|origin|`utf8`|
|issue_counts_by_severity|`json`|
|tags|`json`|
|org_id|`utf8`|