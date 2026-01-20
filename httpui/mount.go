package httpui

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/epxhsid/pginspect/engine"
)

//go:embed ui.html
var uiHTML []byte

func Mount(mux *http.ServeMux, prefix string, eng engine.Engine) {
	mux.HandleFunc(prefix+"/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		html := string(uiHTML)
		html = strings.ReplaceAll(html, "/__db/", prefix+"/")
		w.Write([]byte(html))
	})

	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, prefix+"/", http.StatusPermanentRedirect)
	})

	mux.HandleFunc(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
		if err := eng.Ping(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("OK"))
	})

	mux.HandleFunc(prefix+"/schemas", func(w http.ResponseWriter, r *http.Request) {
		schemas, err := eng.Schemas(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, s := range schemas {
			fmt.Fprintf(w, `<li hx-get="%s/tables?schema=%s" hx-target="#tables">%s</li>`, prefix, s, s)
		}
	})

	mux.HandleFunc(prefix+"/tables", func(w http.ResponseWriter, r *http.Request) {
		schema := r.URL.Query().Get("schema")
		fmt.Println("Tables handler hit for schema:", schema)
		tables, err := eng.Tables(r.Context(), schema)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, t := range tables {
			fmt.Fprintf(w, `<li hx-get="%s/tabledata?schema=%s&table=%s" hx-target="#tabledata">%s</li>`,
				prefix, schema, t, t)
		}
	})

	mux.HandleFunc(prefix+"/tabledata", func(w http.ResponseWriter, r *http.Request) {
		schema := r.URL.Query().Get("schema")
		table := r.URL.Query().Get("table")
		data, err := eng.TableData(r.Context(), schema, table, 100) // 100 = example row limit
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "<table><thead><tr>")
		for _, col := range data.Columns {
			fmt.Fprintf(w, "<th>%s</th>", col)
		}
		fmt.Fprint(w, "</tr></thead><tbody>")
		for _, row := range data.Rows {
			fmt.Fprint(w, "<tr>")
			for _, val := range row {
				fmt.Fprintf(w, "<td>%v</td>", val)
			}
			fmt.Fprint(w, "</tr>")
		}
		fmt.Fprint(w, "</tbody></table>")
	})

}
