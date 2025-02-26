import * as fs from "fs";

export class TestFile {
  constructor() {}

  extractor(rootDirName: string) {
    const files = [] as any;
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

    walkDir(rootDirName);

    return files;
  }
}
