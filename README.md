# CodeCI #

CodeCI is a command line tool used to test your software locally. It works with docker and requires 64 bit machines.

## Requirements ##

* 64 bit machine
* docker
* docker-compose
* golang

## Install ##

* ```bash install.sh```
* ```sudo go build -o /usr/local/bin/codeci```

## How to use it ##

* simply in the root directory of your project run ```codeci```
* attention! It needs a file called ```codeci.yml``` to work

## codeci.yml ###

* it's a simple yml file with 3 properties: os, language and script.
* os -> is the operating system for now it is available only ubuntu14
* language -> is the language of the project
* script -> is the script that the build has to execute
* os types available: ```ubuntu14```
* language types available: ```java```, ```python```, ```php```, ```swiftenv```, ```node```, ```go``` or ```none```
* script you can run whatever you want because you have a docker container with sudo privileges

## Example of codeci.yml ##

* java example
```
os: ubuntu14
language: java
script:
 - echo hello
 - echo java
```

* swiftenv example
```
os:ubuntu14
language: swiftenv
script:
  - swiftenv install 3.0
  - swiftenv global 3.0
  - swift build
  - swift test
```

* none example
```
os: ubuntu14
language: none
script:
  - apt-get update
  - apt-get install build-essential -y
  - gcc -v
  - make -v
```
 