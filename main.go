package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"mocool/interfaceTest"
	"time"
)

func main() {
	fmt.Println("Test func RandomGetDBlistRow")

	rows := `[{"name":"张三","age":"10"},{"name":"李四","age":"12"},{"name":"王五","age":"20"},{"name":"六仔","age":"21"}]`
	//rows := `["name":"张三","age":"10"]`
	fmt.Println("rows:", rows)
	randomRow := GetDBlistRandomRow(rows)
	fmt.Println("randomRow:", randomRow)

	list := `["张三","李四","王五","刘六","刘七","刘八","刘九","11"]`
	listValue := GetListRandomValue(list)
	fmt.Println("listValue:", listValue)

	//sliceTest()
	//interfaceable.PrintStr("aa")
	//c2 := Country{"china"}
	//c1 := interfaceable.Country{"China"}
	//interfaceable.PrintStr(c1)
	//interfaceable.PrintStr(c2)
	c1 := interfaceTest.Country{"China"}
	interfaceTest.PrintStr(c1)

}

//查询数据库返回列表，随时获取一条数据
func GetDBlistRandomRow(args string) string {
	var rowsJson []interface{}
	//if "string" ==  reflect.TypeOf(args).Name() {
	//    json.Unmarshal([]byte(args),&rowsJson)
	//    fmt.Println("rowsJson:",rowsJson)
	//}
	//argsType := fmt.Sprintf("%T \n",args)
	json.Unmarshal([]byte(args), &rowsJson)
	rowlen := len(rowsJson)
	//randomRows := make([]interface{},0)
	if rowlen > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomRowNum := r.Intn(rowlen)
		row := rowsJson[randomRowNum]
		//randomRows = append(randomRows,row)
		newRows, _ := json.Marshal(row)
		return string(newRows)
	} else {
		return args
	}
}

//指定数据随时返回一个值
func GetListRandomValue(args string) string {
	if len(args) < 0 {
		return fmt.Sprintf("请输入合法数组,%s", args)
	}
	var list []string
	json.Unmarshal([]byte(args), &list)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomListNum := r.Intn(len(list))
	listValue := list[randomListNum]
	return string(listValue)
}

//查询数据库返回列表，随时获取一条数据
//func GetDBlistIndexRow(args[] string) string {
//    var rowsJson []interface{}
//    dblist := args[0]
//    dbIndex := args[1]
//
//    err := json.Unmarshal([]byte(dblist),&rowsJson)
//    if err != nil {
//       return "请检查args[0]是否合法"
//    }
//    rowlen := len(rowsJson)
//    //randomRows := make([]interface{},0)
//    if (float64(dbIndex) > 0) &&  (float64(dbIndex) < float64(rowlen)) {
//        row := rowsJson[dbIndex]
//        newRows,_ := json.Marshal(row)
//        return string(newRows)
//    }else if (float64(dbIndex)) < 0 &&  (float64(dbIndex) > -float64(rowlen)){
//        row := rowsJson[dbIndex]
//        newRows,_ := json.Marshal(row)
//        return string(newRows)
//
//    }else {
//        return "请检查args[2]参数，下标值越界args[0]长度"
//    }
//}

func sliceTest() {
	//切片在go中是一个结构休，而非数组
	//slice 变量定义以及赋值
	//切片赋值引用，指向同一个内存地址
	var fo = make([]int, 5)
	fo[3] = 43
	fo[4] = 122

	//slice 索引使用
	//string append追加
	indexNum := -1
	bar := fo[1:4]
	fo = append(fo, 33)
	bar[1] = 99
	fmt.Println("foo:", fo)
	fmt.Println("bar:", bar)
	fmt.Println("-bar:", bar[len(bar)+indexNum])

	//byte 索引获取
	path := []byte("AAAA/BBBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1:", string(dir1))
	fmt.Println("dir2:", string(dir2))

	//利用bytes.NewBufferString将string转化byte
	buffer := bytes.NewBufferString("suffer....")
	buf_dir1 := bytes.NewBuffer(dir1)
	buf_dir1.Write(buffer.Bytes())
	fmt.Println("dir1:", buf_dir1.String())
	fmt.Println("dir2:", string(dir2))

}
