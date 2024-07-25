package comm

import (
	"ebiten/utils"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

// 唯一标识一个卡牌
type CardInfo struct {
	Id int

	Attack     int
	Shield     int
	SelfAttack int
	Cost       int

	Image *ebiten.Image
}

var allCardsMap map[int]CardInfo

func init() {
	allCardBaseinfo, err := readCSVFile(utils.CardInfoPath)
	if err != nil {
		fmt.Printf("failed to read csv, and err is: %s\n", err.Error())
	}
	for _, v := range allCardBaseinfo {
		v.Image = cardImageMap[v.Id]
		allCardsMap[v.Id] = v
	}
}

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
