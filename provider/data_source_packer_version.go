package provider

import (
	"context"
	"os"
	"strings"

	"terraform-provider-packer/packer_interop"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/toowoxx/go-lib-userspace-common/cmds"
)

type dataSourceVersionType struct {
	Version string `tfsdk:"version"`
}

func (r dataSourceVersionType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"version": {
				Description: "Version of embedded Packer",
				Computed:    true,
				Type:        types.StringType,
			},
		},
	}, nil
}

func (r dataSourceVersionType) NewDataSource(_ context.Context, p provider.Provider) (datasource.DataSource, diag.Diagnostics) {
	return dataSourceVersion{
		p: *(p.(*tfProvider)),
	}, nil
}

type dataSourceVersion struct {
	p tfProvider
}

func (r dataSourceVersion) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	resourceState := dataSourceVersionType{}
	exe, _ := os.Executable()
	output, err := cmds.RunCommandWithEnvReturnOutput(
		exe,
		map[string]string{packer_interop.TPPRunPacker: "true"},
		"version")
	if err != nil {
		resp.Diagnostics.AddError("Failed to run packer", err.Error())
		return
	}

	if len(output) == 0 {
		resp.Diagnostics.AddError("Unexpected output", "Packer did not output anything")
		return
	}

	resourceState.Version = strings.TrimPrefix(
		strings.TrimSpace(strings.TrimPrefix(string(output), "Packer")), "v")

	diags := resp.State.Set(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
