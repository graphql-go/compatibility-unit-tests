import { readFileSync } from "fs";
import * as ts from "typescript";
import { TestFile } from "./src/TestFile";
import { TestName } from "./src/TestName";
import { TestFileSaver } from "./src/TestFileSaver";
import * as fs from "fs";

const testName = new TestName();
const testFile = new TestFile();
const testFileSaver = new TestFileSaver();
const allTestsFileName = "unit-tests.txt";
const graphqlJSRootDir = "../repos/graphql-graphql-js";

const testFileResult = testFile.extractor(graphqlJSRootDir);
const testNameResult = testName.extractor(testFileResult);
testFileSaver.save(testNameResult, allTestsFileName);
