# jsonstreamsplit

## what is this?

`jsonstreamsplit` accepts a stream of JSON documents on standard input and
writes each document to separate files. That's it. It does not implement any
interrogation or searching of the JSON documents; instead it is intended to
complement tools like the excellent [`jq`](https://github.com/stedolan/jq).

## why?

Say you're using AWS's commandline tool to describe some kind of resource
and you want them in separate files. You might do something like the below
(reformatted for display purposes; this should be one long command):

```
aws ec2 describe-instances --output=json \
  | jq .Reservations[].Instances[] \
  | jsonstreamsplit --output-prefix=ec2-instance
```

This results in each EC2 instance being described in a separate JSON file.

## build it

```
go build
```

## thanks

This code was lifted and only slightly modified from an excellent
[Golang blog post](https://blog.golang.org/json-and-go). Thanks!
