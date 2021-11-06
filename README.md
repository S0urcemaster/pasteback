# terminal-tool

terminal-tool fills the gap between most common terms used
and rare terms used usually being searched online.

There is one central spot to put every string you need into.
If it's admin stuff like paths or commands or programmer stuff
like build commands, sql or regex strings - if needed twice:
paste it back.

Helper for the terminal.

- Pushes strings into clipboard.
- Executes commands

`[00] exit`  
`[01] commands`     
`[02] path`  
`[__] 01`

## Path Functions
Since you cannot change the current shell's directory from a child process,
instead when selecting an item from that menu, it's copied to the clipboard.

`[00] exit`  
`[01] cd /mnt/c`  
`[02] cd /home/userx`  
`[03] cd /what/ever`  

You can then just paste it to the prompt.
Set an item in `paths.txt` like that:  
`cd /mnt/c`  

(It can be any string, not only a path ofc.)

## Commands Functions
Commands can be executed, though.
Set an item in `commands.txt` like this:  
`[dir <directory>] <cmd name> <args>...`  
The optional workdir of a command is set separately with the "dir" keyword
(to avoid confusion with "cd").  
Example:  
`dir /mnt/c ls -l`

Will print print a detailed list of `/mnt/c`'s contents.
