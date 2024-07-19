package comm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 750
)

const (
	KakaActTag = iota
)

var (
	Lieren = "./pic/lieren.jpg"
	Kaka   = "./pic/kaka.jpg"
)

var (
	LocalCharacter = Character{}
	LocalEnemy     = Enemy{}
)

// 唯一标识一个卡牌
type CardInfo struct {
	Id    int
	Image *ebiten.Image
}

func init() {
	cha, _, err := ebitenutil.NewImageFromFile(Lieren) // 猎人的图片
	if err != nil {
		fmt.Printf("failed to get lieren pic, and err is: %s\n", err.Error())
	}

	LocalCharacter = Character{
		Image:   cha,
		Hp:      99,
		Hplimit: 99,
		Energy:  3,
	}

	allCards := GetCards()

	// gametmp := &models.Game{
	// 	Round: 1,

	// 	Character: Character{
	// 		Image:     cha,
	// 		Hp:        99,
	// 		Hplimit:   99,
	// 		Energy:    99,
	// 		Cards:     allCards,
	// 		DrawCards: allCards,
	// 	},

	// 	DraggingIndex: -1,
	// 	ExpandIndex:   -1,
	// 	IsDragging:    false,
	// }
	LocalCharacter = Character{
		Image:     cha,
		Hp:        99,
		Hplimit:   99,
		Energy:    99,
		Cards:     allCards,
		DrawCards: allCards,
	}
}

func GetLocalCharacter() *Character {
	LocalCharacter.Energy = 3
	return &LocalCharacter
}

// 无用函数

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap = map[string]*ebiten.Image{}

func init() {
	files, ids, err := listDir(cardDir)
	if err != nil {
		fmt.Printf("failed to get files, and err is: %s\n", err.Error())
		return
	}

	for i := range files {
		tmpImage, _, err := ebitenutil.NewImageFromFile(files[i])
		if err != nil {
			fmt.Printf("failed to get image: %s, and err is: %s\n", files[i], err.Error())
			continue
		}
		cardImageMap[ids[i]] = tmpImage
	}
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
