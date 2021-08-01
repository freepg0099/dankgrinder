// Copyright (C) 2021 The Dank Grinder authors.
//
// This source code has been released under the GNU Affero General Public
// License v3.0. A copy of this license is available at
// https://www.gnu.org/licenses/agpl-3.0.en.html

package instance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dankgrinder/dankgrinder/discord"
)

func (in *Instance) search(msg discord.Message) {
	choices := msg.Components[0].Buttons[0].Label
	fmt.Println(choices)
	url := "https://discord.com/api/v9/interactions"

	data := map[string]interface{}{"component_type": msg.Components[0].Buttons[0].Type, "custom_id": msg.Components[0].Buttons[0].CustomID, "hash": msg.Components[0].Buttons[0].Hash}
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
