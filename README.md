# Go-Lang-User-Authentication

# Project Structure
```
$GOPATH/
    └── src/
        └── github.com/
            └── denyshubh/
                └── user-auth-app/
                    ├── cmd/
                    │   └── server/
                    │       └── main.go (creates connection to database and call graphQL endpoint)
                    ├── internal/
                    │   ├── model/
                    │   │   └── user.go (User and Authentication Models)
                    │   │  
                    │   ├── service/
                    │   │   └── email.go (When a user is created, send an email - THIS IS DONE USING GO-ROUTINE)
                    │   └── graph/
                    │       ├── gqlgen.yml
                    │       ├── schema.graphql
                    │       ├── generated.go
                    │       ├── model.go
                    │       └── resolver.go, schema.resolver.go
                    └── go.mod
```

## GraphQL Tool Used
- I have used gqlgen (a Go library) for building GraphQL servers. Passed gqlgen.yml file along with schema.graphql to generate the model, resolver etc.

## About Go Routine. 
We can use go-routine in this example to send email to users, once their account is successfully created. We can implement something like below. 

```go
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	// ... (existing code)

	err = r.db.Transaction(func(tx *gorm.DB) error {
		// ... (existing code)
	})

	if err != nil {
		return nil, err
	}

	// Send a welcome email asynchronously using a Goroutine
	go sendWelcomeEmail(&user)

	return &user, nil
}
```