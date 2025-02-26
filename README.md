This repository describes how to generate Terraform provider files from an
OpenAPI specification, allowing for quick development of a Terraform provider.
See also [this workflow example](https://developer.hashicorp.com/terraform/plugin/code-generation/workflow-example)

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

## Terraform Plugin Framework

Via https://github.com/hashicorp/terraform-plugin-codegen-framework/

```
go install github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest

tfplugingen-framework generate all \
    --input defectdojo-tf-provider-v2.43.3.json \
    --output internal/provider

mkdir -p client/{resources,datasources}

tfplugingen-framework scaffold resource \
    --name products \
    --output-dir client/resources/ \
    --package resources

tfplugingen-framework scaffold resource \
    --name endpoints \
    --output-dir client/resources/ \
    --package resources

tfplugingen-framework scaffold data-source \
    --name products \
    --output-dir client/datasources/ \
    --package datasources

tfplugingen-framework scaffold data-source \
    --name endpoints \
    --output-dir client/datasources/ \
    --package datasources
```

(see `internal/provider` for the generated files)
