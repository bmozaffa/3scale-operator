package template

type TemplateType string

const (
	RedisType            TemplateType = "redis"
	ApicastType          TemplateType = "apicast"
	WildcardRouterType   TemplateType = "wildcardrouter"
	BackendType          TemplateType = "backend"
	SystemType           TemplateType = "system"
	ZyncType             TemplateType = "zync"
	MySQLType            TemplateType = "mysql"
	AmpImagesType        TemplateType = "ampimages"
	MemcachedType        TemplateType = "memcached"
	S3Type               TemplateType = "s3"
	ProductizedType      TemplateType = "productized"
	EvaluationType       TemplateType = "evaluation"
	HighAvailabilityType TemplateType = "ha"

	AmpType       TemplateType = "amp-template"
	AmpS3Type     TemplateType = "amp-s3-template"
	AmpEvalType   TemplateType = "amp-eval-template"
	AmpEvalS3Type TemplateType = "amp-eval-s3-template"
	AmpHAType     TemplateType = "amp-ha-template"
)

func NewTemplate(templateName string, componentOptions []string) *templatev1.Template {
	var tpl *templatev1.Template

	switch TemplateType(componentName) {
	case RedisType:
		tpl = NewRedis(componentOptions)
	case ApicastType:
		tpl = NewApicast(componentOptions)
	case WildcardRouterType:
		tpl = NewWildcardRouter(componentOptions)
	case BackendType:
		tpl = NewBackend(componentOptions)
	case SystemType:
		tpl = NewSystem(componentOptions)
	case ZyncType:
		tpl = NewZync(componentOptions)
	case MySQLType:
		tpl = NewMysql(componentOptions)
	case AmpImagesType:
		tpl = NewAmpImages(componentOptions)
	case MemcachedType:
		tpl = NewMemcached(componentOptions)
	case S3Type:
		tpl = NewS3(componentOptions)
	case ProductizedType:
		tpl = NewProductized(componentOptions)
	case EvaluationType:
		tpl = NewEvaluation(componentOptions)
	case AmpType:
		tpl = NewAmpTemplate(componentOptions)
	case AmpS3Type:
		tpl = NewAmpS3Template(componentOptions)
	case AmpEvalType:
		tpl = NewAmpEvalTemplate(componentOptions)
	case AmpEvalS3Type:
		tpl = NewAmpEvalS3Template(componentOptions)
	case HighAvailabilityType:
		tpl = NewHighAvailability(componentOptions)
	case AmpHAType:
		tpl = NewAmpHATemplate(componentOptions)
	default:
		panic("Error: Template not recognized")
	}

	return template
}
