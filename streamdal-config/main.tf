terraform {
  required_providers {
    streamdal = {
      version = "0.1.2"
      source  = "streamdal/streamdal"
    }
  }
}

provider "streamdal" {
  token              = "1234"
  address            = "localhost:8082"
  connection_timeout = 10
}

resource "streamdal_pipeline" "mask_email" {
  name = "Mask Email"

  step {
    name = "Detect Email Field"

    # We specify abort conditions here since we don't want
    # to continue with the second step if there is nothing
    # to transform.
    on_false {
      abort = "abort_current" # No email found
    }
    on_error {
      abort = "abort_current" # An error occurred
    }
    dynamic = false
    detective {
      type   = "pii_email"
      args   = [] # no args for this type
      negate = false
      path   = "" # No path, we will scan the entire payload
    }
  }

  step {
    name    = "Mask Email Step"
    dynamic = true
    transform {
      mask_value {
        # No path needed since dynamic=true
        # We will use the results from the first detective step
        path = ""

        # Mask the email field(s) we find with asterisks
        mask = "*"
      }
    }
  }
}

resource "streamdal_audience" "billing_sales_report" {
  service_name   = "billing-svc"
  component_name = "kafka"
  operation_name = "gen-sales-report"
  operation_type = "consumer"
  pipeline_ids   = [resource.streamdal_pipeline.mask_email.id]
}