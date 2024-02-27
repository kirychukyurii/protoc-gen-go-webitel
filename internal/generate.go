package internal

import (
	"fmt"
	"regexp"

	"google.golang.org/protobuf/compiler/protogen"
)

var reg = regexp.MustCompile(`^(.*\.)`)

func Generate(gen *protogen.Plugin, f []*protogen.File, version string) error {
	g := gen.NewGeneratedFile("options.go", "")
	g.P("// Code generated by protoc-gen-go-webitel. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-webitel v", version)
	g.P("// - protoc             ", protocVersion(gen))
	g.P()
	g.P("package ", f[0].GoPackageName)
	g.P()

	generateFileContent(gen, f, g)

	return nil
}

func generateFileContent(gen *protogen.Plugin, files []*protogen.File, g *protogen.GeneratedFile) {
	generateType(g)

	g.P("var WebitelAPI = WebitelServicesInfo{")
	for _, f := range files {

		for _, s := range f.Proto.GetService() {
			objclass, err := extractServiceObjClassOption(s)
			if err != nil {
				return
			}

			g.P(`"`, s.GetName(), `": WebitelServices{`)
			g.P("ObjClass: ", `"`, objclass, `"`, ",")
			g.P("WebitelMethods: map[string]WebitelMethod{")
			for _, m := range s.GetMethod() {
				acc, err := extractMethodAccessOption(m)
				if err != nil {
					return
				}

				ht, err := extractMethodHttpOption(m)
				if err != nil {
					return
				}

				g.P(`"`, m.GetName(), `": WebitelMethod{`)
				g.P("Access: ", acc, ",")
				g.P("Input: ", `"`, reg.ReplaceAllString(m.GetInputType(), ""), `",`)
				g.P("Output: ", `"`, reg.ReplaceAllString(m.GetInputType(), ""), `",`)
				g.P("HttpBindings: []*HttpBinding{")
				for _, h := range ht {
					g.P("{")
					g.P("Path: ", `"`, h.Path, `",`)
					g.P("Method: ", `"`, h.Method, `",`)
					g.P("},")
				}
				g.P("},")
				g.P("},")

			}
			g.P("},")
			g.P("},")
		}
	}
	g.P("}")
	g.P()
}

func generateType(g *protogen.GeneratedFile) {
	g.P("// WebitelServicesInfo is the list of services defined in proto files.")
	g.P("type WebitelServicesInfo map[string]WebitelServices")
	g.P()
	g.P("type WebitelServices struct {")
	g.P("ObjClass       string")
	g.P("WebitelMethods map[string]WebitelMethod")
	g.P("}")
	g.P()
	g.P("// WebitelMethod is the list of methods defined in this service.")
	g.P("type WebitelMethod struct {")
	g.P("HttpBindings []*HttpBinding")
	g.P("Access       int")
	g.P("Input        string")
	g.P("Output       string")
	g.P("}")
	g.P()
	g.P("type HttpBinding struct {")
	g.P("Path   string")
	g.P("Method string")
	g.P("}")
	g.P()
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}

	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}

	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}
