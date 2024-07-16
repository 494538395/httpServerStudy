package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type HoleData struct {
	ID          int `json:"id"`
	ObjectTrans struct {
		Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"position"`
		Rotation struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"rotation"`
		Scale struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"scale"`
	} `json:"objectTrans"`
	Status      string `json:"status"`
	AdsLocked   int    `json:"adsLocked"`
	Screw       int    `json:"screw"`
	KeyUnlock   int    `json:"keyUnlock"`
	KeyPosition struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"keyPosition"`
}

type WoodData struct {
	ID           int    `json:"id"`
	Layer        int    `json:"layer"`
	WoodType     string `json:"woodType"`
	ObjTransform struct {
		Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"position"`
		Rotation struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"rotation"`
		Scale struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"scale"`
	} `json:"objTransform"`
	Sliced int     `json:"sliced"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type LevelData struct {
	LevelID    int        `json:"levelId"`
	HoleData   []HoleData `json:"holeData"`
	WoodData   []WoodData `json:"woodData"`
	KeysData   []struct{} `json:"keysData"`
	BoomData   []struct{} `json:"boomData"`
	LevelStage []struct{} `json:"levelStage"`
}

type LevelSummary struct {
	Levels []struct {
		LevelID    int        `json:"levelId"`
		HoleData   []HoleData `json:"holeData"`
		WoodData   []WoodData `json:"woodData"`
		KeysData   []struct{} `json:"keysData"`
		BoomData   []struct{} `json:"boomData"`
		LevelStage []struct{} `json:"levelStage"`
	} `json:"levels"`
}

type Chapter struct {
	LevelID   int      `json:"levelId"`
	LevelStep []string `json:"levelStep"`
}

type NewChapter struct {
	ChapterId int     `json:"chapterId"`
	LevelStep []int32 `json:"levelStep"`
}

type LevelConfig struct {
	LevelID int `json:"levelId"`
}

func main() {

	genLevelData()

	//genChapter()
}

// parseLevelID 解析关卡ID
func parseLevelID(fileName string) (int, error) {
	if strings.HasPrefix(fileName, "localConfig_Level") {
		// 提取数字部分
		levelIDStr := strings.TrimPrefix(fileName, "localConfig_Level")
		levelIDStr = strings.TrimSuffix(levelIDStr, ".json")
		levelID, err := strconv.Atoi(levelIDStr)
		if err != nil {
			return 0, err
		}
		// 关卡Id前拼接一个7
		atoi, err := strconv.Atoi(fmt.Sprintf("%d%d", 7, levelID))
		if err != nil {
			panic(err)
		}
		return atoi, nil
	}

	// 定义匹配不同格式的正则表达式
	var re = regexp.MustCompile(`(?:localConfig_(?:falcon_level_|Level))(\d+)`)

	// 查找匹配项
	matches := re.FindStringSubmatch(fileName)
	if len(matches) == 2 {
		return strconv.Atoi(matches[1])
	}

	return 0, fmt.Errorf("无法从文件名中解析关卡ID: %s", fileName)
}

// genLevelData 汇总关卡json
func genLevelData() {
	falonCnt := 0
	normalCnt := 0
	jsonCnt := 0

	dir, _ := os.Getwd()
	fmt.Println("dir-->", dir)

	// Specify the output file where the summarized JSON will be written
	outputFile := "./game/nuts/level/levelData.json"

	// Specify the input directory where JSON files are located
	//inputDir := "C:\\Users\\青雉\\Documents\\WXWork\\1688855625500491\\Cache\\File\\2024-06\\GameLevel\\GameLevel"
	inputDir := "C:\\Users\\青雉\\Documents\\WXWork\\1688855625500491\\Cache\\File\\2024-07\\localConfig_Level"
	// Read all files in the input directory
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		return
	}
	var levelSummary LevelSummary
	levelSummary.Levels = make([]struct {
		LevelID    int        `json:"levelId"`
		HoleData   []HoleData `json:"holeData"`
		WoodData   []WoodData `json:"woodData"`
		KeysData   []struct{} `json:"keysData"`
		BoomData   []struct{} `json:"boomData"`
		LevelStage []struct{} `json:"levelStage"`
	}, 0)

	// Iterate through each file in the directory
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(inputDir, file.Name())
			jsonFile, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			defer jsonFile.Close()
			jsonCnt++

			byteValue, _ := ioutil.ReadAll(jsonFile)

			var level LevelData
			err = json.Unmarshal(byteValue, &level)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				panic(err)
			}

			//levelID, err := parseLevelID(file.Name())
			//if err != nil {
			//	panic(err)
			//}
			//level.LevelID = levelID

			if level.LevelID == 1841 {
				fmt.Println(1)
			}

			// Add empty levelStage array to each level
			level.LevelStage = make([]struct{}, 0)

			// Append to the summary
			levelSummary.Levels = append(levelSummary.Levels, struct {
				LevelID    int        `json:"levelId"`
				HoleData   []HoleData `json:"holeData"`
				WoodData   []WoodData `json:"woodData"`
				KeysData   []struct{} `json:"keysData"`
				BoomData   []struct{} `json:"boomData"`
				LevelStage []struct{} `json:"levelStage"`
			}{
				LevelID:    level.LevelID,
				HoleData:   level.HoleData,
				WoodData:   level.WoodData,
				KeysData:   level.KeysData,
				BoomData:   level.BoomData,
				LevelStage: level.LevelStage,
			})
		}
	}

	fmt.Println("falonCnt-->", falonCnt)
	fmt.Println("normalCnt-->", normalCnt)
	fmt.Println("jsonCnt-->", jsonCnt)

	// Marshal summary to JSON
	summaryJSON, err := json.Marshal(levelSummary)
	if err != nil {
		fmt.Println("Error marshaling summary:", err)
		return
	}

	// Write summarized JSON to output file
	err = ioutil.WriteFile(outputFile, summaryJSON, 0644)
	if err != nil {
		fmt.Println("Error writing summarized JSON to file:", err)
		return
	}

	fmt.Println("Summary written to", outputFile)
}

