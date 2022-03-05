package factory

import (
	"reflect"
	"testing"
)

func TestCreateRuleConfigParse(t *testing.T) {
	type args struct {
		configFormat string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParse
	}{
		{
			name: "xml",
			args: args{
				configFormat: "xml",
			},
			want: NewXmlRuleConfigParse(),
		},{
			name: "yaml",
			args: args{
				configFormat: "yaml",
			},
			want: NewYamlRuleConfigParse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRuleConfigParse(tt.args.configFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleConfigParse() = %T, want %T", got, tt.want)
			}
		})
	}
}
