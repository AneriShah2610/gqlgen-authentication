//go:generate gorunpkg github.com/99designs/gqlgen

package gqlgen_authentication

import (
	context "context"
	"fmt"
	"math/rand"

	"github.com/aneri/gqlgen-authentication/dal"
	"github.com/aneri/gqlgen-authentication/models"
)

var ctxt context.Context

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateJob(ctx context.Context, input NewJob) (Job, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	jobs := models.Job{
		ID:           fmt.Sprintf("T%d", rand.Int()),
		Name:         input.Name,
		Description:  input.Description,
		JoinLocation: input.JoinLocation,
		CreatedBy:    input.CreatedByID,
	}
	crConn.Db.Create(&jobs)
	return Job{}, nil
}
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateJob(ctx context.Context, id string, input UpdateJob) (Job, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateApplication(ctx context.Context, input NewApplication) (Application, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteApplication(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Jobs(ctx context.Context) ([]Job, error) {
	panic("not implemented")
}
func (r *queryResolver) Applications(ctx context.Context, jobID string) ([]Application, error) {
	panic("not implemented")
}

// func (r *jobResolver) CreatedBy(ctx context.Context, obj *dal.Job) (dal.User, error) {

// 	user, err := userRepository.GetByID(obj.CreatedBy)
// 	if err != nil {
// 		fmt.Printf("Error Fetching user", err)
// 		return dal.User{}, err
// 	}
// 	return user, err
// }
