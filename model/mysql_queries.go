package model

const (
   CheckCredentials = `SELECT username, password
                       FROM logins
                       WHERE username=?
                       AND password=?`

)
