package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	GIT "github.com/swiftwave-org/swiftwave/git_manager"
	"github.com/swiftwave-org/swiftwave/swiftwave_service/core"
	"github.com/swiftwave-org/swiftwave/swiftwave_service/graphql/model"
)

// GitBranches is the resolver for the gitBranches field.
func (r *queryResolver) GitBranches(ctx context.Context, input model.GitBranchesQueryInput) ([]string, error) {
	// Fetch git credential
	var gitCredential = &core.GitCredential{}
	if input.GitCredentialID > 0 {
		tx := r.ServiceManager.DbClient.First(&gitCredential, input.GitCredentialID)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	branches, err := GIT.FetchBranches(input.RepositoryURL, gitCredential.Username, gitCredential.Password, gitCredential.SshPrivateKey)
	return branches, err
}
