package mutations

import (
	"github.com/graphql-go/graphql"
	fields "github.com/smook1980/medialocker/app/server/mutations/fields"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateNotTodo,
	},
})
