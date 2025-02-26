export class TestFileSaver {
  save(testNamesResult: any[], fileName: string) {
    let result = "";
    for (let i = 0; i < testNamesResult.length; i++) {
      const testName = testNamesResult[i];
      result += `${testName}\n`;
    }

    fs.writeFileSync(fileName, result);
  }
}
