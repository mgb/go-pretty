package table

import (
	"testing"

	"github.com/mgb/go-pretty/text"
)

func TestTable_Render_BiDiText(t *testing.T) {
	table := Table{}
	table.AppendHeader(Row{"תאריך", "סכום", "מחלקה", "תגים"})
	table.AppendRow(Row{"2020-01-01", 5.0, "מחלקה1", []string{"תג1", "תג2"}})
	table.AppendRow(Row{"2021-02-01", 5.0, "מחלקה1", []string{"תג1"}})
	table.AppendRow(Row{"2022-03-01", 5.0, "מחלקה2", []string{"תג1"}})
	table.AppendFooter(Row{"סהכ", 30})
	table.SetAutoIndex(true)

	//table.Style().Format.Direction = text.Default
	compareOutput(t, table.Render(), `
+---+------------+------+--------+-----------+
|   | תאריך      | סכום | מחלקה  | תגים      |
+---+------------+------+--------+-----------+
| 1 | 2020-01-01 |    5 | מחלקה1 | [תג1 תג2] |
| 2 | 2021-02-01 |    5 | מחלקה1 | [תג1]     |
| 3 | 2022-03-01 |    5 | מחלקה2 | [תג1]     |
+---+------------+------+--------+-----------+
|   | סהכ        |   30 |        |           |
+---+------------+------+--------+-----------+`)

	table.Style().Format.Direction = text.LeftToRight
	compareOutput(t, table.Render(), `
‪+---+------------+------+--------+-----------+
‪|   | ‪תאריך      | ‪סכום | ‪מחלקה  | ‪תגים      |
‪+---+------------+------+--------+-----------+
‪| 1 | ‪2020-01-01 |    ‪5 | ‪מחלקה1 | ‪[תג1 תג2] |
‪| 2 | ‪2021-02-01 |    ‪5 | ‪מחלקה1 | ‪[תג1]     |
‪| 3 | ‪2022-03-01 |    ‪5 | ‪מחלקה2 | ‪[תג1]     |
‪+---+------------+------+--------+-----------+
‪|   | ‪סהכ        |   ‪30 |        |           |
‪+---+------------+------+--------+-----------+`)

	table.Style().Format.Direction = text.RightToLeft
	compareOutput(t, table.Render(), `
‫+---+------------+------+--------+-----------+
‫|   | ‫תאריך      | ‫סכום | ‫מחלקה  | ‫תגים      |
‫+---+------------+------+--------+-----------+
‫| 1 | ‫2020-01-01 |    ‫5 | ‫מחלקה1 | ‫[תג1 תג2] |
‫| 2 | ‫2021-02-01 |    ‫5 | ‫מחלקה1 | ‫[תג1]     |
‫| 3 | ‫2022-03-01 |    ‫5 | ‫מחלקה2 | ‫[תג1]     |
‫+---+------------+------+--------+-----------+
‫|   | ‫סהכ        |   ‫30 |        |           |
‫+---+------------+------+--------+-----------+`)
}
