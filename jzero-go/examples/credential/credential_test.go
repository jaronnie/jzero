package credential

import (
	"testing"

	"github.com/jzero-io/jzero-go/model/jzero/pb/credentialpb"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/jzero-io/jzero-go"
	"github.com/jzero-io/jzero-go/rest"
	typedjzero "github.com/jzero-io/jzero-go/typed/jzero"
	"github.com/jzero-io/jzero-go/typed/jzero/fake"
)

func TestGetCredentialList(t *testing.T) {
	gomonkey.ApplyFunc(jzero.NewClientWithOptions, func(ops ...rest.Opt) (typedjzero.JzeroInterface, error) {
		return &fake.FakeJzero{}, nil
	})

	fake.FakeReturnCredentialList = &credentialpb.CredentialListResponse{
		Total: 10,
	}

	list, err := GetCredentialList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(list)
}