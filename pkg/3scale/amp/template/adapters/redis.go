package adapters

import (
	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	templatev1 "github.com/openshift/api/template/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Redis struct {
}

func NewRedis(options []string) Adapter {
	return NewAppenderAdapter(&Images{})
}

func (a *Redis) Parameters() []templatev1.Parameter {
	return []templatev1.Parameter{
		{
			Name:        "REDIS_IMAGE",
			Description: "Redis image to use",
			Required:    true,
			Value:       "registry.access.redhat.com/rhscl/redis-32-rhel7:3.2",
		},
	}
}

func (a *Images) Objects() ([]runtime.RawExtension, error) {
	redisOptions, err := options()
	if err != nil {
		return nil, err
	}
	redisComponent := component.NewRedis(redisOptions)
	return redisComponent.Objects(), nil
}

func (a *Images) options() (*ApicastOptions, error) {
	rob := RedisOptionsBuilder{}
	rob.AppLabel("${APP_LABEL}")

	return rob.Build()
}
