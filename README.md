# Go勉強会

## 勉強会の目的

* Go の基本的な文法や機能, ライブラリの使い方を学習する
* Go でプログラムを実装し, 実行できるようになる

## Goとは

* Googleによって開発されているプログラミング言語. Cの開発者が設計に関わっており, Cの文法と似ている
* シンプルで複雑ではない(対人間 & PC)ことを重視している
** 継承がない
** ジェネリックがない
** オーバーロードがない
* 並行処理(concurrency)に関する機能が言語レベルで提供されている
** goroutine
** channel

## Goプログラミング

今日のコード一覧の取得

```sh
$ go get github.com/phelrine/go-practice/...
```

### GOPATHの設定

Goのプログラムを扱うためのワークスペースを設定する. 今回は `${HOME}/go` を対象ディレクトリとする

```sh
$ mkdir ${HOME}/go
$ export GOPATH=${HOME}/go
```

シェルの環境変数の設定項目に記述しておくと, 実行時に毎回設定する必要がなくなる.
Goのプロジェクトは基本的にGOPATH以下に作成するようにする.

### Hello World

標準出力へ文字列を出力するプログラムを作成する

```sh
$ cat go/src/github.com/phelrine/go-practice/hello/main.go
```

```go
package main

import (
        "fmt"
)

func main() {
        fmt.Println("Hello World")
}
```

実行

パッケージのディレクトリまで移動する

```
$ cd go/src/github.com/phelrine/go-practice/hello/
```

コンパイルしたバイナリを直接実行

```sh
$ go run main.go
```

ファイルをコンパイルして実行

```sh
$ go build
$ ./hello
```

プロブラムをインストールして実行

```sh
$ go install
$ $GOPATH/bin/hello
```

### 型と変数

```go
package main

import (
	"fmt"
)

func main() {
	// 宣言と代入
	var a int
	a = 1
	fmt.Printf("int a= %d\n", a)

	// 宣言と初期化を同時に行う
	b := 2
	fmt.Printf("int b = %d\n", b)

	c := "文字列"
	fmt.Printf("string c = %s\n", c)

	d := 1.0
	fmt.Printf("float d = %0.3f\n", d)

	type myInt int
	var e myInt
	e = 1
	fmt.Printf("myInt e = %d\n", e)
}
```

Goのコンパイラは型推論機能をもっているため, 変数宣言の際に型情報を書く必要がない.

### 演算子と関数

Goでは基本的な算術演算子や, 文字列への操作が組み込み関数として提供されている.

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 算術演算
	fmt.Printf("2 + 3 = %d\n", 2+3)
	fmt.Printf("4 %% 2 = %d\n", 4%2)

	// 文字列演算
	str1 := "abc" + "def"
	fmt.Printf("str1 = %s\n", str1)
	str1 += "ghi"
	fmt.Printf("str1 = %s\n", str1)
	ch := str1[2]
	fmt.Printf("str1[2] = %c\n", ch)
	fmt.Printf("len(str1) = %d\n", len(str1))

	// 日本語
	str2 := "テスト"
	fmt.Printf("len(\"テスト\") = %d\n", len(str2))
	fmt.Printf("len(\"テスト\") = %d\n", utf8.RuneCountInString(str2))
}
```

Goはfuncキーワードを使って引数と返り値を指定した関数を定義できる.
Goの関数は多値を返す関数を定義できる.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("calc(2, 3) = %d\n", calc(2, 3))
	quo, remain := calcQuoAndRemain(10, 3)
	fmt.Printf("quo = %d, remain = %d\n", quo, remain)

	// 関数を変数に代入して利用することができる
	f1 := func(x int) int {
		return 5 * x
	}
	fmt.Printf("f1(1) = %d\n", f1(1))
	fmt.Printf("f1(3) = %d\n", f1(3))
}

func calc(a int, b int) int {
	return a + b
}

func calcQuoAndRemain(a, b int) (int, int) {
	return a / b, a % b
}
```

### 制御構文

```go
package main

import (
	"fmt"
)

func main() {
	// if
	a := 5
	if a > 10 {
		fmt.Printf("a > 10\n")
	} else if a > 20 {
		fmt.Printf("a > 20\n")
	} else {
		fmt.Printf("else\n")
	}

	// for
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// 無限ループ
	for {
		fmt.Println("loop")
	}
}
```

#### 練習問題
整数を1から100まで出力するプログラムを作成する.
ただし値が偶数の場合, 数字の後に!追加して出力する.

出力例

```
1
2!
3
4!
...
```

### FizzBuzzの実装

1から100までの数値について、順番に数値を出力する。
ただし、3の倍数の場合は数値ではなく「Fizz」、5の倍数は「Buzz」、15の倍数は「Fizz Buzz」と出力する。

### データ構造(struct, slice, map)

```go
package main

import (
	"fmt"
)

func main() {
	// 構造体
	type myStruct struct {
		field1 string
		field2 int
	}
	s1 := myStruct{
		field1: "f1",
		field2: 2,
	}
	fmt.Printf("s1.field1 = %s\n", s1.field1)
	fmt.Printf("s1.field2 = %d\n", s1.field2)

	// スライス
	var slice1 []int
	// 10個の要素を持つスライスを作成
	slice1 = make([]int, 10)
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i
	}
	// 要素列挙
	for i, e := range slice1 {
		fmt.Printf("slice1[%d] = %d\n", i, e)
	}
	slice1 = append(slice1, 100)
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	fmt.Printf("%d\n", slice1[10])

	slice2 := []int{}
	for i := 0; i < 10; i++ {
		slice2 = append(slice2, i)
	}

	// スライスの連結
	slice3 := append(slice1, slice2...)
	for _, e := range slice3 {
		fmt.Printf("%d\n", e)
	}

	map1 := map[string]int{}
	map1["abc"] = 1
	map1["123"] = 2

	// mapのサイズ取得
	fmt.Printf("len(map1) = %d\n", len(map1))

	// 要素アクセス
	fmt.Printf("map[\"abc\"] = %d\n", map1["abc"])

	// キーチェック
	if e, ok := map1["xxx"]; ok {
		fmt.Printf("key is exists %d\n", e)
	} else {
		fmt.Println("key is not exists")
	}

	// 要素列挙
	for k, v := range map1 {
		fmt.Printf("map1[%s] = %d\n", k, v)
	}

	// キーのみ
	for k := range map1 {
		fmt.Printf("key is %s\n", k)
	}
}

```

### lsコマンドの実装
指定したディレクトリ内のファイル一覧を出力するプログラムを作成する.

* 利用パッケージ
** `fmt`
** `io/ioutil`
** `os`

ヒント
* `ioutil.ReadDir(string) ([]os.FileInfo, error) ` 特定ディレクトリのファイル一覧を出力する
* `os.FileInfo` の `Name()` でエントリー名を取得できる

### goroutineを利用したsleep sortを実装する
[sleep sort](https://qiita.com/snsk/items/33c01951ef27bbd2b093) のアルゴリズムを元にgoroutineを利用してスリープソートを実装する.
