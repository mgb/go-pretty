package table

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/mgb/go-pretty/text"
	"github.com/stretchr/testify/assert"
)

func compareOutput(t *testing.T, out string, expectedOut string) {
	if strings.HasPrefix(expectedOut, "\n") {
		expectedOut = strings.Replace(expectedOut, "\n", "", 1)
	}
	assert.Equal(t, expectedOut, out)
	if out != expectedOut {
		fmt.Printf("Expected:\n%s\nActual:\n%s\n", expectedOut, out)
	} else {
		fmt.Println(out)
	}
}

func compareOutputColored(t *testing.T, out string, expectedOut string) {
	if strings.HasPrefix(expectedOut, "\n") {
		expectedOut = strings.Replace(expectedOut, "\n", "", 1)
	}
	assert.Equal(t, expectedOut, out)
	if out != expectedOut {
		fmt.Printf("Expected:\n%s\nActual:\n%s\n", expectedOut, out)

		// dump formatted output that can be "pasted" into the expectation in
		// the test in case of valid changed behavior
		outLines := strings.Split(out, "\n")
		fmt.Printf("\"\" +\n")
		for idx, line := range outLines {
			if idx < len(outLines)-1 {
				fmt.Printf("%#v +", line+"\n")
			} else {
				fmt.Printf("%#v,", line)
			}
			fmt.Printf("\n")
		}
	} else {
		fmt.Println(out)
	}
}

func generateColumnConfigsWithHiddenColumns(colsToHide []int) []ColumnConfig {
	cc := []ColumnConfig{
		{
			Name: "#",
			Transformer: func(val interface{}) string {
				return fmt.Sprint(val.(int) + 7)
			},
		}, {
			Name: "First Name",
			Transformer: func(val interface{}) string {
				return fmt.Sprintf(">>%s", val)
			},
		}, {
			Name: "Last Name",
			Transformer: func(val interface{}) string {
				return fmt.Sprintf("%s<<", val)
			},
		}, {
			Name: "Salary",
			Transformer: func(val interface{}) string {
				return fmt.Sprint(val.(int) + 13)
			},
		}, {
			Number: 5,
			Transformer: func(val interface{}) string {
				return fmt.Sprintf("~%s~", val)
			},
		},
	}
	for _, colToHide := range colsToHide {
		cc[colToHide].Hidden = true
	}
	return cc
}

func TestTable_Render(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetCaption(testCaption)
	tw.SetStyle(styleTest)
	tw.SetTitle(testTitle2)

	compareOutput(t, tw.Render(), `
(---------------------------------------------------------------------)
[<When you play the Game of Thrones, you win or you die. There is no >]
[<middle ground.                                                     >]
{-----^------------^-----------^--------^-----------------------------}
[<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]
{-----+------------+-----------+--------+-----------------------------}
[<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]
[< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]
[<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]
[<  0>|<Winter    >|<Is       >|<     0>|<Coming.                    >]
[<   >|<          >|<         >|<      >|<The North Remembers!       >]
[<   >|<          >|<         >|<      >|<This is known.             >]
{-----+------------+-----------+--------+-----------------------------}
[<   >|<          >|<TOTAL    >|< 10000>|<                           >]
\-----v------------v-----------v--------v-----------------------------/
A Song of Ice and Fire`)
}

func TestTable_Render_AutoIndex(t *testing.T) {
	tw := NewWriter()
	for rowIdx := 0; rowIdx < 10; rowIdx++ {
		row := make(Row, 10)
		for colIdx := 0; colIdx < 10; colIdx++ {
			row[colIdx] = fmt.Sprintf("%s%d", AutoIndexColumnID(colIdx), rowIdx+1)
		}
		tw.AppendRow(row)
	}
	tw.SetAutoIndex(true)
	tw.SetStyle(StyleLight)

	compareOutput(t, tw.Render(), `
┌────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐
│    │  A  │  B  │  C  │  D  │  E  │  F  │  G  │  H  │  I  │  J  │
├────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┼─────┤
│  1 │ A1  │ B1  │ C1  │ D1  │ E1  │ F1  │ G1  │ H1  │ I1  │ J1  │
│  2 │ A2  │ B2  │ C2  │ D2  │ E2  │ F2  │ G2  │ H2  │ I2  │ J2  │
│  3 │ A3  │ B3  │ C3  │ D3  │ E3  │ F3  │ G3  │ H3  │ I3  │ J3  │
│  4 │ A4  │ B4  │ C4  │ D4  │ E4  │ F4  │ G4  │ H4  │ I4  │ J4  │
│  5 │ A5  │ B5  │ C5  │ D5  │ E5  │ F5  │ G5  │ H5  │ I5  │ J5  │
│  6 │ A6  │ B6  │ C6  │ D6  │ E6  │ F6  │ G6  │ H6  │ I6  │ J6  │
│  7 │ A7  │ B7  │ C7  │ D7  │ E7  │ F7  │ G7  │ H7  │ I7  │ J7  │
│  8 │ A8  │ B8  │ C8  │ D8  │ E8  │ F8  │ G8  │ H8  │ I8  │ J8  │
│  9 │ A9  │ B9  │ C9  │ D9  │ E9  │ F9  │ G9  │ H9  │ I9  │ J9  │
│ 10 │ A10 │ B10 │ C10 │ D10 │ E10 │ F10 │ G10 │ H10 │ I10 │ J10 │
└────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘`)
}

func TestTable_Render_AutoMerge(t *testing.T) {
	rcAutoMerge := RowConfig{AutoMerge: true}

	t.Run("columns only", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE\nEXE", "RCE\nRUN"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"})
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬─────┬─────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE │ RCE │
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │  Y  │  Y  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │
├───┤         │        ├───────────┼───────────┼─────┼─────┤
│ 3 │         │        │ NS 1B     │ C 3       │  N  │  N  │
├───┤         ├────────┼───────────┼───────────┼─────┼─────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │  N  │  N  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │  Y  │  Y  │
├───┤         │        │           ├───────────┼─────┼─────┤
│ 7 │         │        │           │ C 7       │  Y  │  Y  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`)
	})

	t.Run("columns only with hidden columns", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE\nEXE", "RCE\nRUN"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "Y", "Y"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "N"})
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, Hidden: true},
			{Number: 5, Hidden: true, Align: text.AlignCenter},
			{Number: 6, Hidden: true, Align: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌─────────┬────────┬───────────┐
│ NODE IP │ PODS   │ NAMESPACE │
├─────────┼────────┼───────────┤
│ 1.1.1.1 │ Pod 1A │ NS 1A     │
│         │        │           │
│         │        │           │
│         │        ├───────────┤
│         │        │ NS 1B     │
│         ├────────┼───────────┤
│         │ Pod 1B │ NS 2      │
│         │        │           │
│         │        │           │
├─────────┼────────┼───────────┤
│ 2.2.2.2 │ Pod 2  │ NS 3      │
│         │        │           │
│         │        │           │
├─────────┼────────┼───────────┤
│         │        │           │
└─────────┴────────┴───────────┘`)
	})

	t.Run("rows only", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignRight})
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   ├─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 2 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 2       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 3 │ 1.1.1.1 │ Pod 1A │ NS 1B     │ C 3       │  N  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 4 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 5 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼───────────┤
│ 7 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 7       │         Y │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`)
	})

	t.Run("rows and columns", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   │         │        │           │           ├─────┬─────┤
│   │         │        │           │           │ EXE │ RUN │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┤         │        │           ├───────────┼─────┬─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │
├───┤         │        ├───────────┼───────────┼─────┴─────┤
│ 3 │         │        │ NS 1B     │ C 3       │     N     │
├───┤         ├────────┼───────────┼───────────┼───────────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┤         │        │           ├───────────┼─────┬─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┤         │        │           ├───────────┼───────────┤
│ 7 │         │        │           │ C 7       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│   │         │        │           │ 7         │  5  │  3  │
└───┴─────────┴────────┴───────────┴───────────┴─────┴─────┘`)
	})

	t.Run("rows and columns no headers or footers", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignRight})
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌─────────┬────────┬───────┬─────┬───────┐
│ 1.1.1.1 │ Pod 1A │ NS 1A │ C 1 │   Y   │
├─────────┼────────┼───────┼─────┼───┬───┤
│ 1.1.1.1 │ Pod 1A │ NS 1A │ C 2 │ Y │ N │
├─────────┼────────┼───────┼─────┼───┼───┤
│ 1.1.1.1 │ Pod 1A │ NS 1B │ C 3 │ N │ N │
├─────────┼────────┼───────┼─────┼───┴───┤
│ 1.1.1.1 │ Pod 1B │ NS 2  │ C 4 │   N   │
├─────────┼────────┼───────┼─────┼───┬───┤
│ 1.1.1.1 │ Pod 1B │ NS 2  │ C 5 │ Y │ N │
├─────────┼────────┼───────┼─────┼───┴───┤
│ 2.2.2.2 │ Pod 2  │ NS 3  │ C 6 │   Y   │
├─────────┼────────┼───────┼─────┼───────┤
│ 2.2.2.2 │ Pod 2  │ NS 3  │ C 7 │     Y │
└─────────┴────────┴───────┴─────┴───────┘`)
	})

	t.Run("rows and columns no headers or footers with auto-index", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, RowConfig{AutoMerge: true, AutoMergeAlign: text.AlignRight})
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────┬─────┬───┬───┐
│   │    A    │    B   │   C   │  D  │ E │ F │
├───┼─────────┼────────┼───────┼─────┼───┴───┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A │ C 1 │   Y   │
├───┼─────────┼────────┼───────┼─────┼───┬───┤
│ 2 │ 1.1.1.1 │ Pod 1A │ NS 1A │ C 2 │ Y │ N │
├───┼─────────┼────────┼───────┼─────┼───┼───┤
│ 3 │ 1.1.1.1 │ Pod 1A │ NS 1B │ C 3 │ N │ N │
├───┼─────────┼────────┼───────┼─────┼───┴───┤
│ 4 │ 1.1.1.1 │ Pod 1B │ NS 2  │ C 4 │   N   │
├───┼─────────┼────────┼───────┼─────┼───┬───┤
│ 5 │ 1.1.1.1 │ Pod 1B │ NS 2  │ C 5 │ Y │ N │
├───┼─────────┼────────┼───────┼─────┼───┴───┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3  │ C 6 │   Y   │
├───┼─────────┼────────┼───────┼─────┼───────┤
│ 7 │ 2.2.2.2 │ Pod 2  │ NS 3  │ C 7 │     Y │
└───┴─────────┴────────┴───────┴─────┴───────┘`)
	})

	t.Run("rows and columns and footers", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE", "ID"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE", "RUN", ""})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y", 123}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N", 234})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N", 345})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N", 456}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N", 567})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y", 678}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y", 789}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 4, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┬─────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │  ID │
