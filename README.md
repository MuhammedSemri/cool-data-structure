# cool-data-structure

# How to use

`go get github.com/MuhammedSemri/cool-data-structure`

```
func main(){
    vv := 1
    vs := 4
    vvs := "hello,world"
    nt3 := Net{
        Run: strdev,
        Value: &vvs,
    }
    nt2 := Net{
        Link: nt3,
        Run: dev,
        Value: &vs,
    }
    nt := Net{
        Link: nt2,
        Run: runne,
        Value: &vv,
    }
    nt.Set(&vv)
    fmt.Println(vv)
    fmt.Println(vs)
}

func runne(net Net){
    val, _ := net.Value.(*int)
    *val = *val +2
    // fmt.Println(val)
    
    if *val == 11{
        return
    }
}

func dev(net Net){
    val, _ := net.Value.(*int)
    *val = *val * 20
}

func strdev(net Net){
    str, _ := net.Value.(*string)
    lstr := strings.Split(*str,",")
    for _, ele := range lstr{
        fmt.Println(ele)
    }
}

```
