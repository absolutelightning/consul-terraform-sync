# This file is generated by Consul Terraform Sync.
#
# The HCL blocks, arguments, variables, and values are derived from the
# operator configuration for Sync. Any manual changes to this file
# may not be preserved and could be overwritten by a subsequent update.
#
# Task: test
# Description: user description for task named 'test'

# Service definition protocol v0
variable "services" {
  description = "Consul services monitored by Consul Terraform Sync"
  type = map(
    object({
      id        = string
      name      = string
      address   = string
      port      = number
      meta      = map(string)
      tags      = list(string)
      namespace = string
      status    = string

      node                  = string
      node_id               = string
      node_address          = string
      node_datacenter       = string
      node_tagged_addresses = map(string)
      node_meta             = map(string)

      cts_user_defined_meta = map(string)
    })
  )
}

variable "testProvider" {
  default     = null
  description = "Configuration object for testProvider"
  type = object({
    alias = string
    attr  = string
    count = number
    obj = object({
      id       = string
      username = string
    })
  })
}
