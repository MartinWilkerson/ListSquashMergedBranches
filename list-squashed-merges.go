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

	refs := make(map[string][]azuredevops.Ref)
	repoDict := make(map[string]azuredevops.Repository)
	fmt.Println("Repositories:")
	for _, repo := range repos {
		fmt.Printf("%s:\n", repo.Name)
		repoDict[repo.ID] = repo
		repoRefs := azuredevops.GetRefs(*organisationPtr, *projectNamePtr, repo.ID, *apiKeyPtr)
		refs[repo.ID] = repoRefs
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

	for _, pr := range prs {
		if pr.MergeStatus != "succeeded" {
			continue
		}
		sourceRepositoryID := pr.Repository.ID
		sourceCommit := pr.LastMergeSourceCommit.CommitID
		sourceBranch := pr.SourceRefName

		repoRefs := refs[sourceRepositoryID]

		for _, ref := range repoRefs {
			if ref.Name == sourceBranch && ref.ObjectID == sourceCommit {
				fmt.Printf("Repository: %s, Branch: %s", repoDict[sourceRepositoryID].Name, sourceBranch)
			}
		}
	}
}

/*
Example output with a single still-existing branch with a PR that squash-merged it:

Repositories:
squash-merge-test:
36ce32ad4187f2d5f4d6c88acabf53ee3fdbdebb refs/heads/do-something
d8d74bb41d701d2b2167f32c31115e678c7531f8 refs/heads/master

Pull requests:
Source: [refs/heads/do-something] 36ce32ad4187f2d5f4d6c88acabf53ee3fdbdebb

Example output with a single deleted branch with a PR that squash-merged it:
*/
