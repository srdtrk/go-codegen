package codegen

import (
	"github.com/dave/jennifer/jen"

	"github.com/srdtrk/go-codegen/pkg/schemas"
)

func GenerateInstantiateMsg(f *jen.File, schema *schemas.JSONSchema) {
	f.Type().Id("InstantiateMsg").Struct()
}
