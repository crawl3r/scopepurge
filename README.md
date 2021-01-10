# scopepurge

A (probably over complicated) effort at taking two lists and effectively removing list B from list A (if it exists). My use case for this was the following:
  
- Found a few hundred subdomains  
- Needed to ensure that the subdomains did NOT include certain items  
- Bash is hard and Go is not
  
## Installation  
  
```
go get github.com/crawl3r/scopepurge
```
  
## Usage
  
```
./scopepurge urls.txt bad.txt > cleaned_urls.txt
```
  
## License  
I'm just a simple skid. Licensing isn't a big issue to me, I post things that I find helpful online in the hope that others can:  
 A) learn from the code  
 B) find use with the code or   
 C) need to just have a laugh at something to make themselves feel better  
  
Either way, if this helped you - cool :)  
