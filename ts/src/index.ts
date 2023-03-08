import express, { Application } from "express";

interface IExpressHandlerParams {
  port: number;
}

class ExpressHandler {
  private app: Application;

  constructor(params: IExpressHandlerParams) {
    this.app = express();
    this.app.listen(params.port);
  }

  public getApp() {
    return this.app;
  }
}

const handler = new ExpressHandler({ port: 3000 });
