package logging

import (
	"fmt"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type autoLogField struct {
	logKey              string
	fieldMarshallerFunc reflect.Value
}

func marshalLogField(logger *zap.Logger, val interface{}, fieldEntry autoLogField) *zap.Logger {
	expectedType := fieldEntry.fieldMarshallerFunc.Type().In(1)
	valReflect := reflect.ValueOf(val)

	if valReflect.CanConvert(expectedType) {
		args := []reflect.Value{reflect.ValueOf(fieldEntry.logKey), valReflect.Convert(expectedType)}
		result := fieldEntry.fieldMarshallerFunc.Call(args)
		logger = logger.With(result[0].Interface().(zapcore.Field))
	}
	return logger
}

var contextLogFields = make(map[any]autoLogField)

func RegisterFieldForContextLog[T any](logKey string, contextKey any, valueMarshaller func(string, T) zapcore.Field) {
	if existingField, exists := contextLogFields[contextKey]; exists {
		panic(fmt.Errorf("tried to register existing context log field '%v'('%s') - already registered ('%s')",
			contextKey, logKey, existingField.logKey))
	}

	contextLogFields[contextKey] = autoLogField{
		logKey:              logKey,
		fieldMarshallerFunc: reflect.ValueOf(valueMarshaller),
	}
}
