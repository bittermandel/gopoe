package gopoe

// Color holds primary colors.
type Color struct {
	R int
	G int
	B int
}

// Tab describes a tabulation style (background color,
// custom name, style, and so on)
type Tab struct {
	Hidden          bool   `json:"hidden"`
	Selected        bool   `json:"selected"`
	Index           int    `json:"i"`
	BackgroundColor Color  `json:"colour"`
	Name            string `json:"n"`
	ID              string `json:"id"`
	Type            string `json:"type"`
	ImgL            string `json:"srcL"`
	ImgC            string `json:"srcC"`
	ImgR            string `json:"srcR"`
}

// Layout is used for custom layout like currency.
type Layout struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W int     `json:"w"`
	H int     `json:"h"`
}

// CardLayout is used for the divination card layout.
type CardLayout struct {
	Name string `json:"name"`
}

// DivineLayout hold the divination layout information.
type DivineLayout struct {
	Height int          `json:"h"`
	Width  int          `json:"w"`
	PadX   int          `json:"padx"`
	PadY   int          `json:"pady"`
	Scale  float64      `json:"scale"`
	Cards  []CardLayout `json:"cards"`
}

// StashTab holds all stash tabulations (thus all items).
type StashTab struct {
	NumTabs          int               `json:"numTabs"`
	QuadLayout       bool              `json:"quadLayout"`
	Items            []Item            `json:"items"`
	Tabs             []Tab             `json:"tabs"`
	CurrencyLayout   map[string]Layout `json:"currencyLayout"`
	FragmentLayout   map[string]Layout `json:"fragmentLayout"`
	EssenceLayout    map[string]Layout `json:"essenceLayout"`
	DivinationLayout DivineLayout      `json:"divinationLayout"`

	// Can be empty, hence the *, for allowing to be nullable.
	MapLayout *map[string]Layout `json:"mapLayout"`
}
