package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"html"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/rpc"
	"net/smtp"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"testing/quick"
	"text/template"
	"time"
	"unicode"
)

// fmt/xml package
type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

// fmt,"net/http"/"net/http/httptest" package
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, HTTP!")
}

// fmt/context packages
func processRequest(ctx context.Context) {
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println("Request has a deadline:", deadline)
	} else {
		fmt.Println("No deadline for the request.")
	}
}

// fmt/errors package
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divide by Zero")

	}
	return x / y, nil
}

// fmt/"testing/quick"
func Square(x int) int {
	return x * x
}

// fmt/"net/rpc"
type Args struct {
	A, B int
}
type Result int

var name string

func main() {
	//fmt/rand packages
	fmt.Println("My favourite number is:", rand.Intn(10))

	//fmt/time packages
	fmt.Println("The time is:", time.Now())

	// fmt/math packages
	x, y := 4.5, 3.2
	hypotenuse := math.Sqrt(x*x + y*y)
	fmt.Println("hypotenuse", hypotenuse)

	//fmt/ strings package
	fmt.Println(strings.Count("how are you doing", ""))

	//fmt/crypto/md5 package
	data := []byte("Hello, Go!")
	hash := md5.Sum(data)
	fmt.Printf("MD5 Hash: %x\n", hash)

	//log packages
	log.Println("welcome New user", "log in message")
	nname := "lee"
	log.Printf("my name is %s", nname)
	//log.Fatal("This is a fatal error message.")

	//fmt/json
	p := 56
	jsonData, _ := json.Marshal(p)
	fmt.Println("json data:", string(jsonData))

	//fmt/sort package
	xy := []int{45, 75, 84}
	sort.Ints(xy)
	fmt.Println("sort number:", xy)

	// fmt/os/exec package
	out, _ := exec.Command("ls", "-l").Output()
	fmt.Println("Command Output:")
	fmt.Println(string(out))

	// fmt/strconv package
	numberStr := "123"
	numberInt, _ := strconv.Atoi(numberStr)
	fmt.Println("String to Int:", numberInt)

	//fmt/"net/url" packages
	red := "liberty"
	parsedURL, _ := url.Parse(red)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query Params:", parsedURL.Query().Get("q"))

	//fmt/xml packages
	nex := Person{Name: "lee", Age: 30}
	datan, _ := xml.Marshal(nex)
	fmt.Println("XML data:", string(datan))

	//fmt/context packages
	ctx := context.Background()
	processRequest(ctx)

	//sync package
	var gp sync.WaitGroup
	count := 6
	for i := 0; i < count; i++ {
		gp.Add(1)
		go func(i int) {
			fmt.Printf("Goroutine %d is executing.\n", i)
			gp.Done()
		}(i)
	}

	gp.Wait()
	fmt.Println("All goroutines have finished.")

	// fmt/compress/gzip package
	file, _ := os.Create("data.txt.gz")
	writer := gzip.NewWriter(file)
	writer.Write([]byte("Compress this data using gzip!"))
	writer.Close()

	// fmt/regexp package
	str := "Hello world!"
	match, _ := regexp.MatchString("Hello, [A-Za-z]+!", str)
	fmt.Println("Match:", match)

	// "database/sql"/log package
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//fmt/flag package
	var name string
	var age int

	flag.StringVar(&name, "name", "", "Name of the person")
	flag.IntVar(&age, "age", 0, "Age of the person")

	flag.Parse()

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// net/http package
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, HTTP!")
	})
	http.ListenAndServe(":8080", nil)

	//fmt/syscall/os/signal/time package
	signalChannel := make(chan os.Signal, 1)

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C or send SIGTERM to terminate the program.")
	sig := <-signalChannel

	fmt.Printf("Received signal: %v\n", sig)
	time.Sleep(5 * time.Second)

	fmt.Println("Exiting...")

	//fmt/strings/io/ioutil package
	d := strings.NewReader("good afternoon sir")
	r, err := ioutil.ReadAll(d)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("%s", r)

	}

	//fmt/reflect package
	Value := 56
	reflectValue := reflect.ValueOf(Value)
	fmt.Println("Type:", reflectValue.Type())
	fmt.Println("Kind:", reflectValue.Kind())
	fmt.Println("VAlue:", reflectValue.Int())

	//fmt/mime
	filename := "package.jpg"
	mediaType := mime.TypeByExtension(".jgg")
	fmt.Printf("file %s has media type %s\n", filename, mediaType)

	//byte/os/io/mime/multipart package
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, _ := w.CreateFormFile("file", "data.txt")
	//file, _ := os.Open("data.txt")
	io.Copy(fw, file)
	file.Close()

	w.Close()

	fmt.Println("Multiple data:")
	fmt.Println(b.String())

	//fmt/html package
	original := "<b>hello, world<b>"
	escaped := html.EscapeString(original)
	fmt.Println("Escaped HTML:", escaped)

	unescaped := html.UnescapeString(escaped)
	fmt.Println("Unescaped HTML:", unescaped)

	//fmt/sync/atomic
	var counter int32

	atomic.AddInt32(&counter, 10)
	fmt.Println("Count:", atomic.AddInt32(&counter, 10))

	// fmt/path/filepath package
	abspath, _ := filepath.Abs("data.txt")
	fmt.Println("Absolute path", abspath)
	filenamem := filepath.Base(abspath)
	fmt.Println("Filename:", filenamem)
	ext := filepath.Base(abspath)
	fmt.Println("Extention:", ext)

	//fmt/rand/encoding/base64
	randomBytes := make([]byte, 10)
	rand.Read(randomBytes)
	fmt.Println("RandomBytes:", randomBytes)

	randomString := base64.StdEncoding.EncodeToString(randomBytes)
	fmt.Println("Random base64 string:", randomString)

	//fmt/"net/smtp"/log package
	from := "senderliberty234@gmail.com "

	pass := "yourpassword"

	to := "recieverebikade06@gmail.com"

	subject := "Test Email"

	body := "This is an email sent using iphone"

	msg := "From: " + from + "\n" +
		"To : " + to + "\n" +
		"Subject: " + subject + "\n" +
		body

	error := smtp.SendMail("smtp.liberty.com:606",
		smtp.PlainAuth("", from, pass, "smtp.ebikade.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Fatal(error)
	}
	log.Println("Email sent successfully.")

	//fmt/"text/template"/os package
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
    <h1>{{ .Header }}</h1>
    <p>{{ .Content }}</p>
</body>
</html>
`

	date := map[string]string{
		"Title":   "My Page",
		"Header":  "Welcome to My Page",
		"Content": "This is the content of my page.",
	}

	t, _ := template.New("webpage").Parse(tmpl)
	t.Execute(os.Stdout, date)

	//fmt/image/"image/color","image/png"/os package
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	blue := color.RGBA{0, 0, 255, 255}
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, blue)
		}

	}
	fille, _ := os.Create("output.png")
	defer fille.Close()
	png.Encode(fille, img)
	fmt.Println("image created:", "output.png")

	//fmt,"net/http"/"net/http/httptest" package
	req := httptest.NewRequest("Move", " http://booking-app.com", nil)
	web := httptest.NewRecorder()

	handler(web, req)

	resp := web.Result()
	fmt.Println("Status code", resp.StatusCode)

	//fmt/errors package
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//fmt/net
	Addrs, _ := net.LookupHost("liberty@gmail.com")
	fmt.Println("ebikade.IPaddrress.com")
	for _, addr := range Addrs {
		fmt.Println(addr)
	}

	//fmt,os,bufio package
	fiile, _ := os.Open("data.txt")
	defer fiile.Close()

	scanner := bufio.NewScanner(fiile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	//fmt/bytes/"compress/zlib"
	dater := []byte("world go")

	var buf bytes.Buffer
	writerr := zlib.NewWriter(&buf)
	writerr.Write(dater)
	writerr.Close()

	compressed := buf.Bytes()
	fmt.Println("Compressed:", compressed)

	reader, _ := zlib.NewReader(&buf)
	decompressed, _ := io.ReadAll(reader)
	fmt.Println("decompressd:", decompressed)

	//fmt/os/"archive/zip"
	fileToZip := "data.txt"
	zipfile, _ := os.Create("data.zip")
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	fele, _ := os.Open(fileToZip)
	defer fele.Close()

	wriiter, _ := archive.Create(fileToZip)
	io.Copy(wriiter, file)

	fmt.Println("File successfully zipped as data.zip")

	//fmt/"crypto/sha1"
	dataa := []byte("Hello, Go!")
	hashh := sha1.Sum(dataa)
	fmt.Printf("SHA1 Hash: %x\n", hashh)

	//fmt/"net/http"/"net/http/cookiejar"
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}

	respp, err := client.Get("https://httpbin.org/cookies/set?key=value")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Cookies:", jar.Cookies(respp.Request.URL))

	//fmt/"testing/quick"
	rr := 'A'
	fmt.Printf("Character '%c' is a letter: %t\n", r, unicode.IsLetter(rr))

	rr = '3'
	fmt.Printf("Character '%c' is a digit: %t\n", r, unicode.IsDigit(rr))

	f := func(x int) bool {
		return Square(x) == x*x
	}

	if err := quick.Check(f, nil); err != nil {
		fmt.Println("Property does not hold:", err)
	} else {
		fmt.Println("Property holds for all cases.")
	}

	//fmt/"net/rpc"
	clientt, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	args := Args{A: 5, B: 10}
	var resultt Result

	err = clientt.Call("MathService.Multiply", args, &resultt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", result)
}

// fmt/ testing package
func TestAdd(t *testing.T) {
	run := "fast"
	fmt.Println(run)
}
