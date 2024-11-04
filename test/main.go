package main

import (
	"encoding/json"
	"fmt"
)

// 定义要转换的结构
type Price struct {
	Id    int `json:"Id"`
	Type  int `json:"Type"`
	Name  int `json:"Name"`
	Count int `json:"Count"`
}

type NewPrice struct {
	Sele  int `json:"sele"`
	Count int `json:"Count"`
}

type RoomDecorationConfig struct {
	Price []NewPrice `json:"price"` // 保持为 NewPrice 类型
}

type RoomPaintConfig struct {
	Price []NewPrice `json:"price"` // 保持为 NewPrice 类型
}

type Room struct {
	Lvs []struct {
		EntranceFee []Price `json:"entranceFee"`
	} `json:"lvs"`
	RoomDecorationConfigs []RoomDecorationConfig `json:"roomDecorationConfigs"`
	RoomPaintConfigs      []RoomPaintConfig      `json:"roomPaintConfigs"` // 保持为 RoomPaintConfig 类型
}

type Hotel struct {
	Reception struct {
		Lvs []struct {
			Price []Price `json:"price"`
		} `json:"lvs"`
	} `json:"reception"`
	Areas []struct {
		Price []Price `json:"price"`
	} `json:"areas"`
	Cleaners []struct {
		Lvs []struct {
			Price []Price `json:"price"`
		} `json:"lvs"`
	} `json:"cleaners"`
	RoomGuidance []struct {
		Lvs []struct {
			Price []Price `json:"price"`
		} `json:"lvs"`
	} `json:"roomGuidance"`
	Rooms []Room `json:"rooms"` // 保持为 Room 类型
}

type Hotels struct {
	Hotels []Hotel `json:"hotels"`
}

func main() {
	// 原始JSON字符串
	jsonStr := `{
		"hotels": [
			{
				"reception": {
					"lvs": [
						{
							"price": [
								{"Id": 3, "Type": 1, "Name": 3, "Count": 100}
							]
						}
					]
				},
				"areas": [
					{
						"price": [
							{"Id": 3, "Type": 1, "Name": 3, "Count": 100}
						]
					}
				],
				"cleaners": [
					{
						"lvs": [
							{
								"price": [
									{"Id": 3, "Type": 1, "Name": 3, "Count": 10}
								]
							}
						]
					}
				],
				"roomGuidance": [
					{
						"lvs": [
							{
								"price": []
							}
						]
					}
				],
				"rooms": [
					{
						"lvs": [
							{
								"entranceFee": [
									{"Id": 3, "Type": 1, "Name": 3, "Count": 100}
								]
							}
						],
						"roomDecorationConfigs": [
							{
								"price": [
									{"Id": 3, "Type": 1, "Name": 3, "Count": 100}
								]
							}
						],
						"roomPaintConfigs": [
							{
								"price": [
									{"Id": 3, "Type": 1, "Name": 3, "Count": 100}
								]
							}
						]
					}
				]
			}
		]
	}`

	// 解析原始JSON字符串
	var hotels Hotels
	if err := json.Unmarshal([]byte(jsonStr), &hotels); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 遍历并转换数据
	for i := range hotels.Hotels {
		// 转换 reception
		for j := range hotels.Hotels[i].Reception.Lvs {
			hotels.Hotels[i].Reception.Lvs[j].Price = convertPrices(hotels.Hotels[i].Reception.Lvs[j].Price)
		}

		// 转换 areas
		for j := range hotels.Hotels[i].Areas {
			hotels.Hotels[i].Areas[j].Price = convertPrices(hotels.Hotels[i].Areas[j].Price)
		}

		// 转换 cleaners
		for j := range hotels.Hotels[i].Cleaners {
			for k := range hotels.Hotels[i].Cleaners[j].Lvs {
				hotels.Hotels[i].Cleaners[j].Lvs[k].Price = convertPrices(hotels.Hotels[i].Cleaners[j].Lvs[k].Price)
			}
		}

		// 转换 roomGuidance
		for j := range hotels.Hotels[i].RoomGuidance {
			for k := range hotels.Hotels[i].RoomGuidance[j].Lvs {
				hotels.Hotels[i].RoomGuidance[j].Lvs[k].Price = convertPrices(hotels.Hotels[i].RoomGuidance[j].Lvs[k].Price)
			}
		}

		// 转换 rooms
		for j := range hotels.Hotels[i].Rooms {
			for k := range hotels.Hotels[i].Rooms[j].Lvs {
				hotels.Hotels[i].Rooms[j].Lvs[k].EntranceFee = convertPrices(hotels.Hotels[i].Rooms[j].Lvs[k].EntranceFee)
			}
			for k := range hotels.Hotels[i].Rooms[j].RoomDecorationConfigs {
				hotels.Hotels[i].Rooms[j].RoomDecorationConfigs[k].Price = convertPrices(hotels.Hotels[i].Rooms[j].RoomDecorationConfigs[k].Price)
			}
			for k := range hotels.Hotels[i].Rooms[j].RoomPaintConfigs {
				hotels.Hotels[i].Rooms[j].RoomPaintConfigs[k].Price = convertPrices(hotels.Hotels[i].Rooms[j].RoomPaintConfigs[k].Price)
			}
		}
	}

	// 转换后的JSON
	convertedJson, _ := json.MarshalIndent(hotels, "", "  ")
	fmt.Println(string(convertedJson))
}

// convertPrices 用于转换价格数组的结构
func convertPrices(prices []Price) []Price {
	var newPrices []Price
	for _, price := range prices {
		newPrices = append(newPrices, Price{
			Id:    price.Id,
			Type:  price.Type,
			Name:  price.Name,
			Count: price.Count,
		})
	}
	return newPrices
}
