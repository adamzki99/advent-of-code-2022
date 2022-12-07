# advent-of-code-2022
This is my first year of participating in the Advent of Code. I will be using this as an opportunity to learn about the GO programming language.
In this read me I will keep a list of "Lessons learned" for each day of the event (Hopefully so that I can reflect on my progress on the 25th).

## Lessons learned

### Day 7

Today was more usage of pointers! And even a throwback to a computer science lab regarding filesystems.
But today was nothing more that keeping everything in check when laying out the logic behind the program. My solution
might not have ended up as the most efficient one, but it works and I have my stars ;-).

### Day 6
Today was quite easy, so no new lessons learned :).

### Day 5

Today I made use of pointer in GO, a concept I have used before in languages like C/C++. Always a little scary using pointers, but using them in GO seems easy and fun! Hopefully it is something that will be more useful in coming challenges.

### Day 4
Today was nothing new, just smooth sailing and no wrong answers when submitting :)

### Day 3

Today was an interesting day. At first, my implementation got the wrong answer. Even if my logic was sound and the example worked fine, I wasn't able to produce the right answer. I resorted to writing tests, something that I found very easy to do in GO. This didn't find any problems. It was until I explained my solution to a friend that I noticed that there was a mistake when splitting the rugsack content. 

But the tests were able to help me find a problem with my implementation for part 2. So today's lesson was the following command:
```bash
go test
```

I also found it interesting that the type that you get from iterating over a string is "rune". Which holds the ASCII-value of the character according to the VSCode debugging tool.

### Day 2

Today I made use of the "defer" keyword to close the input file. This makes it so that the command after it is executed at the end of the program, making it easier for me to remember closing the file.

Yesterday I made use of the following command to execute my programs:
```bash
go run main.go
```

Which gave slow execution times comparing it to Albin Perssons implementation in COBOL. I forgot that I had to compile my program before running it!
Now me and Albins implementations are neck and neck when it comes to execution times.

Other new stuff is the usage of multidimensional arrays.

In general, it feels like "smooth sailing" when it comes to writing simple programs in GO.

### Day 1

Today I learnt many things, but most importantly of what the ":" does when trying to use variables, see the following snippet.

```go
myVariable := 1 //Initialization

myVariable = 2 //Usage
```

Also, that VSCode deletes any un-used imports when saving. Something that I found very annoying as I am always hitting the "ctrl + s" combo.

It is also good to track how I create my "projects/modules", see the following snippet.
```bash
go mod init parrentFolder/currentFolder
```
