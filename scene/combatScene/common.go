package combatscene

import (
	m "ebiten/scene/models"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
	imageWidth   = 150
	imageHeight  = 200
)

// TODO bugzzhou
// 其实可以直接融合到CombatScene中去，暂时保留，后续优化结构时，统一修改
type Game struct {
	Character m.Character

	Enemy m.Enemy

	ExpandIndex   int
	DraggingIndex int
	IsDragging    bool

	Cards        []CardInfo
	DrawCards    []CardInfo
	HandCards    []CardInfo
	DiscardCards []CardInfo

	Round int //回合数，后续用于计算buff的生效数值
}

func listDir(dir string) (filePaths []string, baseNames []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("err !!!\n")
			return err // 遇到错误时返回
		}
		if !info.IsDir() { // 确保是文件
			filePaths = append(filePaths, path)                                     // 添加完整路径到切片
			baseName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) // 去除文件后缀
			baseNames = append(baseNames, baseName)                                 // 添加去除后缀的文件名到切片
		}
		return nil
	})

	if err != nil {
		return nil, nil, err // 遇到错误时返回
	}

	return filePaths, baseNames, nil
}