│   │         │        │           │           ├─────┬─────┼─────┤
│   │         │        │           │           │ EXE │ RUN │     │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │ 123 │
├───┤         │        │           ├───────────┼─────┬─────┼─────┤
│ 2 │         │        │           │ C 2       │  Y  │  N  │ 234 │
├───┤         │        ├───────────┼───────────┼─────┼─────┼─────┤
│ 3 │         │        │ NS 1B     │ C 3       │  N  │  N  │ 345 │
├───┤         ├────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 4 │         │ Pod 1B │ NS 2      │ C 4       │     N     │ 456 │
├───┤         │        │           ├───────────┼─────┬─────┼─────┤
│ 5 │         │        │           │ C 5       │  Y  │  N  │ 567 │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┼─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │ 678 │
├───┤         │        │           ├───────────┼───────────┼─────┤
│ 7 │         │        │           │ C 7       │     Y     │ 789 │
├───┼─────────┴────────┴───────────┼───────────┼───────────┼─────┤
│   │                              │ 7         │     5     │     │
│   │                              │           ├─────┬─────┼─────┤
│   │                              │           │  5  │  3  │     │
│   │                              │           ├─────┴─────┼─────┤
│   │                              │           │     5     │     │
│   │                              │           ├─────┬─────┼─────┤
│   │                              │           │  5  │  3  │     │
│   │                              │           ├─────┴─────┼─────┤
│   │                              │           │     5     │     │
└───┴──────────────────────────────┴───────────┴───────────┴─────┘`)
	})

	t.Run("samurai sudoku", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendRow(Row{"1.1\n1.1", "1.2\n1.2", "1.3\n1.3", " ", "2.1\n2.1", "2.2\n2.2", "2.3\n2.3"})
		tw.AppendRow(Row{"1.4\n1.4", "1.5\n1.5", "1.6\n1.6", " ", "2.4\n2.4", "2.5\n2.5", "2.6\n2.6"})
		tw.AppendRow(Row{"1.7\n1.7", "1.8\n1.8", "1.9\n0.1", "0.2\n0.2", "2.7\n0.3", "2.8\n2.8", "2.9\n2.9"})
		tw.AppendRow(Row{" ", " ", "0.4\n0.4", "0.5\n0.5", "0.6\n0.6", " ", " "}, rcAutoMerge)
		tw.AppendRow(Row{"3.1\n3.1", "3.2\n3.2", "3.3\n0.7", "0.8\n0.8", "4.1\n0.9", "4.2\n4.2", "4.3\n4.3"})
		tw.AppendRow(Row{"3.4\n3.4", "3.5\n3.5", "3.6\n3.6", " ", "4.4\n4.4", "4.5\n4.5", "4.6\n4.6"})
		tw.AppendRow(Row{"3.7\n3.7", "3.8\n3.8", "3.9\n3.9", " ", "4.7\n4.7", "4.8\n4.8", "4.9\n4.9"})
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 4, AutoMerge: true},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Box.PaddingLeft = ""
		tw.Style().Box.PaddingRight = ""
		tw.Style().Options.DrawBorder = true
		tw.Style().Options.SeparateRows = true
		tw.Style().Options.SeparateColumns = true

		compareOutput(t, tw.Render(), `
┌───┬───┬───┬───┬───┬───┬───┐
│1.1│1.2│1.3│   │2.1│2.2│2.3│
│1.1│1.2│1.3│   │2.1│2.2│2.3│
├───┼───┼───┤   ├───┼───┼───┤
│1.4│1.5│1.6│   │2.4│2.5│2.6│
│1.4│1.5│1.6│   │2.4│2.5│2.6│
├───┼───┼───┼───┼───┼───┼───┤
│1.7│1.8│1.9│0.2│2.7│2.8│2.9│
│1.7│1.8│0.1│0.2│0.3│2.8│2.9│
├───┴───┼───┼───┼───┼───┴───┤
│       │0.4│0.5│0.6│       │
│       │0.4│0.5│0.6│       │
├───┬───┼───┼───┼───┼───┬───┤
│3.1│3.2│3.3│0.8│4.1│4.2│4.3│
│3.1│3.2│0.7│0.8│0.9│4.2│4.3│
├───┼───┼───┼───┼───┼───┼───┤
│3.4│3.5│3.6│   │4.4│4.5│4.6│
│3.4│3.5│3.6│   │4.4│4.5│4.6│
├───┼───┼───┤   ├───┼───┼───┤
│3.7│3.8│3.9│   │4.7│4.8│4.9│
│3.7│3.8│3.9│   │4.7│4.8│4.9│
└───┴───┴───┴───┴───┴───┴───┘`)
	})

	t.Run("long column no merge", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRW", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRH", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRY"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬──────────────────────────┬──────────────────────────┬──────────────────────────┬──────────────────────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │         COLUMN 5         │         COLUMN 6         │         COLUMN 7         │         COLUMN 8         │
├───┼──────────┼──────────┼──────────┼──────────┼──────────────────────────┼──────────────────────────┼──────────────────────────┼──────────────────────────┤
│ 1 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 1      │ 4F8F5CB531E3D49A61CF417C │ 4F8F5CB531E3D49A61CF417C │ 4F8F5CB531E3D49A61CF417C │ 4F8F5CB531E3D49A61CF417C │
│   │          │          │          │          │ D133792CCFA501FD8DA53EE3 │ D133792CCFA501FD8DA53EE3 │ D133792CCFA501FD8DA53EE3 │ D133792CCFA501FD8DA53EE3 │
│   │          │          │          │          │ 68FED20E5FE0248C3A0B64F9 │ 68FED20E5FE0248C3A0B64F9 │ 68FED20E5FE0248C3A0B64F9 │ 68FED20E5FE0248C3A0B64F9 │
│   │          │          │          │          │ 8A6533CEE1DA614C3A8DDEC7 │ 8A6533CEE1DA614C3A8DDEC7 │ 8A6533CEE1DA614C3A8DDEC7 │ 8A6533CEE1DA614C3A8DDEC7 │
│   │          │          │          │          │ 91FF05FEE6D971D57C134832 │ 91FF05FEE6D971D57C134832 │ 91FF05FEE6D971D57C134832 │ 91FF05FEE6D971D57C134832 │
│   │          │          │          │          │         0F4EB42DR        │        0F4EB42DRW        │        0F4EB42DRH        │        0F4EB42DRY        │
├───┼──────────┼──────────┼──────────┼──────────┼──────────────────────────┴──────────────────────────┴──────────────────────────┴──────────────────────────┤
│ 2 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                                                     Y                                                     │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ 3 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                                                     Y                                                     │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────────────────────────────────────────────────────────────────────┘`)
	})

	t.Run("long column partially merged #1", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRR"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬─────────────┬─────────────┬─────────────┬─────────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │   COLUMN 5  │   COLUMN 6  │   COLUMN 7  │   COLUMN 8  │
├───┼──────────┼──────────┼──────────┼──────────┼─────────────┴─────────────┼─────────────┴─────────────┤
│ 1 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 1      │  4F8F5CB531E3D49A61CF417C │  4F8F5CB531E3D49A61CF417C │
│   │          │          │          │          │  D133792CCFA501FD8DA53EE3 │  D133792CCFA501FD8DA53EE3 │
│   │          │          │          │          │  68FED20E5FE0248C3A0B64F9 │  68FED20E5FE0248C3A0B64F9 │
│   │          │          │          │          │  8A6533CEE1DA614C3A8DDEC7 │  8A6533CEE1DA614C3A8DDEC7 │
│   │          │          │          │          │  91FF05FEE6D971D57C134832 │  91FF05FEE6D971D57C134832 │
│   │          │          │          │          │         0F4EB42DR         │         0F4EB42DRR        │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────┴───────────────────────────┤
│ 2 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                           Y                           │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────────────────┤
│ 3 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                           Y                           │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────────────────┘`)
	})

	t.Run("long column partially merged #2", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DRE"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────────────────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │ COLUMN 5 │ COLUMN 6 │ COLUMN 7 │         COLUMN 8         │
├───┼──────────┼──────────┼──────────┼──────────┼──────────┴──────────┴──────────┼──────────────────────────┤
│ 1 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 1      │    4F8F5CB531E3D49A61CF417C    │ 4F8F5CB531E3D49A61CF417C │
│   │          │          │          │          │    D133792CCFA501FD8DA53EE3    │ D133792CCFA501FD8DA53EE3 │
│   │          │          │          │          │    68FED20E5FE0248C3A0B64F9    │ 68FED20E5FE0248C3A0B64F9 │
│   │          │          │          │          │    8A6533CEE1DA614C3A8DDEC7    │ 8A6533CEE1DA614C3A8DDEC7 │
│   │          │          │          │          │    91FF05FEE6D971D57C134832    │ 91FF05FEE6D971D57C134832 │
│   │          │          │          │          │            0F4EB42DR           │        0F4EB42DRE        │
├───┼──────────┼──────────┼──────────┼──────────┼────────────────────────────────┴──────────────────────────┤
│ 2 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                             Y                             │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────────────────────┤
│ 3 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                             Y                             │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────────────────────┘`)
	})

	t.Run("long column fully merged", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Column 1", "Column 2", "Column 3", "Column 4", "Column 5", "Column 6", "Column 7", "Column 8"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR", "4F8F5CB531E3D49A61CF417CD133792CCFA501FD8DA53EE368FED20E5FE0248C3A0B64F98A6533CEE1DA614C3A8DDEC791FF05FEE6D971D57C1348320F4EB42DR"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y", "Y"}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
			{Number: 8, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 24, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┬──────────┐
│   │ COLUMN 1 │ COLUMN 2 │ COLUMN 3 │ COLUMN 4 │ COLUMN 5 │ COLUMN 6 │ COLUMN 7 │ COLUMN 8 │
├───┼──────────┼──────────┼──────────┼──────────┼──────────┴──────────┴──────────┴──────────┤
│ 1 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 1      │          4F8F5CB531E3D49A61CF417C         │
│   │          │          │          │          │          D133792CCFA501FD8DA53EE3         │
│   │          │          │          │          │          68FED20E5FE0248C3A0B64F9         │
│   │          │          │          │          │          8A6533CEE1DA614C3A8DDEC7         │
│   │          │          │          │          │          91FF05FEE6D971D57C134832         │
│   │          │          │          │          │                 0F4EB42DR                 │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────┤
│ 2 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                     Y                     │
├───┼──────────┼──────────┼──────────┼──────────┼───────────────────────────────────────────┤
│ 3 │ 1.1.1.1  │ Pod 1A   │ NS 1A    │ C 2      │                     Y                     │
└───┴──────────┴──────────┴──────────┴──────────┴───────────────────────────────────────────┘`)
	})

	t.Run("headers too", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE", "RCE"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE EXE EXE", "EXE EXE EXE"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"})
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "N"})
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 3}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 6, 4, 4}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬───────────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │    RCE    │
│   ├─────────┴────────┴───────────┴───────────┼───────────┤
│   │                                          │  EXE EXE  │
│   │                                          │    EXE    │
├───┼─────────┬────────┬───────────┬───────────┼───────────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 2 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 2       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┼─────┤
│ 3 │ 1.1.1.1 │ Pod 1A │ NS 1B     │ C 3       │  N  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 4 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 4       │     N     │
├───┼─────────┼────────┼───────────┼───────────┼─────┬─────┤
│ 5 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 5       │  Y  │  N  │
├───┼─────────┼────────┼───────────┼───────────┼─────┴─────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │     Y     │
├───┼─────────┼────────┼───────────┼───────────┼───────────┤
│ 7 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 7       │     Y     │
├───┼─────────┴────────┴───────────┼───────────┼─────┬─────┤
│   │                              │ 7         │  5  │  3  │
│   ├──────────────────────────────┼───────────┼─────┴─────┤
│   │                              │ 6         │     4     │
└───┴──────────────────────────────┴───────────┴───────────┘`)
	})

	t.Run("headers and footers too", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE1", "RCE2"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE EXE EXE", "EXE EXE EXE"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 6, 4, 4}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬──────┬──────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE1 │ RCE2 │
