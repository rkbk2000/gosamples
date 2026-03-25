package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/rkbk2000/samples/access"
)

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

// CommentDemo - This demonstrates golint
func CommentDemo() {

	a := 101
	fmt.Println(a)
	fmt.Print()
}

type JsonStruct struct {
	CapacityVal int
	TagName     string `json:"tagName"`
}

func update(id string) {
	hostProps := make(map[string]interface{})
	hostProps["codepage"] = "s1"
	hostProps["os_description"] = "d1"
	hostProps["language"] = "l1"
	hostProps["arr"] = []string{"s1", "s2"}
	_host, er := json.Marshal(hostProps)
	if er != nil {
		log.Println("ERROR: Unable to update Host details", er)
		return
	}
	ub := fmt.Sprintf(`{"properties":%s}`, string(_host))
	fmt.Println(ub)
}

func printStrings() {
	var tm []*string
	s1 := "123"
	s2 := "345"
	tm = append(tm, &s1)
	tm = append(tm, &s2)
	fmt.Printf(" Timestamps:")
	for _, s := range tm {
		fmt.Printf(" %s ", *s)
	}
	fmt.Println()
}

func main() {
	checkMapCopy()
	fmt.Println("Press enter to continue:")
	var val int
	fmt.Scanf("%d", val)

	s1 := JsonStruct{1, "a"}
	b, e := json.Marshal(&s1)
	if e == nil {
		fmt.Println(string(b))
	}
	netpanic()
	debug.PrintStack()
	printStrings()
	GenericsExample()

	syncMapTest()

	raceDemo()
	matches("monitor:14a6b829-6100-49a3-ba3b-a0ecebedab67:resourceList")
	matches("monitor:14a6b829-6100-49a3-ba3b-a0ecebedab67:selfHealth")
	matches("monitor:14a6b829-6100-49a3-ba3b-a0ecebedab67:other")
	fmt.Printf("Value of exvar: %v", access.ExVar)
	//runTogglingChannel("d1")
	bufferedChannelDemo()
	checkTicker()
	// var c chan bool // nil channel

	update("i1")
	checkResponse()
	checkPassByValue()

	str1 := "17.23.34.56 a.com"
	str2 := " 17.23.34.56 a.com "

	checkTokens(str1)
	checkTokens(str2)
	//checkCopy()
	checkContext(5)

	printAll()

	runBot()

	checkMemForInts(100000)

	checkMemForInts(1000000)

	checkMemForStrings(32, 100000)

	checkMemForStrings(32, 1000000)

	fmt.Println("Press enter to continue:")
	fmt.Scanf("%d", val)

	testCrypto()
}

func checkMapCopy() {
	msg := Msg{}
	msg.Headers = make(map[string]string)
	msg.Headers["h1"] = "v"
	msg.Headers["h2"] = "v2"
	fmt.Printf("before:%v\n", msg)
	handleMsg(&msg)
	fmt.Printf("after:%v\n", msg)
}

func checkTokens(in string) {
	tokens := strings.SplitAfterN(in, " ", 2)

	for idx, v := range tokens {
		fmt.Println("token[" + strconv.Itoa(idx) + "]:\"" + v + "\"")
	}
}

func checkContext(timeout int) {
	log.Println("Running context example")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Context done exiting")
			return
		default:
			toSleep := (time.Duration(timeout) + 5) * time.Second
			log.Println("Sleeping for: ", toSleep)
			time.Sleep(toSleep)
			log.Println("Woken up")
		}
	}(ctx)

	<-ctx.Done()
	log.Println("Context example completed")
	startSimpleReadWrite()
	printSum()
	deferTest()
}
