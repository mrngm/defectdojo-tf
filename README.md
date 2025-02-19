This repository describes how to generate a Terraform provider from an OpenAPI
specification.

We're focusing on the OpenAPI specification of Defect Dojo, version 2.43.3.

## NOTES

- The OpenAPI schema contains circular references, which is not well-supported.
  We ignore the `prefetch` attribute where necessary

## OpenAPI specification

In order to track individual changes between versions, pipe this through
`json_pp` or `jq --format -`

```
curl 'https://demo.defectdojo.org/api/v2/oa3/schema/?format=json' | jq -r '.' > defectdojo-v2.43.3.json
```

## Terraform Provider spec

Via https://github.com/hashicorp/terraform-plugin-codegen-openapi

```
go install github.com/hashicorp/terraform-plugin-codegen-openapi/cmd/tfplugingen-openapi@latest

tfplugingen-openapi generate \
  --config defectdojo-generator.yml \
  --output defectdojo-tf-provider-v2.43.3.json \
  defectdojo-v2.43.3.json
```
