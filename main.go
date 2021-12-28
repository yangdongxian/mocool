package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"mocool/moTest"
	"strconv"
	"strings"
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

	c1 := moTest.Country{"China"}
	c2 := moTest.City{"Shanghai"}
	moTest.PrintStr(c1)
	moTest.PrintStr(c2)

	var name string
	var age int
	age = 22
	name = strconv.Itoa(age)
	fmt.Println("name:",name)
	//update name string to int
	//update name string to int
	name = "23"
	age,_ = strconv.Atoi(name)
	fmt.Println("age:",age)

	//funcOptons
	server,_:=moTest.NewServer("localhost:127.0.0.0.1",8080)
	fmt.Println("server:",server)

	var maplist = []string{"Hao","Cheng","Mess"}
	y := moTest.MapStrToStr(maplist, func(s string) string {
		return strings.ToLower(s)
	})
	fmt.Println("mapable:",y)
	token := "e3f054bdd92b46a0b6f2e0365810f139"
	token_len := GetStringSize(token)
	fmt.Println("token_len:",token_len)


}
func OperationV2(arg []string) string {
	var result float64
	var resultStr string
	//var rem int
	a, err1 := strconv.ParseFloat(arg[0], 64)
	b, err2 := strconv.ParseFloat(arg[1], 64)
	if err1 != nil {
		return "-1"
	}
	if err2 != nil {
		return "-1"
	}
	switch arg[2] {
	case "add":
		result = a + b
	case "minus":
		result = a - b
	case "multi":
		result = a * b
	case "div":
		result = a / b
	case "rem":
		result = float64(int64(a) % int64(b))
	default:
		return "-1"
	}
	if len(arg) >= 4 {
		dout, err := strconv.ParseInt(arg[3], 10, 32)
		if err != nil {
			return "传入参数有误,保存小数点数为正整数或0"
		}
		str := "%." + fmt.Sprintf("%d", dout) + "f"
		resultStr = fmt.Sprintf(str, result)

	} else {
		resultStr = fmt.Sprintf("%.8f", result)
		resultStr = strings.TrimRight(resultStr, "0")
		resultStr = strings.TrimRight(resultStr, ".")
	}
	return resultStr
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

// 定义一个变量
func GetVar(args string) string{
	return args
}

//获取字符串长度
func GetStringSize(args string) string{
	str_len := strconv.Itoa(len([]rune(args)))
	return str_len
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
