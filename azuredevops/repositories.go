package azuredevops

import (
	"encoding/json"
	"fmt"
)

// Repository represented a repo in azure devops
type Repository struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	URL       string  `json:"url"`
	Project   Project `json:"project"`
	RemoteURL string  `json:"remoteUrl"`
}

type repositoryResponse struct {
	count int
	value []Repository
}

// GetRepositories returns a list of repositories in the provided organisation and project
func GetRepositories(organisation string, project string, personalAccessToken string) []Repository {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories?api-version=5.1", organisation, project)

	response := makeRequest(url, personalAccessToken)

	decoder := json.NewDecoder(response.Body)
	repositoriesResponse := repositoryResponse{}
	decoder.Decode(&repositoriesResponse)

	return repositoriesResponse.value
}
