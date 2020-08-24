<p align="center">
  <img src="https://raw.githubusercontent.com/Riku32/Picnic/master/res/banner.png"/>
</p>

## Description

Picnic is a simple framework for making discord bots. It uses Javascript for command and listener scripts. You can easily write and enable/reload scripts live.

## Setup

### Requirements
- Go 1.14
- A computer

### Build
- Clone this repo to your local machine using  `https://github.com/Riku32/Picnic`
- Build the executable `go build .`

## Usage

### Configuration
The `commands` directory  contains command modules. These modules are structured in a folder, each command must have it's own directory, the name does not matter. Each module must have a `commands.js` and `prop.yaml`
```
commands
└── test
    ├── command.js
    └── prop.yaml
```
The `prop.yaml` file for each command must look like this
```YAML
name:  test
description:  Just a testing command
category:  general
alias: [t, info]
```

### Runtime
The Picnic JS runtime is currently very simple and thus prone to change and not documented. A simple command module that works would be the one below
```JS
logger.info(args.author.id)
discord.sendMessage(args.channel.id, args.message.content)
```
There is no console class, it has been replaced with `logger`. The `discord` object is used for most discord detached library methods that do not depend on other data. The `args` object is passed in by default to each module and contains info about the command context.

### The VM
The Javascript virtual machine that Picnic uses is only ECMA5.1 compatible as of right now. So all ECMA6 features are not supported. You can use `require` to require different files in a module. The starting path for require is the picnic executable location, it is not relative to the file.

## Roadmap

- [x] Javascript virtual machine
- [x] Sending messages
  - [ ] Sending embeds
- [ ] Attachables
- [ ] User object
- [x] Message object
- [ ] Categories
- [x] Command system
  - [x] Command arguments
- [ ] Listener system
