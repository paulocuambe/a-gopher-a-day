To run this project just type `go run cat.go text.txt`.  

I built an executable using `go build cat.go`, so now if you run `./cat FILENAME` where you point to the file name you want to cat, you will also see the output in the terminal.

You can also provide a path to a file and the command will also work.  

Minimal exception handling was done, so you should we some helpful messages if the file you mentioned does not exist or if there's a problem reading the contents of the file.  

You should also see an error message when you do not specify a filename.

Example:
```bash
./cat text.txt

#output
I love Jesus
I wrote this is less than 10 min
Forgive me, my go skills are still very basic

This should be streamed and not loaded all at once
```
