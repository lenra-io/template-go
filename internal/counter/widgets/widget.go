package widgets

type Padding struct {
	Top    uint16 `json:"top"`
	Bottom uint16 `json:"bottom"`
	Left   uint16 `json:"left"`
	Right  uint16 `json:"right"`
}

type Decoration struct {
	Color     uint32    `json:"color"`
	BoxShadow BoxShadow `json:"boxShadow"`
}

type BoxShadow struct {
	BlurRadius uint16 `json:"blurRadius"`
	Color      uint32 `json:"color"`
	Offset     Offset `json:"offset"`
}

type Offset struct {
	Dx uint16 `json:"dx"`
	Dy uint16 `json:"dy"`
}

type CounterWidgetProps struct {
	Text string `json:"text"`
}
