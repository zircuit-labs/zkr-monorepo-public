terraform {
  required_providers {
    jetstream = {
      source  = "nats-io/jetstream"
      version = "0.0.35"
    }
  }
}

variable "NATS_ADDRESS" {
  type        = string
  description = "Nats server address."
}

provider "jetstream" {
  servers = var.NATS_ADDRESS
}

resource "jetstream_stream" "prooforchestrator" {
  name     = "prooforchestrator"
  subjects = ["prooforchestrator.>"]
  max_age  = 60 * 60 * 24 * 60 // 60 days
}

resource "jetstream_stream" "raw_l2blocks" {
  name     = "RAW_L2BLOCKS"
  subjects = ["l2_blocks.>"]
  max_age  = 60 * 60 * 24 * 15 // 15 days
}
