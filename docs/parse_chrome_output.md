1. Build tables:
  * Search: ^(?!\|)(.+?):($
  * Replace: |$1|$2|

2. Remove Source:
  * Search: view source\n|view URL encoded\n
  * Replace: 
  
3. Set Headers
  * Search: ^(?![\|#])(.+)
  * Replace: \n### $1\n
  
4. Add table headers
  * Search: \n\n(?!\|Key|#)(.+)
  * Replace: \n\n|Key|Value|\n|---|-----|\n$1
