# Not needed atm, but good thing to add just incase
set dotenv-load := false

# Prefix with _ to make the recepie private, thus not shown in list
# Can alternetacley use [private]
_default:
  @just --list

build:
 go build . -o ./bin/SeSHions
alias b := build

test: build
  go test ./tests/* -v
alias t := test

run:
  go run main.go
alias r := run

execute: build
  ./bin/./SeSHions
alias e := execute

# Adds all files and Commits with string provided to command
[no-cd]
[positional-arguments]
@commit commit:
  git add -A
  git commit -m "$@"
