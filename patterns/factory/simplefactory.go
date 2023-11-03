package factory

type IRuleConfigParse interface {
	parse(content string) RuleConfig
}

type RuleConfig struct {
}

type XmlRuleConfigParse struct {
}

func NewXmlRuleConfigParse() *XmlRuleConfigParse {
	return &XmlRuleConfigParse{}
}

func (p *XmlRuleConfigParse) parse(content string) RuleConfig {
	//TODO
	return RuleConfig{}
}

type YamlRuleConfigParse struct {
}

func NewYamlRuleConfigParse() *YamlRuleConfigParse {
	return &YamlRuleConfigParse{}
}

func (p *YamlRuleConfigParse) parse(content string) RuleConfig {
	//TODO
	return RuleConfig{}
}

// Factory
func CreateRuleConfigParse(configFormat string) IRuleConfigParse {
	var parser IRuleConfigParse
	switch configFormat {
	case "xml":
		parser = NewXmlRuleConfigParse()
	case "yaml":
		parser = NewYamlRuleConfigParse()
	}
	return parser
}
