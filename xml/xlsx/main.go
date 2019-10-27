package main

import (
	"archive/zip"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Worksheet struct {
	XMLName   xml.Name `xml:"worksheet"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	R         string   `xml:"r,attr"`
	Mc        string   `xml:"mc,attr"`
	Ignorable string   `xml:"Ignorable,attr"`
	X14ac     string   `xml:"x14ac,attr"`
	SheetPr   struct {
		Text        string `xml:",chardata"`
		PageSetUpPr struct {
			Text      string `xml:",chardata"`
			FitToPage string `xml:"fitToPage,attr"`
		} `xml:"pageSetUpPr"`
	} `xml:"sheetPr"`
	Dimension struct {
		Text string `xml:",chardata"`
		Ref  string `xml:"ref,attr"`
	} `xml:"dimension"`
	SheetViews struct {
		Text      string `xml:",chardata"`
		SheetView struct {
			Text                     string `xml:",chardata"`
			ShowGridLines            string `xml:"showGridLines,attr"`
			View                     string `xml:"view,attr"`
			ZoomScale                string `xml:"zoomScale,attr"`
			ZoomScaleNormal          string `xml:"zoomScaleNormal,attr"`
			ZoomScaleSheetLayoutView string `xml:"zoomScaleSheetLayoutView,attr"`
			WorkbookViewId           string `xml:"workbookViewId,attr"`
			Selection                struct {
				Text       string `xml:",chardata"`
				ActiveCell string `xml:"activeCell,attr"`
				Sqref      string `xml:"sqref,attr"`
			} `xml:"selection"`
		} `xml:"sheetView"`
	} `xml:"sheetViews"`
	SheetFormatPr struct {
		Text             string `xml:",chardata"`
		DefaultRowHeight string `xml:"defaultRowHeight,attr"`
		DyDescent        string `xml:"dyDescent,attr"`
	} `xml:"sheetFormatPr"`
	Cols struct {
		Text string `xml:",chardata"`
		Col  []struct {
			Text        string `xml:",chardata"`
			Min         string `xml:"min,attr"`
			Max         string `xml:"max,attr"`
			Width       string `xml:"width,attr"`
			Style       string `xml:"style,attr"`
			CustomWidth string `xml:"customWidth,attr"`
		} `xml:"col"`
	} `xml:"cols"`
	SheetData struct {
		Text string `xml:",chardata"`
		Row  []struct {
			Text         string `xml:",chardata"`
			R            string `xml:"r,attr"`
			Spans        string `xml:"spans,attr"`
			Ht           string `xml:"ht,attr"`
			CustomHeight string `xml:"customHeight,attr"`
			DyDescent    string `xml:"dyDescent,attr"`
			ThickBot     string `xml:"thickBot,attr"`
			C            []struct {
				Text string `xml:",chardata"`
				R    string `xml:"r,attr"`
				S    string `xml:"s,attr"`
				T    string `xml:"t,attr"`
				V    string `xml:"v"`
			} `xml:"c"`
		} `xml:"row"`
	} `xml:"sheetData"`
	MergeCells struct {
		Text      string `xml:",chardata"`
		Count     string `xml:"count,attr"`
		MergeCell []struct {
			Text string `xml:",chardata"`
			Ref  string `xml:"ref,attr"`
		} `xml:"mergeCell"`
	} `xml:"mergeCells"`
	PhoneticPr struct {
		Text   string `xml:",chardata"`
		FontId string `xml:"fontId,attr"`
	} `xml:"phoneticPr"`
	PageMargins struct {
		Text   string `xml:",chardata"`
		Left   string `xml:"left,attr"`
		Right  string `xml:"right,attr"`
		Top    string `xml:"top,attr"`
		Bottom string `xml:"bottom,attr"`
		Header string `xml:"header,attr"`
		Footer string `xml:"footer,attr"`
	} `xml:"pageMargins"`
	PageSetup struct {
		Text        string `xml:",chardata"`
		PaperSize   string `xml:"paperSize,attr"`
		Scale       string `xml:"scale,attr"`
		FitToHeight string `xml:"fitToHeight,attr"`
		Orientation string `xml:"orientation,attr"`
		ID          string `xml:"id,attr"`
	} `xml:"pageSetup"`
	HeaderFooter struct {
		Text             string `xml:",chardata"`
		AlignWithMargins string `xml:"alignWithMargins,attr"`
		OddHeader        string `xml:"oddHeader"`
		OddFooter        string `xml:"oddFooter"`
	} `xml:"headerFooter"`
	Drawing struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"drawing"`
}

type SharedStrings struct {
	XMLName     xml.Name `xml:"sst"`
	Text        string   `xml:",chardata"`
	Xmlns       string   `xml:"xmlns,attr"`
	Count       string   `xml:"count,attr"`
	UniqueCount string   `xml:"uniqueCount,attr"`
	Si          []struct {
		Text string `xml:",chardata"`
		T    string `xml:"t"`
		RPh  []struct {
			Text string `xml:",chardata"`
			Sb   string `xml:"sb,attr"`
			Eb   string `xml:"eb,attr"`
			T    string `xml:"t"`
		} `xml:"rPh"`
		PhoneticPr struct {
			Text   string `xml:",chardata"`
			FontId string `xml:"fontId,attr"`
		} `xml:"phoneticPr"`
	} `xml:"si"`
}

func xlsxparse(arg string) error {
	r, err := zip.OpenReader(arg)
	if err != nil {
		return err
	}
	defer r.Close()

	var sst SharedStrings

	for _, f := range r.File {
		if strings.HasPrefix(f.Name, "xl/sharedStrings.xml") {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			b, err := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			err = xml.Unmarshal(b, &sst)
			if err != nil {
				return err
			}
		}
	}

	for _, f := range r.File {
		var ws Worksheet

		if strings.HasPrefix(f.Name, "xl/worksheets/sheet") {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			b, err := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			err = xml.Unmarshal(b, &ws)
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", f.Name)

			for i, r := range ws.SheetData.Row {
				for j, c := range r.C {
					if c.V != "" {
						idx, err := strconv.Atoi(c.V)
						if err != nil {
							return err
						}
						fmt.Printf("(%d, %d) %s:%d:%s\n", i, j, c.R, idx, sst.Si[idx].T)
					}
				}
			}
			fmt.Println()
		}
	}
	return nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if err := xlsxparse(arg); err != nil {
			log.Fatal(err)
		}
	}
}
