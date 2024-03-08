package cmd

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/kirychukyurii/protoc-gen-go-webitel/internal"
)

func Run(version string) func(gen *protogen.Plugin) error {
	return func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		var files []*protogen.File
		for _, f := range gen.Files {
			if f.Generate {
				files = append(files, f)
			}
		}
		if err := internal.Generate(gen, files, version); err != nil {
			return err
		}

		return nil
	}
}
