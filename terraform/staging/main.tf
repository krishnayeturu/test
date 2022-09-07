terraform {
  backend "s3" {
    bucket  = "2w-pe-terraform"
    key     = "platform-automation/ms-wrkr-admissions/staging/terraform.tfstate"
    region  = "us-west-2"
    encrypt = true
    acl     = "bucket-owner-full-control"
  }
}

# region providers
provider "aws" {
  alias               = "main"
  region              = local.main_region
  profile             = "platform_automation_development"
  allowed_account_ids = ["539365964353"]
}
# endregion providers

data "aws_caller_identity" "current" {
  provider = aws.main
}


locals {
  service_name_base = "ms-wrkr-admissions"
  main_region       = "us-west-2"
  replica_region    = "us-west-2"
  env               = "staging"
  tags_as_map = {
    Owner            = "DataPlatform"
    Stage            = "staging"
    Repository       = "https://gitlab.com/2ndwatch/microservices/workers/ms-wrkr-admissions/tree/master/terraform/staging"
    "2W_Environment" = "staging"
    "2W_Workload"    = "ms-wrkr-admissions"
  }
  oidc_principal_providers = [
    "arn:aws:iam::061165946885:oidc-provider/oidc.eks.us-west-2.amazonaws.com/id/62E18F7F4B5A53189D7DE1E83EB3148B", #dev-eng-eks
    "arn:aws:iam::061165946885:oidc-provider/oidc.eks.us-west-2.amazonaws.com/id/314643C241F99EDD32E923193EB60033", #dev-eng-eks2
  ]
}

# region DynamoDB
resource "aws_dynamodb_table" "permissions_table" {
  provider     = aws.main
  name         = format("%s-%s-dynamodb-permissions-table", "${local.env}", "${local.service_name_base}")
  billing_mode = "PAY_PER_REQUEST"
  /* read_capacity  = 3 # initial estimate for test case, changeable as we determine necessary
  write_capacity = 2 # initial estimate for test case, changeable as we determine necessary */
  hash_key  = "ID"
  range_key = "PermissionType"

  point_in_time_recovery {
    enabled = false # defaults to false, will want to consider setting to true for production
  }

  server_side_encryption {
    enabled = false # if false server-side encryption is set to AWS owned CMK (shown as DEFAULT in the AWS console), if true and no kms_key_arn is specified then server-side encryption is set to AWS managed CMK (shown as KMS in the AWS console)
  }

  attribute {
    name = "ID"
    type = "S"
  }

  attribute {
    name = "PermissionEffect"
    type = "N"
  }

  attribute {
    name = "PermissionType"
    type = "N"
  }

  timeouts {
    create = "5m"
    update = "7m"
    delete = "5m"
  }

  global_secondary_index {
    name            = "PermissionTypePermissionEffect"
    hash_key        = "PermissionEffect"
    range_key       = "PermissionType"
    write_capacity  = 3
    read_capacity   = 2
    projection_type = "ALL"
  }

  tags = merge(local.tags_as_map, { "Name" = format("%s-%s-dynamodb-permissions-table", "${local.env}", "${local.service_name_base}") })
}

resource "aws_dynamodb_table" "principal_permissions_table" {
  provider     = aws.main
  name         = format("%s-%s-dynamodb-principal-permissions-table", "${local.env}", "${local.service_name_base}")
  billing_mode = "PAY_PER_REQUEST"
  /* read_capacity  = 3 # initial estimate for test case, changeable as we determine necessary
  write_capacity = 2 # initial estimate for test case, changeable as we determine necessary */
  hash_key = "ID"

  point_in_time_recovery {
    enabled = false # defaults to false, will want to consider setting to true for production
  }

  server_side_encryption {
    enabled = false # if false server-side encryption is set to AWS owned CMK (shown as DEFAULT in the AWS console), if true and no kms_key_arn is specified then server-side encryption is set to AWS managed CMK (shown as KMS in the AWS console)
  }

  attribute {
    name = "ID"
    type = "S"
  }

  timeouts {
    create = "5m"
    update = "7m"
    delete = "5m"
  }

  tags = merge(local.tags_as_map, { "Name" = format("%s-%s-dynamodb-principal-permissions-table", "${local.env}", "${local.service_name_base}") })
}
# endregion DynamoDB

# region IAM
module "dynamodb_access_readonly_role" {
  source = "github.com/2ndWatch/pe-terraform-aws-iam.git?ref=v8.1.0"
  providers = {
    aws = aws.main
  }

  name_prefix = local.env
  role_name   = format("%s-%s-dynamodb-access-ro", "${local.env}", "${local.service_name_base}")
  role_assume_policy_configs = [
    {
      principal_type  = "Federated"
      principal_value = local.oidc_principal_providers
      action          = "sts:AssumeRoleWithWebIdentity"
    }
  ]
  tags = merge(local.tags_as_map, { "Name" = format("%s-%s-dynamodb-access-ro", "${local.env}", "${local.service_name_base}") })

  role_inline_policy_statement = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
              "dynamodb:BatchGetItem",
              "dynamodb:Describe*",
              "dynamodb:List*",
              "dynamodb:GetItem",
              "dynamodb:Query",
              "dynamodb:Scan",
              "dynamodb:PartiQLSelect"
            ],
            "Resource": [
              "${aws_dynamodb_table.principal_permissions_table.arn}",
              "${aws_dynamodb_table.permissions_table.arn}"
            ]
        }
    ]
}
EOF
}

module "dynamodb_access_readwritedelete_role" {
  source = "github.com/2ndWatch/pe-terraform-aws-iam.git?ref=v8.1.0"
  providers = {
    aws = aws.main
  }

  name_prefix = local.env
  role_name   = format("%s-%s-dynamodb-access-rwd", "${local.env}", "${local.service_name_base}")
  role_assume_policy_configs = [
    {
      principal_type  = "Federated"
      principal_value = local.oidc_principal_providers
      action          = "sts:AssumeRoleWithWebIdentity"
    }
  ]
  tags = merge(local.tags_as_map, { "Name" = format("%s-%s-dynamodb-access-rwd", "${local.env}", "${local.service_name_base}") })

  role_inline_policy_statement = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
              "dynamodb:BatchGetItem",
              "dynamodb:BatchWriteItem",
              "dynamodb:Describe*",
              "dynamodb:DeleteItem",
              "dynamodb:List*",
              "dynamodb:GetItem",
              "dynamodb:PartiQLSelect",
              "dynamodb:PutItem",
              "dynamodb:Query",
              "dynamodb:Scan",
              "dynamodb:UpdateItem"
            ],
            "Resource": [
              "${aws_dynamodb_table.principal_permissions_table.arn}",
              "${aws_dynamodb_table.permissions_table.arn}"
            ]
        }
    ]
}
EOF
}
# endregion IAM