package main
import (
    "net/http"
    "io/ioutil"
   // "io"
   // "bufio"
   // "strings"
    "gopkg.in/mgo.v2/bson"
    "encoding/json"
    "time"
    "os"
    //"sort"
    "fmt"
)

type RecordStruct map[string]map[string]interface{}
type RecordItem map[string]interface{}

func SaveRecord(data map[string]interface{}) bool{
    // 将记录存入 RECORD_SET 
    if err := SHORT_TERM_COL.Insert(data); err == nil{
        return true
    } else{
        return false
    }
}

func _showRecord(status string)[]RecordItem{
    var result []RecordItem
    specifyValue := make(bson.M)
    if status == "已完成"{
        specifyValue["$eq"] = status
    }else{
        specifyValue["$ne"] = "已完成"
    }
    query := bson.M{
        "status": specifyValue,
    }
    // Sort 加个 - 号就是代表倒序
    SHORT_TERM_COL.Find(query).Sort("-modifyTime").All(&result)
    return result
}

func ShowRecord(w http.ResponseWriter, r *http.Request){
    params := r.URL.Query()
    result := make([]RecordItem, 0)
    status := params["status"][0]
    result = _showRecord(status)
    ss, _ := json.Marshal(result)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(ss)
}

func ChangeStatus(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.URL)
    params := r.URL.Query()
    mesg := make(map[string]string)
    uuid := params["uuid"][0]
    status := params["status"][0]
    selectorDsl := bson.M{
        "uuid": uuid,
    }
    updateDsl := bson.M{
        "$set": bson.M{
            "status": status,
            "modifyTime": time.Now().Unix(),
        },
    }
    if err := SHORT_TERM_COL.Update(selectorDsl, updateDsl); err == nil{
        mesg["status"] = "success"
    }else{
        mesg["status"] = "error, no such uuid"
    }
    result, _ := json.Marshal(mesg)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(result)

}

func DelTask(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.URL)
    params := r.URL.Query()
    var status string
    uuid := params["uuid"][0]
    if err := SHORT_TERM_COL.Remove(bson.M{"uuid": uuid}); err == nil{
        status = "success"
    }else{
        status = "not such uuid"
    }
    _result, _ := json.Marshal(map[string]string{
        "status": status,
    })
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(_result)
}

func AddTask(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.URL)
    defer r.Body.Close()
    con, _ := ioutil.ReadAll(r.Body)
    var jsonData RecordItem
    json.Unmarshal(con, &jsonData)
    createTime := time.Now().Unix()
    ret_mesg := make(map[string]string)
    jsonData["modifyTime"] = createTime
    jsonData["createTime"] = createTime
    if v, err := jsonData["isStart"].(bool); err{
        if v{
            jsonData["status"] = "进行中"
        } else{
            jsonData["status"] = "计划中"
        }
    }
    if ok := SaveRecord(jsonData); ok{
        ret_mesg["status"] = "success"
    } else {
        ret_mesg["status"] = "error"
    }
    _content, _ := json.Marshal(ret_mesg)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(_content)
}

func init(){
}

func main(){
    path, err := os.Getwd()
    if err != nil{
        fmt.Println(err.Error())
    }

    statcFile := http.FileServer(http.Dir(path + "/static"))
    http.HandleFunc("/add_task", AddTask)
    http.HandleFunc("/del_task", DelTask)
    http.HandleFunc("/change_status", ChangeStatus)
    http.HandleFunc("/getRecord", ShowRecord)
    http.Handle("/static/", http.StripPrefix("/static/", statcFile))
    fmt.Println("Current Path: ", path)
    fmt.Println("http start ! listen on 0.0.0.0:9999")
    http.ListenAndServe(":9999", nil)
}
