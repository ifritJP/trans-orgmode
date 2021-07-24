package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
	"github.com/niklasfasching/go-org/org"
)

const (
	// 通常文字列
	ItemKindText = 0
	// * headline
	ItemKindHeadline = 1
	// | テーブル
	ItemKindTable = 2
	// :
	ItemKindVerb = 3
	// =
	ItemKindExp = 4
	// ブロック
	ItemKindBlock = 5
	// リスト
	ItemKindList = 6
	// ----
	ItemKindRule = 7
	//
	ItemKindBlank = 8
	//
	ItemKindName = 9
	// - term :: descript
	ItemKindDescript = 10
	//
	ItemKindEmphasis = 11
	// #+AAAA:
	ItemKindKeyword = 12
	// #
	ItemKindComment = 13
)

type Item struct {
	Kind LnsInt
	Id   LnsInt
}

func (self *Item) Get_Kind(_env *LnsEnv) LnsInt {
	return self.Kind
}
func (self *Item) Get_Id(_env *LnsEnv) LnsInt {
	return self.Id
}

type Keyword struct {
	Key   string
	TxtId int
}

func (self *Keyword) Get_Keyword(_env *LnsEnv) string {
	return self.Key
}
func (self *Keyword) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type Comment struct {
	TxtId int
}

func (self *Comment) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type Headline struct {
	Level    LnsInt
	TxtId    LnsInt
	CustomId LnsAny
}

func (self *Headline) Get_Level(_env *LnsEnv) LnsInt {
	return self.Level
}
func (self *Headline) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}
func (self *Headline) Get_CustomId(_env *LnsEnv) LnsAny {
	return self.CustomId
}

type TableRow struct {
	TxtIdList []LnsAny
}

func (self *TableRow) Get_TxtIdList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.TxtIdList)
}

type Table struct {
	RowList []LnsAny
}

func (self *Table) Get_RowList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.RowList)
}

type Verb struct {
	TxtId LnsInt
}

func (self *Verb) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type Exp struct {
	Delimit string
	TxtId   LnsInt
}

func (self *Exp) Get_Delimit(_env *LnsEnv) string {
	return self.Delimit
}
func (self *Exp) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type Block struct {
	TxtId LnsInt
}

func (self *Block) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type List struct {
	Level LnsInt
	TxtId LnsInt
}

func (self *List) Get_Level(_env *LnsEnv) LnsInt {
	return self.Level
}
func (self *List) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type Descript struct {
	Level  LnsInt
	TermId LnsInt
	TxtId  LnsInt
}

func (self *Descript) Get_Level(_env *LnsEnv) LnsInt {
	return self.Level
}
func (self *Descript) Get_TermId(_env *LnsEnv) LnsInt {
	return self.TermId
}
func (self *Descript) Get_TxtId(_env *LnsEnv) LnsInt {
	return self.TxtId
}

type Document struct {
	ItemList     []LnsAny
	TextList     []LnsAny
	HeadlineList []LnsAny
	TableList    []LnsAny
	VerbList     []LnsAny
	ExpList      []LnsAny
	BlockList    []LnsAny
	ListList     []LnsAny
	DescriptList []LnsAny
	KeywordList  []LnsAny
	CommentList  []LnsAny
}
type DocumentIF interface {
	GetItemList(_env *LnsEnv) *LnsList
	GetTextList(_env *LnsEnv) *LnsList
	GetHeadlineList(_env *LnsEnv) *LnsList
	GetTableList(_env *LnsEnv) *LnsList
	GetVerbList(_env *LnsEnv) *LnsList
	GetExpList(_env *LnsEnv) *LnsList
	GetBlockList(_env *LnsEnv) *LnsList
	GetListList(_env *LnsEnv) *LnsList
	GetDescriptList(_env *LnsEnv) *LnsList
	GetKeywordList(_env *LnsEnv) *LnsList
	GetCommentList(_env *LnsEnv) *LnsList
}

func (self *Document) GetItemList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.ItemList)
}
func (self *Document) GetTextList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.TextList)
}
func (self *Document) GetHeadlineList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.HeadlineList)
}
func (self *Document) GetTableList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.TableList)
}
func (self *Document) GetVerbList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.VerbList)
}
func (self *Document) GetExpList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.ExpList)
}
func (self *Document) GetBlockList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.BlockList)
}
func (self *Document) GetListList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.ListList)
}
func (self *Document) GetDescriptList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.DescriptList)
}
func (self *Document) GetKeywordList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.KeywordList)
}
func (self *Document) GetCommentList(_env *LnsEnv) *LnsList {
	return NewLnsList(self.CommentList)
}

