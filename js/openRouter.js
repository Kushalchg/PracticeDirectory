import OpenAI from "openai";
import { appendFile, readFileSync } from 'fs';
import { read, utils } from 'xlsx';
import path from 'path';
import { fileURLToPath } from 'url';
import { stringify } from "querystring";

const key = process.env.OPENROUTER_API_KEY
const openai = new OpenAI({
  baseURL: 'https://openrouter.ai/api/v1',
  apiKey: `${key}`,
});

async function CallApi(row) {
  const completion = await openai.chat.completions.create({
    model: 'qwen/qwen3-30b-a3b:free',
    messages: [
      {
        role: 'user',
        content: `
  You are nepali to english translator and you convert everything 
        I given in nepali into english, content may contain the person anme location and other stuff from country nepal, you just convert into english, also give the data in json
  just return the object not the array of object

  ${JSON.stringify(row)}

  `,
      },
    ],
    response_format: {
      "type": "json_schema",
      "json_schema": {
        "name": "weather",
        "strict": true,
        "schema": {
          "type": "object",
          "properties": {
            "data": {
              "type": "object(json) just object not the array of object just object not [] around ob",
              "description": "should contahin the prue json of the result"
            },
          },
          "required": ["data"],
          "additionalProperties": false
        }
      }
    }

  });

  console.log(completion);
  appendFile('./test.json', JSON.stringify(completion.choices[0].message.content) + '\n', err => {
    if (err) {
      console.error(err);
    } else {
      // file written successfully
    }
  });
}




// Resolve __dirname in ES modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Construct the full path to the Excel file
const filePath = path.join(__dirname, 'test.xlsx');

// Read the file as a buffer
const fileBuffer = readFileSync(filePath);

// Parse the buffer into a workbook
const workbook = read(fileBuffer);

// Extract sheet names
const sheets = workbook.SheetNames;

if (!sheets || sheets.length === 0) {
  console.error('No sheets found in the Excel file.');
  process.exit(1);
}


// Iterate over each sheet and extract data
for (let i = 0; i < sheets.length; i++) {
  const sheetName = sheets[i];
  const worksheet = workbook.Sheets[sheetName];
  const temp = utils.sheet_to_json(worksheet);

  for (let i = 0; i < 1; i++) {

    CallApi(temp[i])
  }
  //temp.forEach((row, index) => {
  //  //console.log(`Row ${index + 1} from sheet "${sheetName}":`, row);
  //});
}
