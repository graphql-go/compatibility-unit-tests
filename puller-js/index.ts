import { readFileSync } from "fs";
import * as ts from "typescript";

const fileName = "schema-test.ts";
const sourceFile = ts.createSourceFile(fileName, readFileSync(fileName).toString(), ts.ScriptTarget.ES2015);


const tests = [] as any;

function walk(node: ts.SourceFile | ts.Node) {
  const n = node as any;

  if(n?.kind === ts.SyntaxKind.CallExpression) {
    if(n?.arguments && n?.arguments.length) {
      if(n?.arguments[0].text) {
        if(n?.expression?.escapedText === "describe" || n?.expression?.escapedText === "it") {
          const testName = n?.arguments[0].text;
          tests.push(testName);
        }
      }
    }
  }

  node.forEachChild((subNode: ts.Node) => {
   walk(subNode);
  });
}

walk(sourceFile);
