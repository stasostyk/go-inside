```
 _____          _____           _     _      _
/ ____|        |     |         (_)   | |    | |
| |  __  ___   |     |_ __  ___ _  __| | ___| |
| | |_ |/ _ \  |o    | |_ \/ __| |/ _` |/ _ \ |
| |__| | (_) | |     | | | \__ \ | (_| |  __/_|
\_____| \___/  |_____|_| |_|___/_|\__,_|\___(_)
               //////
              //////
```
# Custom Commands
STOP - Closes both sides of the program.

# Implementation
All files to be ran on the victim computer are located in the "backdoor" folder, and all the files to be ran on the host computer are in the "listener" folder. Either could be run first, backdoor timeout happens after one minute.

# Summary
Go Inside is a TCP connection based backdoor for executing remote and custom commands, written in the programming language Go (Yes, pun was DEFINITELY intended).