│   ├─────────┴────────┴───────────┴───────────┼──────┴──────┤
│   │                                          │   EXE EXE   │
│   │                                          │     EXE     │
├───┼─────────┬────────┬───────────┬───────────┼─────────────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 2 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 2       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 3 │ 1.1.1.1 │ Pod 1A │ NS 1B     │ C 3       │      N      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 4 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 4       │      N      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 5 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 5       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │      Y      │
├───┼─────────┼────────┼───────────┼───────────┼─────────────┤
│ 7 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 7       │      Y      │
├───┼─────────┴────────┴───────────┼───────────┼─────────────┤
│   │                              │ 7         │      5      │
│   ├──────────────────────────────┼───────────┼─────────────┤
│   │                              │ 6         │      4      │
└───┴──────────────────────────────┴───────────┴─────────────┘`)
	})

	t.Run("long header column", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"Node IP", "Pods", "Namespace", "Container", "RCE1", "RCE2", "RCE3"}, rcAutoMerge)
		tw.AppendHeader(Row{"", "", "", "", "EXE EXE EXE", "EXE EXE EXE", "EXE EXE EXE"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 7, 5, 5, 5}, rcAutoMerge)
		tw.AppendFooter(Row{"", "", "", 6, 4, 4, 3}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬─────────┬────────┬───────────┬───────────┬──────┬──────┬──────┐
│   │ NODE IP │ PODS   │ NAMESPACE │ CONTAINER │ RCE1 │ RCE2 │ RCE3 │
│   ├─────────┴────────┴───────────┴───────────┼──────┴──────┴──────┤
│   │                                          │       EXE EXE      │
│   │                                          │         EXE        │
├───┼─────────┬────────┬───────────┬───────────┼────────────────────┤
│ 1 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 1       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 2 │ 1.1.1.1 │ Pod 1A │ NS 1A     │ C 2       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 3 │ 1.1.1.1 │ Pod 1A │ NS 1B     │ C 3       │          N         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 4 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 4       │          N         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 5 │ 1.1.1.1 │ Pod 1B │ NS 2      │ C 5       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 6 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 6       │          Y         │
├───┼─────────┼────────┼───────────┼───────────┼────────────────────┤
│ 7 │ 2.2.2.2 │ Pod 2  │ NS 3      │ C 7       │          Y         │
├───┼─────────┴────────┴───────────┼───────────┼────────────────────┤
│   │                              │ 7         │          5         │
│   ├──────────────────────────────┼───────────┼─────────────┬──────┤
│   │                              │ 6         │      4      │   3  │
└───┴──────────────────────────────┴───────────┴─────────────┴──────┘`)
	})

	t.Run("everything", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(Row{"COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS", "COLUMNS"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 1", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1A", "C 2", "Y", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1A", "NS 1B", "C 3", "N", "N", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 4", "N", "Y", "N"}, rcAutoMerge)
		tw.AppendRow(Row{"1.1.1.1", "Pod 1B", "NS 2", "C 5", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 6", "N", "Y", "Y"}, rcAutoMerge)
		tw.AppendRow(Row{"2.2.2.2", "Pod 2", "NS 3", "C 7", "Y", "Y", "Y"}, rcAutoMerge)
		tw.AppendFooter(Row{"foo", "foo", "foo", "foo", "bar", "bar", "bar"}, rcAutoMerge)
		tw.AppendFooter(Row{7, 7, 7, 7, 7, 7, 7}, rcAutoMerge)
		tw.SetAutoIndex(true)
		tw.SetColumnConfigs([]ColumnConfig{
			{Number: 1, AutoMerge: true},
			{Number: 2, AutoMerge: true},
			{Number: 3, AutoMerge: true},
			{Number: 5, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 6, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
			{Number: 7, Align: text.AlignCenter, AlignFooter: text.AlignCenter, AlignHeader: text.AlignCenter, WidthMax: 7, WidthMaxEnforcer: text.WrapHard},
		})
		tw.SetStyle(StyleLight)
		tw.Style().Options.SeparateRows = true

		compareOutput(t, tw.Render(), `
┌───┬───────────────────────────────────────────────────┐
│   │                      COLUMNS                      │
├───┼─────────┬─────────┬─────────┬─────────┬───────────┤
│ 1 │ 1.1.1.1 │ Pod 1A  │ NS 1A   │ C 1     │     Y     │
├───┤         │         │         ├─────────┼───────┬───┤
│ 2 │         │         │         │ C 2     │   Y   │ N │
├───┤         │         ├─────────┼─────────┼───────┴───┤
│ 3 │         │         │ NS 1B   │ C 3     │     N     │
├───┤         ├─────────┼─────────┼─────────┼───┬───┬───┤
│ 4 │         │ Pod 1B  │ NS 2    │ C 4     │ N │ Y │ N │
├───┤         │         │         ├─────────┼───┴───┴───┤
│ 5 │         │         │         │ C 5     │     Y     │
├───┼─────────┼─────────┼─────────┼─────────┼───┬───────┤
│ 6 │ 2.2.2.2 │ Pod 2   │ NS 3    │ C 6     │ N │   Y   │
├───┤         │         │         ├─────────┼───┴───────┤
│ 7 │         │         │         │ C 7     │     Y     │
├───┼─────────┴─────────┴─────────┴─────────┼───────────┤
│   │                  FOO                  │    BAR    │
│   ├───────────────────────────────────────┴───────────┤
│   │                         7                         │
└───┴───────────────────────────────────────────────────┘`)
	})
}

func TestTable_Render_BorderAndSeparators(t *testing.T) {
	table := Table{}
	table.AppendHeader(testHeader)
	table.AppendRows(testRows)
	table.AppendFooter(testFooter)
	compareOutput(t, table.Render(), `
+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`)

	table.Style().Options = OptionsNoBorders
	compareOutput(t, table.Render(), `
   # | FIRST NAME | LAST NAME | SALARY |                             
-----+------------+-----------+--------+-----------------------------
   1 | Arya       | Stark     |   3000 |                             
  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! 
 300 | Tyrion     | Lannister |   5000 |                             
-----+------------+-----------+--------+-----------------------------
     |            | TOTAL     |  10000 |                             `)

	table.Style().Options.SeparateColumns = false
	compareOutput(t, table.Render(), `
   #  FIRST NAME  LAST NAME  SALARY                              
-----------------------------------------------------------------
   1  Arya        Stark        3000                              
  20  Jon         Snow         2000  You know nothing, Jon Snow! 
 300  Tyrion      Lannister    5000                              
-----------------------------------------------------------------
                  TOTAL       10000                              `)

	table.Style().Options.SeparateFooter = false
	compareOutput(t, table.Render(), `
   #  FIRST NAME  LAST NAME  SALARY                              
-----------------------------------------------------------------
   1  Arya        Stark        3000                              
  20  Jon         Snow         2000  You know nothing, Jon Snow! 
 300  Tyrion      Lannister    5000                              
                  TOTAL       10000                              `)

	table.Style().Options = OptionsNoBordersAndSeparators
	compareOutput(t, table.Render(), `
   #  FIRST NAME  LAST NAME  SALARY                              
   1  Arya        Stark        3000                              
  20  Jon         Snow         2000  You know nothing, Jon Snow! 
 300  Tyrion      Lannister    5000                              
                  TOTAL       10000                              `)

	table.Style().Options.DrawBorder = true
	compareOutput(t, table.Render(), `
+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
|   1  Arya        Stark        3000                              |
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
| 300  Tyrion      Lannister    5000                              |
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`)

	table.Style().Options.SeparateFooter = true
	compareOutput(t, table.Render(), `
+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
|   1  Arya        Stark        3000                              |
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
| 300  Tyrion      Lannister    5000                              |
+-----------------------------------------------------------------+
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`)

	table.Style().Options.SeparateHeader = true
	compareOutput(t, table.Render(), `
+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
+-----------------------------------------------------------------+
|   1  Arya        Stark        3000                              |
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
| 300  Tyrion      Lannister    5000                              |
+-----------------------------------------------------------------+
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`)

	table.Style().Options.SeparateRows = true
	compareOutput(t, table.Render(), `
+-----------------------------------------------------------------+
|   #  FIRST NAME  LAST NAME  SALARY                              |
+-----------------------------------------------------------------+
|   1  Arya        Stark        3000                              |
+-----------------------------------------------------------------+
|  20  Jon         Snow         2000  You know nothing, Jon Snow! |
+-----------------------------------------------------------------+
| 300  Tyrion      Lannister    5000                              |
+-----------------------------------------------------------------+
|                  TOTAL       10000                              |
+-----------------------------------------------------------------+`)

	table.Style().Options.SeparateColumns = true
	compareOutput(t, table.Render(), `
+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
+-----+------------+-----------+--------+-----------------------------+
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`)
}

func TestTable_Render_Colored(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendRow(testRowMultiLine)
		tw.AppendFooter(testFooter)
		tw.SetAutoIndex(true)
		tw.SetStyle(StyleColoredBright)
		tw.Style().Options.DrawBorder = true
		tw.Style().Options.SeparateColumns = true
		tw.Style().Options.SeparateFooter = true
		tw.Style().Options.SeparateHeader = true
		tw.Style().Options.SeparateRows = true

		compareOutputColored(t, tw.Render(), ""+
			"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m------------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m--------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m   # \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m|\x1b[0m\x1b[106;30m                             \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m------------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m--------\x1b[0m\x1b[106;30m+\x1b[0m\x1b[106;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m 1 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[107;30m   1 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m                             \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[107;30m-----\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m------------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m--------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m 2 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m  20 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[47;30m-----\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m------------\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m-----------\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m--------\x1b[0m\x1b[47;30m+\x1b[0m\x1b[47;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m 3 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[107;30m 300 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m|\x1b[0m\x1b[107;30m                             \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[106;30m+\x1b[0m\x1b[106;30m---\x1b[0m\x1b[106;30m+\x1b[0m\x1b[107;30m-----\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m------------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m--------\x1b[0m\x1b[107;30m+\x1b[0m\x1b[107;30m-----------------------------\x1b[0m\x1b[106;30m+\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m 4 \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m   0 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Winter     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Is        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m      0 \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m Coming.                     \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m            \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m           \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m The North Remembers!        \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[106;30m|\x1b[0m\x1b[106;30m   \x1b[0m\x1b[106;30m|\x1b[0m\x1b[47;30m     \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m            \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m           \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m        \x1b[0m\x1b[47;30m|\x1b[0m\x1b[47;30m This is known.              \x1b[0m\x1b[106;30m|\x1b[0m\n"+
			"\x1b[46;30m+\x1b[0m\x1b[46;30m---\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m------------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m--------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------------------------\x1b[0m\x1b[46;30m+\x1b[0m\n"+
			"\x1b[46;30m|\x1b[0m\x1b[46;30m   \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m     \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m|\x1b[0m\x1b[46;30m                             \x1b[0m\x1b[46;30m|\x1b[0m\n"+
			"\x1b[46;30m+\x1b[0m\x1b[46;30m---\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m------------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m--------\x1b[0m\x1b[46;30m+\x1b[0m\x1b[46;30m-----------------------------\x1b[0m\x1b[46;30m+\x1b[0m",
		)
	})

	t.Run("with borders", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendRow(testRowMultiLine)
		tw.AppendFooter(testFooter)
		tw.SetTitle(testTitle1)
		tw.Style().Title.Colors = text.Colors{text.FgYellow}
		tw.Style().Color = ColorOptions{
			Header:       text.Colors{text.FgRed},
			Row:          text.Colors{text.FgGreen},
			RowAlternate: text.Colors{text.FgHiGreen},
			Footer:       text.Colors{text.FgBlue},
		}

		compareOutputColored(t, tw.Render(), ""+
			"\x1b[33m+\x1b[0m\x1b[33m---------------------------------------------------------------------\x1b[0m\x1b[33m+\x1b[0m\n"+
			"\x1b[33m|\x1b[0m\x1b[33m Game of Thrones                                                     \x1b[0m\x1b[33m|\x1b[0m\n"+
			"\x1b[31m+\x1b[0m\x1b[31m-----\x1b[0m\x1b[31m+\x1b[0m\x1b[31m------------\x1b[0m\x1b[31m+\x1b[0m\x1b[31m-----------\x1b[0m\x1b[31m+\x1b[0m\x1b[31m--------\x1b[0m\x1b[31m+\x1b[0m\x1b[31m-----------------------------\x1b[0m\x1b[31m+\x1b[0m\n"+
			"\x1b[31m|\x1b[0m\x1b[31m   # \x1b[0m\x1b[31m|\x1b[0m\x1b[31m FIRST NAME \x1b[0m\x1b[31m|\x1b[0m\x1b[31m LAST NAME \x1b[0m\x1b[31m|\x1b[0m\x1b[31m SALARY \x1b[0m\x1b[31m|\x1b[0m\x1b[31m                             \x1b[0m\x1b[31m|\x1b[0m\n"+
			"\x1b[31m+\x1b[0m\x1b[31m-----\x1b[0m\x1b[31m+\x1b[0m\x1b[31m------------\x1b[0m\x1b[31m+\x1b[0m\x1b[31m-----------\x1b[0m\x1b[31m+\x1b[0m\x1b[31m--------\x1b[0m\x1b[31m+\x1b[0m\x1b[31m-----------------------------\x1b[0m\x1b[31m+\x1b[0m\n"+
			"\x1b[32m|\x1b[0m\x1b[32m   1 \x1b[0m\x1b[32m|\x1b[0m\x1b[32m Arya       \x1b[0m\x1b[32m|\x1b[0m\x1b[32m Stark     \x1b[0m\x1b[32m|\x1b[0m\x1b[32m   3000 \x1b[0m\x1b[32m|\x1b[0m\x1b[32m                             \x1b[0m\x1b[32m|\x1b[0m\n"+
			"\x1b[92m|\x1b[0m\x1b[92m  20 \x1b[0m\x1b[92m|\x1b[0m\x1b[92m Jon        \x1b[0m\x1b[92m|\x1b[0m\x1b[92m Snow      \x1b[0m\x1b[92m|\x1b[0m\x1b[92m   2000 \x1b[0m\x1b[92m|\x1b[0m\x1b[92m You know nothing, Jon Snow! \x1b[0m\x1b[92m|\x1b[0m\n"+
			"\x1b[32m|\x1b[0m\x1b[32m 300 \x1b[0m\x1b[32m|\x1b[0m\x1b[32m Tyrion     \x1b[0m\x1b[32m|\x1b[0m\x1b[32m Lannister \x1b[0m\x1b[32m|\x1b[0m\x1b[32m   5000 \x1b[0m\x1b[32m|\x1b[0m\x1b[32m                             \x1b[0m\x1b[32m|\x1b[0m\n"+
			"\x1b[92m|\x1b[0m\x1b[92m   0 \x1b[0m\x1b[92m|\x1b[0m\x1b[92m Winter     \x1b[0m\x1b[92m|\x1b[0m\x1b[92m Is        \x1b[0m\x1b[92m|\x1b[0m\x1b[92m      0 \x1b[0m\x1b[92m|\x1b[0m\x1b[92m Coming.                     \x1b[0m\x1b[92m|\x1b[0m\n"+
			"\x1b[92m|\x1b[0m\x1b[92m     \x1b[0m\x1b[92m|\x1b[0m\x1b[92m            \x1b[0m\x1b[92m|\x1b[0m\x1b[92m           \x1b[0m\x1b[92m|\x1b[0m\x1b[92m        \x1b[0m\x1b[92m|\x1b[0m\x1b[92m The North Remembers!        \x1b[0m\x1b[92m|\x1b[0m\n"+
			"\x1b[92m|\x1b[0m\x1b[92m     \x1b[0m\x1b[92m|\x1b[0m\x1b[92m            \x1b[0m\x1b[92m|\x1b[0m\x1b[92m           \x1b[0m\x1b[92m|\x1b[0m\x1b[92m        \x1b[0m\x1b[92m|\x1b[0m\x1b[92m This is known.              \x1b[0m\x1b[92m|\x1b[0m\n"+
			"\x1b[34m+\x1b[0m\x1b[34m-----\x1b[0m\x1b[34m+\x1b[0m\x1b[34m------------\x1b[0m\x1b[34m+\x1b[0m\x1b[34m-----------\x1b[0m\x1b[34m+\x1b[0m\x1b[34m--------\x1b[0m\x1b[34m+\x1b[0m\x1b[34m-----------------------------\x1b[0m\x1b[34m+\x1b[0m\n"+
			"\x1b[34m|\x1b[0m\x1b[34m     \x1b[0m\x1b[34m|\x1b[0m\x1b[34m            \x1b[0m\x1b[34m|\x1b[0m\x1b[34m TOTAL     \x1b[0m\x1b[34m|\x1b[0m\x1b[34m  10000 \x1b[0m\x1b[34m|\x1b[0m\x1b[34m                             \x1b[0m\x1b[34m|\x1b[0m\n"+
			"\x1b[34m+\x1b[0m\x1b[34m-----\x1b[0m\x1b[34m+\x1b[0m\x1b[34m------------\x1b[0m\x1b[34m+\x1b[0m\x1b[34m-----------\x1b[0m\x1b[34m+\x1b[0m\x1b[34m--------\x1b[0m\x1b[34m+\x1b[0m\x1b[34m-----------------------------\x1b[0m\x1b[34m+\x1b[0m",
		)
	})

	t.Run("with borders and separators not colored", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendRow(testRowMultiLine)
		tw.AppendFooter(testFooter)
		tw.SetTitle(testTitle1)
		tw.Style().Title.Colors = text.Colors{text.FgYellow}
		tw.Style().Color = ColorOptions{
			Header:       text.Colors{text.FgRed},
			Row:          text.Colors{text.FgGreen},
			RowAlternate: text.Colors{text.FgHiGreen},
			Footer:       text.Colors{text.FgBlue},
		}
		tw.Style().Options.DoNotColorBordersAndSeparators = true

		compareOutputColored(t, tw.Render(), ""+
			"+---------------------------------------------------------------------+\n"+
			"|\x1b[33m Game of Thrones                                                     \x1b[0m|\n"+
			"+-----+------------+-----------+--------+-----------------------------+\n"+
			"|\x1b[31m   # \x1b[0m|\x1b[31m FIRST NAME \x1b[0m|\x1b[31m LAST NAME \x1b[0m|\x1b[31m SALARY \x1b[0m|\x1b[31m                             \x1b[0m|\n"+
			"+-----+------------+-----------+--------+-----------------------------+\n"+
			"|\x1b[32m   1 \x1b[0m|\x1b[32m Arya       \x1b[0m|\x1b[32m Stark     \x1b[0m|\x1b[32m   3000 \x1b[0m|\x1b[32m                             \x1b[0m|\n"+
			"|\x1b[92m  20 \x1b[0m|\x1b[92m Jon        \x1b[0m|\x1b[92m Snow      \x1b[0m|\x1b[92m   2000 \x1b[0m|\x1b[92m You know nothing, Jon Snow! \x1b[0m|\n"+
			"|\x1b[32m 300 \x1b[0m|\x1b[32m Tyrion     \x1b[0m|\x1b[32m Lannister \x1b[0m|\x1b[32m   5000 \x1b[0m|\x1b[32m                             \x1b[0m|\n"+
			"|\x1b[92m   0 \x1b[0m|\x1b[92m Winter     \x1b[0m|\x1b[92m Is        \x1b[0m|\x1b[92m      0 \x1b[0m|\x1b[92m Coming.                     \x1b[0m|\n"+
			"|\x1b[92m     \x1b[0m|\x1b[92m            \x1b[0m|\x1b[92m           \x1b[0m|\x1b[92m        \x1b[0m|\x1b[92m The North Remembers!        \x1b[0m|\n"+
			"|\x1b[92m     \x1b[0m|\x1b[92m            \x1b[0m|\x1b[92m           \x1b[0m|\x1b[92m        \x1b[0m|\x1b[92m This is known.              \x1b[0m|\n"+
			"+-----+------------+-----------+--------+-----------------------------+\n"+
			"|\x1b[34m     \x1b[0m|\x1b[34m            \x1b[0m|\x1b[34m TOTAL     \x1b[0m|\x1b[34m  10000 \x1b[0m|\x1b[34m                             \x1b[0m|\n"+
			"+-----+------------+-----------+--------+-----------------------------+",
		)
	})

	t.Run("column customizations", func(t *testing.T) {
		tw := NewWriter()
		tw.AppendHeader(testHeader)
		tw.AppendRows(testRows)
		tw.AppendRow(testRowMultiLine)
		tw.AppendFooter(testFooter)
		tw.SetCaption(testCaption)
		tw.SetColumnConfigs([]ColumnConfig{
			{Name: "#", Colors: testColor, ColorsHeader: testColorHiRedBold},
			{Name: "First Name", Colors: testColor, ColorsHeader: testColorHiRedBold},
			{Name: "Last Name", Colors: testColor, ColorsHeader: testColorHiRedBold, ColorsFooter: testColorHiBlueBold},
			{Name: "Salary", Colors: testColor, ColorsHeader: testColorHiRedBold, ColorsFooter: testColorHiBlueBold},
			{Number: 5, Colors: text.Colors{text.FgCyan}},
		})
		tw.SetStyle(StyleRounded)

		compareOutputColored(t, tw.Render(), ""+
			"╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮\n"+
			"│\x1b[91;1m   # \x1b[0m│\x1b[91;1m FIRST NAME \x1b[0m│\x1b[91;1m LAST NAME \x1b[0m│\x1b[91;1m SALARY \x1b[0m│                             │\n"+
			"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n"+
			"│\x1b[32m   1 \x1b[0m│\x1b[32m Arya       \x1b[0m│\x1b[32m Stark     \x1b[0m│\x1b[32m   3000 \x1b[0m│\x1b[36m                             \x1b[0m│\n"+
			"│\x1b[32m  20 \x1b[0m│\x1b[32m Jon        \x1b[0m│\x1b[32m Snow      \x1b[0m│\x1b[32m   2000 \x1b[0m│\x1b[36m You know nothing, Jon Snow! \x1b[0m│\n"+
			"│\x1b[32m 300 \x1b[0m│\x1b[32m Tyrion     \x1b[0m│\x1b[32m Lannister \x1b[0m│\x1b[32m   5000 \x1b[0m│\x1b[36m                             \x1b[0m│\n"+
			"│\x1b[32m   0 \x1b[0m│\x1b[32m Winter     \x1b[0m│\x1b[32m Is        \x1b[0m│\x1b[32m      0 \x1b[0m│\x1b[36m Coming.                     \x1b[0m│\n"+
			"│\x1b[32m     \x1b[0m│\x1b[32m            \x1b[0m│\x1b[32m           \x1b[0m│\x1b[32m        \x1b[0m│\x1b[36m The North Remembers!        \x1b[0m│\n"+
			"│\x1b[32m     \x1b[0m│\x1b[32m            \x1b[0m│\x1b[32m           \x1b[0m│\x1b[32m        \x1b[0m│\x1b[36m This is known.              \x1b[0m│\n"+
			"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n"+
			"│     │            │\x1b[94;1m TOTAL     \x1b[0m│\x1b[94;1m  10000 \x1b[0m│                             │\n"+
			"╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯\n"+
			"A Song of Ice and Fire",
		)
	})

	t.Run("colored table within table", func(t *testing.T) {
		table := Table{}
		table.AppendHeader(testHeader)
		table.AppendRows(testRows)
		table.AppendFooter(testFooter)
		table.SetStyle(StyleColoredBright)
		table.SetIndexColumn(1)

		// colored is simple; render the colored table into another table
		tableOuter := Table{}
		tableOuter.AppendRow(Row{table.Render()})
		tableOuter.SetStyle(StyleRounded)

		compareOutputColored(t, tableOuter.Render(), ""+
			"╭───────────────────────────────────────────────────────────────────╮\n"+
			"│ \x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m │\n"+
			"│ \x1b[106;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m │\n"+
			"│ \x1b[106;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m │\n"+
			"│ \x1b[106;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m │\n"+
			"│ \x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m │\n"+
			"╰───────────────────────────────────────────────────────────────────╯",
		)
	})

	t.Run("colored table within colored table", func(t *testing.T) {
		table := Table{}
		table.AppendHeader(testHeader)
		table.AppendRows(testRows)
		table.AppendFooter(testFooter)
		table.SetStyle(StyleColoredBright)
		table.SetIndexColumn(1)

		// colored is simple; render the colored table into another colored table
		tableOuter := Table{}
		tableOuter.AppendHeader(Row{"Colored Table within a Colored Table"})
		tableOuter.AppendRow(Row{"\n" + table.Render() + "\n"})
		tableOuter.SetColumnConfigs([]ColumnConfig{{Number: 1, AlignHeader: text.AlignCenter}})
		tableOuter.SetStyle(StyleColoredBright)

		compareOutputColored(t, tableOuter.Render(), ""+
			"\x1b[106;30m                COLORED TABLE WITHIN A COLORED TABLE               \x1b[0m\n"+
			"\x1b[107;30m                                                                   \x1b[0m\n"+
			"\x1b[107;30m \x1b[106;30m   # \x1b[0m\x1b[107;30m\x1b[106;30m FIRST NAME \x1b[0m\x1b[107;30m\x1b[106;30m LAST NAME \x1b[0m\x1b[107;30m\x1b[106;30m SALARY \x1b[0m\x1b[107;30m\x1b[106;30m                             \x1b[0m\x1b[107;30m \x1b[0m\n"+
			"\x1b[107;30m \x1b[106;30m   1 \x1b[0m\x1b[107;30m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m\x1b[107;30m                             \x1b[0m\x1b[107;30m \x1b[0m\n"+
			"\x1b[107;30m \x1b[106;30m  20 \x1b[0m\x1b[107;30m\x1b[47;30m Jon        \x1b[0m\x1b[107;30m\x1b[47;30m Snow      \x1b[0m\x1b[107;30m\x1b[47;30m   2000 \x1b[0m\x1b[107;30m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\x1b[107;30m \x1b[0m\n"+
			"\x1b[107;30m \x1b[106;30m 300 \x1b[0m\x1b[107;30m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m\x1b[107;30m                             \x1b[0m\x1b[107;30m \x1b[0m\n"+
			"\x1b[107;30m \x1b[46;30m     \x1b[0m\x1b[107;30m\x1b[46;30m            \x1b[0m\x1b[107;30m\x1b[46;30m TOTAL     \x1b[0m\x1b[107;30m\x1b[46;30m  10000 \x1b[0m\x1b[107;30m\x1b[46;30m                             \x1b[0m\x1b[107;30m \x1b[0m\n"+
			"\x1b[107;30m                                                                   \x1b[0m",
		)
	})

	t.Run("colored table with auto-index", func(t *testing.T) {
		table := Table{}
		table.AppendHeader(testHeader)
		table.AppendRows(testRows)
		table.AppendFooter(testFooter)
		table.SetAutoIndex(true)
		table.SetStyle(StyleColoredDark)
		table.SetTitle(testTitle2)

		compareOutputColored(t, table.Render(), ""+
			"\x1b[106;30;1m When you play the Game of Thrones, you win or you die. There is no \x1b[0m\n"+
			"\x1b[106;30;1m middle ground.                                                     \x1b[0m\n"+
			"\x1b[96;100m   \x1b[0m\x1b[96;100m   # \x1b[0m\x1b[96;100m FIRST NAME \x1b[0m\x1b[96;100m LAST NAME \x1b[0m\x1b[96;100m SALARY \x1b[0m\x1b[96;100m                             \x1b[0m\n"+
			"\x1b[96;100m 1 \x1b[0m\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n"+
			"\x1b[96;100m 2 \x1b[0m\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n"+
			"\x1b[96;100m 3 \x1b[0m\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n"+
			"\x1b[36;100m   \x1b[0m\x1b[36;100m     \x1b[0m\x1b[36;100m            \x1b[0m\x1b[36;100m TOTAL     \x1b[0m\x1b[36;100m  10000 \x1b[0m\x1b[36;100m                             \x1b[0m",
		)
	})
}

