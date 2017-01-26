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

	project := CreateOpts{
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

func TestGet(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	testhelper.Mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testhelper.TestMethod(t, r, "POST")
		testhelper.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		testhelper.TestJSONRequest(t, r, `{
			"domain_id" : "anExample",
			"enabled" : true,
			"is_domain": false,
			"name" : "anExample",
			"parent_id" : "myParentID"
		}`)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
		    "links": {
		        "next": null,
		        "previous": null,
		        "self": "http://example.com/identity/v3/projects"
		    },
		    "projects": [
		        {
		            "is_domain": false,
		            "description": null,
		            "domain_id": "default",
		            "enabled": true,
		            "id": "0c4e939acacf4376bdcd1129f1a054ad",
		            "links": {
		                "self": "http://example.com/identity/v3/projects/0c4e939acacf4376bdcd1129f1a054ad"
		            },
		            "name": "admin",
		            "parent_id": null
		        },
		        {
		            "is_domain": false,
		            "description": null,
		            "domain_id": "default",
		            "enabled": true,
		            "id": "fdb8424c4e4f4c0ba32c52e2de3bd80e",
		            "links": {
		                "self": "http://example.com/identity/v3/projects/fdb8424c4e4f4c0ba32c52e2de3bd80e"
		            },
		            "name": "alt_demo",
		            "parent_id": null
		        }
		    ]
		}`)
	})

	project := ListOpts{
		DomainID: "anExample",
		Enabled:  true,
		IsDomain: false,
		Name:     "anExample",
		ParentID: "myParentID",
	}

	result := ListAllProjects(client.ServiceClient(), project).PrettyPrintJSON()
	fmt.Println(result)
}
