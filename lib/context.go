package rat

import (
	"fmt"
	"strings"
)

type Context map[string]string

func NewContextFromAnnotations(annotations []Annotation) Context {
	ctx := Context{}

	for _, a := range annotations {
		ctx[a.Class()] = a.Val()
	}

	return ctx
}

func InterpolateContext(str string, ctx Context) string {
	for k, v := range ctx {
		str = strings.Replace(str, fmt.Sprintf("%%(%s)", k), v, -1)
	}

	return str
}

func MergeContext(orig, extra Context) Context {
	merged := Context{}

	for k, v := range orig {
		merged[k] = v
	}

	for k, v := range extra {
		merged[k] = v
	}

	return merged
}
