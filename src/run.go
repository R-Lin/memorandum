package main
import (
    "html/template"
    "net/http"
    "io/ioutil"
    "io"
    "bufio"
    "strings"
    "encoding/json"
    "time"
    "os"
    "sort"
    "fmt"
)

const (
    TEMPLATE_DIR = "./html/"
    DATA_BASEDIR = "./data/"
    RECORD_FILE = DATA_BASEDIR + "working/records.txt"
)

type RecordStruct map[string]map[string]interface{}
type RecordItem map[string]interface{}

// 排序
type RecordItemSort []RecordItem
func (r RecordItemSort) Len() int{
    return len(r)
}
func (r RecordItemSort) Swap(i, j int){
    r[i], r[j] = r[j], r[i]
}
func (r RecordItemSort) Less(i, j int) bool{
    jValue, _:= r[j]["modifyTime"].(float64)
    iValue, _:= r[i]["modifyTime"].(float64)
    return jValue > iValue
}

var RECORD_SET RecordStruct = make(RecordStruct)

func Index(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles(TEMPLATE_DIR + "index.html")
    if err != nil{
        fmt.Println(err.Error())
    }
    s := map[string]string{
        "name": "sd",
        "fuck": "fuck",
    }
    t.Execute(w, s)
}
func CronUpdateFile(){
    // 定时将内存的数据刷入硬盘
    ticker := time.NewTicker(60 * 30 * time.Second)
    for _ = range ticker.C{
        fmt.Println("Record Save Cron running....")
        file, err := os.OpenFile(RECORD_FILE, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
        defer file.Close()
        if err != nil{
            fmt.Println(err.Error())
        }
        for _, v := range RECORD_SET{
            _content, _ := json.Marshal(v)
            file.Write(_content)
            file.Write([]byte("\n"))
        }
    }
}

func SaveRecord(data map[string]interface{}) bool{
    // 将记录存入 RECORD_SET 
    if v, err := data["uuid"].(string); err{
        RECORD_SET[v] = data
        return true
    } else{
        return false
    }
}

func _showRecord(status string)[]RecordItem{
    result := make([]RecordItem, 0)
    for _, v := range RECORD_SET{
        if s_value, ok := v["status"].(string); ok{
            if status != "all" && s_value != status{
                continue
            }
            result = append(result, v)
        }
    }

    sort.Sort(RecordItemSort(result))
    return result
}

func ShowRecord(w http.ResponseWriter, r *http.Request){
    params := r.URL.Query()
    result := make([]RecordItem, 0)
    status := params["status"][0]
    result = _showRecord(status)
    ss, _ := json.Marshal(result)
    w.Write(ss)
}

func ChangeStatus(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.URL)
    params := r.URL.Query()
    mesg := make(map[string]string)
    uuid := params["uuid"][0]
    status := params["status"][0]
    if v, ok := RECORD_SET[uuid]; ok{
        v["status"] = status
        v["modifyTime"] = time.Now().Unix()
        mesg["status"] = "success"
    }else{
        mesg["status"] = "error, no such uuid"
    }
    result, _ := json.Marshal(mesg)
    w.Write(result)

}

func DelTask(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.URL)
    params := r.URL.Query()
    var status string
    uuid := params["uuid"][0]
    if _, ok := RECORD_SET[uuid]; ok{
        delete(RECORD_SET, uuid)
        status = "success"
    }else{
        status = "not such uuid"
    }
    _result, _ := json.Marshal(map[string]string{
        "status": status,
    })
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
    jsonData["createTime"] = createTime
    if v, err := jsonData["isStart"].(string); err{
        if v == "true"{
            jsonData["modifyTime"] = createTime
            jsonData["status"] = "进行中"
        } else{
            jsonData["modifyTime"] = -1
            jsonData["status"] = "计划中"
        }
    }
    if ok := SaveRecord(jsonData); ok{
        ret_mesg["status"] = "success"
    } else {
        ret_mesg["status"] = "error"
    }
    _content, _ := json.Marshal(ret_mesg)
    w.Write(_content)
}


func GetInlineContent(w http.ResponseWriter, r *http.Request){
    pathSlice := strings.Split(r.URL.String(), "/")
    path := pathSlice[len(pathSlice) - 1]
    fmt.Println(r.URL)
    targetFile := TEMPLATE_DIR + path + ".html"
    if _, err := os.Stat(targetFile); err == nil{
        content, _ := ioutil.ReadFile(targetFile)
        w.Write(content)
    }else{
        w.Write([]byte("Not Such " + targetFile))
    }
}

func LoadLocalRecord(){
    f, err := os.Open(RECORD_FILE)
    if err != nil{
        fmt.Println(err.Error())
    }
    defer f.Close()
    rb := bufio.NewReader(f)
    for{
        line, err := rb.ReadString('\n')
        if err != nil{
            if err == io.EOF {
                fmt.Println("Record load completed!")
            }else{
                fmt.Println(err.Error())
            }
            break
        }
        line = strings.TrimSpace(line)
        _jsonData := make(RecordItem, 0)
        _ = json.Unmarshal([]byte(line), &_jsonData)
        SaveRecord(_jsonData)
    }
}

func init(){
    // 建立目录
    PLANS_DIRNAMES := []string{ "working", "template"}
    for _, name := range PLANS_DIRNAMES{
        _targetPath := DATA_BASEDIR + name
        if _, err := os.Stat(_targetPath); err != nil{
            if err := os.Mkdir(_targetPath, os.ModePerm); err != nil{
                fmt.Println(err.Error())
            }else{
                fmt.Printf("已创建目录：%s\n", _targetPath)
            }
        }
    }
    LoadLocalRecord()
    go CronUpdateFile()
}
func main(){
    path, err := os.Getwd()
    if err != nil{
        fmt.Println(err.Error())
    }

    statcFile := http.FileServer(http.Dir(path + "/static"))
    http.HandleFunc("/index", Index)
    http.HandleFunc("/add_task", AddTask)
    http.HandleFunc("/del_task", DelTask)
    http.HandleFunc("/change_status", ChangeStatus)
    http.HandleFunc("/inline_content/", GetInlineContent)
    http.HandleFunc("/getRecord", ShowRecord)
    http.Handle("/static/", http.StripPrefix("/static/", statcFile))
    fmt.Println("Current Path: ", path)
    fmt.Println("http start ! listen on 0.0.0.0:9999")
    http.ListenAndServe(":9999", nil)
}
