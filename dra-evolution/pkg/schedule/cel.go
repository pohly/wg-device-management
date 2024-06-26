package cel

import (
	"context"
	"fmt"

	"k8s.io/apiserver/pkg/cel/environment"
)

func MeetsConstraints(filter *string, input DeviceAttributes) (bool, error) {
	if filter == nil || *filter == "" {
		return true, nil
	}

	expr := Compiler.CompileCELExpression(*filter, environment.StoredExpressions)
	if expr.Error != nil {
		return false, fmt.Errorf("compile CEL expression: %v", expr.Error)
	}
	return expr.DeviceMatches(context.TODO(), input)
}