func NewDocument() *Document {
	doc := Document{}
	doc.ItemList = make([]LnsAny, 0)
	doc.TextList = make([]LnsAny, 0)
	doc.HeadlineList = make([]LnsAny, 0)
	doc.TableList = make([]LnsAny, 0)
	doc.VerbList = make([]LnsAny, 0)
	doc.ExpList = make([]LnsAny, 0)
	doc.BlockList = make([]LnsAny, 0)
	doc.ListList = make([]LnsAny, 0)
	doc.DescriptList = make([]LnsAny, 0)
	doc.KeywordList = make([]LnsAny, 0)
	doc.CommentList = make([]LnsAny, 0)
	return &doc
}

func (self *Document) addItem(kind, id LnsInt) {
	self.ItemList = append(self.ItemList, &Item{kind, id})
}

func (self *Document) addText(txt string) LnsInt {
	id := len(self.TextList) + 1
	self.TextList = append(self.TextList, txt)
	return id
}
func (self *Document) addExp(delimit string, txt string) LnsInt {
	id := self.addText(txt)
	self.ExpList = append(self.ExpList, &Exp{delimit, id})
	return len(self.ExpList)
}

type DescriptDepth struct {
	level int
	term  bool
}

type Writer struct {
	validLog      bool
	depth         LnsInt
	space         string
	lineList      []string
	txt           strings.Builder
	emphasis      string
	lastEmphasis  string
	rawNo         LnsInt
	headlineLv    LnsInt
	listLv        LnsInt
	descriptDepth []*DescriptDepth
	tableRow      *TableRow
	doc           *Document
}

func NewWriter(enableLog bool) *Writer {
	writer := Writer{}
	writer.validLog = enableLog
	writer.doc = NewDocument()
	return &writer
}

func isOnlyAscii(txt string) bool {
	for index := 0; index < len(txt); index++ {
		if txt[index] >= 0x7f {
			return false
		}
	}
	return true
}

func (self *Writer) getNextRawId() string {
	self.rawNo++
	return fmt.Sprintf("sy_%d_", self.rawNo)
}

func (self *Writer) printf(format string, vals ...LnsAny) {
	if self.validLog {
		fmt.Printf(self.space)
		fmt.Printf(format, vals...)
	}
}

func (self *Writer) push() {
	self.depth++
	self.space += "   "
}
func (self *Writer) pop() {
	self.depth--
	self.space = self.space[:len(self.space)-3]
}

func (self *Writer) write(node org.Node) {
	self.push()
	org.WriteNodes(self, node)
	self.pop()
}

func (self *Writer) addItem(kind, id LnsInt) {
	self.doc.addItem(kind, id)
	self.lastEmphasis = ""
}

func (self *Writer) addTxt(txt string) {
	self.txt.WriteString(txt)
}
func (self *Writer) addCr() {
	if self.txt.Len() != 0 {
		self.txt.WriteString("\\n")
	} else {
		self.addItem(ItemKindBlank, 0)
	}
}

func (self *Writer) fixLine() {
	if self.txt.Len() != 0 {
		line := self.txt.String()
		self.txt.Reset()
		id := self.doc.addText(line)
		if self.tableRow != nil {
			self.tableRow.TxtIdList = append(self.tableRow.TxtIdList, id)
		} else if self.headlineLv != 0 {
			self.doc.HeadlineList = append(
				self.doc.HeadlineList, &Headline{self.headlineLv, id, nil})
			self.addItem(ItemKindHeadline, len(self.doc.HeadlineList))
		} else if self.listLv != 0 {
			depth := &DescriptDepth{}
			if len(self.descriptDepth) > 0 {
				depth = self.descriptDepth[len(self.descriptDepth)-1]
			}

			if depth.level == self.listLv {
				if depth.term {
					self.doc.DescriptList = append(
						self.doc.DescriptList, &Descript{self.listLv, id, 0})
					self.addItem(ItemKindDescript, len(self.doc.DescriptList))
				} else {
					descript := self.doc.DescriptList[len(self.doc.DescriptList)-1]
					(descript.(*Descript)).TxtId = id
				}
			} else {
				self.doc.ListList = append(
					self.doc.ListList, &List{self.listLv, id})
				self.addItem(ItemKindList, len(self.doc.ListList))
			}
		} else {
			if self.lastEmphasis == "*" || self.lastEmphasis == "/" {
				self.doc.ExpList = append(self.doc.ExpList, &Exp{self.lastEmphasis, id})
				self.addItem(ItemKindEmphasis, len(self.doc.ExpList))
			} else {
				self.addItem(ItemKindText, id)
			}
		}
	}
}

func (self *Writer) fixDoc() {
	if self.txt.Len() != 0 {
		self.fixLine()
	}
}

func (self *Writer) GetDocument() *Document {
	return self.doc
}

