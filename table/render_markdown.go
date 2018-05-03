package table

import "strings"

// RenderMarkdown renders the Table in Markdown format. Example:
//  | # | First Name | Last Name | Salary |  |
//  | ---:| --- | --- | ---:| --- |
//  | 1 | Arya | Stark | 3000 |  |
//  | 20 | Jon | Snow | 2000 | You know nothing, Jon Snow! |
//  | 300 | Tyrion | Lannister | 5000 |  |
//  |  |  | Total | 10000 |  |
func (t *Table) RenderMarkdown() string {
	t.init()

	var out strings.Builder
	t.markdownRenderRows(&out, t.rowsHeader, true, false)
	t.markdownRenderRows(&out, t.rows, false, false)
	t.markdownRenderRows(&out, t.rowsFooter, false, true)
	if t.caption != "" {
		out.WriteString("\n_")
		out.WriteString(t.caption)
		out.WriteRune('_')
	}
	return t.render(&out)
}

func (t *Table) markdownRenderRow(out *strings.Builder, row Row, isSeparator bool) {
	if len(row) > 0 {
		// when working on line number 2 or more, insert a newline first
		if out.Len() > 0 {
			out.WriteRune('\n')
		}

		// render each column up to the max. columns seen in all the rows
		out.WriteRune('|')
		for colIdx := range t.maxColumnLengths {
			if isSeparator {
				out.WriteString(t.getAlign(colIdx).MarkdownProperty())
			} else {
				var colStr string
				if colIdx < len(row) {
					colStr = row[colIdx].(string)
				}
				out.WriteRune(' ')
				if strings.Contains(colStr, "|") {
					colStr = strings.Replace(colStr, "|", "\\|", -1)
				}
				if strings.Contains(colStr, "\n") {
					colStr = strings.Replace(colStr, "\n", "<br>", -1)
				}
				out.WriteString(colStr)
				out.WriteRune(' ')
			}
			out.WriteRune('|')
		}
	}
}

func (t *Table) markdownRenderRows(out *strings.Builder, rows []Row, isHeader bool, isFooter bool) {
	if len(rows) > 0 {
		for idx, row := range rows {
			t.markdownRenderRow(out, row, false)
			if idx == len(rows)-1 && isHeader {
				t.markdownRenderRow(out, t.rowSeparator, true)
			}
		}
	}
}