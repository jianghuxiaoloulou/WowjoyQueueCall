package general

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 基础函数
// action:对象操作类型：上传/下载/删除
// file :获取的对象名
// remotefile：数据库中保存的远端下载的key

type UploadFile struct {
	// 表单名称
	Name string
	// 文件全路径
	Filepath string
}

// 检查文件路径
func CheckPath(path string) {
	dir, _ := filepath.Split(path)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
}

// io.copy()来复制
// 参数说明：
// src: 源文件路径
// dest: 目标文件路径
// key :值不为空是更新instance表中的localtion_code值
func CopyFile(src, dest string) (int64, error) {
	// 判断路径文件夹是否存在，不存在，创建文件夹
	CheckPath(dest)
	file1, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	file2, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file2.Close()
	return io.Copy(file2, file1)
}

// 判断文件是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 将Unicode转换为string
func U2S(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(rune(r))
	}
	return
}

// 去重操作
func RemoveDuplicate(arr []string) []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range arr {
		//判断主键为val的map是否存在
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = nil
		}
	}
	return resArr
}

func File2Base64(file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Errorf(err.Error())
		return ""
	}

	bytedata, err := io.ReadAll(f)
	if err != nil {
		fmt.Errorf(err.Error())
		return ""
	}
	encodeString := base64.StdEncoding.EncodeToString(bytedata)
	return encodeString
}

func Base64ToFile(str, file string) {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Errorf("base64 失败")
	}
	os.WriteFile(file, decodeBytes, 0666)
}
