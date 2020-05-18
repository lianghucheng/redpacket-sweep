package common

import (
	"fmt"
	"math"
	"math/rand"
	"redpacket-sweep/conf"
	"strconv"
	"strings"
	"time"

	"github.com/name5566/leaf/timer"
)

var (
	str = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z", "A", "B", "C", "D",
		"E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z", "~", "!", "@", "#", "$", "%", "^", "&",
		"*", "(", ")", "-", "_", "=", "+", "[", "]", "{",
		"}", "|", "<", ">", "?", "/", ".", ",", ";", ":",
	}

	numberStr = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 停止定时器的同时将定时器置空
func StopTimer(t *timer.Timer) *timer.Timer {
	if t != nil {
		t.Stop()
		t = nil
	}
	return t
}

func GetRandomString(n int) string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, str[rand.Intn(90)]) // 90 是 str 的长度
	}
	return strings.Join(s, "")
}

// 四舍五入，保留n位小数
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}

func Round2(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

/*20190813*/
func TimeFormat() string {
	return time.Now().Format("20060102")
}

func TranferChipRate(chips int64) float64 {
	return Round(float64(chips) / float64(conf.Server.ChipGactRate), 4)
}

func Int64ToFloat64ByRate(args []int64) []float64{
	rt := []float64{}
	for _, v := range args {
		rt = append(rt, TranferChipRate(v))
	}

	return rt
}