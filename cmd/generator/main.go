package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

type opType struct {
	Op                    string
	Code                  string
	Indent                func() string
	Escaped               func() string
	HeadToPtrHead         func() string
	HeadToNPtrHead        func() string
	HeadToAnonymousHead   func() string
	HeadToOmitEmptyHead   func() string
	HeadToStringTagHead   func() string
	HeadToOnlyHead        func() string
	PtrHeadToHead         func() string
	FieldToEnd            func() string
	FieldToOmitEmptyField func() string
	FieldToStringTagField func() string
}

func (t opType) IsEscaped() bool {
	return t.Op != t.Escaped()
}

func (t opType) IsHeadToPtrHead() bool {
	return t.Op != t.HeadToPtrHead()
}

func (t opType) IsHeadToNPtrHead() bool {
	return t.Op != t.HeadToNPtrHead()
}

func (t opType) IsHeadToAnonymousHead() bool {
	return t.Op != t.HeadToAnonymousHead()
}

func (t opType) IsHeadToOmitEmptyHead() bool {
	return t.Op != t.HeadToOmitEmptyHead()
}

func (t opType) IsHeadToStringTagHead() bool {
	return t.Op != t.HeadToStringTagHead()
}

func (t opType) IsPtrHeadToHead() bool {
	return t.Op != t.PtrHeadToHead()
}

func (t opType) IsHeadToOnlyHead() bool {
	return t.Op != t.HeadToOnlyHead()
}

func (t opType) IsFieldToEnd() bool {
	return t.Op != t.FieldToEnd()
}

func (t opType) IsFieldToOmitEmptyField() bool {
	return t.Op != t.FieldToOmitEmptyField()
}

func (t opType) IsFieldToStringTagField() bool {
	return t.Op != t.FieldToStringTagField()
}

func createOpType(op, code string) opType {
	return opType{
		Op:                    op,
		Code:                  code,
		Indent:                func() string { return fmt.Sprintf("%sIndent", op) },
		Escaped:               func() string { return op },
		HeadToPtrHead:         func() string { return op },
		HeadToNPtrHead:        func() string { return op },
		HeadToAnonymousHead:   func() string { return op },
		HeadToOmitEmptyHead:   func() string { return op },
		HeadToStringTagHead:   func() string { return op },
		HeadToOnlyHead:        func() string { return op },
		PtrHeadToHead:         func() string { return op },
		FieldToEnd:            func() string { return op },
		FieldToOmitEmptyField: func() string { return op },
		FieldToStringTagField: func() string { return op },
	}
}

