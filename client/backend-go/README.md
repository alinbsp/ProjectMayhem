#This is the backend which does most of the heavy work

#Install
```
go get github.com/alinbsp/ProjectMayhem/client/backend-go
```

#Test
Right now this application only has the ability to compare two local project directories.
The best way to test it, is to create two project directories with similar content.
Let's say these two projects have the following structure:

```
Project1:
/home/mayhem/Project1/README.md
/home/mayhem/Project1/main.go

Project2:
/home/mayhem/Project2/README.md
/home/mayhem/Project2/main.go
/home/mayhem/Project2/test.go
```

You can fiddle with the file contents as you please.

Then, you will run:
```
cd $GOPATH/src/github.com/ProjectMayhem/client/backend-go

go run main.go /home/mayhem/Project1 /home/mayhem/Project2
```
