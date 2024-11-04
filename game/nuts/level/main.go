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

// 关卡
var (
	// 源文件。客户端给的 GameLevelData.zip
	levelDataSourceDir = "C:\\Users\\青雉\\Documents\\WXWork\\1688855625500491\\Cache\\File\\2024-08\\GameLevel\\GameLevel"
	// 生成文件
	newLevelDataOutputDir = "D:\\development\\goProject\\study\\httpServerStudy\\game\\nuts\\level\\output\\levelData.json"
)

// 章节
var (
	// 源文件。客户端给的 all_levels_tmp.json 关卡顺序文件
	chapterSourceDir = "D:\\development\\goProject\\study\\httpServerStudy\\game\\nuts\\level\\all_levels_tmp.json"
	// 生成文件
	chapterOutputDir = "D:\\development\\goProject\\study\\httpServerStudy\\game\\nuts\\level\\output\\all_levels.json"
)

// 规范章节
var (
	formatChapterSourceDir = "D:\\development\\goProject\\study\\httpServerStudy\\game\\nuts\\level\\format_chapter.json"
)

func main() {

	// 基于客户端给的 GameLevelData.zip 生成 levelData.json 文件，将生成到 levelData.json 上传到S3，然后替换配置后台上的关卡数据url
	genLevelData() // 生成关卡

	fmt.Println()

	// 基于客户端给到 all_levels_tmp.json 生成 all_levels.json，将生成到 all_levels.json 导入到配置后台到 chapter_config 然后发布
	//genChapter() // 生成章节

	fmt.Println()

	// 将章节配置文件中的 levelIds 改成 levelStep
	//formatChapterConfig()
}

// region --------------------------- 生成关卡数据文件 -------------------------

// genLevelData 生成关卡数据文件
func genLevelData() {
	log.Println("-----------开始生关卡数据-----------")
	log.Println("生成关卡数据文件路径:", newLevelDataOutputDir)
	defer log.Println("-----------生成关卡数据结束-----------")
	//falonCnt := 0
	//normalCnt := 0
	jsonCnt := 0

	//dir, _ := os.Getwd()
	//fmt.Println("dir-->", dir)

	// Read all files in the input directory
	files, err := ioutil.ReadDir(levelDataSourceDir)
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
			filePath := filepath.Join(levelDataSourceDir, file.Name())
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

	//fmt.Println("falonCnt-->", falonCnt)
	//fmt.Println("normalCnt-->", normalCnt)
	fmt.Println("jsonCnt-->", jsonCnt)

	// Marshal summary to JSON
	summaryJSON, err := json.Marshal(levelSummary)
	if err != nil {
		fmt.Println("Error marshaling summary:", err)
		return
	}

	// Write summarized JSON to output file
	err = ioutil.WriteFile(newLevelDataOutputDir, summaryJSON, 0644)
	if err != nil {
		fmt.Println("Error writing summarized JSON to file:", err)
		return
	}

	//fmt.Println("Summary written to", newLevelDataOutputDir)
}

// endregion --------------------------- 生成关卡数据文件 -------------------------

// region ------------------------- 生成章节配置文件 ------------------------------

// genChapter 生成章节配置文件
func genChapter() {
	log.Println("-----------开始生成章节配置-----------")
	log.Println("生成章节配置文件路径:", chapterOutputDir)
	defer log.Println("-----------生成章节配结束-----------")

	// 读取章节配置文件
	//chaptersFile := "./game/nuts/level/all_levels_tmp.json"
	chaptersData, err := ioutil.ReadFile(chapterSourceDir)
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
	//allLevelsFile := "./game/nuts/level/all_levels.json"

	allLevels := make(map[string]interface{}, 0)

	// 添加章节数据
	allLevels["chapters"] = newChapters

	// 输出更新后的章节配置
	updatedAllLevelsData, err := json.MarshalIndent(allLevels, "", "  ")
	if err != nil {
		log.Fatalf("序列化更新后的章节配置失败: %v", err)
	}

	err = ioutil.WriteFile(chapterOutputDir, updatedAllLevelsData, 0644)
	if err != nil {
		log.Fatalf("写入更新后的章节配置失败: %v", err)
	}

	fmt.Println("章节配置更新成功")
}

// stringsToInts 字符串数组转 int32 数组
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

// findFileInCurrentDir 判断文件是否存在
func findFileInCurrentDir(fileName string) {
	// 构建完整的文件路径
	filePath := levelDataSourceDir + "/" + fileName

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

// endregion ------------------------- 生成章节配置文件 ------------------------------

// region ------------------------------------------------------- 规章节配置 ---------------------------------

func formatChapterConfig() {
	// 读取 JSON 文件
	file, err := ioutil.ReadFile(formatChapterSourceDir)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 使用通用的 map 结构来解析 JSON
	var data map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 遍历 chapters 并修改 levelIds 为 levelStep
	if chapters, ok := data["chapters"].([]interface{}); ok {
		for _, chapter := range chapters {
			if chapterMap, ok := chapter.(map[string]interface{}); ok {
				// 将 levelIds 的值赋给 levelStep
				if levelIds, ok := chapterMap["levelIds"]; ok {
					chapterMap["levelStep"] = levelIds
					delete(chapterMap, "levelIds") // 删除原来的 levelIds 字段
				}
			}
		}
	}

	// 将修改后的数据重新编码为 JSON
	modifiedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// 将修改后的数据写回文件
	err = ioutil.WriteFile(formatChapterSourceDir, modifiedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("JSON file updated successfully!")
}

// endregion ----------------------------------------------------------------------------------------------------

// region ---------------------------------- 结构体 ----------------------------------------

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

// endregion ---------------------------------- 结构体 ----------------------------------------
