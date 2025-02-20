package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*defectdojoProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &defectdojoProvider{}
	}
}

type defectdojoProvider struct{}

func (p *defectdojoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

}

func (p *defectdojoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *defectdojoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "defectdojo"
}

func (p *defectdojoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *defectdojoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
