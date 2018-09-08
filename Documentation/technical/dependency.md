om_admin dependency injection package :
-	"github.com/aws/aws-sdk-go/aws"
-	"github.com/aws/aws-sdk-go/service/s3"
-	"github.com/aws/aws-sdk-go/service/sns"
-	"github.com/aws/aws-sdk-go/aws/session"
-	"github.com/aws/aws-sdk-go/service/dynamodb"
-	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
-	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
-	"github.com/astaxie/beego/orm"
-	"github.com/davecgh/go-spew/spew"
-	"github.com/douglasmakey/go-fcm"
-	"github.com/fatih/structs"
-	"github.com/go-sql-driver/mysql"
-	"github.com/golang/protobuf/proto"
-	"github.com/golang/protobuf/protoc-gen-go/descriptor"
-	"github.com/golang/protobuf/ptypes"
-	"github.com/golang/protobuf/ptypes/any"
-	"github.com/golang/protobuf/ptypes/empty"
-	"github.com/golang/protobuf/ptypes/struct"
-	"github.com/golang/protobuf/ptypes/timestamp"
-	"github.com/golang/protobuf/ptypes/wrappers"
-	"github.com/googleapis/gax-go"
-	"github.com/kataras/go-sessions"
-	"github.com/lib/pq"
-	"github.com/mattbaird/gochimp"
-	"github.com/rs/cors"
-	"github.com/spf13/cast"
-	"github.com/spf13/viper"


