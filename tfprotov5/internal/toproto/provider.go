package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func GetProviderSchema_Request(in tfprotov5.GetProviderSchemaRequest) (tfplugin5.GetProviderSchema_Request, error) {
	return tfplugin5.GetProviderSchema_Request{}, nil
}

func GetProviderSchema_Response(in tfprotov5.GetProviderSchemaResponse) (tfplugin5.GetProviderSchema_Response, error) {
	var resp tfplugin5.GetProviderSchema_Response
	if in.Provider != nil {
		schema := Schema(*in.Provider)
		resp.Provider = &schema
	}
	if in.ProviderMeta != nil {
		schema := Schema(*in.ProviderMeta)
		resp.ProviderMeta = &schema
	}
	resp.ResourceSchemas = make(map[string]*tfplugin5.Schema, len(in.ResourceSchemas))
	for k, v := range in.ResourceSchemas {
		if v == nil {
			resp.ResourceSchemas[k] = nil
			continue
		}
		schema := Schema(*v)
		resp.ResourceSchemas[k] = &schema
	}
	resp.DataSourceSchemas = make(map[string]*tfplugin5.Schema, len(in.DataSourceSchemas))
	for k, v := range in.DataSourceSchemas {
		if v == nil {
			resp.DataSourceSchemas[k] = nil
			continue
		}
		schema := Schema(*v)
		resp.DataSourceSchemas[k] = &schema
	}
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func PrepareProviderConfig_Request(in tfprotov5.PrepareProviderConfigRequest) (tfplugin5.PrepareProviderConfig_Request, error) {
	resp := tfplugin5.PrepareProviderConfig_Request{}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	return resp, nil
}

func PrepareProviderConfig_Response(in tfprotov5.PrepareProviderConfigResponse) (tfplugin5.PrepareProviderConfig_Response, error) {
	var resp tfplugin5.PrepareProviderConfig_Response
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.PreparedConfig != nil {
		config, err := Cty(*in.PreparedConfig)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.PreparedConfig = &config
	}
	return resp, nil
}

func Configure_Request(in tfprotov5.ConfigureProviderRequest) (tfplugin5.Configure_Request, error) {
	resp := tfplugin5.Configure_Request{
		TerraformVersion: in.TerraformVersion,
	}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	return resp, nil
}

func Configure_Response(in tfprotov5.ConfigureProviderResponse) (tfplugin5.Configure_Response, error) {
	var resp tfplugin5.Configure_Response
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func Stop_Request(in tfprotov5.StopProviderRequest) (tfplugin5.Stop_Request, error) {
	return tfplugin5.Stop_Request{}, nil
}

func Stop_Response(in tfprotov5.StopProviderResponse) (tfplugin5.Stop_Response, error) {
	return tfplugin5.Stop_Response{
		Error: in.Error,
	}, nil
}