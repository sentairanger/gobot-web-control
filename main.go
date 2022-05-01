package main

import (
    "net/http"
    "html/template"
    "github.com/stianeikeland/go-rpio/v4"
    "github.com/sirupsen/logrus"
    "os"
    "fmt"
    
)

// Define the pins for the robot
var (

    pin1 = rpio.Pin(7)
    pin2 = rpio.Pin(8)
    pin3 = rpio.Pin(9)
    pin4 = rpio.Pin(10)
)

// Define and load the template file for the app
func motorControl(w http.ResponseWriter, r *http.Request){
    t := template.Must(template.ParseFiles("index.html"))
    t.Execute(w, nil)
}

// Robot moves forward
func forward(w http.ResponseWriter, r *http.Request) {
     pin1.Low()
     pin2.High()
     pin3.High()
     pin4.Low()

}

// Robot moves backward
func backward(w http.ResponseWriter, r *http.Request) {
    pin1.High()
    pin2.Low()
    pin3.Low()
    pin4.High()
}

// Robot moves left
func left(w http.ResponseWriter, r *http.Request) {
    pin1.High()
    pin2.Low()
    pin3.High()
    pin4.Low()


}

// Robot moves right
func right(w http.ResponseWriter, r *http.Request) {
    pin1.Low()
    pin2.High()
    pin3.Low()
    pin4.High()

}

// Robot stops
func stop(w http.ResponseWriter, r *http.Request) {
    pin1.Low()
    pin2.Low()
    pin3.Low()
    pin4.Low()

}


func main() {
    // Detect any errors in the gpio
    if err := rpio.Open(); err != nil {
            fmt.Println(err)
            os.Exit(1)
    
    }
    
    // Unmap gpio memory when done
    defer rpio.Close()
    
    // Set pins to output mode
    pin1.Output()
    pin2.Output()
    pin3.Output()
    pin4.Output()
    // Create and handle the routes for each direction
    logrus.Println("Robot ready to go")
    http.HandleFunc("/", motorControl)
    http.HandleFunc("/forward", forward)
    http.HandleFunc("/backward", backward)
    http.HandleFunc("/left", left)
    http.HandleFunc("/right", right)
    http.HandleFunc("/stop", stop )
    // Handle the CSS and jQuery files here
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    // Run app in port 6111 across all hosts
    http.ListenAndServe("0.0.0.0:6111", nil)
}
