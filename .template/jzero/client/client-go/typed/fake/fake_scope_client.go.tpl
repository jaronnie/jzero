// Code generated by jzero. DO NOT EDIT.
// type: fake_{{.ScopeVersion}}_client

package fake

import (
	"{{.GoModule}}/rest"
	"{{$.GoModule}}/typed/{{.ScopeVersion}}"
)

type Fake{{.UpScopeVersion}} struct {}

func (f *Fake{{.UpScopeVersion}}) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

{{range $v := .UpResources}}func (f *Fake{{$.UpScopeVersion}}) {{$v}}() {{$.ScopeVersion}}.{{$v}}Interface {
	return &Fake{{$v}}{Fake: f}
}

{{end}}