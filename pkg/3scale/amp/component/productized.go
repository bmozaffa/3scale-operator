package component

import (
	"fmt"

	imagev1 "github.com/openshift/api/image/v1"
	templatev1 "github.com/openshift/api/template/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Productized struct {
	options []string
	Options *ProductizedOptions
}

func NewProductized(options []string) *Productized {
	productized := &Productized{
		options: options,
	}
	return productized
}

type ProductizedOptions struct {
	productizedNonRequiredOptions
	productizedRequiredOptions
}

type productizedRequiredOptions struct {
	ampRelease   string
	apicastImage string
	backendImage string
	routerImage  string
	systemImage  string
	zyncImage    string
}

type productizedNonRequiredOptions struct {
}

type ProductizedOptionsProvider interface {
	GetProductizedOptions() *ProductizedOptions
}
type CLIProductizedOptionsProvider struct {
}

func (o *CLIProductizedOptionsProvider) GetProductizedOptions() (*ProductizedOptions, error) {
	pob := ProductizedOptionsBuilder{}
	pob.ApicastImage("${AMP_APICAST_IMAGE}")
	pob.BackendImage("${AMP_BACKEND_IMAGE}")
	pob.RouterImage("${AMP_ROUTER_IMAGE}")
	pob.SystemImage("${AMP_SYSTEM_IMAGE}")
	pob.ZyncImage("${AMP_ZYNC_IMAGE}")
	res, err := pob.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Productized Options - %s", err)
	}
	return res, nil
}

func (productized *Productized) AssembleIntoTemplate(template *templatev1.Template, otherComponents []Component) {
}

func (productized *Productized) PostProcess(template *templatev1.Template, otherComponents []Component) {
	// TODO move this outside this specific method
	optionsProvider := CLIProductizedOptionsProvider{}
	productizedOpts, err := optionsProvider.GetProductizedOptions()
	_ = err
	productized.Options = productizedOpts
	res := template.Objects
	productized.updateAmpImagesParameters(template)
	template.Objects = res
}

func (productized *Productized) PostProcessObjects(objects []runtime.RawExtension) []runtime.RawExtension {
	res := objects
	res = productized.updateAmpImagesURIs(res)

	return res
}

func (productized *Productized) updateAmpImagesParameters(template *templatev1.Template) {
	for paramIdx := range template.Parameters {
		param := &template.Parameters[paramIdx]
		switch param.Name {
		case "AMP_SYSTEM_IMAGE":
			param.Value = "registry.access.redhat.com/3scale-amp25/system"
		case "AMP_BACKEND_IMAGE":
			param.Value = "registry.access.redhat.com/3scale-amp25/backend"
		case "AMP_APICAST_IMAGE":
			param.Value = "registry.access.redhat.com/3scale-amp25/apicast-gateway"
		case "AMP_ROUTER_IMAGE":
			param.Value = "registry.access.redhat.com/3scale-amp22/wildcard-router"
		case "AMP_ZYNC_IMAGE":
			param.Value = "registry.access.redhat.com/3scale-amp25/zync"
		}
	}
}

func (productized *Productized) updateAmpImagesURIs(objects []runtime.RawExtension) []runtime.RawExtension {
	res := objects

	for _, rawExtension := range res {
		obj := rawExtension.Object
		is, ok := obj.(*imagev1.ImageStream)
		if ok {
			for tagIdx := range is.Spec.Tags {
				// Only change the ImageStream tag name that has the ampRelease
				// value. We do not modify the latest tag
				if is.Spec.Tags[tagIdx].Name == productized.Options.ampRelease {
					switch is.Name {
					case "amp-apicast":
						is.Spec.Tags[tagIdx].From.Name = productized.Options.apicastImage
					case "amp-system":
						is.Spec.Tags[tagIdx].From.Name = productized.Options.systemImage
					case "amp-backend":
						is.Spec.Tags[tagIdx].From.Name = productized.Options.backendImage
					case "amp-wildcard-router":
						is.Spec.Tags[tagIdx].From.Name = productized.Options.routerImage
					case "amp-zync":
						is.Spec.Tags[tagIdx].From.Name = productized.Options.zyncImage
					}
				}
			}
		}
	}

	return res
}
