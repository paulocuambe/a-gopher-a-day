This time I try to develop the `ls` unix command, with a small caveat. **It is recursive**.  
The recursion stops at 4 levels deep by default, but **you can customize that behavior by setting the depth value in the environment variable `DIR_DEPTH`**.

**Example**:
```bash
# reminder that you can go as dee as you want or as low as you want
export DIR_DEPTH=2
```

To test it run `./pc ls path [...paths]`  
`go run . ls path [...paths]` is also valid.

**Example**:
```bash
./pc ls dummy

# output
📂 dummy
    📄 awesome.txt
    📄 big.txt
    📂 cats
        📄 bony.txt
        📂 cute
            📄 morgan.txt
        📄 giga.txt
    📂 secret
        📄 huu.txt
```

This command behaves can also list multiple directories at the same time, just like the UNIX `ls` that you now.  
**Example**:
```bash
./pc ls dummy ../07-03-22/

# output
📂 dummy
    📄 awesome.txt
    📄 big.txt
    📂 cats
        📄 bony.txt
        📂 cute
            📄 morgan.txt
        📄 giga.txt
    📂 secret
        📄 huu.txt

📂 07-03-22
    📄 README.md
    📄 go.mod
    📄 main.go
    📄 use
```