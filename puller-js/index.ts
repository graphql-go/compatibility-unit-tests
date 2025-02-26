import { readFileSync } from "fs";
import * as ts from "typescript";
import { TestFile } from "./src/TestFile";
import { TestName } from "./src/TestName";
import * as fs from "fs";

const fileName = "schema-test.ts";

const testNames = new TestName();
const testFiles = new TestFile();

const testFilesResult = testFiles.extractor();
const testNamesResult = testNames.extractor(testFilesResult);

let result = "";
for (let i = 0; i < testNamesResult.length; i++) {
  const testName = testNamesResult[i];
  result += `${testName}\n`;
}

fs.writeFileSync("unit-tests.txt", result);
