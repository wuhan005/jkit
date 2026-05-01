package format

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "simple object",
			input: `{"a":1,"b":"x"}`,
			want: `{
    "a": 1,
    "b": "x"
}`,
		},
		{
			name:  "preserve large integer precision",
			input: `{"id":9223372036854775807}`,
			want: `{
    "id": 9223372036854775807
}`,
		},
		{
			name:  "nested array",
			input: `[1,[2,[3]]]`,
			want: `[
    1,
    [
        2,
        [
            3
        ]
    ]
]`,
		},
		{
			name:    "invalid json",
			input:   `{not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Fatalf("Format() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
				t.Errorf("Format() got:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}