1. AWS SDK (https://github.com/aws/aws-sdk-go):
	Once the client is created you can make an API request to the service. Each API method takes a input parameter, and returns the service response and an 	error. The SDK provides methods for making the API call in multiple ways.

Install :
go get github.com/aws/aws-sdk-go

Configuring AWS Region:
	In addition to the credentials you'll need to specify the region the SDK will use to make AWS API requests to. In the SDK you can specify the region either with an environment variable, or directly in code when a Session or service client is created. The last value specified in code wins if the region is specified multiple ways.

	To set the region via the environment variable set the "AWS_REGION" to the region you want to the SDK to use. Using this method to set the region will allow you to run your application in multiple regions without needing additional code in the application to select the region.

	The SDK includes the Go types and utilities you can use to make requests to AWS service APIs. Within the service folder at the root of the SDK you'll find a package for each AWS service the SDK supports. All service clients follows a common pattern of creation and usage.

	When creating a client for an AWS service you'll first need to have a Session value constructed. The Session provides shared configuration that can be shared between your service clients. When service clients are created you can pass in additional configuration via the aws.Config type to override configuration provided by in the Session to create service client instances with custom configuration.

	Once the service's client is created you can use it to make API requests the AWS service. These clients are safe to use concurrently.
	-	"github.com/aws/aws-sdk-go/aws"
	-	"github.com/aws/aws-sdk-go/service/s3"
	-	"github.com/aws/aws-sdk-go/service/sns"
	-	"github.com/aws/aws-sdk-go/aws/session"
	-	"github.com/aws/aws-sdk-go/service/dynamodb"
	-	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	-	"github.com/aws/aws-sdk-go/service/cloudwatchevents"


2. beego orm (https://github.com/astaxie/beego/tree/master/orm):
Support Database :
	MySQL: github.com/go-sql-driver/mysql
	PostgreSQL: github.com/lib/pq
	Sqlite3: github.com/mattn/go-sqlite3
	Passed all test, but need more feedback.

Features :
	full go type support easy for usage, simple CRUD operation auto join with relation table cross DataBase compatible query Raw SQL query / mapper without orm model full test keep stable and strong more features please read the docs

Install :
go get github.com/astaxie/beego/orm

Changelog:
	support table auto create
	update test for database types
	go type support, such as int8, uint8, byte, rune
	date / datetime timezone support very well

3. go-spew (https://github.com/davecgh/go-spew):
	Go-spew implements a deep pretty printer for Go data structures to aid in debugging. A comprehensive suite of tests with 100% test coverage is provided to ensure proper functionality. See test_coverage.txt for the gocov coverage report. Go-spew is licensed under the liberal ISC license, so it may be used in open source or commercial projects.

Install :
	go get -u github.com/davecgh/go-spew/spew


4. go-fcm (https://github.com/douglasmakey/go-fcm):
	Firebase Cloud Messaging ( FCM ) Library using golang ( Go )
	This library uses HTTP/JSON Firebase Cloud Messaging connection server protocol

Install :
	go get github.com/douglasmakey/go-fcm


5. Structs (https://github.com/fatih/structs):
	Structs contains various utilities to work with Go (Golang) structs. It was initially used by me to convert a struct into a map[string]interface{}. With time I've added other utilities for structs. It's basically a high level package based on primitives from the reflect package. Feel free to add new functions or improve the existing code.

Install :
	go get github.com/fatih/structs

6. Go-MySQL-Driver (https://github.com/go-sql-driver/mysql) :
	A MySQL-Driver for Go's database/sql package

Features :
	Lightweight and fast
	Native Go implementation. No C-bindings, just pure Go
	Connections over TCP/IPv4, TCP/IPv6, Unix domain sockets or custom protocols
	Automatic handling of broken connections
	Automatic Connection Pooling (by database/sql package)
	Supports queries larger than 16MB
	Full sql.RawBytes support.
	Intelligent LONG DATA handling in prepared statements
	Secure LOAD DATA LOCAL INFILE support with file Whitelisting and io.Reader support
	Optional time.Time parsing
	Optional placeholder interpolation

Install :
	go get -u github.com/go-sql-driver/mysql


7. Go support for Protocol Buffers - Google's data interchange format (https://github.com/golang/protobuf ):
	This package and the code it generates requires at least Go 1.6.
	This software implements Go bindings for protocol buffers. For information about protocol buffers themselves, see https://developers.google.com/protocol-buffers/

	-	"github.com/golang/protobuf/proto"
	-	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	-	"github.com/golang/protobuf/ptypes"
	-	"github.com/golang/protobuf/ptypes/any"
	-	"github.com/golang/protobuf/ptypes/empty"
	-	"github.com/golang/protobuf/ptypes/struct"
	-	"github.com/golang/protobuf/ptypes/timestamp"
	-	"github.com/golang/protobuf/ptypes/wrappers"

Using protocol buffers with Go:
	Once the software is installed, there are two steps to using it. First you must compile the protocol buffer definitions and then import them, with the support library, into your program.
	To compile the protocol buffer definition, run protoc with the --go_out parameter set to the directory you want to output the Go code to.

	protoc --go_out=. *.proto

gRPC Support :
	If a proto file specifies RPC services, protoc-gen-go can be instructed to generate code compatible with gRPC (http://www.grpc.io/). To do this, pass the plugins parameter to protoc-gen-go; the usual way is to insert it into the --go_out argument to protoc:

	protoc --go_out=plugins=grpc:. *.proto

Compatibility :
	The library and the generated code are expected to be stable over time. However, we reserve the right to make breaking changes without notice for the following reasons:

	Security. A security issue in the specification or implementation may come to light whose resolution requires breaking compatibility. We reserve the right to address such security issues.
	
	Unspecified behavior. There are some aspects of the Protocol Buffers specification that are undefined. Programs that depend on such unspecified behavior may break in future releases.
	
	Specification errors or changes. If it becomes necessary to address an inconsistency, incompleteness, or change in the Protocol Buffers specification, resolving the issue could affect the meaning or legality of existing programs. We reserve the right to address such issues, including updating the implementations.
	
	Bugs. If the library has a bug that violates the specification, a program that depends on the buggy behavior may break if the bug is fixed. We reserve the right to fix such bugs.
	
	Adding methods or fields to generated structs. These may conflict with field names that already exist in a schema, causing applications to break. When the code generator encounters a field in the schema that would collide with a generated field or method name, the code generator will append an underscore to the generated field or method name.
	
	Adding, removing, or changing methods or fields in generated structs that start with XXX. These parts of the generated code are exported out of necessity, but should not be considered part of the public API.
	
	Adding, removing, or changing unexported symbols in generated code.

Install :
	go get github.com/golang/protobuf/proto

8. Google API Extensions for Go (https://github.com/googleapis/gax-go ):
	Google API Extensions for Go (gax-go) is a set of modules which aids the development of APIs for clients and servers based on gRPC and Google API conventions.
	Note: Application code will rarely need to use this library directly, but the code generated automatically from API definition files can use it to simplify code generation and to provide more convenient and idiomatic API surface

Install :
	go get github.com/googleapis/gax-go

9. go-sessions (https://github.com/kataras/go-sessions):
	Fast http sessions manager for Go.
	Simple API, while providing robust set of features such as immutability, expiration time (can be shifted), databases like badger and redis as back-end storage.

Features :
	Focus on simplicity and performance.
	Flash messages.
	Supports any type of external database.
	Works with both net/http and valyala/fasthttp.

Install :
	go get -u github.com/kataras/go-sessions

10. pq - A pure Go postgres driver for Go's database/sql package (https://github.com/lib/pq ):
	Similarly to libpq, when establishing a connection using pq you are expected to supply a connection string containing zero or more parameters. A subset of the connection parameters supported by libpq are also supported by pq. Additionally, pq also lets you specify run-time parameters (such as search_path or work_mem) directly in the connection string. This is different from libpq, which does not allow run-time parameters in the connection string, instead requiring you to supply them in the options parameter.

Features :
	SSL
	Handles bad connections for database/sql
	Scan time.Time correctly (i.e. timestamp[tz], time[tz], date)
	Scan binary blobs correctly (i.e. bytea)
	Package for hstore support
	COPY FROM support
	pq.ParseURL for converting urls to connection strings for sql.Open.
	Many libpq compatible environment variables
	Unix socket support
	Notifications: LISTEN/NOTIFY
	pgpass support

Install :
	go get github.com/lib/pq


11. gochimp (https://github.com/mattbaird/gochimp ):
	Go based API for Mailchimp, starting with Mandrill.
	To run tests, set a couple env variables. (replacing values with your own mandrill credentials):
		$ export MANDRILL_KEY=111111111-1111-1111-1111-111111111
		$ export MANDRILL_USER=user@domain.com
Mandrill Status :
	API Feature complete on Oct 26/2012
	Adding tests, making naming conventions consistent, and refactoring error handling

Install : 
	github.com/mattbaird/gochimp

12. Go CORS handler (https://github.com/rs/cors ):
	CORS is a net/http handler implementing Cross Origin Resource Sharing W3 specification in Golang.

Install :
	go get github.com/rs/cors

13. cast (https://github.com/spf13/cast ):
	Easy and safe casting from one type to another in Go

Install :
	go get github.com/spf13/cast

14. viper (https://github.com/spf13/viper):
	Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:

	setting defaults
	reading from JSON, TOML, YAML, HCL, and Java properties config files
	live watching and re-reading of config files (optional)
	reading from environment variables
	reading from remote config systems (etcd or Consul), and watching changes
	reading from command line flags
	reading from buffer
	setting explicit values
	Viper can be thought of as a registry for all of your applications configuration needs.

Install :
	go get github.com/spf13/viper

