package project_util

func MakefileTemplate () string {
    return `clean:
	rm -rf ../bin ../pkg

compile:
	go build -o ../bin/app bootstrap/bootstrap.go

all: clean compile
	mv ../bin/app /opt
	mv app.ini /opt
`
}