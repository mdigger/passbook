package passbook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/kr/pretty"
)

func TestPassRead(t *testing.T) {
	var filenames = []string{"Generic" /*"BoardingPass",*/, "Coupon", "Event", "StoreCard"}
	for _, filename := range filenames {
		fmt.Println("---", filename)
		data, err := ioutil.ReadFile(fmt.Sprintf("samples/%s.raw/pass.json", filename))
		if err != nil {
			t.Error("Read error:", err)
			continue
		}
		var pass Pass
		if err := json.Unmarshal(data, &pass); err != nil {
			pretty.Println(err)
			t.Error("Parse Error:", err)
			continue
		}
		pretty.Println(pass)
	}
}
