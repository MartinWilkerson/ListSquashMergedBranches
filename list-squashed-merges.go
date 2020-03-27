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

	fmt.Println("Repositories:")
	for _, repo := range repos {
		fmt.Printf("%s:\n", repo.Name)
		repoRefs := azuredevops.GetRefs(*organisationPtr, *projectNamePtr, repo.ID, *apiKeyPtr)
		for _, ref := range repoRefs {
			fmt.Printf("%s %s\n", ref.ObjectID, ref.Name)
		}
	}
	fmt.Println()

	fmt.Println("Pull requests:")
	for _, pr := range prs {
		if pr.MergeStatus != "succeeded" {
			continue
		}
		sourceCommit := pr.LastMergeSourceCommit.CommitID
		sourceBranch := pr.SourceRefName
		fmt.Printf("Source: [%s] %s\n", sourceBranch, sourceCommit)
	}

}
