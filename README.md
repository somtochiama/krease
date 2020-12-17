## KREASE TOOL

This is a small tool to help with tracking enhancements and docs pr. 
Currently, it can comment on issues(you can filter by labels and pull request).

# Prerequisites: Go version 1.7 or greater.

1. Get the code
```
go get github.com/SomtochiAma/krease
``` 

2. Build
```
cd $GOPATH/src/github.com/SomtochiAma/krease
go install
```

# Example Usage
To comment on issues with particular labels for a particular milestone. You have to pass in a file containing a template for the message.
```
krease issue <name-of-repository> \
    --labels <labels seperated by commas> \
    --milestone <milestone-number> \
    --token <personal-access-token> \
    --name <owner of repository>
    --file <path to file containing the template for message>
```

To comment on issues that are labelled `tracked/yes`  for `1.20` milestone in the enhancements repository
```
krease issue enhancements \
    --labels tracked/yes \
    --milestone 1(milestone number-should be an interger) \
    --token <personal-access-token> \
    --name kubernetes
    --file ./comment.txt
```