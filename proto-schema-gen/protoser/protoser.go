package protoser

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yoheimuta/go-protoparser"
	"github.com/yoheimuta/go-protoparser/parser"
	"io"
	"proto-schema-gen/schema"
)

func SerialiseBytes(b []byte) (*schema.Proto, error) {
	buf := bytes.NewBuffer(b)
	return Serialise(buf)
}

func Serialise(reader io.Reader) (*schema.Proto, error) {
	proto, err := protoparser.Parse(reader)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse the given schema")
	}

	return fromParserProto(proto), nil
}

func fromParserProto(proto *parser.Proto) *schema.Proto {
	var r schema.Proto

	r.Syntax = proto.Syntax.ProtobufVersion
	r.Body = protoBody(proto.ProtoBody)

	return &r
}

func protoBody(body []parser.Visitee) []*schema.ProtoBodyElement {
	r := make([]*schema.ProtoBodyElement, len(body))
	for i, v := range body {
		r[i] = protoBodyElement(v)
	}
	return r
}

func protoBodyElement(i parser.Visitee) *schema.ProtoBodyElement {
	switch v := i.(type) {

	case *parser.Import:
		return protoBodyElement_Import(import_(v))
	case *parser.Package:
		return protoBodyElement_Package(package_(v))
	default:
		fmt.Printf("Discarding unknown type %#t\n", v)
		return nil
	}
}

func protoBodyElement_Import(imp *schema.Import) *schema.ProtoBodyElement {
	return &schema.ProtoBodyElement{Elem: &schema.ProtoBodyElement_Import{Import: imp}}
}

func import_(imp *parser.Import) *schema.Import {
	isPublic := imp.Modifier&parser.ImportModifierPublic != 0
	isWeak := imp.Modifier&parser.ImportModifierWeak != 0

	return &schema.Import{
		Import:   imp.Location,
		IsPublic: isPublic,
		IsWeak:   isWeak,
	}
}

func protoBodyElement_Package(imp *schema.Package) *schema.ProtoBodyElement {
	return &schema.ProtoBodyElement{Elem: &schema.ProtoBodyElement_Package{Package: imp}}
}

func package_(pack *parser.Package) *schema.Package {
	return &schema.Package{Name: pack.Name}
}
