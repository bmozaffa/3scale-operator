package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
)

func (o *OperatorProductizedOptionsProvider) GetProductizedOptions() (*component.ProductizedOptions, error) {
	pob := component.ProductizedOptionsBuilder{}
	pob.AmpRelease(o.AmpSpec.AmpRelease)
	pob.ApicastImage("registry.access.redhat.com/3scale-amp24/apicast-gateway")
	pob.BackendImage("registry.access.redhat.com/3scale-amp24/backend")
	pob.RouterImage("registry.access.redhat.com/3scale-amp22/wildcard-router")
	pob.SystemImage("registry.access.redhat.com/3scale-amp24/system")
	pob.ZyncImage("registry.access.redhat.com/3scale-amp24/zync")
	pob.MemcachedImage("registry.access.redhat.com/3scale-amp20/memcached")
	res, err := pob.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Productized Options - %s", err)
	}
	return res, nil
}