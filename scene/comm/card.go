package comm

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

// 唯一标识一个卡牌
type CardInfo struct {
	Id    int
	Image *ebiten.Image

	Attack     int
	Shield     int
	SelfAttack int
	Cost       int
}

// 用于存放卡牌的图片
// key:value = 卡牌id:图片
var cardImageMap map[int]*ebiten.Image
var allCardsMap map[int]CardInfo

func readCSVFile(filename string) ([]CardInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // 读取并忽略表头
	if err != nil {
		return nil, err
	}

	var result []CardInfo
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		id, _ := strconv.Atoi(record[0])
		attack, _ := strconv.Atoi(record[1])
		shield, _ := strconv.Atoi(record[2])
		selfAttack, _ := strconv.Atoi(record[3])
		cost, _ := strconv.Atoi(record[4])

		item := CardInfo{
			Id:         id,
			Attack:     attack,
			Shield:     shield,
			SelfAttack: selfAttack,
			Cost:       cost,
		}

		result = append(result, item)
	}

	return result, nil
}
