package templatex

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/jzero-io/jzero/pkg/stringx"
)

// ParseTemplate parse template with register func
func ParseTemplate(data interface{}, tplT []byte, funcMaps ...template.FuncMap) ([]byte, error) {
	t := template.New("production").Funcs(sprig.TxtFuncMap()).Funcs(RegisterTxtFuncMap())
	for _, fn := range funcMaps {
		t = t.Funcs(fn)
	}
	t, err := t.Parse(string(tplT))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func RegisterTxtFuncMap() template.FuncMap {
	return RegisterFuncMap()
}

func RegisterFuncMap() template.FuncMap {
	gfm := make(map[string]interface{}, len(registerFuncMap))
	for k, v := range registerFuncMap {
		gfm[k] = v
	}
	return gfm
}

var registerFuncMap = map[string]interface{}{
	"FirstUpper": stringx.FirstUpper,
	"ToCamel":    stringx.ToCamel,
}
