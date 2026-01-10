const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
    this.data = {};
  }

  async readData() {
    try {
      const content = await fs.promises.readFile(this.filePath, 'utf8');
      this.data = JSON.parse(content);
    } catch (error) {
      console.error(`Error reading file: ${error.message}`);
    }
  }

  async writeData() {
    try {
      const content = JSON.stringify(this.data, null, 2);
      await fs.promises.writeFile(this.filePath, content);
    } catch (error) {
      console.error(`Error writing file: ${error.message}`);
    }
  }

  getData() {
    return this.data;
  }

  setData(data) {
    this.data = data;
  }
}

module.exports = Parser;