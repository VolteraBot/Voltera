package twitch

import (
	"bytes"
	"fmt"
	"github.com/VolteraBot/Voltera/backend/types"
	"github.com/VolteraBot/Voltera/backend/utils"
)

const HOST string = "http://localhost:8080/mock/"
const CLIENT_ID string = "e454c274333d0b61280a9079e82cb9"
const TOKEN string = "f5f535547888efc"

func StartCommercial(broadcasterID int, length int) string {
	data := bytes.NewReader([]byte(fmt.Sprintf("{\"broadcaster_id\": \"%d\",\"length\": %d}", broadcasterID, length)))
	return utils.Post(HOST+"channels/commercial", types.Header{"Client-Id": CLIENT_ID, "Authorization": "Bearer " + TOKEN}, data)
}
