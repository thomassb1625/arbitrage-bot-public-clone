package arbitrage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func InitBlackList(filename string) map[common.Address]bool {
	blacklist := make(map[common.Address]bool)
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &blacklist)
	return blacklist
}

func SaveBlackList(list map[common.Address]bool, filename string) {
	data, _ := json.Marshal(list)
	ioutil.WriteFile(filename, data, 0666)
}

func AddWhiteListed(list map[common.Address]int) map[common.Address]int {
	jsonFile, _ := os.Open("active_tokens.json")
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	resp, _ := UnmarshalResponse(bytes)
	for _, t := range resp.Tokens {
		a := common.HexToAddress(t.Address)
		list[a] = -100000
	}
	return list

}

// Check if any token needs to be blacklisted. Returns true if trading paths should be updated afterwards
func BlackistWasUpdated(toBlackList *map[common.Address]int, BLACKLIST *map[common.Address]bool, BLACKLIST_THRESHOLD int, logger log.Logger) bool {
	highest := 0
	badToken := common.HexToAddress("")
	for k, v := range *toBlackList {
		if v > BLACKLIST_THRESHOLD {
			(*toBlackList)[k] = 0
			if v > highest {
				highest = v
				badToken = k
			}
		}
	}

	if highest != 0 {
		logger.Printf("Blacklisted %v \n", badToken)
		(*BLACKLIST)[badToken] = true
		delete((*toBlackList), badToken)
		SaveBlackList((*BLACKLIST), "blacklist.json")
		return true
	}
	return false
}

func (bot *ArbBot) BlacklistWasUpdated(BLACKLIST_THRESHOLD int) bool {
	highest := 0
	badToken := common.HexToAddress("")
	for k, v := range bot.TokenReverts {
		if v > BLACKLIST_THRESHOLD {
			bot.TokenReverts[k] = 0
			if v > highest {
				highest = v
				badToken = k
			}
		}
	}

	if badToken != common.HexToAddress("") {
		bot.genLog.Printf("Blacklisted %v \n", badToken)
		bot.BLACKLIST[badToken] = true
		delete((bot.TokenReverts), badToken)
		SaveBlackList(bot.BLACKLIST, "blacklist.json")
		return true
	}
	return false
}
