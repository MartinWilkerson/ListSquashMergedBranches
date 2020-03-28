# List non-deleted squash merged branches in Azure Devops

Ever found yourself stuck on a project that uses squash merges? Found that often branches don't get deleted when merging? If you're using Azure Devops, this project will list the branches that were the source of a completed pull request that have not been deleted and still point to the same commit as when the pull request was completed.

## Usage

```bash
list-squashed-merges -organisation <organisation name> -projectName<project name> -apikey <personal access token with read access to code>
```

| Argument | Description |
| -------- | ----------- |
| `-organisation` | The name of the organisation as it appears in the devops url |
| `-projectName` | The name of the project as it appears in the devops url |
| `-apikey` | A personal access token for Azure Devops that has 'read code' access |
