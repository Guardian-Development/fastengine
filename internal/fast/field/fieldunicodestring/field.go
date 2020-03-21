package fieldunicodestring

import (
	"bytes"

	"github.com/Guardian-Development/fastengine/client/fix"
	"github.com/Guardian-Development/fastengine/internal/fast"
	"github.com/Guardian-Development/fastengine/internal/fast/dictionary"
	"github.com/Guardian-Development/fastengine/internal/fast/field/properties"
	"github.com/Guardian-Development/fastengine/internal/fast/operation"
	"github.com/Guardian-Development/fastengine/internal/fast/presencemap"
	"github.com/Guardian-Development/fastengine/internal/fast/value"
)

// FieldUnicodeString represents a FAST template <string charset="unicode"/> type
type FieldUnicodeString struct {
	FieldDetails properties.Properties
	Operation    operation.Operation
}

// Deserialise a <string charset="unicode"/> from the input source
func (field FieldUnicodeString) Deserialise(inputSource *bytes.Buffer, pMap *presencemap.PresenceMap, dictionary *dictionary.Dictionary) (fix.Value, error) {
	previousValue := dictionary.GetValue(field.FieldDetails.Name)
	if field.Operation.ShouldReadValue(pMap) {
		var stringValue value.Value
		var err error

		if field.FieldDetails.Required {
			stringValue, err = fast.ReadByteVector(inputSource)
		} else {
			stringValue, err = fast.ReadOptionalByteVector(inputSource)
		}

		if err != nil {
			return nil, err
		}

		switch t := stringValue.(type) {
		case value.ByteVector:
			stringValue = value.StringValue{Value: string(t.Value)}
		}

		transformedValue, err := field.Operation.Apply(stringValue, previousValue)
		if err != nil {
			return nil, err
		}
		dictionary.SetValue(field.FieldDetails.Name, transformedValue)
		return transformedValue, nil
	}

	transformedValue, err := field.Operation.GetNotEncodedValue(pMap, field.FieldDetails.Required, previousValue)
	if err != nil {
		return nil, err
	}
	dictionary.SetValue(field.FieldDetails.Name, transformedValue)
	return transformedValue, nil
}

// GetTagId for this field
func (field FieldUnicodeString) GetTagId() uint64 {
	return field.FieldDetails.ID
}

// RequiresPmap returns whether the underlying operation for this field requires a pmap bit being set
func (field FieldUnicodeString) RequiresPmap() bool {
	return field.Operation.RequiresPmap(field.FieldDetails.Required)
}

// New <string charset="unicode"/> field with the given properties and no operation
func New(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation:    operation.None{},
	}

	return field
}

// NewConstantOperation <string charset="unicode"/> field with the given properties and <constant value="constantValue"/> operator
func NewConstantOperation(properties properties.Properties, constantValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Constant{
			ConstantValue: fix.NewRawValue(constantValue),
		},
	}

	return field
}

// NewDefaultOperation <string charset="unicode"/> field with the given properties and <default /> operator
func NewDefaultOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Default{
			DefaultValue: fix.NullValue{},
		},
	}

	return field
}

// NewDefaultOperationWithValue <string charset="unicode"/> field with the given properties and <default value="constantValue"/> operator
func NewDefaultOperationWithValue(properties properties.Properties, defaultValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Default{
			DefaultValue: fix.NewRawValue(defaultValue),
		},
	}

	return field
}

// NewCopyOperation <string charset="unicode"/> field with the given properties and <copy/> operator
func NewCopyOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Copy{
			InitialValue: fix.NullValue{},
		},
	}

	return field
}

// NewCopyOperationWithInitialValue <string charset="unicode"/> field with the given properties and <copy value="initialValue"/> operator
func NewCopyOperationWithInitialValue(properties properties.Properties, initialValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Copy{
			InitialValue: fix.NewRawValue(initialValue),
		},
	}

	return field
}

// NewTailOperation <string charset="unicode"/> field with the given properties and <tail/> operator
func NewTailOperation(properties properties.Properties) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Tail{
			InitialValue: fix.NullValue{},
			BaseValue: fix.NewRawValue(""),
		},
	}
	return field
}

// NewTailOperationWithInitialValue <string charset="unicode"/> field with the given properties and <tail value="initialValue"/> operator
func NewTailOperationWithInitialValue(properties properties.Properties, initialValue string) FieldUnicodeString {
	field := FieldUnicodeString{
		FieldDetails: properties,
		Operation: operation.Tail{
			InitialValue: fix.NewRawValue(initialValue),
			BaseValue: fix.NewRawValue(""),
		},
	}
	return field
}
