package azure

import (
	"sync"

	"github.com/Azure/azure-sdk-for-go/arm/resources"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/mitchellh/packer/common/json"
)

// Config is the configuration structure used to instantiate a
// new Azure management client.
type Config struct {
	Settings       []byte
	ClientID       string
	ClientSecret   string
	SubscriptionID string
	TenantID       string
}

// Client contains all the handles required for managing Azure services.
type Client struct {
	servicePrincipalToken *azure.ServicePrincipalToken
	resourceGroupsClient  resources.ResourceGroupsClient
	mutex                 *sync.Mutex
}

func (c *Config) NewClientFromSettingsData() (*Client, error) {

	config := Config{}
	err := json.Unmarshal(c.Settings,&config)

	if(err == nil){
		return	config.NewClient()
	}

	return nil, err
}

func (c *Config) NewClient() (*Client, error) {

	token, err := azure.NewServicePrincipalToken(c.ClientSecret,c.ClientSecret,c.TenantID,azure.AzureResourceManagerScope)
	if(err != nil) {
		return &Client{
			servicePrincipalToken: token,
			resourceGroupsClient:  resources.NewResourceGroupsClient(c.SubscriptionID),
			mutex:                 &sync.Mutex{},
		}, nil
	} else {
		return nil,err
	}
}
