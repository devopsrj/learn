

// A general program to maintain the catalog of items in book store Comic Mischief. Challenged by CodeAcademy Trainning 

package main

import "fmt"

func main(){
  
  var publisher,writer,artist,title string
  var year,pageNumber int
  var grade float32
  
  //Assign Variables to First STDOUT
  title = "Mr. GotoSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  
  // First STDOUT
  
  fmt.Println(title, "written by",writer,"drawn by",artist, "Publisher: ",publisher, "Year: ",year,"pageNumber: ",pageNumber,"Grade: ",grade)
  
  //Assign Variables to Secondary STDOUT
  title = "Epic Vol. 1"
  writer= "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0
  
  
  // Secondary STDOUT
  
    fmt.Println(title, "written by",writer,"drawn by",artist, "Publisher: ",publisher, "Year: ",year,"pageNumber: ",pageNumber,"Grade: ",grade)
}
