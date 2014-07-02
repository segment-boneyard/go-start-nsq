
# start-nsq

  Little helper program to boot nsqd / nsqlookupd / nsqadmin nodes for local development.

## Installation

```
$ go get github.com/segmentio/go-start-nsq
```

## Usage

 One NSQD:

```
$ go-start-nsq
```

 Three NSQD nodes:

```
$ go-start-nsq -n 3
```

# License

  MIT
