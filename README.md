# Ganzi

A simple program that manages and generates welcome messages when you connect to a server.

## Motivation

When accessing multiple servers, sometimes you don't know which server you are accessing. You can tell which server you
are currently connected to by displaying the text you set through this program.

## Command
```shell
Usage:
  ganzi [command]

Available Commands:
  help        Help about any command
  set         Typing text message which displayed shell activate message
  clean       Remove current set text message(it wll be delete ~/.banner.txt file)
  reset       Reset and remove all configuration

Flags:
  -h, --help            help for ganzi

Additional help topics:
  ganzi show  Showing current welcome message

Use "ganzi [command] --help" for more information about a command.

```