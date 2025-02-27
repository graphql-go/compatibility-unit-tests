import * as ts from "typescript";
import { readFileSync } from "fs";
import { camelCase, upperFirst } from "lodash";

const tests = [] as any;

export class TestName {
  constructor() {}

  extractor(files: string[]) {
    const result = [];

    for (let i = 0; i < files.length; i++) {
      const fileName = files[i];

      const check = fileName.includes("__tests__");
      if (!check) {
        continue;
      }

      const sourceFile = ts.createSourceFile(
        fileName,
        readFileSync(fileName).toString(),
        ts.ScriptTarget.ES2015,
      );

      const testNames = this.walk(sourceFile);

      result.push(...testNames);
    }

    return result;
  }

  private isTestNode(node: ts.Node) {
    const n = node as any;

    if (n?.kind !== ts.SyntaxKind.CallExpression) {
      return false;
    }

    if (!n?.arguments && !n?.arguments.length) {
      return false;
    }

    if (n?.arguments[0] && !n?.arguments[0].text) {
      return false;
    }
    if (
      n?.expression?.escapedText !== "describe" &&
      n?.expression?.escapedText !== "it"
    ) {
      return false;
    }

    return true;
  }

  walk(node: ts.SourceFile | ts.Node) {
    node.forEachChild((subNode: ts.Node) => {
      const n = subNode as any;
      const isTestNodeCheck = this.isTestNode(subNode);
      if (this.isTestNode(subNode)) {
        const testName = n?.arguments[0].text;
        tests.push(testName);
      }
      this.walk(subNode);
    });

    const result = [];
    for (let i = 0; i < tests.length; i++) {
      const testName = upperFirst(camelCase(tests[i]));
      result.push(testName);
    }

    return result;
  }
}
