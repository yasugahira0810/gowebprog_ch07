// リスト7.6
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// 関数NewDecoderにio.Readerを渡す（この場合はpost.xml）
	decoder := xml.NewDecoder(xmlFile)
	// decoder内のXMLデータを順次処理（Unmarshalはioutil.ReadAllでファイルの中身全部読み込んでいた）
	for {
		// decoderからトークン（Token）を取得。この場合、トークンはXML要素を表すインタフェース
		t, err := decoder.Token()
		// トークンがなくなるまでデコーダから取り出し続けたい。トークンがなくなるとerrに構造体io.EOFが代入される
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		// トークンの型を表示
		fmt.Printf("Token: %T\n", t)

		// tはdecoderから取得したトークン。トークンの型をチェックする。
		switch se := t.(type) {
		case xml.StartElement:
			// デコーダからトークンを取り出すたびに、XML要素の開始タグがチェックする
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment.Id)
				fmt.Println(comment.Content)
				fmt.Println(comment.Author)
			}
		}
	}
}
