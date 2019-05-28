package template

import (
	templatev1 "github.com/openshift/api/template/v1"
)

func NewAmpTemplate(options []string) *templatev1.Template {
	//components := []Component{
	//	NewMemcached(options),
	//	NewSystem(options),
	//	NewZync(options),
	//	NewWildcardRouter(options),
	//}

	adapterList := []adapters.Adapter{
		adapters.NewImagesAdapter(options),
		adapters.NewRedisAdapter(options),
		adapters.NewApicastAdapter(options),
		adapters.NewBackendAdapter(options),
		adapters.NewMysqlAdapter(options),
	}

	return Generator(adapterList)
}

func (ampTemplate *AmpTemplate) setTemplateFields(template *templatev1.Template) {
	template.ObjectMeta.Name = "3scale-api-management"
	template.ObjectMeta.Annotations = ampTemplate.buildAmpTemplateMetaAnnotations()
	template.Message = "Login on https://${TENANT_NAME}-admin.${WILDCARD_DOMAIN} as ${ADMIN_USERNAME}/${ADMIN_PASSWORD}"
}
