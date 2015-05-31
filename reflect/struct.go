package reflect

import (
	"go/ast"
	"path/filepath"
	"strings"
)

// Struct is a type that represents information about a specific struct,
// its fields, and comments group.
type Struct struct {
	Comments []string // Comments right above the struct declaration.
	Fields   []Arg    // A list of fields that belong to this struct.
	Name     string   // Name of the struct, e.g. "Application".
}

// processStructTypeSpec expects ast type spec as input parameter
// that is transformed into *Struct representation and returned.
func processTypeSpec(spec *ast.TypeSpec) *Struct {
	// Make sure it is a structure type. Return nil if not.
	structType, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil
	}

	// Compose a structure and return it.
	return &Struct{
		Fields: processFieldList(structType.Fields),
		Name:   spec.Name.Name,
	}
}

// processImportSpec gets ast import spec as input parameter
// and transforms it into a pair of import's name and its value.
func processImportSpec(spec *ast.ImportSpec) (name, value string) {
	// Remove quote signes, etc. from the left and right sides of import.
	value = strings.Trim(spec.Path.Value, "\"`")

	// By-default, use the last (base) part of import as a name.
	name = filepath.Base(value)

	// However, if name value is presented in spec, use it instead.
	if n := spec.Name; n != nil {
		name = n.Name
	}

	return
}
