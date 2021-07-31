// Copyright (C) 2021 The Dank Grinder authors.
//
// This source code has been released under the GNU Affero General Public
// License v3.0. A copy of this license is available at
// https://www.gnu.org/licenses/agpl-3.0.en.html

package instance

import (
	"fmt"

	"github.com/dankgrinder/dankgrinder/discord"
)

func (in *Instance) search(msg discord.Message) {
	choices := msg.Components
	fmt.Println(choices)
}
