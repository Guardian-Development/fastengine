package loadbytevector

import (
	"fmt"
	"github.com/Guardian-Development/fastengine/pkg/template/structure"

	"github.com/Guardian-Development/fastengine/internal/xml"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/fieldbytevector"
	"github.com/Guardian-Development/fastengine/pkg/fast/field/properties"
	"github.com/Guardian-Development/fastengine/pkg/template/loader/converter"
)

// Load an <bytevector /> tag with supported operation
func Load(tagInTemplate *xml.Tag, fieldDetails properties.Properties) (fieldbytevector.FieldByteVector, error) {
	if len(tagInTemplate.NestedTags) != 1 {
		return fieldbytevector.New(fieldDetails), nil
	}

	operationTag := tagInTemplate.NestedTags[0]
	operationType := operationTag.Type
	hasOperationValue := structure.HasValue(&operationTag)

	switch operationType {
	case structure.DefaultOperation:
		if !hasOperationValue {
			return fieldbytevector.NewDefaultOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToByteVector(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldbytevector.FieldByteVector{}, err
		}
		return fieldbytevector.NewDefaultOperationWithValue(fieldDetails, operationValue), nil
	case structure.ConstantOperation:
		if !hasOperationValue {
			return fieldbytevector.FieldByteVector{}, fmt.Errorf("no value specified for constant operation")
		}

		operationValue, err := converter.ToByteVector(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldbytevector.FieldByteVector{}, err
		}
		return fieldbytevector.NewConstantOperation(fieldDetails, operationValue), nil
	case structure.CopyOperation:
		if !hasOperationValue {
			return fieldbytevector.NewCopyOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToByteVector(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldbytevector.FieldByteVector{}, err
		}
		return fieldbytevector.NewCopyOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.TailOperation:
		if !hasOperationValue {
			return fieldbytevector.NewTailOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToByteVector(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldbytevector.FieldByteVector{}, err
		}
		return fieldbytevector.NewTailOperationWithInitialValue(fieldDetails, operationValue), nil
	case structure.DeltaOperation:
		if !hasOperationValue {
			return fieldbytevector.NewDeltaOperation(fieldDetails), nil
		}

		operationValue, err := converter.ToByteVector(operationTag.Attributes[structure.ValueAttribute])
		if err != nil {
			return fieldbytevector.FieldByteVector{}, err
		}
		return fieldbytevector.NewDeltaOperationWithInitialValue(fieldDetails, operationValue), nil
	default:
		return fieldbytevector.FieldByteVector{}, fmt.Errorf("unsupported operation type: %s", operationTag)
	}
}
