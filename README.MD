<p align="center">
 <img width=200px height=200px src="https://upload.wikimedia.org/wikipedia/commons/thumb/8/82/Telegram_logo.svg/480px-Telegram_logo.svg.png" alt="Project logo"></a>
 <h3 align="center">teleBrat</h3>
 <h5 align="center">A simple RAT written fully in go using telegram as a C2</h3>
</p>

---


###### *⚠️ For education purposes only. I do not take and responsibility in any unauthorized or malicious use of this software!*⚠️


## Why?
I used this as an opportuinity to learn golang and try out new stuff, also using telegram as a c2<br>
was an old idea i had a while ago that i wanted to test out and see how can you abuse legitimate services<br>
to work as your c2.
## Features
* Command Execution
* Download files from remote PC
* builtin wget like option (upload files using a 3d party hosting service)
* Screenshot


## Setup

##### Install go dependencies:

```
go install github.com/go-telegram-bot-api/telegram-bot-api@latest
go install github.com/kbinani/screenshot@latest

```

##### Get your bot up and running:
in general you should follow the official telegram tutorial on how to set up your first bot using the link:
```
https://core.telegram.org/bots
```
##### Setup API key for your bot:
```
In project file change your apiKey var in the main.go file to your api key, should look something like
var apiKey = "1234567890:AAAAAAA-AA-AAAAAAAAAAAAAA_AAAAA-AAA"
```

## Complie
You should probably compile using the nogui flag in go
```
go build -ldflags -H=windowsgui telebrat
```

## First run
After you run your compiled program you should start a conversation with your bot, only then you can interact with it.<br>
you can also try the <b>/test</b> to check if everything is working and the <b>/help</b> command to see whats command are available

# Tested on
* Windows 7
* windows 8.1
* windows 10
* windows server 2008
## Todo
* Add options for a costume timeout for command execution. 
* Add more execution options, PS, LOLBAS etc...
* Get an autostart system
* Use system proxy
* Maybe add some AV evasion etc... probably a different project...
---


###### *⚠️ For education purposes only. I do not take and responsibility in any unauthorized or malicious use of this software!*⚠️
