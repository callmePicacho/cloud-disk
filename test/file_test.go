package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

const chunkSize = 100 * 1024 * 1024 // 100 MB

// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileName := "test.mp4"
	// 获取文件信息
	stat, err := os.Stat(fileName)
	if err != nil {
		t.Fatal(err)
	}
	targetFile, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer targetFile.Close()
	// 根据文件大小确认分片个数
	chunkNum := math.Ceil(float64(stat.Size()) / float64(chunkSize))
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		// 最后一点可能不够整的切片大小
		if chunkSize > stat.Size()-int64(i*chunkSize) {
			b = make([]byte, stat.Size()-int64(i*chunkSize))
		}
		targetFile.ReadAt(b, int64(i*chunkSize))
		err = os.WriteFile("./"+strconv.Itoa(i)+".chunk", b, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// 分片文件合并
func TestMergeChunkFile(t *testing.T) {
	// 合并文件
	mergeFile, err := os.OpenFile("test2.mp4", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer mergeFile.Close()
	stat, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	// 根据文件大小反推分片个数
	chunkNum := math.Ceil(float64(stat.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNum); i++ {
		mergeFileName := "./" + strconv.Itoa(i) + ".chunk"
		b, err := ioutil.ReadFile(mergeFileName)
		if err != nil {
			t.Fatal(err)
		}
		mergeFile.Write(b)
		defer func(filename string) {
			os.Remove(filename)
		}(mergeFileName)
	}
}

// 文件一致性校验
func TestCheckFile(t *testing.T) {
	b1, err := os.ReadFile("test.mp4")
	if err != nil {
		t.Fatal(t)
	}
	b2, err := os.ReadFile("test2.mp4")
	if err != nil {
		t.Fatal(t)
	}
	hash1 := fmt.Sprintf("%x", md5.Sum(b1))
	hash2 := fmt.Sprintf("%x", md5.Sum(b2))

	fmt.Println(hash1)
	fmt.Println(hash2)
	fmt.Println(hash1 == hash2)
}

// 打印文件的 md5 值
func TestFileMd5(t *testing.T) {
	b, err := os.ReadFile("2222.png")
	if err != nil {
		t.Fatal(t)
	}
	fmt.Println(fmt.Sprintf("%x", md5.Sum(b)))
}