func _main() error {
	tmpl, err := template.New("").Parse(`// Code generated by cmd/generator. DO NOT EDIT!
package json

type codeType int

const (
{{- range $index, $type := .CodeTypes }}
  code{{ $type }} codeType = {{ $index }}
{{- end }}
)

type opType int

const (
{{- range $index, $type := .OpTypes }}
  op{{ $type.Op }} opType = {{ $index }}
{{- end }}
)

func (t opType) String() string {
  if int(t) >= {{ .OpLen }} {
    return t.toNotIndent().String() + "Indent"
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
  case op{{ $type.Op }}:
    return "{{ $type.Op }}"
{{- end }}
  }
  return ""
}

func (t opType) codeType() codeType {
  if int(t) >= {{ .OpLen }} {
    return t.toNotIndent().codeType()
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
  case op{{ $type.Op }}:
    return code{{ $type.Code }}
{{- end }}
  }
  return codeOp
}

func (t opType) toNotIndent() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t) - {{ .OpLen }})
  }
  return t
}

func (t opType) toIndent() opType {
  if int(t) >= {{ .OpLen }} {
    return t
  }
  return opType(int(t) + {{ .OpLen }})
}

func (t opType) toEscaped() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().toEscaped()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsEscaped }}
  case op{{ $type.Op }}:
    return op{{ call $type.Escaped }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) headToPtrHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().headToPtrHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsHeadToPtrHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.HeadToPtrHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) headToNPtrHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().headToNPtrHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsHeadToNPtrHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.HeadToNPtrHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) headToAnonymousHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().headToAnonymousHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsHeadToAnonymousHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.HeadToAnonymousHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) headToOmitEmptyHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().headToOmitEmptyHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsHeadToOmitEmptyHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.HeadToOmitEmptyHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) headToStringTagHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().headToStringTagHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsHeadToStringTagHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.HeadToStringTagHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) headToOnlyHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().headToOnlyHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsHeadToOnlyHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.HeadToOnlyHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) ptrHeadToHead() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().ptrHeadToHead()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsPtrHeadToHead }}
  case op{{ $type.Op }}:
    return op{{ call $type.PtrHeadToHead }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) fieldToEnd() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().fieldToEnd()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsFieldToEnd }}
  case op{{ $type.Op }}:
    return op{{ call $type.FieldToEnd }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) fieldToOmitEmptyField() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().fieldToOmitEmptyField()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsFieldToOmitEmptyField }}
  case op{{ $type.Op }}:
    return op{{ call $type.FieldToOmitEmptyField }}
{{- end }}
{{- end }}
  }
  return t
}

func (t opType) fieldToStringTagField() opType {
  if int(t) >= {{ .OpLen }} {
    return opType(int(t.toNotIndent().fieldToStringTagField()) + {{ .OpLen }})
  }

  switch t {
{{- range $type := .OpNotIndentTypes }}
{{- if $type.IsFieldToStringTagField }}
  case op{{ $type.Op }}:
    return op{{ call $type.FieldToStringTagField }}
{{- end }}
{{- end }}
  }
  return t
}

`)
	if err != nil {
		return err
	}
	codeTypes := []string{
		"Op",
		"ArrayHead",
		"ArrayElem",
		"SliceHead",
		"SliceElem",
		"MapHead",
		"MapKey",
		"MapValue",
		"MapEnd",
		"StructFieldRecursive",
		"StructField",
		"StructEnd",
	}
	primitiveTypes := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool", "string", "escapedString", "bytes",
		"array", "map", "mapLoad", "slice", "struct", "MarshalJSON", "MarshalText", "recursive",
		"intString", "int8String", "int16String", "int32String", "int64String",
		"uintString", "uint8String", "uint16String", "uint32String", "uint64String",
		"intPtr", "int8Ptr", "int16Ptr", "int32Ptr", "int64Ptr",
		"uintPtr", "uint8Ptr", "uint16Ptr", "uint32Ptr", "uint64Ptr",
		"float32Ptr", "float64Ptr", "boolPtr", "stringPtr", "escapedStringPtr", "bytesPtr",
		"intNPtr", "int8NPtr", "int16NPtr", "int32NPtr", "int64NPtr",
		"uintNPtr", "uint8NPtr", "uint16NPtr", "uint32NPtr", "uint64NPtr",
		"float32NPtr", "float64NPtr", "boolNPtr", "stringNPtr", "escapedStringNPtr", "bytesNPtr",
	}
	primitiveTypesUpper := []string{}
	for _, typ := range primitiveTypes {
		primitiveTypesUpper = append(primitiveTypesUpper, strings.ToUpper(string(typ[0]))+typ[1:])
	}
	opTypes := []opType{
		createOpType("End", "Op"),
		createOpType("Interface", "Op"),
		createOpType("InterfaceEnd", "Op"),
		createOpType("Ptr", "Op"),
		createOpType("NPtr", "Op"),
		createOpType("SliceHead", "SliceHead"),
		createOpType("RootSliceHead", "SliceHead"),
		createOpType("SliceElem", "SliceElem"),
		createOpType("RootSliceElem", "SliceElem"),
		createOpType("SliceEnd", "Op"),
		createOpType("ArrayHead", "ArrayHead"),
		createOpType("ArrayElem", "ArrayElem"),
		createOpType("ArrayEnd", "Op"),
		createOpType("MapHead", "MapHead"),
		createOpType("MapHeadLoad", "MapHead"),
		createOpType("MapKey", "MapKey"),
		createOpType("MapValue", "MapValue"),
		createOpType("MapEnd", "Op"),
		createOpType("StructFieldRecursiveEnd", "Op"),
		createOpType("StructAnonymousEnd", "StructEnd"),
	}
	for _, typ := range primitiveTypesUpper {
		typ := typ
		optype := createOpType(typ, "Op")
		switch typ {
		case "String", "StringPtr", "StringNPtr":
			optype.Escaped = func() string {
				return fmt.Sprintf("Escaped%s", typ)
			}
		}
		opTypes = append(opTypes, optype)
	}
	for _, escapedOrNot := range []string{"", "Escaped"} {
		for _, ptrOrNot := range []string{"", "Ptr", "NPtr"} {
			for _, headType := range []string{"", "Anonymous"} {
				for _, opt := range []string{"", "OmitEmpty", "StringTag"} {
					for _, typ := range append(primitiveTypesUpper, "") {
						escapedOrNot := escapedOrNot
						ptrOrNot := ptrOrNot
						headType := headType
						opt := opt
						typ := typ

						op := fmt.Sprintf(
							"Struct%sField%s%sHead%s%s",
							escapedOrNot,
							ptrOrNot,
							headType,
							opt,
							typ,
						)
						opTypes = append(opTypes, opType{
							Op:     op,
							Code:   "StructField",
							Indent: func() string { return fmt.Sprintf("%sIndent", op) },
							Escaped: func() string {
								switch typ {
								case "String", "StringPtr", "StringNPtr":
									return fmt.Sprintf(
										"StructEscapedField%s%sHead%sEscaped%s",
										ptrOrNot,
										headType,
										opt,
										typ,
									)
								}
								return fmt.Sprintf(
									"StructEscapedField%s%sHead%s%s",
									ptrOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToPtrHead: func() string {
								return fmt.Sprintf(
									"Struct%sFieldPtr%sHead%s%s",
									escapedOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToNPtrHead: func() string {
								return fmt.Sprintf(
									"Struct%sFieldNPtr%sHead%s%s",
									escapedOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToAnonymousHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%sAnonymousHead%s%s",
									escapedOrNot,
									ptrOrNot,
									opt,
									typ,
								)
							},
							HeadToOmitEmptyHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%s%sHeadOmitEmpty%s",
									escapedOrNot,
									ptrOrNot,
									headType,
									typ,
								)
							},
							HeadToStringTagHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%s%sHeadStringTag%s",
									escapedOrNot,
									ptrOrNot,
									headType,
									typ,
								)
							},
							HeadToOnlyHead: func() string {
								switch typ {
								case "", "Array", "Map", "MapLoad", "Slice",
									"Struct", "Recursive", "MarshalJSON", "MarshalText",
									"IntString", "Int8String", "Int16String", "Int32String", "Int64String",
									"UintString", "Uint8String", "Uint16String", "Uint32String", "Uint64String":
									return op
								}
								return fmt.Sprintf(
									"Struct%sField%s%sHead%s%sOnly",
									escapedOrNot,
									ptrOrNot,
									headType,
									opt,
									typ,
								)
							},
							PtrHeadToHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%sHead%s%s",
									escapedOrNot,
									headType,
									opt,
									typ,
								)
							},
							FieldToEnd:            func() string { return op },
							FieldToOmitEmptyField: func() string { return op },
							FieldToStringTagField: func() string { return op },
						})
					}
				}
			}
		}
	}
	for _, escapedOrNot := range []string{"", "Escaped"} {
		for _, ptrOrNot := range []string{"", "Ptr", "NPtr"} {
			for _, headType := range []string{"", "Anonymous"} {
				for _, opt := range []string{"", "OmitEmpty", "StringTag"} {
					for _, typ := range []string{
						"", "Int", "Int8", "Int16", "Int32", "Int64",
						"Uint", "Uint8", "Uint16", "Uint32", "Uint64",
						"Float32", "Float64", "Bool", "String", "EscapedString", "Bytes",
						"IntPtr", "Int8Ptr", "Int16Ptr", "Int32Ptr", "Int64Ptr",
						"UintPtr", "Uint8Ptr", "Uint16Ptr", "Uint32Ptr", "Uint64Ptr",
						"Float32Ptr", "Float64Ptr", "BoolPtr", "StringPtr", "EscapedStringPtr", "BytesPtr",
						"IntNPtr", "Int8NPtr", "Int16NPtr", "Int32NPtr", "Int64NPtr",
						"UintNPtr", "Uint8NPtr", "Uint16NPtr", "Uint32NPtr", "Uint64NPtr",
						"Float32NPtr", "Float64NPtr", "BoolNPtr", "StringNPtr", "EscapedStringNPtr", "BytesNPtr",
					} {
						escapedOrNot := escapedOrNot
						ptrOrNot := ptrOrNot
						headType := headType
						opt := opt
						typ := typ

						op := fmt.Sprintf(
							"Struct%sField%s%sHead%s%sOnly",
							escapedOrNot,
							ptrOrNot,
							headType,
							opt,
							typ,
						)
						opTypes = append(opTypes, opType{
							Op:     op,
							Code:   "StructField",
							Indent: func() string { return fmt.Sprintf("%sIndent", op) },
							Escaped: func() string {
								switch typ {
								case "String", "StringPtr", "StringNPtr":
									return fmt.Sprintf(
										"StructEscapedField%s%sHead%sEscaped%sOnly",
										ptrOrNot,
										headType,
										opt,
										typ,
									)
								}
								return fmt.Sprintf(
									"StructEscapedField%s%sHead%s%sOnly",
									ptrOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToPtrHead: func() string {
								return fmt.Sprintf(
									"Struct%sFieldPtr%sHead%s%sOnly",
									escapedOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToNPtrHead: func() string {
								return fmt.Sprintf(
									"Struct%sFieldNPtr%sHead%s%sOnly",
									escapedOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToAnonymousHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%sAnonymousHead%s%sOnly",
									escapedOrNot,
									ptrOrNot,
									opt,
									typ,
								)
							},
							HeadToOmitEmptyHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%s%sHeadOmitEmpty%sOnly",
									escapedOrNot,
									ptrOrNot,
									headType,
									typ,
								)
							},
							HeadToStringTagHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%s%sHeadStringTag%sOnly",
									escapedOrNot,
									ptrOrNot,
									headType,
									typ,
								)
							},
							PtrHeadToHead: func() string {
								return fmt.Sprintf(
									"Struct%sField%sHead%s%sOnly",
									escapedOrNot,
									headType,
									opt,
									typ,
								)
							},
							HeadToOnlyHead:        func() string { return op },
							FieldToEnd:            func() string { return op },
							FieldToOmitEmptyField: func() string { return op },
							FieldToStringTagField: func() string { return op },
						})
					}
				}
			}
		}
	}
	for _, escapedOrNot := range []string{"", "Escaped"} {
		for _, opt := range []string{"", "OmitEmpty", "StringTag"} {
			for _, typ := range append(primitiveTypesUpper, "") {
				escapedOrNot := escapedOrNot
				opt := opt
				typ := typ

				op := fmt.Sprintf(
					"Struct%sField%s%s",
					escapedOrNot,
					opt,
					typ,
				)
				opTypes = append(opTypes, opType{
					Op:     op,
					Code:   "StructField",
					Indent: func() string { return fmt.Sprintf("%sIndent", op) },
					Escaped: func() string {
						switch typ {
						case "String", "StringPtr", "StringNPtr":
							return fmt.Sprintf(
								"StructEscapedField%sEscaped%s",
								opt,
								typ,
							)
						}
						return fmt.Sprintf(
							"StructEscapedField%s%s",
							opt,
							typ,
						)
					},
					HeadToPtrHead:       func() string { return op },
					HeadToNPtrHead:      func() string { return op },
					HeadToAnonymousHead: func() string { return op },
					HeadToOmitEmptyHead: func() string { return op },
					HeadToStringTagHead: func() string { return op },
					HeadToOnlyHead:      func() string { return op },
					PtrHeadToHead:       func() string { return op },
					FieldToEnd: func() string {
						switch typ {
						case "", "Array", "Map", "MapLoad", "Slice", "Struct", "Recursive":
							return op
						}
						return fmt.Sprintf(
							"Struct%sEnd%s%s",
							escapedOrNot,
							opt,
							typ,
						)
					},
					FieldToOmitEmptyField: func() string {
						return fmt.Sprintf(
							"Struct%sFieldOmitEmpty%s",
							escapedOrNot,
							typ,
						)
					},
					FieldToStringTagField: func() string {
						return fmt.Sprintf(
							"Struct%sFieldStringTag%s",
							escapedOrNot,
							typ,
						)
					},
				})
			}
		}
	}
	for _, escapedOrNot := range []string{"", "Escaped"} {
		for _, opt := range []string{"", "OmitEmpty", "StringTag"} {
			for _, typ := range append(primitiveTypesUpper, "") {
				escapedOrNot := escapedOrNot
				opt := opt
				typ := typ

				op := fmt.Sprintf(
					"Struct%sEnd%s%s",
					escapedOrNot,
					opt,
					typ,
				)
				opTypes = append(opTypes, opType{
					Op:     op,
					Code:   "StructEnd",
					Indent: func() string { return fmt.Sprintf("%sIndent", op) },
					Escaped: func() string {
						switch typ {
						case "String", "StringPtr", "StringNPtr":
							return fmt.Sprintf(
								"StructEscapedEnd%sEscaped%s",
								opt,
								typ,
							)
						}
						return fmt.Sprintf(
							"StructEscapedEnd%s%s",
							opt,
							typ,
						)
					},
					HeadToPtrHead:         func() string { return op },
					HeadToNPtrHead:        func() string { return op },
					HeadToAnonymousHead:   func() string { return op },
					HeadToOmitEmptyHead:   func() string { return op },
					HeadToStringTagHead:   func() string { return op },
					HeadToOnlyHead:        func() string { return op },
					PtrHeadToHead:         func() string { return op },
					FieldToEnd:            func() string { return op },
					FieldToOmitEmptyField: func() string { return op },
					FieldToStringTagField: func() string { return op },
				})
			}
		}
	}
	indentOpTypes := []opType{}
	for _, typ := range opTypes {
		typ := typ
		indentOpTypes = append(indentOpTypes, opType{
			Op:      fmt.Sprintf("%sIndent", typ.Op),
			Code:    typ.Code,
			Indent:  func() string { return fmt.Sprintf("%sIndent", typ.Op) },
			Escaped: func() string { return fmt.Sprintf("%sIndent", typ.Escaped()) },
			HeadToPtrHead: func() string {
				return fmt.Sprintf("%sIndent", typ.HeadToPtrHead())
			},
			HeadToNPtrHead: func() string {
				return fmt.Sprintf("%sIndent", typ.HeadToNPtrHead())
			},
			HeadToAnonymousHead: func() string {
				return fmt.Sprintf("%sIndent", typ.HeadToAnonymousHead())
			},
			HeadToOmitEmptyHead: func() string {
				return fmt.Sprintf("%sIndent", typ.HeadToOmitEmptyHead())
			},
			HeadToStringTagHead: func() string {
				return fmt.Sprintf("%sIndent", typ.HeadToStringTagHead())
			},
			HeadToOnlyHead: func() string {
				return fmt.Sprintf("%sIndent", typ.HeadToOnlyHead())
			},
			PtrHeadToHead: func() string {
				return fmt.Sprintf("%sIndent", typ.PtrHeadToHead())
			},
			FieldToOmitEmptyField: func() string {
				return fmt.Sprintf("%sIndent", typ.FieldToOmitEmptyField())
			},
			FieldToStringTagField: func() string {
				return fmt.Sprintf("%sIndent", typ.FieldToStringTagField())
			},
			FieldToEnd: func() string {
				return fmt.Sprintf("%sIndent", typ.FieldToEnd())
			},
		})
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, struct {
		CodeTypes        []string
		OpTypes          []opType
		OpNotIndentTypes []opType
		OpLen            int
		OpIndentLen      int
	}{
		CodeTypes:        codeTypes,
		OpTypes:          append(opTypes, indentOpTypes...),
		OpNotIndentTypes: opTypes,
		OpLen:            len(opTypes),
		OpIndentLen:      len(indentOpTypes),
	}); err != nil {
		return err
	}
	path := filepath.Join(repoRoot(), "encode_optype.go")
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, buf, 0644)
}

func repoRoot() string {
	_, file, _, _ := runtime.Caller(0)
	relativePathFromRepoRoot := filepath.Join("cmd", "generator")
	return strings.TrimSuffix(filepath.Dir(file), relativePathFromRepoRoot)
}

func main() {
	if err := _main(); err != nil {
		panic(err)
	}
}
