package directdebit_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

func TestStructToURLValues(t *testing.T) {
	type Nested struct {
		InnerField string `url:"inner_field"`
	}

	type DoubleNested struct {
		FirstLevel Nested `url:"first_level"`
	}

	type Sample struct {
		FieldA string `url:"field_a"`
		FieldB string `url:"field_b"`
		FieldC string
	}

	type NoTagSample struct {
		FieldA string
		FieldB string
	}

	type NonStringSample struct {
		FieldA int `url:"field_a"`
	}

	type NestedUnMarshalable struct {
		InnerField func() `url:"inner_field"`
	}

	type UnmarshallableSample struct {
		FirstLevel NestedUnMarshalable `url:"FirstLevel"`
	}

	tests := []struct {
		name        string
		input       interface{}
		expected    url.Values
		wantErr     bool
		expectedErr string
	}{
		{
			name: "Typical use-case",
			input: Sample{
				FieldA: "valueA",
				FieldB: "valueB",
				FieldC: "valueC",
			},
			expected: url.Values{
				"field_a": []string{"valueA"},
				"field_b": []string{"valueB"},
			},
		},
		{
			name:  "Empty struct",
			input: Sample{},
			expected: url.Values{
				"field_a": []string{""},
				"field_b": []string{""},
			},
		},
		{
			name:     "No tags",
			input:    NoTagSample{FieldA: "valueA", FieldB: "valueB"},
			expected: url.Values{},
		},
		{
			name:        "Non-string field",
			input:       NonStringSample{FieldA: 123},
			expected:    nil,
			wantErr:     true,
			expectedErr: "non-string field with 'url' tag encountered",
		},
		{
			name: "Pointer to struct",
			input: &Sample{
				FieldA: "valueA",
				FieldB: "valueB",
				FieldC: "valueC",
			},
			expected: url.Values{
				"field_a": []string{"valueA"},
				"field_b": []string{"valueB"},
			},
		},
		{
			name:  "Nested struct",
			input: Nested{InnerField: "nestedValue"},
			expected: url.Values{
				"inner_field": []string{"nestedValue"},
			},
		},
		{
			name: "Double nested struct",
			input: DoubleNested{
				FirstLevel: Nested{InnerField: "nestedValue"},
			},
			expected: url.Values{
				"first_level": []string{"{\"InnerField\":\"nestedValue\"}"},
			},
		},
		{
			name:        "Unmarshallable field",
			input:       UnmarshallableSample{FirstLevel: NestedUnMarshalable{InnerField: func() {}}},
			expected:    nil,
			wantErr:     true,
			expectedErr: "json: unsupported type: func()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := directdebit.StructToURLValues(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("StructToURLValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("StructToURLValues() error message = %v, expectedErr %v", err.Error(), tt.expectedErr)
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, got)
			}
		})
	}
}
