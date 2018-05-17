# GoSql
A NoSQL database built in Go
--------------------------------

Commands to the database should be structured in the following manner and sent to the socket on localhost:333 using the TCP/IP Protocol:

* Create: "CREATE [Id] [Value String]"
* Update: "UPDATE [Id] [New Value String]"
* Get: "GET [Id]"
* Delete : "DELETE [Id]"
