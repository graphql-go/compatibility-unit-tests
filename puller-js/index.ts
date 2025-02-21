import { readFileSync } from "fs";
import * as ts from "typescript";

function walk(sourceFile: ts.SourceFile) {
  innerWalk(sourceFile);

  function innerWalk(node: ts.Node) {
    switch (node.kind) {
        case ts.SyntaxKind.FunctionDeclaration:
          default:
            const n = node as any;
            if(n?.name?.escapedText === "describe") {
              console.log(n?.name?.escapedText);
            }
    }

    ts.forEachChild(node, innerWalk);
  }
}

const fileName = "schema-test.ts";

const sourceFile = ts.createSourceFile(
  fileName,
  readFileSync(fileName).toString(),
  ts.ScriptTarget.ES2015,
);

walk(sourceFile)
