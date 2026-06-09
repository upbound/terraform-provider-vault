package xpprovider

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	tfsdkschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-vault/internal/provider/fwprovider"
)

func FrameworkProvider(sdkProvider *tfsdkschema.Provider) provider.Provider {
	return fwprovider.New(sdkProvider)
}
