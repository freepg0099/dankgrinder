package instance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/dankgrinder/dankgrinder/discord"
)

type Database struct {
	Database []TriviaDetail `json:"database"`
}

type TriviaDetail struct {
	Question string `json:"question"`
	Answer   string `json:"correct_answer"`
}

func (in *Instance) trivia(msg discord.Message) {

	details := exp.trivia.FindStringSubmatch(msg.Embeds[0].Description)[1:]
	question := details[0]

	ex, _ := os.Executable()
	ex = filepath.ToSlash(ex)
	jsonFile, err := os.Open((path.Join(path.Dir(ex), "trivia.json")))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	bytevalue, _ := ioutil.ReadAll(jsonFile)

	var database Database
	json.Unmarshal(bytevalue, &database)

	for i := 0; i < len(database.Database); i++ {
		if question == html.UnescapeString(database.Database[i].Question) {
			var answer = html.UnescapeString(database.Database[i].Answer)
			for i := 0; i < 4; i++ {
				if answer == msg.Components[0].Buttons[i].Label {
					url := "https://discord.com/api/v9/interactions"

					data := map[string]interface{}{"component_type": msg.Components[0].Buttons[i].Type, "custom_id": msg.Components[0].Buttons[i].CustomID, "hash": msg.Components[0].Buttons[i].Hash}
					values := map[string]interface{}{"application_id": "270904126974590976", "channel_id": in.ChannelID, "type": "3", "data": data, "guild_id": msg.GuildID, "message_flags": 0, "message_id": msg.ID}
					json_data, err := json.Marshal(values)

					if err != nil {
						fmt.Println(err)
					}
					req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
					req.Header.Set("authorization", in.Client.Token)
					req.Header.Set("Content-Type", "application/json")

					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()

					fmt.Println("response Status:", resp.Status)
					fmt.Println("response Headers:", resp.Header)
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("response Body:", string(body))
				}

			}
		}
	}
}
