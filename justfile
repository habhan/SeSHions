# Not needed atm, but good thing to add just incase
set dotenv-load := false

# Prefix with _ to make the recepie private, thus not shown in list
# Can alternetacley use [private]
_default:
  @just --list

build:
 go build -o ./bin/SeSHions
alias b := build

[no-cd]
test-all: build
  go test ./tests/* -v
alias t := test-all

[no-cd]
[positional-arguments]
@test specificTest: build
  go test ./tests/$1 -v

#go run
run:
  go run main.go
alias r := run

# Runs the built executable
execute: build
  ./bin/./SeSHions
alias e := execute

# Adds all files and Commits with commitMessage, member"" 
[no-cd]
[positional-arguments]
@commit commitMessage:
  git add -A
  git commit -m "$1"

# Pushes to github
push:
  git push

#Pulls from github
pull:
  git pull
