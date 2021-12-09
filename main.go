package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	aws "github.com/pulumi/pulumi-aws/provider/v4"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	vault "github.com/pulumi/pulumi-vault/provider/v4"
	"github.com/clstokes/pulumi-generate-import-file/pkg/state"
)

func main() {

	// Get type mapping of `provider_module_resource` to `provider:module/resource:Resource`
	terraformToPulumiTypeMapping := make(map[string]string)

	vaultProvider := vault.Provider()
	getTypeMapping(terraformToPulumiTypeMapping, vaultProvider)
	// repeat the getTypeMapping() to handle additional providers - e.g.
	awsProvider := aws.Provider()
	getTypeMapping(terraformToPulumiTypeMapping, awsProvider)

	// Parse terraform state file
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Missing argument - the path to a terraform state file must be provided as the first argument")
		os.Exit(1)
	}
	importFromStateFile := os.Args[1]

	terraformResources, err := parseTerraformState(importFromStateFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	// Create file to `pulumi import --file ...` from
	pulumiImportMapping := make([]pulumiFileResource, 0)
	for _, tResource := range terraformResources.Resources {
		for _, tResourceInstance := range tResource.Instances {
			// fmt.Println(fmt.Sprintf("Resource found [%s:%s:%v]...", tResource.Type, tResource.Name, tResourceInstance.IndexKey))

			foundType := terraformToPulumiTypeMapping[tResource.Type]
			// use indexKey '0' if one does not exist
			foundIndexKey := tResourceInstance.IndexKey
			if foundIndexKey == nil {
				foundIndexKey = 0
			}
			if foundType != "" {
				// Have to adjust the resource name due to https://github.com/pulumi/pulumi/issues/6032
				// This produces a name in the format {module_name}_{resource_type}_{resource_name}_{index}
				pResourceVariableName := fmt.Sprintf("%s_%s_%s%v", strings.TrimPrefix(tResource.Module, "module."), tResource.Type, tResource.Name, foundIndexKey)
				pResource := pulumiFileResource{
					Type: foundType,
					Name: pResourceVariableName,
					ID:   strings.ToLower(tResourceInstance.AttributesFlat["id"]),
				}
				pulumiImportMapping = append(pulumiImportMapping, pResource)
			} else {
				fmt.Fprintln(os.Stderr, fmt.Sprintf("Omitting Terraform resource [%s:%s:%v] from import output. No mapping found. Do you need to add this provider to the getTypeMapping() implementation.", tResource.Type, tResource.Name, tResourceInstance.IndexKey))
			}
		}
	}

	pulumiFile := pulumiFile{
		Resources: pulumiImportMapping,
	}

	prettyPrintJSON(pulumiFile)
}

func getTypeMapping(mapToAddTo map[string]string, provider tfbridge.ProviderInfo) {
	for key, element := range provider.Resources {
		pulumiType := element.Tok
		mapToAddTo[key] = string(pulumiType)
	}
}

func parseTerraformState(importFromStateFile string) (*tfstate.StateV4, error) {
	terraformState, err := ioutil.ReadFile(importFromStateFile)
	if err != nil {
		return nil, err
	}

	err = tfstate.CheckTerraformStateVersion(terraformState)
	if err != nil {
		return nil, err
	}

	var terraformResources tfstate.StateV4
	err = json.Unmarshal(terraformState, &terraformResources)

	return &terraformResources, nil
}

func prettyPrintJSON(object interface{}) {
	jsonData, err := json.Marshal(object)
	if err != nil {
		fmt.Fprintln(os.Stderr, "JSON parse error: ", err)
		os.Exit(1)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, jsonData, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "JSON pretty indent error: ", err)
		os.Exit(1)
	}

	fmt.Println(prettyJSON.String())
}

type pulumiFile struct {
	Resources []pulumiFileResource `json:"resources"`
}

type pulumiFileResource struct {
	Type string `json:"type"`
	Name string `json:"name"`
	ID   string `json:"id"`
}
