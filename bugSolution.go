func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    if r.Method == "POST" {
        // ... other code ...
        defer r.Body.Close()
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // ... process request body ...
    }
}