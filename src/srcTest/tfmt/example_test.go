package tfmt

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

// The Errorf function lets us use formatting features
// to create descriptive error messages.
func TestExampleErrorf(t *testing.T) {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)			//%q %s都可以
	fmt.Println(err.Error())

	// Output: user "bueller" (id 17) not found
}


func TestExampleFscanf(t *testing.T) {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s )		//按照format格式从r中读取   %t布尔值
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
	// Output:
	// 5 true gophers
	// 3
}


func TestExampleFscanln(t *testing.T) {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
	// Output:
	// 3: dmr, 1771, 1.618034
	// 3: ken, 271828, 3.141590
}


func TestExampleSscanf(t *testing.T) {
	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

	// Output:
	// 2: Kim, 22
}


func ExampleSprint() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")
	s1 := fmt.Sprintln(name, "is", age, "years old.")			//有\n

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.
	io.WriteString(os.Stdout, s1)

	// Output:
	// Kim is 22 years old.
	// Kim is 22 years old.
}

func ExampleFprint() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")		//n返回字符串长度
	//n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")		//自动加\n

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)		//注意这个用法
	}
	fmt.Print(n, " bytes written.\n")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}


func Example_formats() {

	integer := 23
	// Each of these prints "23" (without the quotes).
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	fmt.Printf("%T %T\n", integer, &integer)
	// Result: int *int

	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// Result: true true

	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// Result: 42 42 2a 52 101010

	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// Result: 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00

	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)


	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// Result: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// Result: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// Result: &{Kim 22} 0x0

	// Arrays and slices are formatted by applying the format to each element.
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// Result: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// Result: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// Result: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

	now := time.Unix(123456789, 0).UTC() // time.Time implements fmt.Stringer.
	fmt.Printf("%v %q\n", now, now)
	// Result: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
}
