package main

import (
    "gopkg.in/mgo.v2"
    "fmt"
)

const (
    MONGO_URL = "localhost:27017"
    DBNAME = "popo"
)

var SESSION *mgo.Session
var DATABASE *mgo.Database
var LONG_TERM_COL *mgo.Collection
var SHORT_TERM_COL *mgo.Collection
var err error

type LongTermStruct struct {
    Title     string   `json: :title"`
    Desc      string   `json: "desc"`
    TimeStamp float64  `json: "timestamp"`
}

type ShortTermStruct struct{
    TaskType   string  `json:",omitempty"`
    Uuid       string  `json:",omitempty"`
    TaskDesc   string  `json:",omitempty"`
    Priority   string  `json:",omitempty"`
    IsStart    string  `json:",omitempty"`
    Status     string  `json:",omitempty"`
    CreateTime float64 `json:",omitempty"`
    ModifyTime float64 `json:",omitempty"`
}

func Test(){
    fmt.Println("Mongo test ok")
}

func init(){
    SESSION, err = mgo.Dial(MONGO_URL)
    if err != nil{
        panic(err.Error())
    }
    DATABASE := SESSION.DB(DBNAME)
    LONG_TERM_COL= DATABASE.C("long_term_task")
    SHORT_TERM_COL= DATABASE.C("short_term_task")
}