func TestTable_Render_ColumnConfigs(t *testing.T) {
	generatePrefixTransformer := func(prefix string) text.Transformer {
		return func(val interface{}) string {
			return fmt.Sprintf("%s%v", prefix, val)
		}
	}
	generateSuffixTransformer := func(suffix string) text.Transformer {
		return func(val interface{}) string {
			return fmt.Sprintf("%v%s", val, suffix)
		}
	}
	salaryTransformer := text.Transformer(func(val interface{}) string {
		if valInt, ok := val.(int); ok {
			return fmt.Sprintf("$ %.2f", float64(valInt)+0.03)
		}
		return strings.Replace(fmt.Sprint(val), "ry", "riii", -1)
	})

	tw := NewWriter()
	tw.AppendHeader(testHeaderMultiLine)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooterMultiLine)
	tw.SetAutoIndex(true)
	tw.SetColumnConfigs([]ColumnConfig{
		{
			Name:              fmt.Sprint(testHeaderMultiLine[1]), // First Name
			Align:             text.AlignRight,
			AlignFooter:       text.AlignRight,
			AlignHeader:       text.AlignRight,
			Colors:            text.Colors{text.BgBlack, text.FgRed},
			ColorsHeader:      text.Colors{text.BgRed, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgRed, text.FgBlack},
			Transformer:       generatePrefixTransformer("(r_"),
			TransformerFooter: generatePrefixTransformer("(f_"),
			TransformerHeader: generatePrefixTransformer("(h_"),
			VAlign:            text.VAlignTop,
			VAlignFooter:      text.VAlignTop,
			VAlignHeader:      text.VAlignTop,
			WidthMax:          10,
		}, {
			Name:              fmt.Sprint(testHeaderMultiLine[2]), // Last Name
			Align:             text.AlignLeft,
			AlignFooter:       text.AlignLeft,
			AlignHeader:       text.AlignLeft,
			Colors:            text.Colors{text.BgBlack, text.FgGreen},
			ColorsHeader:      text.Colors{text.BgGreen, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgGreen, text.FgBlack},
			Transformer:       generateSuffixTransformer("_r)"),
			TransformerFooter: generateSuffixTransformer("_f)"),
			TransformerHeader: generateSuffixTransformer("_h)"),
			VAlign:            text.VAlignMiddle,
			VAlignFooter:      text.VAlignMiddle,
			VAlignHeader:      text.VAlignMiddle,
			WidthMax:          10,
		}, {
			Number:            4, // Salary
			Colors:            text.Colors{text.BgBlack, text.FgBlue},
			ColorsHeader:      text.Colors{text.BgBlue, text.FgBlack, text.Bold},
			ColorsFooter:      text.Colors{text.BgBlue, text.FgBlack},
			Transformer:       salaryTransformer,
			TransformerFooter: salaryTransformer,
			TransformerHeader: salaryTransformer,
			VAlign:            text.VAlignBottom,
			VAlignFooter:      text.VAlignBottom,
			VAlignHeader:      text.VAlignBottom,
			WidthMin:          16,
		}, {
			Name:   "Non-existent Column",
			Colors: text.Colors{text.BgYellow, text.FgHiRed},
		},
	})
	tw.SetStyle(styleTest)

	compareOutputColored(t, tw.Render(), ""+
		"(---^-----^-----------^------------^------------------^-----------------------------)\n"+
		"[< >|<  #>|\x1b[41;30;1m< (H_FIRST>\x1b[0m|\x1b[42;30;1m<LAST      >\x1b[0m|\x1b[44;30;1m<                >\x1b[0m|<                           >]\n"+
		"[< >|<   >|\x1b[41;30;1m<     NAME>\x1b[0m|\x1b[42;30;1m<NAME_H)   >\x1b[0m|\x1b[44;30;1m<        SALARIII>\x1b[0m|<                           >]\n"+
		"{---+-----+-----------+------------+------------------+-----------------------------}\n"+
		"[<1>|<  1>|\x1b[40;31m<  (r_Arya>\x1b[0m|\x1b[40;32m<Stark_r)  >\x1b[0m|\x1b[40;34m<       $ 3000.03>\x1b[0m|<                           >]\n"+
		"[<2>|< 20>|\x1b[40;31m<   (r_Jon>\x1b[0m|\x1b[40;32m<Snow_r)   >\x1b[0m|\x1b[40;34m<       $ 2000.03>\x1b[0m|<You know nothing, Jon Snow!>]\n"+
		"[<3>|<300>|\x1b[40;31m<(r_Tyrion>\x1b[0m|\x1b[40;32m<Lannister_>\x1b[0m|\x1b[40;34m<                >\x1b[0m|<                           >]\n"+
		"[< >|<   >|\x1b[40;31m<         >\x1b[0m|\x1b[40;32m<r)        >\x1b[0m|\x1b[40;34m<       $ 5000.03>\x1b[0m|<                           >]\n"+
		"[<4>|<  0>|\x1b[40;31m<(r_Winter>\x1b[0m|\x1b[40;32m<          >\x1b[0m|\x1b[40;34m<                >\x1b[0m|<Coming.                    >]\n"+
		"[< >|<   >|\x1b[40;31m<         >\x1b[0m|\x1b[40;32m<Is_r)     >\x1b[0m|\x1b[40;34m<                >\x1b[0m|<The North Remembers!       >]\n"+
		"[< >|<   >|\x1b[40;31m<         >\x1b[0m|\x1b[40;32m<          >\x1b[0m|\x1b[40;34m<          $ 0.03>\x1b[0m|<This is known.             >]\n"+
		"{---+-----+-----------+------------+------------------+-----------------------------}\n"+
		"[< >|<   >|\x1b[41;30m<      (F_>\x1b[0m|\x1b[42;30m<TOTAL     >\x1b[0m|\x1b[44;30m<                >\x1b[0m|<                           >]\n"+
		"[< >|<   >|\x1b[41;30m<         >\x1b[0m|\x1b[42;30m<SALARY_F) >\x1b[0m|\x1b[44;30m<      $ 10000.03>\x1b[0m|<                           >]\n"+
		"\\---v-----v-----------v------------v------------------v-----------------------------/",
	)
}

