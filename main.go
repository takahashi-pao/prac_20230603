package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var count int

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		/*
			command => curl 'localhost:8080/greet?name={name}}'
		*/

		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		handlerHello(w, r, name)
	})
	http.HandleFunc("/afternoon", afternoonAPIHandler)
	http.HandleFunc("/afternoonPage", afternoonPageHandler)
	http.HandleFunc("/src/script.js", scriptHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/file", requestHtmlFileOnServer)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	response, err := ioutil.ReadFile("./src/script.js")
	if err != nil {
		// エラー処理
		fmt.Fprintln(w, "Error:", err)
		return
	}

	fmt.Fprintf(w, string(response))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	imagePath := "./favicon.ico" // 画像ファイルのパス

	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		fmt.Fprintln(w, "Error:", err)
		return
	}
	defer file.Close()

	// MIMEタイプの設定
	w.Header().Set("Content-Type", "image/ico")

	// ファイルデータをレスポンスに書き込む
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write image data", http.StatusInternalServerError)
		return
	}
}

func afternoonAPIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "ok"}`)
}

func afternoonPageHandler(w http.ResponseWriter, r *http.Request) {
	response, err := ioutil.ReadFile("./src/afternoon.html")
	if err != nil {
		// エラー処理
		fmt.Fprintln(w, "Error:", err)
		return
	}

	fmt.Fprintf(w, string(response))
}

func handlerHello(w http.ResponseWriter, r *http.Request, name string) {
	/*
		message := strings.Replace("Hello, {name}", "{name}", name, 1)
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s\n", name)
}

// githubサーバー上に保存されたHtmlファイルの取得😃
func requestHtmlFileOnServer(w http.ResponseWriter, r *http.Request) {
	// github上のファイルデータをGetリクエストで取得
	//response, err := http.Get("https://takahashi-pao.github.io/oretachi-omaetachi/index.html")
	// if err != nil {
	// 	// エラー処理
	// 	fmt.Fprintln(w, "Error:", err)
	// 	return
	// } else if response.StatusCode == http.StatusNotFound {
	// 	// 404処理
	// 	fmt.Fprintln(w, "ファイルが見つかりませんでした。")
	// 	return
	// }
	// defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil || response.StatusCode == http.StatusNotFound {
	// 	// エラー処理
	// 	fmt.Fprintln(w, "Error:", err)
	// 	return
	// }

	response, err := ioutil.ReadFile("./src/index.html")
	if err != nil {
		// エラー処理
		fmt.Fprintln(w, "Error:", err)
		return
	}

	fmt.Fprintf(w, string(response))
}
