package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	aws "github.com/pulumi/pulumi-aws/provider/v4"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	vault "github.com/pulumi/pulumi-vault/provider/v4"
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
		fmt.Println("Missing argument - the path to a terraform state file must be provided as the first argument")
		os.Exit(1)
	}
	importFromStateFile := os.Args[1]

	terraformResources, err := parseTerraformState(importFromStateFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Create file to `pulumi import --file ...` from
	pulumiImportMapping := make([]pulumiFileResource, 0)
	for _, tResource := range terraformResources.Resources {
		for _, tResourceInstance := range tResource.Instances {
			// fmt.Println(fmt.Sprintf("Resource found [%s:%s:%v]...", tResource.Type, tResource.Name, tResourceInstance.IndexKey))

			pResource := pulumiFileResource{
				Type: terraformToPulumiTypeMapping[tResource.Type],
				ID:   tResourceInstance.AttributesFlat["id"],
				// Have to adjust the resource name due to https://github.com/pulumi/pulumi/issues/6032
				Name: fmt.Sprintf("%s_%s%v", tResource.Type, tResource.Name, tResourceInstance.IndexKey),
			}
			pulumiImportMapping = append(pulumiImportMapping, pResource)
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

func parseTerraformState(importFromStateFile string) (*stateV4, error) {
	terraformState, err := ioutil.ReadFile(importFromStateFile)
	if err != nil {
		return nil, err
	}

	err = checkTerraformStateVersion(terraformState)
	if err != nil {
		return nil, err
	}

	var terraformResources stateV4
	err = json.Unmarshal(terraformState, &terraformResources)

	return &terraformResources, nil
}

func prettyPrintJSON(object interface{}) {
	jsonData, err := json.Marshal(object)
	if err != nil {
		fmt.Println("JSON parse error: ", err)
		os.Exit(1)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, jsonData, "", "    ")
	if err != nil {
		fmt.Println("JSON pretty indent error: ", err)
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