func TestTable_Render_Empty(t *testing.T) {
	tw := NewWriter()
	assert.Empty(t, tw.Render())
}

func TestTable_Render_HiddenColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)

	// ensure sorting is done before hiding the columns
	tw.SortBy([]SortBy{
		{Name: "Salary", Mode: DscNumeric},
	})

	t.Run("no columns hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns(nil))

		compareOutput(t, tw.Render(), `
+-----+------------+-------------+--------+-------------------------------+
|   # | FIRST NAME | LAST NAME   | SALARY |                               |
+-----+------------+-------------+--------+-------------------------------+
| 307 | >>Tyrion   | Lannister<< |   5013 |                               |
|   8 | >>Arya     | Stark<<     |   3013 |                               |
|  27 | >>Jon      | Snow<<      |   2013 | ~You know nothing, Jon Snow!~ |
+-----+------------+-------------+--------+-------------------------------+
|     |            | TOTAL       |  10000 |                               |
+-----+------------+-------------+--------+-------------------------------+`)
	})

	t.Run("every column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0, 1, 2, 3, 4}))

		compareOutput(t, tw.Render(), "")
	})

	t.Run("some columns hidden (1)", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1, 2, 3, 4}))

		compareOutput(t, tw.Render(), `
+-----+
|   # |
+-----+
| 307 |
|   8 |
|  27 |
+-----+
|     |
+-----+`)
	})

	t.Run("some columns hidden (2)", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1, 2, 3}))

		compareOutput(t, tw.Render(), `
+-----+-------------------------------+
|   # |                               |
+-----+-------------------------------+
| 307 |                               |
|   8 |                               |
|  27 | ~You know nothing, Jon Snow!~ |
+-----+-------------------------------+
|     |                               |
+-----+-------------------------------+`)
	})

	t.Run("some columns hidden (3)", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0, 4}))

		compareOutput(t, tw.Render(), `
+------------+-------------+--------+
| FIRST NAME | LAST NAME   | SALARY |
+------------+-------------+--------+
| >>Tyrion   | Lannister<< |   5013 |
| >>Arya     | Stark<<     |   3013 |
| >>Jon      | Snow<<      |   2013 |
+------------+-------------+--------+
|            | TOTAL       |  10000 |
+------------+-------------+--------+`)
	})

	t.Run("first column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{0}))

		compareOutput(t, tw.Render(), `
+------------+-------------+--------+-------------------------------+
| FIRST NAME | LAST NAME   | SALARY |                               |
+------------+-------------+--------+-------------------------------+
| >>Tyrion   | Lannister<< |   5013 |                               |
| >>Arya     | Stark<<     |   3013 |                               |
| >>Jon      | Snow<<      |   2013 | ~You know nothing, Jon Snow!~ |
+------------+-------------+--------+-------------------------------+
|            | TOTAL       |  10000 |                               |
+------------+-------------+--------+-------------------------------+`)
	})

	t.Run("column hidden in the middle", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{1}))

		compareOutput(t, tw.Render(), `
+-----+-------------+--------+-------------------------------+
|   # | LAST NAME   | SALARY |                               |
+-----+-------------+--------+-------------------------------+
| 307 | Lannister<< |   5013 |                               |
|   8 | Stark<<     |   3013 |                               |
|  27 | Snow<<      |   2013 | ~You know nothing, Jon Snow!~ |
+-----+-------------+--------+-------------------------------+
|     | TOTAL       |  10000 |                               |
+-----+-------------+--------+-------------------------------+`)
	})

	t.Run("last column hidden", func(t *testing.T) {
		tw.SetColumnConfigs(generateColumnConfigsWithHiddenColumns([]int{4}))

		compareOutput(t, tw.Render(), `
+-----+------------+-------------+--------+
|   # | FIRST NAME | LAST NAME   | SALARY |
+-----+------------+-------------+--------+
| 307 | >>Tyrion   | Lannister<< |   5013 |
|   8 | >>Arya     | Stark<<     |   3013 |
|  27 | >>Jon      | Snow<<      |   2013 |
+-----+------------+-------------+--------+
|     |            | TOTAL       |  10000 |
+-----+------------+-------------+--------+`)
	})
}

