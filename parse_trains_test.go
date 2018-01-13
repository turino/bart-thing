package bart

import (
	"testing"
)

const testJSON = `{"root":{"station":[{"name":"Montgomery St.","abbr":"MONT","etd":[{"destination":"SF Airport","abbreviation":"SFIA","limited":"0","estimate":[{"minutes":"5","platform":"1","direction":"South","length":"10","color":"YELLOW","hexcolor":"#ffff33","bikeflag":"1","delay":"0"}]},{"destination":"Warm Springs","abbreviation":"WARM","limited":"0","estimate":[{"minutes":"6","platform":"2","direction":"North","length":"10","color":"GREEN","hexcolor":"#339933","bikeflag":"1","delay":"0"},{"minutes":"21","platform":"2","direction":"North","length":"10","color":"GREEN","hexcolor":"#339933","bikeflag":"1","delay":"0"}]}]}]}}`

func TestParseTrains(t *testing.T) {
	trains := ParseTrains([]byte(testJSON))

	if len(trains) != 3 {
		t.Error("message: ", trains[0].Location())
	}
}
