package projects

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

func TestCreateSuccessful(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "POST")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		testhelper.TestJSONRequest(t, r, `{
		  "project": {
		    "description": "My new project",
		    "domain_id": "default",
		    "enabled": true,
		    "is_domain": true,
		    "name": "myNewProject"
		  }
		}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
            "project": {
                "description": "My new project",
                "domain_id": "default",
                "enabled": true,
                "is_domain": false,
                "name": "myNewProject",
                "id": "1234567",
                "links": {
                    "self": "http://os.test.com/v3/identity/projects/1234567"
                }
            }
        }`)
	})

	project := &CreateOpts{
		IsDomain:    true,
		Description: "My new project",
		DomainID:    "default",
		Enabled:     true,
		Name:        "myNewProject",
	}

	result, err := Create(client.ServiceClient(), project).Extract()
	if err != nil {
		t.Fatalf("Unexpected error from Create: %v", err)
	}

	if result.Description != "My new project" {
		t.Errorf("Project description was unexpected [%s]", result.Description)
	}
}
