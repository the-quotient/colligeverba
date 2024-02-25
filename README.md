# ColligeVerba

ColligeVerba is a simple and blazingly fast Online-Tool that assists you while transcribing handwritten latin source texts searching latin word forms using Regular Expressions. Visit [cv.jn2p.de](https://cv.jn2p.de) and try it. 

## When to Use It 

Transcribing handwritten latin source texts you may encounter words that are very difficult to read. Especially if the the first letter is indecipherable, it is almost impossible to look for possible candidates in a normal dictionary.

## How to Use It 

Vist [cv.jn2p.de](https://cv.jn2p.de). You can enter the letters that you can read and mark the others with a question mark. If you have a suggestions for the missing letters, you can add several candidates for the question marks. The tool then returns words and also inflected forms that match the entered pattern. 

If at any point you are unsure whether a sign represents one or two letters you tell this to the program with two question marks.

In order to accelerate the process of checking wether a form makes sense in the context of the source, the tool also provides some translation suggestions and a form analysis powered by [latein.me](https://latein.me). 

## Remarks on the Implementation

### Engine 

The engine (implemented in Go) builds a regular expression from your input and matches it against a list of latin word forms. The form analysis and translation suggestions is powered by [latein.me](https://latein.me)

### Form list 

The basic form list is extracted from a spell check dictionary for Vim (ca. 890,000 forms). We expanded this list with the forms of the [Classical Language Toolkit](https://github.com/cltk) project corpus (now ca. 1,098,000 forms). Currently, we are working on expanding the list further with common spelling deviation of Medieval and Early Modern (Ecclesiastical) Latin. 

### Web Interface  

The web interface uses a combination of HTMX and Go. 


