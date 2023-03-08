import { buildSchema } from "graphql";

const schema = buildSchema(`
  input CreateThemeInput {
    themeName: String!
    backgroundColor: String!
    foregroundColor: String!
  }

  input UpdateThemeInput {
    themeName: String!
    backgroundColor: String!
    foregroundColor: String!
  }

  input DeleteThemeInput {
    themeName: String!
  }

  input GetThemeInput {
    themeName: String!
  }

  type Theme {
    themeName: String!
    backgroundColor: String!
    foregroundColor: String!
  }

  type Mutation {
    createTheme(input: CreateThemeInput!): Theme!
    updateTheme(input: UpdateThemeInput!): Boolean
    deleteTheme(input: DeleteThemeInput!): Boolean
  }

  type Query {
    getTheme(input: GetThemeInput!): Theme
    getThemes: [Theme]!
  }
`);

export { schema };