func TestTable_Render_Paged(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(Row{"", "", "Total", 10000})
	tw.SetPageSize(1)

	compareOutput(t, tw.Render(), `
+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   1 | Arya       | Stark     |   3000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
| 300 | Tyrion     | Lannister |   5000 |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|   0 | Winter     | Is        |      0 | Coming.                     |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            |           |        | The North Remembers!        |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+

+-----+------------+-----------+--------+-----------------------------+
|   # | FIRST NAME | LAST NAME | SALARY |                             |
+-----+------------+-----------+--------+-----------------------------+
|     |            |           |        | This is known.              |
+-----+------------+-----------+--------+-----------------------------+
|     |            | TOTAL     |  10000 |                             |
+-----+------------+-----------+--------+-----------------------------+`)
}

func TestTable_Render_Reset(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)

	compareOutput(t, tw.Render(), `
┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │ Stark     │   3000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`)

	tw.ResetFooters()
	compareOutput(t, tw.Render(), `
┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │ Stark     │   3000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`)

	tw.ResetHeaders()
	compareOutput(t, tw.Render(), `
┌─────┬────────┬───────────┬──────┬─────────────────────────────┐
│   1 │ Arya   │ Stark     │ 3000 │                             │
│  20 │ Jon    │ Snow      │ 2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion │ Lannister │ 5000 │                             │
└─────┴────────┴───────────┴──────┴─────────────────────────────┘`)

	tw.ResetRows()
	assert.Empty(t, tw.Render())
}

func TestTable_Render_RowPainter(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(testRowMultiLine)
	tw.AppendFooter(testFooter)
	tw.SetIndexColumn(1)
	tw.SetRowPainter(func(row Row) text.Colors {
		if salary, ok := row[3].(int); ok {
			if salary > 3000 {
				return text.Colors{text.BgYellow, text.FgBlack}
			} else if salary < 2000 {
				return text.Colors{text.BgRed, text.FgBlack}
			}
		}
		return nil
	})
	tw.SetStyle(StyleLight)
	tw.SortBy([]SortBy{{Name: "Salary", Mode: AscNumeric}})

	expectedOutLines := []string{
		"┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐",
		"│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│   0 │\x1b[41;30m Winter     \x1b[0m│\x1b[41;30m Is        \x1b[0m│\x1b[41;30m      0 \x1b[0m│\x1b[41;30m Coming.                     \x1b[0m│",
		"│     │\x1b[41;30m            \x1b[0m│\x1b[41;30m           \x1b[0m│\x1b[41;30m        \x1b[0m│\x1b[41;30m The North Remembers!        \x1b[0m│",
		"│     │\x1b[41;30m            \x1b[0m│\x1b[41;30m           \x1b[0m│\x1b[41;30m        \x1b[0m│\x1b[41;30m This is known.              \x1b[0m│",
		"│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │",
		"│   1 │ Arya       │ Stark     │   3000 │                             │",
		"│ 300 │\x1b[43;30m Tyrion     \x1b[0m│\x1b[43;30m Lannister \x1b[0m│\x1b[43;30m   5000 \x1b[0m│\x1b[43;30m                             \x1b[0m│",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │ TOTAL     │  10000 │                             │",
		"└─────┴────────────┴───────────┴────────┴─────────────────────────────┘",
	}
	expectedOut := strings.Join(expectedOutLines, "\n")
	assert.Equal(t, expectedOut, tw.Render())

	tw.SetStyle(StyleColoredBright)
	tw.Style().Color.RowAlternate = tw.Style().Color.Row
	expectedOutLines = []string{
		"\x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m",
		"\x1b[106;30m   0 \x1b[0m\x1b[41;30m Winter     \x1b[0m\x1b[41;30m Is        \x1b[0m\x1b[41;30m      0 \x1b[0m\x1b[41;30m Coming.                     \x1b[0m",
		"\x1b[106;30m     \x1b[0m\x1b[41;30m            \x1b[0m\x1b[41;30m           \x1b[0m\x1b[41;30m        \x1b[0m\x1b[41;30m The North Remembers!        \x1b[0m",
		"\x1b[106;30m     \x1b[0m\x1b[41;30m            \x1b[0m\x1b[41;30m           \x1b[0m\x1b[41;30m        \x1b[0m\x1b[41;30m This is known.              \x1b[0m",
		"\x1b[106;30m  20 \x1b[0m\x1b[107;30m Jon        \x1b[0m\x1b[107;30m Snow      \x1b[0m\x1b[107;30m   2000 \x1b[0m\x1b[107;30m You know nothing, Jon Snow! \x1b[0m",
		"\x1b[106;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m",
		"\x1b[106;30m 300 \x1b[0m\x1b[43;30m Tyrion     \x1b[0m\x1b[43;30m Lannister \x1b[0m\x1b[43;30m   5000 \x1b[0m\x1b[43;30m                             \x1b[0m",
		"\x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m",
	}
	expectedOut = strings.Join(expectedOutLines, "\n")
	assert.Equal(t, expectedOut, tw.Render())
}

