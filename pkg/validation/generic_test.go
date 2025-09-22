package validation

import "testing"

func TestGeneric(t *testing.T) {
	type Struct struct {
		value string
		fn    string
		want  bool
	}

	tests := []Struct{
		{"my-name", "IsQualifiedName", true},
		{"#", "IsQualifiedName", false},
		{"123", "IsValidLabel", true},
		{"my-name", "IsValidLabelDNS1123", true},
		{"example.com", "IsValidsubdomainDNS1123", true},
		{"9.1", "IsValidIP", false},
		{"Wto1260644864!", "IsValidPassword", true},
	}

	for _, tt := range tests {
		switch tt.fn {
		case "IsQualifiedName":
			if (len(IsQualifiedName(tt.value)) == 0) != tt.want {
				t.Errorf("IsQualifiedName(string) []string has an error: want: %v got: %v\n", tt.want, !tt.want)
			}
		case "IsValidLabel":
			if (len(IsValidLabel(tt.value)) == 0) != tt.want {
				t.Errorf("IsValidLabel(string) []string has an error: want: %v got: %v\n", tt.want, !tt.want)
			}
		case "IsValidLabelDNS1123":
			if (len(IsValidLabelDNS1123(tt.value)) == 0) != tt.want {
				t.Errorf("IsValidLabelDNS1123(string) []string has an error: want: %v got: %v\n", tt.want, !tt.want)
			}
		case "IsValidsubdomainDNS1123":
			if (len(IsValidsubdomainDNS1123(tt.value)) == 0) != tt.want {
				t.Errorf("IsValidsubdomainDNS1123(string) []string has an error: want: %v got: %v\n", tt.want, !tt.want)
			}
		case "IsValidIP":
			if (len(IsValidIP(tt.value)) == 0) != tt.want {
				t.Errorf("IsValidIP(string) []string has an error: want: %v got: %v\n", tt.want, !tt.want)
			}
		case "IsValidPassword":
			if (len(IsValidPassword(tt.value)) == 0) != tt.want {
				t.Errorf("IsValidPassword(string) []string has an error: want: %v got: %v\n", tt.want, !tt.want)
			}
		}
	}
}
