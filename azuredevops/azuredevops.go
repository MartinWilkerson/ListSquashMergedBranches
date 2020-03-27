package azuredevops

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Project is an azure devops project
type Project struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

// Repository is an azure devops repository
type Repository struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	URL     string  `json:"url"`
	Project Project `json:"project"`
}

// User is an azure devops user
type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	UniqueName  string `json:"uniqueName"`
	URL         string `json:"url"`
	ImageURL    string `json:"imageUrl"`
}

// Commit is a commit in an azure devops git repository
type Commit struct {
	CommitID string `json:"commitId"`
	URL      string `json:"url"`
}

// Reviewer is an azure devops user assigne do be a reviewer of a pull request
type Reviewer struct {
	User        `json:"user"`
	ReviewerURL string `json:"reviewerUrl"`
	Vote        int    `json:"vote"`
}

// PullRequest is an azure devops pull request
type PullRequest struct {
	Repository            Repository `json:"repository"`
	PullRequestID         int        `json:"pullRequestId"`
	CodeReviewID          int        `json:"codeReviewId"`
	Status                string     `json:"status"`
	CreatedBy             User       `json:"createdBy"`
	CreationDate          time.Time  `json:"creationDate"`
	Title                 string     `json:"title"`
	Description           string     `json:"description"`
	SourceRefName         string     `json:"sourceRefName"`
	TargetRefName         string     `json:"targetRefName"`
	MergeStatus           string     `json:"mergeStatus"`
	MergeID               string     `json:"mergeId"`
	LastMergeSourceCommit Commit     `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit Commit     `json:"lastMergeTargetCommit"`
	LastMergeCommit       Commit     `json:"lastMergeCommit"`
	Reviewers             []Reviewer `json:"reviewers"`
	URL                   string     `json:"url"`
	SupportsIterations    bool       `json:"supportsIterations"`
}

type pullRequestResponse struct {
	Value []PullRequest `json:"value"`
}

// GetPullRequests returns all pull requests for the specified project
func GetPullRequests(organisation string, project string, personalAccessToken string) []PullRequest {
	url := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/pullrequests?searchCriteria.status=all&api-version=5.1", organisation, project)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(":%s", personalAccessToken))))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(response.Body)
	pullRequestResponse := pullRequestResponse{}
	decoder.Decode(&pullRequestResponse)

	return pullRequestResponse.Value
}

type Ref struct {
	Name     string `json:"Name"`
	ObjectID string `json:"objectId"`
	Creator  User   `json:"creator"`
}

func GetRefs(organisation string, project string, apikey string) {

}
