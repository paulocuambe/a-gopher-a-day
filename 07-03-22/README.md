Do you want to know details about a user in your operating system?  
This program helps you do that.

Run `./user` or `go run .` and answer the prompted question with a username in your OS.  

If you are out of ideas, you can try these common usernames in Unix systems `root, bin, daemon, backup` or you can just type in your username.

The output looks like this:
```bash
# username input: bin
User
        Uid:    2 
        Gid:    2 
        Name:   bin 
        Username:       bin 
        HomeDir:        /bin

# username input: root
User
        Uid:    0 
        Gid:    0 
        Name:   root 
        Username:       root 
        HomeDir:        /root

# username input: daemon
User
        Uid:    1 
        Gid:    1 
        Name:   daemon 
        Username:       daemon 
        HomeDir:        /usr/sbin

# username input: myusername
User
        Uid:    1000 
        Gid:    1000 
        Name:   Paulo 
        Username:       gosenx 
        HomeDir:        /home/gosenx
```