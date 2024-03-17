package factory

// define
type JsonRuleConfigParse struct {
}

func (p *JsonRuleConfigParse) parse(content string) RuleConfig {
	return RuleConfig{}
}

type IRuleConfigParserFactory interface {
	createParser() IRuleConfigParse
}

type XmlRuleConfigParserFactory struct {
}

func (t *XmlRuleConfigParserFactory) createParser() IRuleConfigParse {
	return &XmlRuleConfigParse{}
}

type YamlRuleConfigParserFactory struct {
}

func (t *YamlRuleConfigParserFactory) createParser() IRuleConfigParse {
	return &YamlRuleConfigParse{}
}

type JsonRuleConfigParserFactory struct {
}

func (t *JsonRuleConfigParserFactory) createParser() IRuleConfigParse {
	return &JsonRuleConfigParse{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "xml":
		return &XmlRuleConfigParserFactory{}
	case "json":
		return &JsonRuleConfigParserFactory{}
	case "yaml":
		return &YamlRuleConfigParserFactory{}
	}
	return nil
}
