package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/kozigh01/go_yt_DivRhino/fruitful-pdf/data"
)

func main() {
	fmt.Printf("Hello from fruitful-pdf\n")

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildFooter(m)
	buildFruitList(m)
	buildSignature(m)

	if err := m.OutputFileAndClose("pdfs/sample01.pdf"); err != nil {
		log.Fatalf("Could not save PDF: %v\n", err)
	}

	log.Println("PDF saved successfully.")
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/logo_div_rhino.jpg", props.Rect{
					Center:  true,
					Percent: 75,
				})
				if err != nil {
					log.Printf("Image file was not loaded: %v\n", err)
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Prepared for you by the Div Rhino Fruit Company", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func buildFooter(m pdf.Maroto) {
	begin := time.Now()
	m.SetAliasNbPages("{nb}")
	m.SetFirstPageNb(1)

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(6, func() {
				m.Text(begin.Format("01/02/2006"), props.Text{
					Top: 10,
					Size: 8,
					Color: getGreyColor(),
					Align: consts.Left,
				})
			})

			m.Col(6, func() {
				m.Text(fmt.Sprintf("Page %v of {nb}", strconv.Itoa(m.GetCurrentPage())), props.Text{
					Top: 10,
					Size: 8,
					Style: consts.Italic,
					Color: getGreyColor(),
					Align: consts.Right,
				})
			})
		})
	})
}

func buildFruitList(m pdf.Maroto) {
	tableHeadings := []string{"Fruit", "Description", "Price"}
	// contents := [][]string{
	// 	{"Apple", "Red and juicy", "2.00"},
	// 	{"Orange", "Orange and juicy", "3.00"},
	// }
	contents := data.FruitList(20)
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurpleColor,
	})
}

func buildSignature(m pdf.Maroto) {
	m.Row(15, func() {
		m.Col(5, func() {
			m.QrCode("https://divrhino.com", props.Rect{
				Left: 0,
				Top: 5,
				Center: false,
				Percent: 100,
			})
		})

		m.ColSpace(2)

		m.Col(5, func() {
			m.Signature("Signed by", props.Font{
				Size: 8,
				Style: consts.Italic,
				Family: consts.Courier,
			})
		})
	})
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getGreyColor() color.Color {
	return color.Color{
		Red:   206,
		Green: 206,
		Blue:  206,
	}
}
