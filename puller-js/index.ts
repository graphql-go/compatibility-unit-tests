import { readFileSync } from "fs";
import * as ts from "typescript";
import { camelCase, upperFirst } from "lodash";
import * as fs from "fs";

const fileName = "schema-test.ts";

class TestFiles {
  constructor() {}

  extractor() {
    const files = [] as any;
    const rootDir = "../repos/graphql-graphql-js";
    const walkDir = (dirName: any) => {
      const dirNames = fs.readdirSync(dirName, { withFileTypes: true });

      for (let i = 0; i < dirNames.length; i++) {
        const item = dirNames[i];
        const filePath = `${item.path}/${item.name}`;

        if (item.name.startsWith(".")) {
          continue;
        }

        if (!fs.existsSync(filePath)) {
          continue;
        }

        if (item.isDirectory()) {
          walkDir(filePath);
          continue;
        }

        files.push(filePath);
      }
    };

    walkDir(rootDir);

    return files;
  }
}

const tests = [] as any;

class TestNames {
  constructor() {}

  extractor(files: string[]) {
    const result = [];

    for (let i = 0; i < 100; i++) {
      const fileName0 = files[i];
      const fileName = "schema-test.ts";

      const check = fileName0.includes("__tests__");
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

  walk(node: ts.SourceFile | ts.Node) {
    const n = node as any;

    if (n?.kind === ts.SyntaxKind.CallExpression) {
      if (n?.arguments && n?.arguments.length) {
        if (n?.arguments[0].text) {
          if (
            n?.expression?.escapedText === "describe" ||
            n?.expression?.escapedText === "it"
          ) {
            const testName = n?.arguments[0].text;
            tests.push(testName);
          }
        }
      }
    }

    node.forEachChild((subNode: ts.Node) => {
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

const testNames = new TestNames();
const testFiles = new TestFiles();

const testFilesResult = testFiles.extractor();
const testNamesResult = testNames.extractor(testFilesResult);

let result = "";
for (let i = 0; i < testNamesResult.length; i++) {
  const testName = testNamesResult[i];
  result += `${testName}\n`;
}

fs.writeFileSync("unit-tests.txt", result);
