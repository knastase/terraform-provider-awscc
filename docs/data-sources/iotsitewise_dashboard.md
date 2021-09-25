---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iotsitewise_dashboard Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::IoTSiteWise::Dashboard
---

# awscc_iotsitewise_dashboard (Data Source)

Data Source schema for AWS::IoTSiteWise::Dashboard



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **dashboard_arn** (String) The ARN of the dashboard.
- **dashboard_definition** (String) The dashboard definition specified in a JSON literal.
- **dashboard_description** (String) A description for the dashboard.
- **dashboard_id** (String) The ID of the dashboard.
- **dashboard_name** (String) A friendly name for the dashboard.
- **project_id** (String) The ID of the project in which to create the dashboard.
- **tags** (Attributes List) A list of key-value pairs that contain metadata for the dashboard. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