func stringsToInts(in []string) []int32 {
	out := make([]int32, 0, len(in))
	for _, s := range in {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			continue
		}
		out = append(out, int32(num))
	}
	return out
}

// extractLevelID 从文件名中提取关卡ID
func extractLevelID(filename string) (int, error) {
	if strings.HasPrefix(filename, "Level") {
		// 提取数字部分
		levelIDStr := strings.TrimPrefix(filename, "Level")
		levelID, err := strconv.Atoi(levelIDStr)
		if err != nil {
			return 0, err
		}
		// 关卡Id前拼接一个7
		//atoi, err := strconv.Atoi(fmt.Sprintf("%d%d", 7, levelID))
		atoi, err := strconv.Atoi(fmt.Sprintf("%d", levelID))

		if err != nil {
			panic(err)
		}
		return atoi, nil
	}

	// 定义正则表达式匹配不同的文件名格式
	patterns := []string{
		`falcon_level_(\d+)`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(filename)
		if len(matches) == 2 {
			levelID, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, fmt.Errorf("failed to convert level ID to integer: %v", err)
			}
			return levelID, nil
		}
	}

	return 0, fmt.Errorf("no matching pattern found for filename: %s", filename)
}

func genChapter() {
	// 读取章节配置文件
	chaptersFile := "./game/nuts/level/all_levels_tmp.json"
	chaptersData, err := ioutil.ReadFile(chaptersFile)
	if err != nil {
		log.Fatalf("无法读取章节配置文件: %v", err)
	}

	var chapters []Chapter
	err = json.Unmarshal(chaptersData, &chapters)
	if err != nil {
		log.Fatalf("解析章节配置文件失败: %v", err)
	}

	newChapters := make([]*NewChapter, 0, len(chapters))

	// 处理每个章节
	for i, chapter := range chapters {
		for j, levelFile := range chapter.LevelStep {

			levelID, err := extractLevelID(levelFile)
			if err != nil {
				panic(err)
			}

			findFileInCurrentDir(convertToLocalConfigFileName(levelFile))

			// 用关卡ID替换文件名
			chapters[i].LevelStep[j] = fmt.Sprintf("%d", levelID)
		}
		newChapters = append(newChapters, &NewChapter{
			ChapterId: chapter.LevelID,
			LevelStep: stringsToInts(chapter.LevelStep),
		})
	}

	// 读取现有的 `all_levels.json` 文件内容
	allLevelsFile := "./game/nuts/level/all_levels.json"

	allLevels := make(map[string]interface{}, 0)

	// 添加章节数据
	allLevels["chapters"] = newChapters

	// 输出更新后的章节配置
	updatedAllLevelsData, err := json.MarshalIndent(allLevels, "", "  ")
	if err != nil {
		log.Fatalf("序列化更新后的章节配置失败: %v", err)
	}

	err = ioutil.WriteFile(allLevelsFile, updatedAllLevelsData, 0644)
	if err != nil {
		log.Fatalf("写入更新后的章节配置失败: %v", err)
	}

	fmt.Println("章节配置更新成功")
}

func findFileInCurrentDir(fileName string) {
	// 构建完整的文件路径
	filePath := "C:\\Users\\青雉\\Documents\\WXWork\\1688855625500491\\Cache\\File\\2024-07\\localConfig_Level\\" + fileName

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("文件 %s 不存在", filePath)
	} else if err != nil {
		log.Fatalf("检查文件 %s 时出错: %v", filePath, err)
	} else {
		//log.Printf("文件 %s 存在", filePath)
	}
}

// convertToLocalConfigFileName 将文件名转换为 localConfig 格式的文件名
func convertToLocalConfigFileName(fileName string) string {
	// 提取数字部分
	parts := strings.Split(fileName, "Level")
	if len(parts) != 2 {
		log.Fatalf("无效的文件名格式: %s", fileName)
	}

	// 生成新文件名
	return fmt.Sprintf("localConfig_Level%s.json", parts[1])
}
