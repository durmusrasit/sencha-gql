package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	themes := Themes{}
	themes.ReadThemes()

	themeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Theme",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if theme, ok := p.Source.(*Theme); ok {
						return theme.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if theme, ok := p.Source.(*Theme); ok {
						return theme.Name, nil
					}
					return nil, nil
				},
			},
			"foreground": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if theme, ok := p.Source.(*Theme); ok {
						return theme.Foreground, nil
					}
					return nil, nil
				},
			},
			"background": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if theme, ok := p.Source.(*Theme); ok {
						return theme.Background, nil
					}
					return nil, nil
				},
			},
			"enabled": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if theme, ok := p.Source.(*Theme); ok {
						return theme.Enabled, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"themes": &graphql.Field{
				Type: graphql.NewList(themeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return themes.ListThemes()
				},
			},
			"theme": &graphql.Field{
				Type: themeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return themes.GetTheme(p.Args["id"].(int))
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatal(err)
	}

	query := `
		{
			themes {
				name, foreground, background
			}
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("An error occured while execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
