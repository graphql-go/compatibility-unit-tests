import { readFileSync } from "fs";
import * as ts from "typescript";

const fileName = "schema-test.ts";
const sourceFile = ts.createSourceFile(fileName, readFileSync(fileName).toString(), ts.ScriptTarget.ES2015);

function walk(node: ts.SourceFile | ts.Node) {
  const n = node as any;
  console.log(n?.name?.escapedText);

  node.forEachChild((subNode: ts.Node) => {
   walk(subNode);
  });
}

walk(sourceFile);