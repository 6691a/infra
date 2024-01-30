locals {
    communication = data.terraform_remote_state.communication.outputs.communication
    smtp_role = data.terraform_remote_state.communication.outputs.smtp_role
}