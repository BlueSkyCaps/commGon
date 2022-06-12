package commGon

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
)

var ce rune

// DebugPrint 输出错误 退出程序
func DebugPrint(err error) {
	println(err.Error())
	fmt.Scanf("%c", &ce)
	debug.PrintStack()
	os.Exit(1)
}

// MatchRegexString 按正则表达式p则匹配字符串v
func MatchRegexString(p, v string) bool {
	var e error
	var m bool
	if m, e = regexp.MatchString(p, v); m && (e == nil) {
		return true
	}
	if e != nil {
		println("error in MatchRegexString", e.Error())
	}
	return false
}

// UuidGenerator 生成16字符组成的唯一标识符
func UuidGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// CreateFolder 在指定路径下创建一个新文件夹 存在则删除再创建
func CreateFolder(createPath string) {
	// 有没有存在此文件夹
	_, err := os.Stat(createPath)
	// 存在则先删除
	if err == nil {
		// 先删除此文件夹中的所有子数据，若有
		err = os.RemoveAll(createPath)
		// 删除子数据发生错误
		if err != nil {
			panic("removeAll tmp folder data failed.")
		}
		// 再删除此文件夹本身(此处是多余的 因为本身也已删除 )
		_ = os.Remove(createPath)
		// 再次判断此文件夹是否存在
		_, err = os.Stat(createPath)
		// 若nil,表示仍存在，预期不理想。中止
		if err == nil {
			panic("remove tmp folder failed.")
		}
	}
	err = os.MkdirAll(createPath, os.ModePerm)
	if err != nil {
		println("error in CreateFolder:", err.Error())
		var c int
		_, _ = fmt.Scanf("%c", &c)
		os.Exit(1)
	}
}

// SortStringSlice 按数字标准排序一个序列，指定desc为true则降序
func SortStringSlice(sr []string, desc bool) {
	sort.Slice(sr, func(i, j int) bool {
		formatEndI := strings.LastIndex(sr[i], ".")
		formatEndJ := strings.LastIndex(sr[j], ".")
		fileNameI := sr[i][0:formatEndI]
		fileNameJ := sr[j][0:formatEndJ]
		ci, errI := strconv.ParseInt(fileNameI, 10, 32)
		cj, errJ := strconv.ParseInt(fileNameJ, 10, 32)
		if errI == nil && errJ == nil {
			if ci == 10 || cj == 10 {
				return ci < cj
			}
			return ci < cj
		} else {
			return sr[i] < sr[j]
		}
	})
	// 降序，逆转序列
	if desc {
		Reverse(sr)
	}
}

// RemoveFolder 删除一个文件夹。p必须是文件夹绝对路径
func RemoveFolder(p string) {
	// 有没有存在此文件夹
	_, err := os.Stat(p)
	// 存在则先删除
	if err == nil {
		// 先删除此文件夹中的所有子数据，若有
		err = os.RemoveAll(p)
		// 删除子数据发生错误
		if err != nil {
			panic("removeAll tmp folder data failed.")
		}
		// 再删除此文件夹本身(此处是多余的 因为本身也已删除 )
		_ = os.Remove(p)
	}
}

// RemoveFolderChildren 删除一个文件夹下的内容，保留文件夹。p必须是文件夹绝对路径
func RemoveFolderChildren(p string) {
	// 有没有存在此文件夹
	_, err := os.Stat(p)
	// 存在则删除
	if err == nil {
		dys, _ := os.ReadDir(p)
		for _, d := range dys {
			if d.IsDir() {
				// 删除需绝对路径
				err := os.RemoveAll(path.Join(p, d.Name()))
				if err != nil {
					return
				}
			} else {
				err := os.Remove(path.Join(p, d.Name()))
				if err != nil {
					return
				}
			}
		}
	}
}

// Reverse 将一个序列逆转(倒转)
func Reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
