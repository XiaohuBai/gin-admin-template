/*
 * @Author: XiaohuBai@outlook.com
 * @Date: 2020-12-01 23:37:32
 * @LastEditors: XiaohuBai
 * @LastEditTime: 2020-12-01 23:42:23
 * @Description: 描述
 */

package utils

import (
	"io/ioutil"
	"os"
	"strconv"
)

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

const breakpointDir = "./breakpointDir/"
const finishDir = "./fileDir/"

// BreakPointContinue 断点续传
func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathc, err := makeFileContent(content, fileName, path, contentNumber)
	return pathc, err

}

// CheckMd5 检查Md5
func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true // "可以继续上传"
	}
	return false // "切片不完整，废弃"

}

//makeFileContent 创建切片内容
func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	return path, nil
}

//MakeFile 创建切片文件
func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := ioutil.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, _ := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	for k := range rd {
		content, _ := ioutil.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}
	defer fd.Close()
	return finishDir + fileName, nil
}

// RemoveChunk 移除切片
func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
