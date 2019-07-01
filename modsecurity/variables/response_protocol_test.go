package variables

import (
	"testing"

	"github.com/senghoo/modsecurity-go/modsecurity"
)

func TestVariableResponseProtocol(t *testing.T) {
	v := NewVariableResponseProtocol()
	if v.Name() != "RESPONSE_PROTOCOL" {
		t.Errorf("got unexcepted variable name %s", v.Name())
		return
	}
	tr, err := modsecurity.NewTransaction(modsecurity.NewEngine(), modsecurity.NewSecRuleSet())
	if err != nil {
		t.Error(err)
		return
	}
	tr.ProcessResponseHeaders(200, "HTTP/1.1", nil)
	vars := v.Fetch(tr)

	if len(vars) != 1 {
		t.Errorf("unexcepted count %d", len(vars))
		return
	}
	if vars[0] != "HTTP/1.1" {
		t.Errorf("variable args get fail got %q", vars)
		return
	}
}
