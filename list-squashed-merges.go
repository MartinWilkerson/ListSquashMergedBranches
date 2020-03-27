package main

import (
	"flag"
	"fmt"

	"github.com/MartinWilkerson/list-squashed-merges/azuredevops"
)

func main() {
	organisationPtr := flag.String("organisation", "", "The azure devops organisation to check")
	projectNamePtr := flag.String("projectname", "", "The azure devops project to check")
	apiKeyPtr := flag.String("apikey", "", "API key to use")

	flag.Parse()

	repos := azuredevops.GetRepositories(*organisationPtr, *projectNamePtr, *apiKeyPtr)
	prs := azuredevops.GetPullRequests(*organisationPtr, *projectNamePtr, *apiKeyPtr)

	for _, repo := range repos {
		repoRefs := azuredevops.GetRefs(*organisationPtr, *projectNamePtr, repo.ID, *apiKeyPtr)
		for _, ref := range repoRefs {
			fmt.Printf("%s %s", ref.ObjectID, ref.Name)
		}
	}

	for _, pr := range prs {
		if pr.MergeStatus != "succeeded" {
			continue
		}
		sourceCommit := pr.LastMergeSourceCommit.CommitID
		sourceBranch := pr.SourceRefName
		fmt.Printf("Source: [%s] %s\n", sourceBranch, sourceCommit)
	}

}
