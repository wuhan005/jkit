package get

import "testing"

func TestGet(t *testing.T) {
	const blob = `{
		"data": {
			"items": [
				{"id": 1, "name": "alpha"},
				{"id": 2, "name": "beta"}
			],
			"count": 2,
			"active": true,
			"meta": null
		}
	}`

	tests := []struct {
		name string
		path string
		want string
	}{
		{name: "string leaf returns raw", path: "data.items.0.name", want: "alpha"},
		{name: "number leaf", path: "data.count", want: "2"},
		{name: "bool leaf", path: "data.active", want: "true"},
		{name: "null leaf", path: "data.meta", want: "null"},
		{
			name: "object subtree",
			path: "data.items.0",
			want: `{
    "id": 1,
    "name": "alpha"
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get([]byte(blob), tt.path)
			if err != nil {
				t.Fatalf("Get(%q) unexpected error: %v", tt.path, err)
			}
			if got != tt.want {
				t.Errorf("Get(%q) got:\n%s\nwant:\n%s", tt.path, got, tt.want)
			}
		})
	}
}

func TestGet_errors(t *testing.T) {
	const blob = `{"items":[1,2,3],"name":"x"}`

	tests := []struct {
		name string
		path string
	}{
		{name: "missing key", path: "missing"},
		{name: "non-integer index", path: "items.abc"},
		{name: "out of range", path: "items.99"},
		{name: "descend into scalar", path: "name.foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Get([]byte(blob), tt.path); err == nil {
				t.Errorf("Get(%q) expected error, got nil", tt.path)
			}
		})
	}
}
