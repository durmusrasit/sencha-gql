import express, { Application } from "express";
import cors from "cors";
import { graphqlHTTP } from "express-graphql";
import { root } from "./resolver";
import { schema } from "./schema";

interface IExpressHandlerParams {
  port: number;
}

class ExpressHandler {
  private app: Application;

  constructor(params: IExpressHandlerParams) {
    this.app = express();
    this.app.use(cors());
    this.app.listen(params.port);

    console.log(`GraphQL server API running at ${params.port}/graphql`);
  }

  public getApp() {
    return this.app;
  }
}

const handler = new ExpressHandler({ port: 3000 });

const app = handler.getApp();

app.use(
  "/graphql",
  graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true,
  })
);
