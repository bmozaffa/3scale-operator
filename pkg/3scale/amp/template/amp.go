package template

import (
	templatev1 "github.com/openshift/api/template/v1"
)

func NewAmpTemplate(options []string) *templatev1.Template {
	//components := []Component{
	//	NewAmpImages(options),
	//	NewRedis(options),
	//	NewBackend(options),
	//	NewMysql(options),
	//	NewMemcached(options),
	//	NewSystem(options),
	//	NewZync(options),
	//	NewApicast(options),
	//	NewWildcardRouter(options),
	//}

	adapterList := []adapters.Adapter{
		adapters.NewImagesAdapter(options),
		adapters.NewRedisAdapter(options),
		adapters.NewApicastAdapter(options),
		adapters.NewBackendAdapter(options),
	}

	return Generator(adapterList)
}

func (ampTemplate *AmpTemplate) AssembleIntoTemplate(template *templatev1.Template, otherComponents []Component) {
	for componentIdx := range ampTemplate.components {
		ampTemplate.components[componentIdx].AssembleIntoTemplate(template, otherComponents)
	}
	ampTemplate.setTemplateFields(template)
}

func (ampTemplate *AmpTemplate) PostProcess(template *templatev1.Template, otherComponents []Component) {
	for componentIdx := range ampTemplate.components {
		ampTemplate.components[componentIdx].PostProcess(template, otherComponents)
	}
}

func (ampTemplate *AmpTemplate) setTemplateFields(template *templatev1.Template) {
	template.ObjectMeta.Name = "3scale-api-management"
	template.ObjectMeta.Annotations = ampTemplate.buildAmpTemplateMetaAnnotations()
	template.Message = "Login on https://${TENANT_NAME}-admin.${WILDCARD_DOMAIN} as ${ADMIN_USERNAME}/${ADMIN_PASSWORD}"
}

func (ampTemplate *AmpTemplate) buildAmpTemplateMetaAnnotations() map[string]string {
	annotations := map[string]string{
		"openshift.io/display-name":          "3scale API Management",
		"openshift.io/provider-display-name": "Red Hat, Inc.",
		"iconClass":                          "icon-3scale",
		"description":                        "3scale API Management main system",
		"tags":                               "integration, api management, 3scale",
	}

	return annotations
}
