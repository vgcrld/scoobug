package main

import (
	"context"

	"github.com/openai/openai-go"
)

func main() {
	client := openai.NewClient() // defaults to OPENAI_API_KEY env var
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(`
here is a list of medications:

Aspirin 81 mg enteric-coated tablet  
Atorvastatin 80 mg tablet (Lipitor)  
Biotene Dry Mouth Oral Rinse  
Carboxymethylcellulose 0.5% dropperette (Refresh Plus)  
Clopidogrel 75 mg tablet (Plavix)  
Colchicine 0.6 mg tablet (Colcrys)  
Doxazosin 4 mg tablet (Cardura)  
Famotidine 40 mg tablet (Pepcid)  
Ferrous sulfate 325 mg (65 mg iron) tablet  
Finasteride 5 mg tablet (Proscar)  
Fluoxetine 20 mg capsule (Prozac)  
Folic acid 1 mg tablet (Folvite)  
Fluticasone propionate spray (Flonase)  
Furosemide 40 mg tablet (Lasix)  
Isosorbide mononitrate 30 mg 24 hr tablet (Imdur)  
Losartan 25 mg tablet (Cozaar)  
Metoprolol succinate XL 25 mg 24 hr tablet (Toprol-XL)  
Nitroglycerin 0.4 mg SL tablet (Nitrostat)  
Ocuvite oral (eye supplement)  
Pantoprazole 40 mg EC tablet (Protonix)  
Phenytoin 100 mg ER capsule (Dilantin)  
Polyethylene glycol 17 gram packet (Miralax)  
Potassium chloride 10 mEq CR tablet (Klor-Con M)  
Prednisone 5 mg tablet (Deltasone)  
Senna 8.6 mg tablet (Senokot)  
Sennosides-Docusate Sodium 8.6-50 mg tablet (Senokot-S)  
Silver Sulfadiazine 1% cream (Silvadene)  
Sodium Chloride 0.65% nasal spray (Ocean)

which of these can effect muscle strength and tone?
Also, which of these can cause dizzyness and lightheadedness?

			
`),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	println(chatCompletion.Choices[0].Message.Content)
}

/*

Certain medications in the list can potentially affect muscle strength and tone, as well as cause dizziness and lightheadedness. Here’s a general idea of which medications might have these effects:

### Affect Muscle Strength and Tone:

1. **Atorvastatin (Lipitor)**: Statins can, in some cases, cause muscle pain, weakness, or a condition known as rhabdomyolysis.
2. **Prednisone (Deltasone)**: Long-term use of corticosteroids can lead to muscle weakness.

### Cause Dizziness and Lightheadedness:

1. **Clopidogrel (Plavix)**: May cause dizziness as a side effect.
2. **Colchicine (Colcrys)**: Can cause dizziness.
3. **Doxazosin (Cardura)**: An alpha-blocker that can cause dizziness and lightheadedness, especially when changing positions (orthostatic hypotension).
4. **Furosemide (Lasix)**: A diuretic that may cause dizziness due to changes in blood pressure or electrolyte imbalances.
5. **Isosorbide mononitrate (Imdur)**: Nitrates can cause dizziness and lightheadedness due to blood vessel dilation.
6. **Losartan (Cozaar)**: An angiotensin receptor blocker that can cause dizziness, especially in the initial treatment period or due to blood pressure changes.
7. **Metoprolol succinate (Toprol-XL)**: A beta-blocker that can cause dizziness or lightheadedness, especially when starting the medication.
8. **Nitroglycerin (Nitrostat)**: Can cause dizziness and lightheadedness due to its vasodilatory effects, which may lead to a temporary drop in blood pressure.

Note that these are general associations and individual experiences may vary. It is important to consult a healthcare professional if you have concerns regarding side effects or interactions of medications.

*/
