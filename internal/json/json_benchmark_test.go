package json

import (
	"testing"
	"encoding/json"
	"github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/mailru/easyjson"
	"github.com/ugorji/go/codec"
	"bufio"
	"strings"
	"bytes"
)

// 运行性能测试
// go test -bench=. *

type Book struct {
	BookId int     `json:"id"`
	Title  string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Hot    bool    `json:"hot"`
	Weight int     `json:"-"`
}

func BenchmarkMarshalStdJson(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(&book)
	}
}

func BenchmarkMarshalJsonIterator(b *testing.B) {
	book := Book{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		jsonIterator.Marshal(&book)
	}
}

func BenchmarkMarshalFfjson(b *testing.B) {
	book := FBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		ffjson.Marshal(&book)
	}
}

func BenchmarkMarshalEasyjson(b *testing.B) {
	book := EBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	for i := 0; i < b.N; i++ {
		easyjson.Marshal(&book)
	}
}

func BenchmarkMarshalCodecJson(b *testing.B) {
	book := EBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	buf := make([]byte, 0, 1024)
	jsonHandler := &codec.JsonHandle{}
	encoder := codec.NewEncoderBytes(&buf, jsonHandler)
	for i := 0; i < b.N; i++ {
		encoder.Encode(&book)
	}
}

func BenchmarkMarshalCodecJsonWithBufio(b *testing.B) {
	book := EBook{
		BookId: 12125924,
		Title:  "人类简史-从动物到上帝",
		Author: "尤瓦尔·赫拉利",
		Price:  40.8,
		Hot:    true,
		Weight: 100,
	}

	jsonHandler := &codec.JsonHandle{}
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, 1024))
		writer := bufio.NewWriter(buf)
		encoder := codec.NewEncoder(writer, jsonHandler)
		encoder.Encode(&book)
		writer.Flush()
	}
}

func BenchmarkUnMarshalStdJson(b *testing.B) {
	data := []byte(`{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	for i := 0; i < b.N; i++ {
		json.Unmarshal(data, &book)
	}
}

func BenchmarkUnMarshalJsonIterator(b *testing.B) {
	data := []byte(`{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for i := 0; i < b.N; i++ {
		jsonIterator.Unmarshal(data, &book)
	}
}

func BenchmarkUnMarshalFfjson(b *testing.B) {
	data := []byte(`{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book FBook

	for i := 0; i < b.N; i++ {
		ffjson.Unmarshal(data, &book)
	}
}

func BenchmarkUnMarshalEasyjson(b *testing.B) {
	data := []byte(`{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book EBook

	for i := 0; i < b.N; i++ {
		easyjson.Unmarshal(data, &book)
	}
}

func BenchmarkUnMarshalCodecJson(b *testing.B) {
	data := []byte(`{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	jsonHandler := &codec.JsonHandle{}
	decoder := codec.NewDecoderBytes(data, jsonHandler)
	for i := 0; i < b.N; i++ {
		decoder.Decode(&book)
	}
}

func BenchmarkUnMarshalCodecJsonWithBufio(b *testing.B) {
	data := []byte(`{"id":12125925,"name":"未来简史-从智人到智神","author":"尤瓦尔·赫拉利","price":40.8,"hot":true}`)
	var book Book

	jsonHandler := &codec.JsonHandle{}
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(string(data)))
		decoder := codec.NewDecoder(reader, jsonHandler)
		decoder.Decode(&book)
	}
}
