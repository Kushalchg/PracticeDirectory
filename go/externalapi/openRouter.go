package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func OpenRouter() {
	key := os.Getenv("OPENROUTER_API_KEY")
	url := "https://openrouter.ai/api/v1/completions"
	body := map[string]string{
		"model": "qwen/qwen3-30b-a3b:free",
		"prompt": `
			You are nepali to english translator and you convert everything 
      I given in nepali into english, content may contain the person anme location and other stuff from country nepal, you just convert into english, also give the data in json\n

ताप्लेजुंग	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	तिल कुमार मेन्याङवो [लिम्बु]	पुरुष	57	नेपाल कम्युनिष्ट पार्टी (एमाले)	12772	निर्वाचित
ताप्लेजुंग	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं २	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	खगेन सिङ हाङगाम	पुरुष	40	नेपाली काँग्रेस	10496	निर्वाचित
पाँचथर	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	कमल प्रसाद जबेगु	पुरुष	64	नेपाल कम्युनिष्ट पार्टी (एकिकृत समाजबादी) 	16849	निर्वाचित


    `,
	}
	josnBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error while converting into json::")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(josnBody))
	if err != nil {
		fmt.Println("Error while converting into json::")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", key))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while calling api::")
	}
	defer resp.Body.Close()

	// Read and decode the JSON response
	var response any
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Use the parsed JSON
	fmt.Println("Name:", response)
}
