package network

import (
	"strings"
	"fmt"
)


type Net struct{
    Link interface{}
    Run func(Net)
    Value interface{}
}

func (nt Net) OnChange(){
    nt.Run(nt)
    if nt.Link != nil{
        nt = nt.Link.(Net)
        nt.OnChange()
    }
    
}

func (nt Net) Set(value interface{}){
    nt.Value = value
    nt.OnChange()
}



// func main(){
//     vv := 1
//     vs := 4
//     vvs := "hello,world"
//     nt3 := Net{
//         Run: strdev,
//         Value: &vvs,
//     }
//     nt2 := Net{
//         Link: nt3,
//         Run: dev,
//         Value: &vs,
//     }
//     nt := Net{
//         Link: nt2,
//         Run: runne,
//         Value: &vv,
//     }
//     nt.Set(&vv)
//     fmt.Println(vv)
//     fmt.Println(vs)
// }

// func runne(net Net){
//     val, _ := net.Value.(*int)
//     *val = *val +2
//     // fmt.Println(val)
    
//     if *val == 11{
//         return
//     }
// }

// func dev(net Net){
//     val, _ := net.Value.(*int)
//     *val = *val * 20
// }

// func strdev(net Net){
//     str, _ := net.Value.(*string)
//     strr := strings.Split(*str,",")
//     for _, ele := range strr{
//         fmt.Println(ele)
//     }
// }