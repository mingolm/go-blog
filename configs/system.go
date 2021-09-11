package configs

var SystemConfig = struct {
	TemplateHtmlPrefix   string `env:"TEMPLATE_HTML_PREFIX" flag:"template-html-prefix" flagUsage:"http template prefix"`
	TemplateHtmlCommon   string `env:"TEMPLATE_HTML_COMMON" flag:"template-html-common" flagUsage:"http template common"`
	TemplateNotFoundPage string `env:"TEMPLATE_NOT_FOUND_PAGE" flag:"template-not-found-page" flagUsage:"http template for 404"`
	TemplateBladeType    string `env:"TEMPLATE_BLADE_TYPE" flag:"template-blade-type" flagUsage:"http template type"`
}{
	TemplateHtmlPrefix:   "resources/html/",
	TemplateHtmlCommon:   "resources/html/common/",
	TemplateNotFoundPage: "resources/html/404.html",
	TemplateBladeType:    "html",
}
