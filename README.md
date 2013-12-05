ProtoML
=======

What is ProtoML?
----------------
An engine to enable easy data analytics. At its base it is a task processing graph exposed through a REST api. At a higher level there is a type system over different types of data which encompasses all data and tasks in the engine. The tasks called **transforms** are made of a JSON description tied with some execution context. Transforms are chained together and validated according to constraints in the transform and type constraints. Tasks and data can be searched through using [elasticsearch](http://www.elasticsearch.org/overview/). Task execution is handled through [Luigi](https://github.com/spotify/luigi).

ProtoML is still under heavy development stay tuned for our first full release! 

Installation
-------------
### Prerequisites
* [Go](http://golang.org/doc/install)
* [Python 2.7](http://www.python.org/download/)
* [elasticsearch](http://www.elasticsearch.org/download/)
* [Luigi](https://github.com/spotify/luigi)

### Install
* Make sure Luigi's `luigid` and elasticsearch's `elasticsearch` commands can be found on the system path
* Setup a [go workspace](http://golang.org/doc/code.html#Organization)
* Then run the following commands to install all needed packages:
  * `go get https://github.com/mattbaird/elastigo`
  * `go get https://github.com/ant0ine/go-json-rest`
  * `go get https://github.com/ProtoML/ProtoML`
  * `go get https://github.com/ProtoML/ProtoML-persist`
  * `go get https://github.com/ProtoML/ProtoML-transforms`
  * `go get https://github.com/ProtoML/ProtoML-dashboard`
* Be sure to set your `$PROTOMLDIR` enviroment variable to `$GOPATH/src/github.com/ProtoML`
* Make sure `$GOPATH/bin` is on the system path

Running
-------
### ProtoML Server
Create a directory with a protoml configuration file `ProtoML.json`. See [here](https://github.com/ProtoML/ProtoML/blob/master/tests/testsets/synthetic/ProtoML_Startup.json) for an example. 

Run `ProtoML` in the directory.

### ProtoML Dashboard
Run `ProtoMLDashboard`.
