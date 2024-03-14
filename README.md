# Terraform blog post demo code

This repository contains the code for the blog post demoing the Streamdal terraform module: https://medium.com/streamdal/configuring-streamdal-with-terraform-d07e6cf9dfe3

## Usage

1. Ensure Terraform is installed
2. Ensure the Streamdal server is running locally :`curl -sSL https://sh.streamdal.com | bash`
3. cd `streamdal-config`
4. Run `terraform init && terraform apply`
5. Navigate to `http://localhost:8080` in your browser to see the Streamdal UI and that the audience and pipeline are configured
6. cd `../demo-app`
7. Run the demo app: `go run main.go`
8. The resulting output should contain the masked email in the payload:
```json
{
  "customer": {
    "email": "john******************",
    "first_name": "John",
    "last_name": "Doe"
  }
}
```
