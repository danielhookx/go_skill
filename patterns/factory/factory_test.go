package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		{
			name: "xml",
			args: args{
				t: "xml",
			},
			want: &XmlRuleConfigParserFactory{},
		}, {
			name: "yaml",
			args: args{
				t: "yaml",
			},
			want: &YamlRuleConfigParserFactory{},
		}, {
			name: "json",
			args: args{
				t: "json",
			},
			want: &JsonRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %T, want %T", got, tt.want)
			}
		})
	}
}
