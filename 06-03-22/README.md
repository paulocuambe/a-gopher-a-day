This program executes native Unix commands via Go, using the `os/exec` package.  
In this basic experiment, we run the `ls` and `cat` commands.  
I found it entertaining to use a different way to run the `go version` command inside the program so I did it using the `os/exec.Cmd` struct other than the provided builder that is mostly used and well documented. 

Run the project with `go run .`