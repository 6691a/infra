locals {
    resource_group_name = data.terraform_remote_state.resource_group.outputs.playhub.name
}
