一个go版时间格式转换封装
安装: go get "github.com/ooodjc/day"
```
func main() {
	d := new(day.Day).New(time.Now())
	fmt.Printf("TimeToString:%s\n", d.TimeToString())
	fmt.Printf("TimeToInt:%d\n", d.TimeToInt())
}
```