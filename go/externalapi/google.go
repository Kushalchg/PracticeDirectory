package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

func GoogleApiCall(ctx context.Context) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GOOGLE_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		fmt.Println("Error occured in creating client::", err)
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		SystemInstruction: genai.NewContentFromText(
			`You are nepali to english translator and you convert everything 
      I given in nepali into english, content may contain the person anme location and other stuff from country nepal, you just convert into english`,
			genai.RoleUser),
	}

	res, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(` 

ताप्लेजुंग	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	तिल कुमार मेन्याङवो [लिम्बु]	पुरुष	57	नेपाल कम्युनिष्ट पार्टी (एमाले)	12772	निर्वाचित
ताप्लेजुंग	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं २	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	खगेन सिङ हाङगाम	पुरुष	40	नेपाली काँग्रेस	10496	निर्वाचित
पाँचथर	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	कमल प्रसाद जबेगु	पुरुष	64	नेपाल कम्युनिष्ट पार्टी (एकिकृत समाजबादी) 	16849	निर्वाचित
पाँचथर	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं २	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	ईन्द्र बहादुर आङवो	पुरुष	40	नेपाल कम्युनिष्ट पार्टी (माओवादी केन्द्र)	15685	निर्वाचित
इलाम	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं २	गोविन्द गिरी	पुरुष	62	नेपाली काँग्रेस	16658	निर्वाचित
इलाम	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	खिनु लङवा [लिम्बू]	महिला	52	नेपाल कम्युनिष्ट पार्टी (एकिकृत समाजबादी) 	13371	निर्वाचित
इलाम	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं २	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं २	राम बहादुर मगर	पुरुष	60	नेपाल कम्युनिष्ट पार्टी (एमाले)	13754	निर्वाचित
इलाम	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं २	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	शम्शेर राई	पुरुष	51	नेपाली काँग्रेस	12780	निर्वाचित
झापा	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं १	गोपाल तामाङ्ग	पुरुष	50	नेपाली काँग्रेस	23045	निर्वाचित
झापा	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं ५	हिक्मत कुमार कार्की	पुरुष	58	नेपाल कम्युनिष्ट पार्टी (एमाले)	22530	निर्वाचित
झापा	प्रदेश सभा सदस्य निर्वाचन क्षेत्र नं १	प्रतिनिधि सभा सदस्य निर्वाचन क्षेत्र नं ३	भुमी प्रसाद राजवंशी	पुरुष	57	नेपाली काँग्रेस	20770	निर्वाचित
`),
		config,
	)
	fmt.Println(res.Text())

	f, err := os.OpenFile("./google.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error while opeinig file::", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	value, err := w.WriteString(res.Text())
	fmt.Println(value)

	if err != nil {
		fmt.Println("error while writing into file::", err)
	}

	err = w.Flush()
	if err != nil {
		fmt.Println("error while flushing buffer:", err)
	}
}
