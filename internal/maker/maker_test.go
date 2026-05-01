package maker

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		unique bool
		want   string
	}{
		{
			name:  "basic split, blank lines dropped, trims whitespace",
			input: "a\n  b  \n\nc\n",
			want: `[
    "a",
    "b",
    "c"
]`,
		},
		{
			name:  "windows line endings",
			input: "a\r\nb\r\nc",
			want: `[
    "a",
    "b",
    "c"
]`,
		},
		{
			name:   "unique drops duplicates",
			input:  "a\nb\na\nc\nb",
			unique: true,
			want: `[
    "a",
    "b",
    "c"
]`,
		},
		{
			name:   "unique false keeps duplicates",
			input:  "a\na",
			unique: false,
			want: `[
    "a",
    "a"
]`,
		},
		{
			name:  "all blank yields empty array",
			input: "\n\n   \n",
			want:  `[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Make(tt.input, tt.unique)
			if err != nil {
				t.Fatalf("Make() unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Make() got:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}