func (self *Writer) Before(*org.Document) {
}
func (self *Writer) After(*org.Document) {
}
func (self *Writer) String() string {
	return ""
}
func (self *Writer) WriterWithExtensions() org.Writer {
	return self
}
func (self *Writer) WriteNodesAsString(...org.Node) string {
	return "asstring "
}
func (self *Writer) WriteKeyword(node org.Keyword) {
	self.fixLine()

	self.printf("keyword #+%s: %s\n", node.Key, node.Value)

	id := self.doc.addText(node.Value)
	self.doc.KeywordList = append(self.doc.KeywordList, &Keyword{node.Key, id})
	self.addItem(ItemKindKeyword, len(self.doc.KeywordList))
}
func (self *Writer) WriteInclude(node org.Include) {
	self.printf("inc %v\n", node)
}
func (self *Writer) WriteComment(node org.Comment) {
	self.fixLine()

	self.printf("comment %v\n", node.Content)
	id := self.doc.addText(node.Content)
	self.doc.CommentList = append(self.doc.CommentList, &Comment{id})
	self.addItem(ItemKindComment, len(self.doc.CommentList))

}
func (self *Writer) WriteNodeWithMeta(node org.NodeWithMeta) {
	self.printf("meta %v\n", node)
}
func (self *Writer) WriteNodeWithName(node org.NodeWithName) {
	self.fixLine()

	self.printf("name %v\n", node)

	id := self.doc.addText(node.String())
	self.doc.BlockList = append(self.doc.BlockList, &Block{id})
	self.addItem(ItemKindName, len(self.doc.BlockList))
}
func (self *Writer) WriteHeadline(node org.Headline) {

	self.fixLine()

	self.printf("headline %d\n", node.Lvl)
	var customID LnsAny = nil
	if node.Properties != nil {
		for _, profList := range node.Properties.Properties {
			if len(profList) == 2 {
				if profList[0] == "CUSTOM_ID" {
					customID = profList[1]
				}
			}
		}
	}

	self.headlineLv = node.Lvl
	for _, title := range node.Title {
		self.write(title)
	}
	self.fixLine()
	headline := self.doc.HeadlineList[len(self.doc.HeadlineList)-1].(*Headline)
	headline.CustomId = customID

	self.headlineLv = 0

	self.printf("headline child\n")

	for _, child := range node.Children {
		self.write(child)
	}
	// self.printf( "headline %v\n", node )
	self.fixLine()

}

func trimNodeString(node org.Node) string {
	txt := node.String()
	if strings.HasSuffix(txt, "\n") {
		return txt[:len(txt)-1]
	}
	return txt
}