func TestTable_Render_Sorted(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)
	tw.SortBy([]SortBy{{Name: "Last Name", Mode: Asc}, {Name: "First Name", Mode: Asc}})

	compareOutput(t, tw.Render(), `┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│   1 │ Arya       │ Stark     │   3000 │                             │
│  11 │ Sansa      │ Stark     │   6000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`)
}

func TestTable_Render_Separator(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendRows(testRows)
	tw.AppendSeparator()
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendRow(testRowMultiLine)
	tw.AppendSeparator()
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendRow(Row{11, "Sansa", "Stark", 6000})
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendSeparator() // doesn't make any difference
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)

	compareOutput(t, tw.Render(), `
┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │ Stark     │   3000 │                             │
│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │ Lannister │   5000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   0 │ Winter     │ Is        │      0 │ Coming.                     │
│     │            │           │        │ The North Remembers!        │
│     │            │           │        │ This is known.              │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│  11 │ Sansa      │ Stark     │   6000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`)
}

func TestTable_Render_Styles(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetStyle(StyleLight)

	styles := map[*Style]string{
		&StyleDefault:                    "+-----+------------+-----------+--------+-----------------------------+\n|   # | FIRST NAME | LAST NAME | SALARY |                             |\n+-----+------------+-----------+--------+-----------------------------+\n|   1 | Arya       | Stark     |   3000 |                             |\n|  20 | Jon        | Snow      |   2000 | You know nothing, Jon Snow! |\n| 300 | Tyrion     | Lannister |   5000 |                             |\n+-----+------------+-----------+--------+-----------------------------+\n|     |            | TOTAL     |  10000 |                             |\n+-----+------------+-----------+--------+-----------------------------+",
		&StyleBold:                       "┏━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n┃   # ┃ FIRST NAME ┃ LAST NAME ┃ SALARY ┃                             ┃\n┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n┃   1 ┃ Arya       ┃ Stark     ┃   3000 ┃                             ┃\n┃  20 ┃ Jon        ┃ Snow      ┃   2000 ┃ You know nothing, Jon Snow! ┃\n┃ 300 ┃ Tyrion     ┃ Lannister ┃   5000 ┃                             ┃\n┣━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫\n┃     ┃            ┃ TOTAL     ┃  10000 ┃                             ┃\n┗━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛",
		&StyleColoredBlackOnBlueWhite:    "\x1b[104;30m   # \x1b[0m\x1b[104;30m FIRST NAME \x1b[0m\x1b[104;30m LAST NAME \x1b[0m\x1b[104;30m SALARY \x1b[0m\x1b[104;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[44;30m     \x1b[0m\x1b[44;30m            \x1b[0m\x1b[44;30m TOTAL     \x1b[0m\x1b[44;30m  10000 \x1b[0m\x1b[44;30m                             \x1b[0m",
		&StyleColoredBlackOnCyanWhite:    "\x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m",
		&StyleColoredBlackOnGreenWhite:   "\x1b[102;30m   # \x1b[0m\x1b[102;30m FIRST NAME \x1b[0m\x1b[102;30m LAST NAME \x1b[0m\x1b[102;30m SALARY \x1b[0m\x1b[102;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[42;30m     \x1b[0m\x1b[42;30m            \x1b[0m\x1b[42;30m TOTAL     \x1b[0m\x1b[42;30m  10000 \x1b[0m\x1b[42;30m                             \x1b[0m",
		&StyleColoredBlackOnMagentaWhite: "\x1b[105;30m   # \x1b[0m\x1b[105;30m FIRST NAME \x1b[0m\x1b[105;30m LAST NAME \x1b[0m\x1b[105;30m SALARY \x1b[0m\x1b[105;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[45;30m     \x1b[0m\x1b[45;30m            \x1b[0m\x1b[45;30m TOTAL     \x1b[0m\x1b[45;30m  10000 \x1b[0m\x1b[45;30m                             \x1b[0m",
		&StyleColoredBlackOnRedWhite:     "\x1b[101;30m   # \x1b[0m\x1b[101;30m FIRST NAME \x1b[0m\x1b[101;30m LAST NAME \x1b[0m\x1b[101;30m SALARY \x1b[0m\x1b[101;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[41;30m     \x1b[0m\x1b[41;30m            \x1b[0m\x1b[41;30m TOTAL     \x1b[0m\x1b[41;30m  10000 \x1b[0m\x1b[41;30m                             \x1b[0m",
		&StyleColoredBlackOnYellowWhite:  "\x1b[103;30m   # \x1b[0m\x1b[103;30m FIRST NAME \x1b[0m\x1b[103;30m LAST NAME \x1b[0m\x1b[103;30m SALARY \x1b[0m\x1b[103;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[43;30m     \x1b[0m\x1b[43;30m            \x1b[0m\x1b[43;30m TOTAL     \x1b[0m\x1b[43;30m  10000 \x1b[0m\x1b[43;30m                             \x1b[0m",
		&StyleColoredBlueWhiteOnBlack:    "\x1b[94;100m   # \x1b[0m\x1b[94;100m FIRST NAME \x1b[0m\x1b[94;100m LAST NAME \x1b[0m\x1b[94;100m SALARY \x1b[0m\x1b[94;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[34;100m     \x1b[0m\x1b[34;100m            \x1b[0m\x1b[34;100m TOTAL     \x1b[0m\x1b[34;100m  10000 \x1b[0m\x1b[34;100m                             \x1b[0m",
		&StyleColoredBright:              "\x1b[106;30m   # \x1b[0m\x1b[106;30m FIRST NAME \x1b[0m\x1b[106;30m LAST NAME \x1b[0m\x1b[106;30m SALARY \x1b[0m\x1b[106;30m                             \x1b[0m\n\x1b[107;30m   1 \x1b[0m\x1b[107;30m Arya       \x1b[0m\x1b[107;30m Stark     \x1b[0m\x1b[107;30m   3000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[47;30m  20 \x1b[0m\x1b[47;30m Jon        \x1b[0m\x1b[47;30m Snow      \x1b[0m\x1b[47;30m   2000 \x1b[0m\x1b[47;30m You know nothing, Jon Snow! \x1b[0m\n\x1b[107;30m 300 \x1b[0m\x1b[107;30m Tyrion     \x1b[0m\x1b[107;30m Lannister \x1b[0m\x1b[107;30m   5000 \x1b[0m\x1b[107;30m                             \x1b[0m\n\x1b[46;30m     \x1b[0m\x1b[46;30m            \x1b[0m\x1b[46;30m TOTAL     \x1b[0m\x1b[46;30m  10000 \x1b[0m\x1b[46;30m                             \x1b[0m",
		&StyleColoredCyanWhiteOnBlack:    "\x1b[96;100m   # \x1b[0m\x1b[96;100m FIRST NAME \x1b[0m\x1b[96;100m LAST NAME \x1b[0m\x1b[96;100m SALARY \x1b[0m\x1b[96;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[36;100m     \x1b[0m\x1b[36;100m            \x1b[0m\x1b[36;100m TOTAL     \x1b[0m\x1b[36;100m  10000 \x1b[0m\x1b[36;100m                             \x1b[0m",
		&StyleColoredDark:                "\x1b[96;100m   # \x1b[0m\x1b[96;100m FIRST NAME \x1b[0m\x1b[96;100m LAST NAME \x1b[0m\x1b[96;100m SALARY \x1b[0m\x1b[96;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[36;100m     \x1b[0m\x1b[36;100m            \x1b[0m\x1b[36;100m TOTAL     \x1b[0m\x1b[36;100m  10000 \x1b[0m\x1b[36;100m                             \x1b[0m",
		&StyleColoredGreenWhiteOnBlack:   "\x1b[92;100m   # \x1b[0m\x1b[92;100m FIRST NAME \x1b[0m\x1b[92;100m LAST NAME \x1b[0m\x1b[92;100m SALARY \x1b[0m\x1b[92;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[32;100m     \x1b[0m\x1b[32;100m            \x1b[0m\x1b[32;100m TOTAL     \x1b[0m\x1b[32;100m  10000 \x1b[0m\x1b[32;100m                             \x1b[0m",
		&StyleColoredMagentaWhiteOnBlack: "\x1b[95;100m   # \x1b[0m\x1b[95;100m FIRST NAME \x1b[0m\x1b[95;100m LAST NAME \x1b[0m\x1b[95;100m SALARY \x1b[0m\x1b[95;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[35;100m     \x1b[0m\x1b[35;100m            \x1b[0m\x1b[35;100m TOTAL     \x1b[0m\x1b[35;100m  10000 \x1b[0m\x1b[35;100m                             \x1b[0m",
		&StyleColoredRedWhiteOnBlack:     "\x1b[91;100m   # \x1b[0m\x1b[91;100m FIRST NAME \x1b[0m\x1b[91;100m LAST NAME \x1b[0m\x1b[91;100m SALARY \x1b[0m\x1b[91;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[31;100m     \x1b[0m\x1b[31;100m            \x1b[0m\x1b[31;100m TOTAL     \x1b[0m\x1b[31;100m  10000 \x1b[0m\x1b[31;100m                             \x1b[0m",
		&StyleColoredYellowWhiteOnBlack:  "\x1b[93;100m   # \x1b[0m\x1b[93;100m FIRST NAME \x1b[0m\x1b[93;100m LAST NAME \x1b[0m\x1b[93;100m SALARY \x1b[0m\x1b[93;100m                             \x1b[0m\n\x1b[97;40m   1 \x1b[0m\x1b[97;40m Arya       \x1b[0m\x1b[97;40m Stark     \x1b[0m\x1b[97;40m   3000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[37;40m  20 \x1b[0m\x1b[37;40m Jon        \x1b[0m\x1b[37;40m Snow      \x1b[0m\x1b[37;40m   2000 \x1b[0m\x1b[37;40m You know nothing, Jon Snow! \x1b[0m\n\x1b[97;40m 300 \x1b[0m\x1b[97;40m Tyrion     \x1b[0m\x1b[97;40m Lannister \x1b[0m\x1b[97;40m   5000 \x1b[0m\x1b[97;40m                             \x1b[0m\n\x1b[33;100m     \x1b[0m\x1b[33;100m            \x1b[0m\x1b[33;100m TOTAL     \x1b[0m\x1b[33;100m  10000 \x1b[0m\x1b[33;100m                             \x1b[0m",
		&StyleDouble:                     "╔═════╦════════════╦═══════════╦════════╦═════════════════════════════╗\n║   # ║ FIRST NAME ║ LAST NAME ║ SALARY ║                             ║\n╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣\n║   1 ║ Arya       ║ Stark     ║   3000 ║                             ║\n║  20 ║ Jon        ║ Snow      ║   2000 ║ You know nothing, Jon Snow! ║\n║ 300 ║ Tyrion     ║ Lannister ║   5000 ║                             ║\n╠═════╬════════════╬═══════════╬════════╬═════════════════════════════╣\n║     ║            ║ TOTAL     ║  10000 ║                             ║\n╚═════╩════════════╩═══════════╩════════╩═════════════════════════════╝",
		&StyleLight:                      "┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐\n│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│   1 │ Arya       │ Stark     │   3000 │                             │\n│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │\n│ 300 │ Tyrion     │ Lannister │   5000 │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│     │            │ TOTAL     │  10000 │                             │\n└─────┴────────────┴───────────┴────────┴─────────────────────────────┘",
		&StyleRounded:                    "╭─────┬────────────┬───────────┬────────┬─────────────────────────────╮\n│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│   1 │ Arya       │ Stark     │   3000 │                             │\n│  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │\n│ 300 │ Tyrion     │ Lannister │   5000 │                             │\n├─────┼────────────┼───────────┼────────┼─────────────────────────────┤\n│     │            │ TOTAL     │  10000 │                             │\n╰─────┴────────────┴───────────┴────────┴─────────────────────────────╯",
		&styleTest:                       "(-----^------------^-----------^--------^-----------------------------)\n[<  #>|<FIRST NAME>|<LAST NAME>|<SALARY>|<                           >]\n{-----+------------+-----------+--------+-----------------------------}\n[<  1>|<Arya      >|<Stark    >|<  3000>|<                           >]\n[< 20>|<Jon       >|<Snow     >|<  2000>|<You know nothing, Jon Snow!>]\n[<300>|<Tyrion    >|<Lannister>|<  5000>|<                           >]\n{-----+------------+-----------+--------+-----------------------------}\n[<   >|<          >|<TOTAL    >|< 10000>|<                           >]\n\\-----v------------v-----------v--------v-----------------------------/",
	}
	var mismatches []string
	for style, expectedOut := range styles {
		tw.SetStyle(*style)
		out := tw.Render()
		assert.Equal(t, expectedOut, out)
		if expectedOut != out {
			mismatches = append(mismatches, fmt.Sprintf("&%s: %#v,", style.Name, out))
			fmt.Printf("// %s renders a Table like below:\n", style.Name)
			for _, line := range strings.Split(out, "\n") {
				fmt.Printf("//  %s\n", line)
			}
			fmt.Println()
		}
	}
	sort.Strings(mismatches)
	for _, mismatch := range mismatches {
		fmt.Println(mismatch)
	}
}

