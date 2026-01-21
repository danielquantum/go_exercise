package main

import "fmt"
import "strconv"
import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            count, _ := strconv.Atoi(r.FormValue("counter"))
            count++
           fmt.Fprintf(w, `
            <body>
                <form action="/" method="POST">
                    <label>Counter</label>
                    <input name="counter" value="%d" readonly>
                    <button type="submit">Add</button>
                </form>
                <a href="/">Reset</a>
            </body>
        `, count)
        return
        }

        fmt.Fprintf(w, `
            <body>
                <form action="/" method="POST">
                    <label>Counter</label>
                    <input name="counter" value="1" readonly>
                    <button type="submit">Add</button>
                </form>
            </body>
            <a href="/">Reset</a>
        `)
        })

    http.ListenAndServe(":8088", nil)
}
