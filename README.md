# advent-of-code-2022
This is my first year of participating in the Advent of Code. I will be using this as a oppertunity to learn about the GO programming language.
In this read me I will keep a list of "Lessons learned" for each day of the event (Hopefully so that I can reflect on my progress on the 25th).

## Lessons learned

### Day 1

Today I learnt many things, but most importantly of what the ":" does when trying to use variables, see the following snippet.

```go
myVariable := 1 //Initilization

myVariable = 2 //Usage
```

Also that VSCode deletes any un-used imports when saving. Something that I found very annoying as I am always hitting the "ctrl + s" combo.

It is also good to track how I create my "projects/modules", see the following snippet.
```bash
go mod init parrentFolder/currentFolder
```
### Day 2

Today I made use of the "defer" keyword to close the input file. This makes it so that the command after it is executed at the end of the program, making it easier for me to remember closing the file.

Yesterday I made use of the following command to execute my programs:
```bash
go run main.go
```

Wich gave slow execution times comparing it to Albin Perssons implementation in COBOL. I forgot that I had to compile my program before running it!
Now me and Albins implementations are neck and neck when it comes to execution times.

Other new stuff is the usage of multi-dimensional arrays.

In general it feels like "smooth sailing" when it comes to writing simple programs in GO.