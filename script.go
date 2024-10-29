package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
   )

   type Task struct {
    Text      string
    Completed bool
   }

func showMenu(){
 fmt.Println("\nMenu:")
 fmt.Println("1. Show Tasks")
 fmt.Println("2. Add Task")
 fmt.Println("3. Mark Task as Completed")
 fmt.Println("4. Save Tasks to File")
 fmt.Println("5. Exit")
} 

func showTasks(tasks []Task){
    if len(tasks) == 0 {
        fmt.Println("No task available");
        return;
    }

   for i,s := range tasks {
       status:= "not comopleted"
       if s.Completed == true {
         status = "completed"
       }
       fmt.Printf("Task is number %d => %s  and status [%s] \n",i+1,s.Text,status)
   }

}

func getUserInput(prompt string) string {
  reader :=  bufio.NewReader(os.Stdin);
   fmt.Printf(prompt)
  input, _ := reader.ReadString('\n')
   return strings.TrimSpace(input)

}

func addTask(tasks *[]Task){
    text := getUserInput("Enter task description:")
    *tasks = append(*tasks, Task{Text: text})
   fmt.Println("Task added");
}

func markTaskCompleted(tasks *[]Task){
    showTasks(*tasks);
    num:= getUserInput("Mark task as completed by entering number:")
    taskIndex,err := strconv.Atoi(num)
    if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
   fmt.Println("Invalid task number. Please try again.")
   return
    }

    (*tasks)[taskIndex-1].Completed = true
    fmt.Println("Task marked completed") 

}

func saveTasksToFile(tasks []Task){
    file, err := os.Create("tasks.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
       }
       defer file.Close()
       for _, task := range tasks {
        status := "not completed"
        if task.Completed {
         status = "completed"
        }
        file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
       }
       fmt.Println("Tasks saved to file 'tasks.txt'.")
}


func main() {
 tasks := []Task{}

 for {
  showMenu()
  option := getUserInput("Enter your choice number: ")

  switch option {
   case "1":
    showTasks(tasks)
   case "2":
    addTask(&tasks)
   case "3":
    markTaskCompleted(&tasks)
   case "4":
    saveTasksToFile(tasks)
   case "5":
    fmt.Println("Exiting the ToDo application.")
    return
   default:
    fmt.Println("Invalid choice. Please try again.")
  }
 }
}