func TestTable_Render_SuppressEmptyColumns(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows([]Row{
		{1, "Arya", "", 3000},
		{20, "Jon", "", 2000, "You know nothing, Jon Snow!"},
		{300, "Tyrion", "", 5000},
	})
	tw.AppendRow(Row{11, "Sansa", "", 6000})
	tw.AppendFooter(Row{"", "", "TOTAL", 10000})
	tw.SetStyle(StyleLight)

	compareOutput(t, tw.Render(), `
┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ LAST NAME │ SALARY │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │           │   3000 │                             │
│  20 │ Jon        │           │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │           │   5000 │                             │
│  11 │ Sansa      │           │   6000 │                             │
├─────┼────────────┼───────────┼────────┼─────────────────────────────┤
│     │            │ TOTAL     │  10000 │                             │
└─────┴────────────┴───────────┴────────┴─────────────────────────────┘`)

	tw.SuppressEmptyColumns()
	compareOutput(t, tw.Render(), `
┌─────┬────────────┬────────┬─────────────────────────────┐
│   # │ FIRST NAME │ SALARY │                             │
├─────┼────────────┼────────┼─────────────────────────────┤
│   1 │ Arya       │   3000 │                             │
│  20 │ Jon        │   2000 │ You know nothing, Jon Snow! │
│ 300 │ Tyrion     │   5000 │                             │
│  11 │ Sansa      │   6000 │                             │
├─────┼────────────┼────────┼─────────────────────────────┤
│     │            │  10000 │                             │
└─────┴────────────┴────────┴─────────────────────────────┘`)
}

func TestTable_Render_TableWithinTable(t *testing.T) {
	twInner := NewWriter()
	twInner.AppendHeader(testHeader)
	twInner.AppendRows(testRows)
	twInner.AppendFooter(testFooter)
	twInner.SetStyle(StyleLight)

	twOuter := NewWriter()
	twOuter.AppendHeader(Row{"Table within a Table"})
	twOuter.AppendRow(Row{twInner.Render()})
	twOuter.SetColumnConfigs([]ColumnConfig{{Number: 1, AlignHeader: text.AlignCenter}})
	twOuter.SetStyle(StyleDouble)

	compareOutput(t, twOuter.Render(), `
╔═════════════════════════════════════════════════════════════════════════╗
║                           TABLE WITHIN A TABLE                          ║
╠═════════════════════════════════════════════════════════════════════════╣
║ ┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐ ║
║ │   # │ FIRST NAME │ LAST NAME │ SALARY │                             │ ║
║ ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤ ║
║ │   1 │ Arya       │ Stark     │   3000 │                             │ ║
║ │  20 │ Jon        │ Snow      │   2000 │ You know nothing, Jon Snow! │ ║
║ │ 300 │ Tyrion     │ Lannister │   5000 │                             │ ║
║ ├─────┼────────────┼───────────┼────────┼─────────────────────────────┤ ║
║ │     │            │ TOTAL     │  10000 │                             │ ║
║ └─────┴────────────┴───────────┴────────┴─────────────────────────────┘ ║
╚═════════════════════════════════════════════════════════════════════════╝`)
}

func TestTable_Render_TableWithTransformers(t *testing.T) {
	bolden := func(val interface{}) string {
		return text.Bold.Sprint(val)
	}
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetColumnConfigs([]ColumnConfig{{
		Name:              "Salary",
		Transformer:       bolden,
		TransformerFooter: bolden,
		TransformerHeader: bolden,
	}})
	tw.SetStyle(StyleLight)

	expectedOut := []string{
		"┌─────┬────────────┬───────────┬────────┬─────────────────────────────┐",
		"│   # │ FIRST NAME │ LAST NAME │ \x1b[1mSALARY\x1b[0m │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│   1 │ Arya       │ Stark     │   \x1b[1m3000\x1b[0m │                             │",
		"│  20 │ Jon        │ Snow      │   \x1b[1m2000\x1b[0m │ You know nothing, Jon Snow! │",
		"│ 300 │ Tyrion     │ Lannister │   \x1b[1m5000\x1b[0m │                             │",
		"├─────┼────────────┼───────────┼────────┼─────────────────────────────┤",
		"│     │            │ TOTAL     │  \x1b[1m10000\x1b[0m │                             │",
		"└─────┴────────────┴───────────┴────────┴─────────────────────────────┘",
	}
	out := tw.Render()
	assert.Equal(t, strings.Join(expectedOut, "\n"), out)
	if strings.Join(expectedOut, "\n") != out {
		for _, line := range strings.Split(out, "\n") {
			fmt.Printf("%#v,\n", line)
		}
	}
}

func TestTable_Render_SetWidth_Title(t *testing.T) {
	tw := NewWriter()
	tw.AppendHeader(testHeader)
	tw.AppendRows(testRows)
	tw.AppendFooter(testFooter)
	tw.SetTitle("Game Of Thrones")

	t.Run("length 20", func(t *testing.T) {
		tw.SetAllowedRowLength(20)

		expectedOut := []string{
			"+------------------+",
			"| Game Of Thrones  |",
			"+-----+----------- ~",
			"|   # | FIRST NAME ~",
			"+-----+----------- ~",
			"|   1 | Arya       ~",
			"|  20 | Jon        ~",
			"| 300 | Tyrion     ~",
			"+-----+----------- ~",
			"|     |            ~",
			"+-----+----------- ~",
		}

		assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
	})

	t.Run("length 30", func(t *testing.T) {
		tw.SetAllowedRowLength(30)

		expectedOut := []string{
			"+----------------------------+",
			"| Game Of Thrones            |",
			"+-----+------------+-------- ~",
			"|   # | FIRST NAME | LAST NA ~",
			"+-----+------------+-------- ~",
			"|   1 | Arya       | Stark   ~",
			"|  20 | Jon        | Snow    ~",
			"| 300 | Tyrion     | Lannist ~",
			"+-----+------------+-------- ~",
			"|     |            | TOTAL   ~",
			"+-----+------------+-------- ~",
		}

		assert.Equal(t, strings.Join(expectedOut, "\n"), tw.Render())
	})
}

func TestTable_Render_WidthEnforcer(t *testing.T) {
	tw := NewWriter()
	tw.AppendRows([]Row{
		{"U2", "Hey", "2021-04-19 13:37", "Yuh yuh yuh"},
		{"S12", "Uhhhh", "2021-04-19 13:37", "Some dummy data here"},
		{"R123", "Lobsters", "2021-04-19 13:37", "I like lobsters"},
		{"R123", "Some big name here and it's pretty big", "2021-04-19 13:37", "Abcdefghijklmnopqrstuvwxyz"},
		{"R123", "Small name", "2021-04-19 13:37", "Abcdefghijklmnopqrstuvwxyz"},
	})
	tw.SetColumnConfigs([]ColumnConfig{
		{Number: 2, WidthMax: 20, WidthMaxEnforcer: text.Trim},
	})

	compareOutput(t, tw.Render(), `
+------+----------------------+------------------+----------------------------+
| U2   | Hey                  | 2021-04-19 13:37 | Yuh yuh yuh                |
| S12  | Uhhhh                | 2021-04-19 13:37 | Some dummy data here       |
| R123 | Lobsters             | 2021-04-19 13:37 | I like lobsters            |
| R123 | Some big name here a | 2021-04-19 13:37 | Abcdefghijklmnopqrstuvwxyz |
| R123 | Small name           | 2021-04-19 13:37 | Abcdefghijklmnopqrstuvwxyz |
+------+----------------------+------------------+----------------------------+`)
}
