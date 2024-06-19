package generator

import (
	"bytes"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jzero-io/jzero/embeded"
	"github.com/jzero-io/jzero/internal/gen"
	"github.com/jzero-io/jzero/internal/gen/gensdk/config"
	"github.com/jzero-io/jzero/internal/gen/gensdk/jparser"
	"github.com/jzero-io/jzero/pkg/templatex"
	apiparser "github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"os"
	"path/filepath"
)

func init() {
	Register("ts", func(config config.Config) (Generator, error) {
		return &Typescript{
			Config: &config,
		}, nil
	})
}

type Typescript struct {
	Config *config.Config

	wd string
}

func (t *Typescript) Gen() ([]*GeneratedFile, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	t.wd = wd

	// parse api
	var apiSpecs []*spec.ApiSpec

	mainApiFilePath := gen.GetMainApiFilePath(t.Config.ApiDir)
	apiSpec, err := apiparser.Parse(mainApiFilePath)
	if err != nil {
		return nil, err
	}
	if mainApiFilePath != filepath.Join(t.Config.ApiDir, "main.api") {
		os.Remove(mainApiFilePath)
	}

	apiSpecs = append(apiSpecs, apiSpec)

	protoFiles, err := gen.GetProtoFilenames(t.Config.ProtoDir)
	if err != nil {
		return nil, err
	}

	var fds []*desc.FileDescriptor

	// parse proto
	var protoParser protoparse.Parser
	if len(protoFiles) > 0 {
		protoParser.ImportPaths = []string{t.Config.ProtoDir}
		fds, err = protoParser.ParseFiles(protoFiles...)
		if err != nil {
			return nil, err
		}
	}

	rhis, err := jparser.Parse(t.Config, fds, apiSpecs)
	if err != nil {
		return nil, err
	}

	var files []*GeneratedFile

	clientSetFile, err := t.genClientSet(GetScopes(rhis))
	if err != nil {
		return nil, err
	}
	files = append(files, clientSetFile)
	return files, nil
}

func (t *Typescript) genClientSet(scopes []string) (*GeneratedFile, error) {
	clientBytes, err := templatex.ParseTemplate(map[string]interface{}{
		"APP":    t.Config.APP,
		"Scopes": scopes,
	}, embeded.ReadTemplateFile(filepath.Join("client", "client-ts", "index.ts.tpl")))
	if err != nil {
		return nil, err
	}
	return &GeneratedFile{
		Path:    "index.ts",
		Content: *bytes.NewBuffer(clientBytes),
	}, nil
}