func (self *Writer) WriteBlock(node org.Block) {
	self.fixLine()

	self.printf("block %v\n", node)

	id := self.doc.addText(node.String())
	self.doc.BlockList = append(self.doc.BlockList, &Block{id})
	self.addItem(ItemKindBlock, len(self.doc.BlockList))
}
func (self *Writer) WriteResult(node org.Result) {
	self.printf("result %v\n", node)
}
func (self *Writer) WriteInlineBlock(node org.InlineBlock) {
	self.printf("inline %v\n", node)
}
func (self *Writer) WriteExample(node org.Example) {
	self.fixLine()

	self.printf("example %v\n", node)

	id := self.doc.addText(node.String())
	self.doc.VerbList = append(self.doc.VerbList, &Verb{id})
	self.addItem(ItemKindVerb, len(self.doc.VerbList))
}
func (self *Writer) WriteDrawer(node org.Drawer) {
	self.printf("drawrer %v\n", node)
}
func (self *Writer) WritePropertyDrawer(node org.PropertyDrawer) {
	self.printf("property %v\n", node)
}
func (self *Writer) WriteList(node org.List) {
	self.fixLine()
	self.printf("list\n")
	self.listLv++
	for _, item := range node.Items {
		self.write(item)
		self.fixLine()
	}
	self.listLv--
}
func (self *Writer) WriteListItem(node org.ListItem) {
	self.printf("list item\n")
	for _, child := range node.Children {
		self.write(child)
	}
}
func (self *Writer) WriteDescriptiveListItem(node org.DescriptiveListItem) {
	self.fixLine()

	self.printf("descript %v\n", node)

	depth := DescriptDepth{self.listLv, true}
	self.descriptDepth = append(self.descriptDepth, &depth)
	for _, term := range node.Term {
		self.write(term)
	}
	self.fixLine()

	depth.term = false

	for _, detail := range node.Details {
		self.write(detail)
	}
	self.fixLine()

	self.descriptDepth = self.descriptDepth[:len(self.descriptDepth)-1]
}
func (self *Writer) WriteTable(node org.Table) {
	self.printf("table\n")

	table := &Table{make([]LnsAny, 0)}
	for index, row := range node.Rows {
		self.tableRow = &TableRow{make([]LnsAny, 0)}
		self.printf("IsSpecial %d %v\n", index, row.IsSpecial)
		for _, colum := range row.Columns {
			for _, child := range colum.Children {
				self.write(child)
				self.fixLine()
			}
		}
		table.RowList = append(table.RowList, self.tableRow)
	}
	self.doc.TableList = append(self.doc.TableList, table)
	self.addItem(ItemKindTable, len(self.doc.TableList))

	self.tableRow = nil
}
func (self *Writer) WriteHorizontalRule(node org.HorizontalRule) {
	self.fixLine()

	self.printf("rule %v\n", node)

	self.addItem(ItemKindRule, 0)
}
func (self *Writer) WriteParagraph(node org.Paragraph) {
	self.fixLine()
	self.printf("para\n")

	// if len(node.Children) > 0 {
	// 	if _, ok := node.Children[0].(org.LineBreak); ok {
	// 		self.addItem(ItemKindBlank, 0)
	// 	}
	// }

	for _, child := range node.Children {
		self.write(child)
	}
}
func (self *Writer) WriteText(node org.Text) {
	self.printf("text %v\n", node)
	txt := node.String()
	if self.emphasis != "" {
		if self.emphasis == "*" {
			if strings.Index(txt, "*") == 0 {
				txt = txt[1:]
			}
		}
		id := self.doc.addExp(self.emphasis, txt)
		if self.emphasis != "*" {
			self.addTxt(fmt.Sprintf("SyM_%d", id))
		} else {
			self.addTxt(txt)
		}
	} else {
		self.addTxt(txt)
	}
}
func (self *Writer) WriteEmphasis(node org.Emphasis) {
	self.printf("emphasis %s\n", node.Kind)
	self.emphasis = node.Kind
	self.lastEmphasis = node.Kind
	for _, child := range node.Content {
		self.write(child)
	}
	self.emphasis = ""
}
func (self *Writer) WriteLatexFragment(node org.LatexFragment) {
	self.printf("frag %v\n", node)
}
func (self *Writer) WriteStatisticToken(node org.StatisticToken) {
	self.printf("stastis %v\n", node)
}
func (self *Writer) WriteExplicitLineBreak(node org.ExplicitLineBreak) {
	self.printf("explicit %v\n", node)
}
func (self *Writer) WriteLineBreak(node org.LineBreak) {
	self.printf("break\n")
	self.addCr()
}
func (self *Writer) WriteRegularLink(node org.RegularLink) {
	self.printf("regular %v\n", node)

	self.addTxt(node.String())
}
func (self *Writer) WriteMacro(node org.Macro) {
	self.printf("macro %v\n", node)
}
func (self *Writer) WriteTimestamp(node org.Timestamp) {
	self.printf("time %v\n", node)
}
func (self *Writer) WriteFootnoteLink(node org.FootnoteLink) {
	self.printf("link %v\n", node)
}
func (self *Writer) WriteFootnoteDefinition(node org.FootnoteDefinition) {
	self.printf("footnote %v\n", node)
}

func dump(doc *Document) {
	txtList := doc.TextList

	for index, work := range doc.ItemList {
		item := work.(Item)
		fmt.Printf("item-%d: %d %d\n", index, item.Kind, item.Id)
	}
	for index, work := range doc.HeadlineList {
		headline := work.(Headline)
		fmt.Printf(
			"headline-%d: %d %s\n", index, headline.Level,
			doc.TextList[headline.TxtId])
	}
	for index, work := range doc.ExpList {
		exp := work.(Exp)
		fmt.Printf(
			"exp-%d: %s%s\n", index, exp.Delimit,
			doc.TextList[exp.TxtId])
	}
	for index, work := range doc.ListList {
		list := work.(List)
		fmt.Printf(
			"list-%d: %d %s\n", index, list.Level,
			doc.TextList[list.TxtId])
	}
	for index, work := range doc.TableList {
		table := work.(Table)
		for subIndex, workrow := range table.RowList {
			row := workrow.(TableRow)
			for rowIndex, id := range row.TxtIdList {
				fmt.Printf(
					"tbl-%d-%d-%d: %s\n",
					index, subIndex, rowIndex, txtList[id.(LnsInt)])
			}
		}
	}
}

func LoadOrg(_env *LnsEnv, path string, enableLog bool) *Document {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	d := org.New().Parse(bytes.NewReader(bs), path)

	write := func(w org.Writer) {
		_, err := d.Write(w)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.FprLnsInt(os.Stdout, out)
	}
	writer := NewWriter(enableLog)
	write(writer)

	return writer.GetDocument()
}
