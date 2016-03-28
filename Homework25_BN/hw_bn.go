package main

import (
  "net/http"
  "github.com/nu7hatch/gouuid"
  "html/template"
  "log"
  "encoding/json"
  "crypto/sha256"
  "crypto/hmac"
  "encoding/base64"
  "strings"
  "io"
  "fmt"
)


type User struct {
  Name string
  Age string
}


func set_user(req *http.Request, user *User) string {
  user.Name = req.FormValue("user_name")
  user.Age = req.FormValue("user_age")

  bs, err := json.Marshal(user)
  if err != nil {
    fmt.Println("error: ", err)
  }

  b64 := base64.URLEncoding.EncodeToString(bs)

  return b64
}


func getCode(data string) string {
  h := hmac.New(sha256.New, []byte("ourkey"))
  io.WriteString(h, data)
  return fmt.Sprintf("%x", h.Sum(nil))
}


func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
  tpl, err := template.ParseFiles("hw_bn.html")
  if err != nil {
    log.Fatalln(err)
  }

  user_data := set_user(req, new(User))
  hmac_code := getCode(user_data)

  cookie, err := req.Cookie("session-fino")
  if err != nil {
    id, _ := uuid.NewV4()
    cookie = &http.Cookie{
      Name: "session-fino",
      Value: id.String() + "|" + user_data + "|" + hmac_code,
      // Secure: true,
      HttpOnly: true,
    }
  }

  if req.FormValue("user_name") != "" {
    cookie_id := strings.Split(cookie.Value, "|")
    cookie.Value = cookie_id[0] + "|" + user_data + "|" + hmac_code
  }
  xs := strings.Split(cookie.Value, "|")
  data := xs[1] + "tampered this!"
  code := xs[2]

  encode := getCode(data)
  msg := ""
  if code == encode {
    msg = "code valid: " + code + " == " + encode
  } else {
   msg = "code invalid: " + code + " != " + encode
  }

  http.SetCookie(res, cookie)

  res.Header().Set("Content-Type", "text/html")
  io.WriteString(res, msg)

  tpl.Execute(res, nil)
}


func main() {
  http.HandleFunc("/", serve_the_webpage)
  http.Handle("/favicon.ico", http.NotFoundHandler())

  log.Println("Listening...")
  http.ListenAndServe(":9000", nil)
}
