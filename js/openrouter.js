const https = require('https');

const key = process.env.OPENROUTER_API_KEY;

const data = JSON.stringify({
  "model": "qwen/qwen3-30b-a3b:free",
  "messages": [
    {
      "role": "user",
      "content": `
You are nepali to english translator and you convert everything 
      I given in nepali into english, content may contain the person anme location and other stuff from country nepal, you just convert into english, also give the data in json\n

ताप्लेजुंग	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	तिल कुमार मेन्याङवो [लिम्बु]	पुरुष	57	नेपाल कम्युनिष्ट पार्टी (एमाले)	12772	निर्वाचित
ताप्लेजुंग	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं २	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	खगेन सिङ हाङगाम	पुरुष	40	नेपाली काँग्रेस	10496	निर्वाचित
पाँचथर	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	कमल प्रसाद जबेगु	पुरुष	64	नेपाल कम्युनिष्ट पार्टी (एकिकृत समाजबादी) 	16849	निर्वाचित

 ` }
  ],
  "response_format": {
    "type": "json_schema",
    "json_schema": {
      "name": "weather",
      "strict": true,
      "schema": {
        "type": "object",
        "properties": {
          "location": {
            "type": "string",
            "description": "City or location name"
          },
          "temperature": {
            "type": "number",
            "description": "Temperature in Celsius"
          },
          "conditions": {
            "type": "string",
            "description": "Weather conditions description"
          }
        },
        "required": ["location", "temperature", "conditions"],
        "additionalProperties": false
      }
    }
  }
}
);

const options = {
  hostname: 'https://openrouter.ai/api/v1/completions',
  path: '',
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${key}`
  }
};

const req = https.request(options, res => {
  let responseData;

  res.on('data', chunk => {
    responseData += chunk;
  });

  res.on('end', () => {
    console.log('Response:', JSON.parse(responseData));
  });
});

req.on('error', error => {
  console.error('Errorwhile:', error.message);
});

// Send the request body
req.write(data);
req.end();
