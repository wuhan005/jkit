package cut

import (
	"math"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxDepth int
		want     string
	}{
		{
			name:     "depth 0 collapses root object",
			input:    `{"a":1,"b":2}`,
			maxDepth: 0,
			want:     `{ 2 items dict }`,
		},
		{
			name:     "depth 0 collapses root array",
			input:    `[1,2,3]`,
			maxDepth: 0,
			want:     `[ 3 items array ]`,
		},
		{
			name:     "object keys are sorted",
			input:    `{"b":1,"a":2,"c":3}`,
			maxDepth: math.MaxInt,
			want: `{
    "a": 2,
    "b": 1,
    "c": 3
}`,
		},
		{
			name:     "depth 1 keeps top object, folds nested",
			input:    `{"x":{"y":{"z":1}}}`,
			maxDepth: 1,
			want: `{
    "x": { 1 items dict }
}`,
		},
		{
			name:     "preserve integer precision",
			input:    `{"id":9223372036854775807}`,
			maxDepth: math.MaxInt,
			want: `{
    "id": 9223372036854775807
}`,
		},
		{
			name:     "string with quote and newline is escaped",
			input:    `{"s":"a\"b\nc"}`,
			maxDepth: math.MaxInt,
			want: `{
    "s": "a\"b\nc"
}`,
		},
		{
			name:     "null lowercase",
			input:    `{"n":null}`,
			maxDepth: math.MaxInt,
			want: `{
    "n": null
}`,
		},
		{
			name:     "boolean",
			input:    `[true,false]`,
			maxDepth: math.MaxInt,
			want: `[
    true,
    false
]`,
		},
		{
			name:     "empty object stays compact",
			input:    `{"a":{}}`,
			maxDepth: math.MaxInt,
			want: `{
    "a": {}
}`,
		},
		{
			name:     "empty array stays compact",
			input:    `{"a":[]}`,
			maxDepth: math.MaxInt,
			want: `{
    "a": []
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cut([]byte(tt.input), tt.maxDepth)
			if err != nil {
				t.Fatalf("Cut() unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Cut() got:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}

func TestCut_invalidJSON(t *testing.T) {
	if _, err := Cut([]byte("{not json"), 1); err == nil {
		t.Fatal("expected error, got nil")
	}
}
