# Go Clean Architechture

### Install

```
go get -u github.com/prongbang/goclean
```

<img src="http://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg">
<center>ภาพจาก blog.cleancoder.com</center>

### Go module feature ด้วยคำสั่ง

```
export GO111MODULE=on
```

### Test

- 1
```
go test ./...
```

- 2

```
go test -v ./...
```

- 3

```
go test -cover ./...
```

### Test coverage report

```
make cover
```

or

```
go test -cover ./... -coverprofile=cover.out
go tool cover -html=cover.out -o coverage.html
```

### Benchmark

Benchmark สำหรับวัดประสิทธิภาพการทำงานของ code ทั้งความเร็ว ทั้งการใช้งาน resource ต่าง ๆ ทำการ run หลาย ๆ ครั้ง โดยที่ค่า b.N จะเพิ่มขึ้นเรื่อย ๆ จนกว่าผลการทำงานจะมีค่าที่เสถียร function ที่ต้องทำการทดสอบต้องอยู่ใน loop ลักษณะนี้เท่านั้น

#### อธิบายผลการทดสอบ
- xx ns/op ตัวเลขต่ำยิ่งเร็ว
- xx จำนวนสูงสุดที่ run ใน 1 วินาที

Reference: [มาใช้งาน Benchmark ในภาษา Go กัน](http://www.somkiat.cc/benchmark-in-golang/)


```go
func BenchmarkXxxx(b *testing.B) {
	// run function b.N times
	for n := 0; n < b.N; n++ {
		// Call function
	}
}
```

- Run Benchmar

```
go test -bench=. หรือ go test -bench=. -benchmem
```

Or

```
go test -bench=. ./...
```

- CPU profiling

```
go test -bench=. -benchmem -cpuprofile cpu.out
```

Or

```
go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile cpu.out
```

- วิเคราะห์ CPU profile ด้วย pprof

```
go tool pprof cpu.out
(pprof) top20 --cum
```
Reference: [profiling-go-programs](https://blog.golang.org/profiling-go-programs)


เปรียบเทียบผลการทดสอบด้วย [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp)

```
$go test -bench=. > old
$go test -bench=. > new

$benchcmp old new 
```


### REST API

#### ADD

- Request

```
POST http://localhost:1323/api/v1/promotion
```

Body

```
{
    "id": 1,
    "code": "sd-promo",
    "name": "Sunday promotion",
    "priority": 4,
    "exclusive": false,
    "used": 0,
    "couponBased": false,
    "rules": [],
    "actions": [],
    "createdAt": "2017-02-28T12:05:12+0100",
    "updatedAt": "2017-02-28T12:05:13+0100",
    "channels": [],
    "_links": {
        "self": {
            "href": "\/api\/v1\/promotions\/sd-promo"
        }
    }
}
```

#### GET ALL

- Request

```
GET http://localhost:1323/api/v1/promotion

```

#### GET BY ID

- Request

```
GET http://localhost:1323/api/v1/promotion/1

```
