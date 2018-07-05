package golog

import (
	"os"
	"errors"
	"fmt"
)
var OutputFileName string
var DoubleOutput bool
func SetOutPut(file string){
	OutputFileName = file
}
func SetDoubleOutput(flag bool){
	DoubleOutput = flag
}
func Output(data ...interface{})error{
	if DoubleOutput{
		if OutputFileName == ""{
			return errors.New("No OutputFile")
		}
		fmt.Println(data)
		//输出到文件
		OutputFile(data)
		return nil
	}
	if OutputFileName != ""{
		OutputFile(data)
	}else{
		fmt.Println(data)
	}
	return nil

}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}
func OutputFile(data ...interface{})error{
	if s,_ := exists(OutputFileName);!s{
		//不存在
		os.Create(OutputFileName)
	}
	file,err := os.Open(OutputFileName)
	if err != nil{
		return err
	}
	for _,data := range data{
		file.Write(data.([]byte))
	}
	//file.Write()
	return nil
}