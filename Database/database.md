## General Introduction

* BSON or Binary JSON data (fast internal access) stored in the database, but it is just JSON from the developers stand point.
* A document is just an instance of a collection => format/table, documents in the same collection can have unrelated fields. unlike relational DBs, a schema is not imposed on the documentations / analogous to table entries /,

* listCollection command to show all the collections in a DB
* **Replica Sets:** more than one copy of ur data. Done automatically by the system. At least 3 mongo DB instances

* **Shrading:** handling very fast queries over bigger clusters of small machine. Distributing data across multiple servers for the sake of workload. document distribution among different machines / servers.

* **indexes:** queries reading the index instead of going thru all the documentation. Just as relational DBS.
* **Aggregation Pipelines:** 


## Mongo Driver Tutorial for Go lang

### Important commands
* go get go.mongodb.org/mongo-driver
* dep ensure --add go.mongodb.org/mongo-driver/mongo \
go.mongodb.org/mongo-driver/bson \
go.mongodb.org/mongo-driver/mongo/options

"context"
"fmt"
"log"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"

### creating a wireframe (fancy for collections, documents and schema)
### connect to mongo db using Go driver


//this is in the main function.

clientOPtions = options.Client().ApplyURI(ur mongo db string)

client, err := mongo.Connect(context.TODO(), clientOptions) CONTEXT AND CLIENT OPTION

if err != nil {
    fmt.printLn("error")
}

err = client.ping(context, nil => pinging data)
if err => not true

else :"connected"

to get a collection for example

collection := client.database(db_name).collection(collection_name)

collecton.findOne mnamn ahun

and after all operations we can close the connection

err = client.Disconnect(context.TODO())
if no err close.....else log.Fatal()

## THE BSON DATA FORMAT AND ALL ITS BULLSHIT

//fast access and processing
two family types to represent the BSON the D type and the Raw type

D is used in Go. to concisely build BSON objects

#the insert operation on Mongo

insertResult, err := collection.InsertOne(context.TODO(), tobeInserted struct)

if err = > do something
insertResult.InsertedID

same is true for insert mant

trainers := []interface{}{misty, brook} => both are structs pf type Trainers

res, err := Collections.InsertMany(context.TODO(), trainers)
.InsertedIDs => as a slice

//update documents

u need a filted for this tho

filer := bson.D{{"name" : "ash"}} => search by name ash

res,err := collections.UpdateOne(context.TODO(),fiter, update)

//a little controversy here tho

//Find documents
u have to prepare a receiver here tho

var result Trainer

collections.FindOne(context.TODO(), filter).Decode(&result) => this is just like the bind json

//to find multiple => collections .Find() => returns a cursor


use this cursor to navigate thru all the returned results and append If u want to

example finding all documents bson.D{{}} as a filter, then findOptions

cur, err := collections.Find(context.TODO(), bson.D{{}})

if err => handle it

else
trainer_list := []*Trainer


for cur.next(context.TODO()) {

    var elem Trainer //my cursor now on the first trainer on DB
    err := cur.Decode(&elem)
    trainer_list = append(trainer_list, &elem)

}

if err := cur.Err(); err != nil {
    do something
    log.Fatal(err)
}

cur.close(context.TODO())

delete one basically the same as what u did before := filter

to delete many

delRes, err := collections.DeleteMany(context.TODO(),bson.D{{}})
if err => do smth

delRes.DeletedCount




















