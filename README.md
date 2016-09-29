# CodeCI #

[![Build Status](https://travis-ci.org/itjustworksteam/codeci.svg?branch=master)](https://travis-ci.org/itjustworksteam/codeci)

CodeCI is a command line tool used to test your software locally. It works with [docker](http://www.docker.com/) and requires 64-bit machines.

## Requirements ##

* 64-bit machine
* docker
* docker-compose
* golang

## Install ##

* ```bash install.sh```
* ```sudo go build -o /usr/local/bin/codeci```

#### Testing Installation ####

* ```mkdir test```
* ```cd test```
* create this file and name it ```codeci.yml```:
```
image: docker/whalesay
script:
   - cowsay Hello CodeCI!
```
* run ```codeci``` and you will see a whale that say ```Hello CodeCI!```


## How to use it ##

* simply in the root directory of your project run ```codeci```
* attention! It needs a file called ```codeci.yml``` to work

## Commands ##

* ```codeci --version``` show the version
* ```codeci images``` show default images
* ```codeci --help``` show an help
* ```codeci -f codeci.anotherconf.yml``` specify another name for your codeci.yml file. The file has to be named with a prefix of ```codeci``` and suffix of ```.yml```

## codeci.yml ###

* it's a simple yml file with 3 properties: os, language and script.
* os -> is the operating system for now it is available only ubuntu14
* language -> is the language of the project
* image -> the image from dockerhub or local ( replace os and language )
* script -> is the script that the build has to execute
* os types available: ```ubuntu14```
* language types available: ```java```, ```python```, ```php```, ```swiftenv```, ```node```, ```go```, ```cpp```, ```scala```, ```ruby``` or ```none```
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
* image example
```
image: ubuntu:14.04
script:
  - apt-get update
  - apt-get install vim -y
```

* cpp example: ( used for c and c++ )
```
os: ubuntu14
language: cpp
script:
  - make -v
  - gcc -v
  - cc -v
  - g++ -v
```
