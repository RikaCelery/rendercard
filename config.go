package rendercard

import (
	"hash/crc64"
	"unsafe"
)

const (
	// DefaultWidth 默认宽度
	DefaultWidth = 1272.0
)

// Title 标题配置
type Title struct {
	// Line 行数
	Line int
	// IsEnabled 状态
	IsEnabled bool
	// LeftTitle 左侧标题
	LeftTitle string
	// LeftSubtitle 左侧副标题
	LeftSubtitle string
	// RightTitle 右侧标题
	RightTitle string
	// RightSubtitle 右侧副标题
	RightSubtitle string
	// ImagePath 图片路径
	ImagePath string
	// TitleFont 标题字体数据
	TitleFontData []byte
	// TextFont 正文字体数据
	TextFontData []byte
	// OffsetX 文字X坐标偏移 向右为正方向
	OffsetX float64
	// OffsetX 文字Y坐标偏移 向下为正方向
	OffsetY float64
	// TitleFontOffsetPoint 标题字体大小偏移
	TitleFontOffsetPoint float64
	// TextFontOffsetPoint 正文字体大小偏移
	TextFontOffsetPoint float64
	Width               float64
}

// Sum64 struct 的摘要
func (t *Title) Sum64() uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	sz := unsafe.Sizeof(Title{})
	var raw []byte
	s := (*slice)(unsafe.Pointer(&raw))
	s.data = unsafe.Pointer(t)
	s.len = int(sz)
	s.cap = int(sz)
	_, err := h.Write(raw)
	if err != nil {
		return 0
	}
	return h.Sum64()
}

// Alignment 对齐规则
type Alignment uint8

const (
	NilAlign    Alignment = iota // NilAlign ..
	AlignLeft                    // AlignLeft 左对齐
	AlignCenter                  // AlignCenter 居中对齐
	AlignRight                   // AlignRight 右对齐
)

// Card 卡片配置
type Card struct {
	// Width 宽度,默认600
	Width int
	// Height 高度,默认由Title+Text内容决定
	Height int
	// BackgroundImage 背景图
	BackgroundImage string
	// TitleFontData 标题字体数据
	TitleFontData []byte
	// TextFontData 正文字体数据
	TextFontData []byte
	// Title 标题内容
	Title string
	// CanTitleShown 是否显示标题
	CanTitleShown bool
	// IsTextSplitPerElement true为每个元素按行显示,false按空格分割显示;
	IsTextSplitPerElement bool
	// TitleAlign 标题布局[Left|Center|Right],默认Left
	TitleAlign Alignment
	// Text 正文内容
	Text []string
}

// Sum64 struct 的摘要
func (c *Card) Sum64() uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	sz := unsafe.Sizeof(Card{})
	var raw []byte
	s := (*slice)(unsafe.Pointer(&raw))
	s.data = unsafe.Pointer(c)
	s.len = int(sz)
	s.cap = int(sz)
	_, err := h.Write(raw)
	if err != nil {
		return 0
	}
	return h.Sum64()
}
