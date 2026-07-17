# ColligeVerba

ColligeVerba is a simple Online-Tool that assists you while transcribing handwritten Latin and historical German source texts by searching for word forms using Regular Expressions.

While transcribing handwritten source texts, you may encounter words that are very difficult to read.
Especially if the first letter is indecipherable, it is almost impossible to look for possible candidates in a standard physical or digital dictionary.
ColligeVerba bridges the gap between transcription and dictionary lookup by letting you search using only the letters you can actually read.

## How to Use It

Visit [colligeverba.jn2p.de](https://colligeverba.jn2p.de). 

1. **Select your language and dictionary:** You can choose between Latin and German. For German, the tool dynamically fetches a comprehensive list of historical dictionaries (like the *Grimmsches Wörterbuch*, *Mittelhochdeutsches Wörterbuch*, etc.) categorized by epoch from [woerterbuchnetz.de](https://woerterbuchnetz.de).
2. **Enter your query:** Enter the letters that you can read and mark the unknown ones with a question mark (`?`).
3. **Refine with candidates:** If you have a suggestion for a missing letter, you can add several possible character candidates for the question marks using the "or" button in the UI.
4. **Search:** The tool then returns words and (momentarily only in the case of Latin) inflected forms that match the entered pattern.

If at any point you are unsure whether a sign represents one or two letters, you can tell the program this by using two question marks (`??`).

**Exploring Results:**
* **Latin:** To accelerate the process of checking whether a form makes sense in the context of the source, clicking a Latin result provides translation suggestions and a form analysis powered by [latein.me](https://latein.me). 
* **German:** Clicking a German result will open a new tab taking you directly to that exact w:ord's entry in your selected dictionary on the [woertbuchnetz.de](https://woerterbuchnetz.de).
