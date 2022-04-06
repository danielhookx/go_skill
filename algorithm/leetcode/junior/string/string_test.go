package string

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				s: []byte("hello"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			t.Logf("%s", tt.args.s)
		})
	}
}

func Test_prefixTable(t *testing.T) {
	s := "aabaabaaa"
	table := prefixTable(s)
	t.Log(table)
}

func Test_kmp(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				haystack: "aaa",
				needle:   "aaa",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
