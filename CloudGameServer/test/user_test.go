package test

import (
	"CloudGameServer/db"
	"CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"runtime"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	user1 := user.User{
		Username: "",
		Mobile:   "15598870762",
		Avatar:   "",
		Password: "",
		Uid:      bilicoin.CreateMD5("12"),
	}

	println(user1.CreateToken())
}

func TestCheckToken(t *testing.T) {
	println(user.CheckToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTcxODk1NDgsImp0aSI6ImMyMGFkNGQ3NmZlOTc3NTlhYTI3YTBjOTliZmY2NzEwIiwiaXNzIjoicjNpbmIiLCJuYmYiOjE2MTcxMDMxNDh9.LPb20xTbGBtIu7CXi9g8m_eq8LY4-ZFVhIaJKZSN1OE"))
}

func TestExist(t *testing.T) {
	db.InitMongo()
	user1 := user.User{
		Username: "",
		Mobile:   "r3inbowari",
		Avatar:   "",
		Password: "",
	}
	println(user1.IsExist())
}

func TestLogin(t *testing.T) {
	db.InitMongo()

	var user = user.User{
		Mobile:   "15598870762",
		Password: "15598870762",
	}
	token, err := user.Login()
	//d := fmt.Sprintf("%s", unsafe.Pointer(token))
	fmt.Println(err)
	println(token)
}

func TestMem(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n",m.Alloc/1024)
	fmt.Printf("%d Kb\n",m.Sys/1024)

	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total/1024/1024, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)


	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Println("获取cpu信息出错")
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// cpu使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Println("cpu使用率", percent)


	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}